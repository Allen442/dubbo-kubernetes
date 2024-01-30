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

package matchers

import (
	stdErrors "errors"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"strings"
)

import (
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/gstruct/errors"
	"github.com/onsi/gomega/types"
)

type allKeysMatcher struct {
	keyMatcher types.GomegaMatcher
	failures   []error
}

// AllKeys allows you to specify a matcher which all map's keys needs to fulfill
// to make the test statement successful
func AllKeys(matcher types.GomegaMatcher) types.GomegaMatcher {
	return &allKeysMatcher{
		keyMatcher: matcher,
	}
}

func (m *allKeysMatcher) Match(actual interface{}) (bool, error) {
	if reflect.TypeOf(actual).Kind() != reflect.Map {
		return false, fmt.Errorf("%v is type %T, expected map", actual, actual)
	}

	m.failures = m.matchKeys(actual)
	if len(m.failures) > 0 {
		return false, nil
	}
	return true, nil
}

func (m *allKeysMatcher) matchKeys(actual interface{}) []error {
	var errs []error
	actualValue := reflect.ValueOf(actual)
	keys := map[interface{}]bool{}
	for _, keyValue := range actualValue.MapKeys() {
		key := keyValue.Interface()
		keys[key] = true

		err := func() (err error) {
			// This test relies heavily on reflect, which tends to panic.
			// Recover here to provide more useful error messages in that case.
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf(
						"panic checking %+v: %v\n%s",
						actual,
						r,
						debug.Stack(),
					)
				}
			}()

			match, err := m.keyMatcher.Match(key)
			if err != nil {
				return err
			}

			if !match {
				if nesting, ok := m.keyMatcher.(errors.NestingMatcher); ok {
					return errors.AggregateError(nesting.Failures())
				}

				return stdErrors.New(m.keyMatcher.FailureMessage(key))
			}

			return nil
		}()
		if err != nil {
			errs = append(errs, errors.Nest(fmt.Sprintf(".%#v", key), err))
		}
	}

	return errs
}

func (m *allKeysMatcher) FailureMessage(actual interface{}) string {
	failures := make([]string, len(m.failures))

	for i := range m.failures {
		failures[i] = m.failures[i].Error()
	}

	return format.Message(
		actual,
		fmt.Sprintf(
			"to consits of only valid ports <1,%d>: {\n%v\n}\n",
			math.MaxUint16,
			strings.Join(failures, "\n"),
		))
}

func (m *allKeysMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to match keys")
}

func (m *allKeysMatcher) Failures() []error {
	return m.failures
}
