/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package manager

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

import (
	"github.com/sethvargo/go-retry"
)

import (
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/store"
)

type ReadOnlyResourceManager interface {
	Get(context.Context, model.Resource, ...store.GetOptionsFunc) error
	List(context.Context, model.ResourceList, ...store.ListOptionsFunc) error
}

type ResourceManager interface {
	ReadOnlyResourceManager
	Create(context.Context, model.Resource, ...store.CreateOptionsFunc) error
	Update(context.Context, model.Resource, ...store.UpdateOptionsFunc) error
	Delete(context.Context, model.Resource, ...store.DeleteOptionsFunc) error
	DeleteAll(context.Context, model.ResourceList, ...store.DeleteAllOptionsFunc) error
}

func NewResourceManager(store store.ResourceStore) ResourceManager {
	return &resourcesManager{
		Store: store,
	}
}

var _ ResourceManager = &resourcesManager{}

type resourcesManager struct {
	Store store.ResourceStore
}

func (r *resourcesManager) Get(ctx context.Context, resource model.Resource, fs ...store.GetOptionsFunc) error {
	return r.Store.Get(ctx, resource, fs...)
}

func (r *resourcesManager) List(ctx context.Context, list model.ResourceList, fs ...store.ListOptionsFunc) error {
	return r.Store.List(ctx, list, fs...)
}

func (r *resourcesManager) Create(ctx context.Context, resource model.Resource, fs ...store.CreateOptionsFunc) error {
	if err := model.Validate(resource); err != nil {
		return err
	}
	return r.Store.Create(ctx, resource)
}

func (r *resourcesManager) Delete(ctx context.Context, resource model.Resource, fs ...store.DeleteOptionsFunc) error {
	return r.Store.Delete(ctx, resource, fs...)
}

func (r *resourcesManager) DeleteAll(ctx context.Context, list model.ResourceList, fs ...store.DeleteAllOptionsFunc) error {
	return DeleteAllResources(r, ctx, list, fs...)
}

func DeleteAllResources(manager ResourceManager, ctx context.Context, list model.ResourceList, fs ...store.DeleteAllOptionsFunc) error {
	opts := store.NewDeleteAllOptions(fs...)
	if err := manager.List(ctx, list, store.ListByMesh(opts.Mesh)); err != nil {
		return err
	}
	for _, item := range list.GetItems() {
		if err := manager.Delete(ctx, item, store.DeleteBy(model.MetaToResourceKey(item.GetMeta()))); err != nil && !store.IsResourceNotFound(err) {
			return err
		}
	}
	return nil
}

func (r *resourcesManager) Update(ctx context.Context, resource model.Resource, fs ...store.UpdateOptionsFunc) error {
	if err := model.Validate(resource); err != nil {
		return err
	}
	return r.Store.Update(ctx, resource, append(fs, store.ModifiedAt(time.Now()))...)
}

type ConflictRetry struct {
	BaseBackoff   time.Duration
	MaxTimes      uint
	JitterPercent uint
}

type UpsertOpts struct {
	ConflictRetry ConflictRetry
	Transactions  store.Transactions
}

type UpsertFunc func(opts *UpsertOpts)

func WithConflictRetry(baseBackoff time.Duration, maxTimes uint, jitterPercent uint) UpsertFunc {
	return func(opts *UpsertOpts) {
		opts.ConflictRetry.BaseBackoff = baseBackoff
		opts.ConflictRetry.MaxTimes = maxTimes
		opts.ConflictRetry.JitterPercent = jitterPercent
	}
}

func WithTransactions(transactions store.Transactions) UpsertFunc {
	return func(opts *UpsertOpts) {
		opts.Transactions = transactions
	}
}

func NewUpsertOpts(fs ...UpsertFunc) UpsertOpts {
	opts := UpsertOpts{
		Transactions: store.NoTransactions{},
	}
	for _, f := range fs {
		f(&opts)
	}
	return opts
}

var ErrSkipUpsert = errors.New("don't do upsert")

func Upsert(ctx context.Context, manager ResourceManager, key model.ResourceKey, resource model.Resource, fn func(resource model.Resource) error, fs ...UpsertFunc) error {
	opts := NewUpsertOpts(fs...)
	upsert := func(ctx context.Context) error {
		return store.InTx(ctx, opts.Transactions, func(ctx context.Context) error {
			create := false
			err := manager.Get(ctx, resource, store.GetBy(key), store.GetConsistent())
			if err != nil {
				if store.IsResourceNotFound(err) {
					create = true
				} else {
					return err
				}
			}
			if err := fn(resource); err != nil {
				if err == ErrSkipUpsert { // Way to skip inserts when there are no change
					return nil
				}
				return err
			}
			if create {
				return manager.Create(ctx, resource, store.CreateBy(key))
			} else {
				return manager.Update(ctx, resource)
			}
		})
	}

	if opts.ConflictRetry.BaseBackoff <= 0 || opts.ConflictRetry.MaxTimes == 0 {
		return upsert(ctx)
	}
	backoff := retry.NewExponential(opts.ConflictRetry.BaseBackoff)
	backoff = retry.WithMaxRetries(uint64(opts.ConflictRetry.MaxTimes), backoff)
	backoff = retry.WithJitterPercent(uint64(opts.ConflictRetry.JitterPercent), backoff)
	return retry.Do(ctx, backoff, func(ctx context.Context) error {
		resource.SetMeta(nil)
		specType := reflect.TypeOf(resource.GetSpec()).Elem()
		zeroSpec := reflect.New(specType).Interface().(model.ResourceSpec)
		if err := resource.SetSpec(zeroSpec); err != nil {
			return err
		}
		err := upsert(ctx)
		if errors.Is(err, &store.ResourceConflictError{}) {
			return retry.RetryableError(err)
		}
		return err
	})
}

type MeshNotFoundError struct {
	Mesh string
}

func (m *MeshNotFoundError) Error() string {
	return fmt.Sprintf("mesh of name %s is not found", m.Mesh)
}

func MeshNotFound(meshName string) error {
	return &MeshNotFoundError{meshName}
}

func IsMeshNotFound(err error) bool {
	_, ok := err.(*MeshNotFoundError)
	return ok
}
