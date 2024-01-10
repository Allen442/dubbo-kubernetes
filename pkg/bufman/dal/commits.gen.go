// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/apache/dubbo-kubernetes/pkg/bufman/model"
)

func newCommit(db *gorm.DB, opts ...gen.DOOption) commit {
	_commit := commit{}

	_commit.commitDo.UseDB(db, opts...)
	_commit.commitDo.UseModel(&model.Commit{})

	tableName := _commit.commitDo.TableName()
	_commit.ALL = field.NewAsterisk(tableName)
	_commit.ID = field.NewInt64(tableName, "id")
	_commit.UserID = field.NewString(tableName, "user_id")
	_commit.UserName = field.NewString(tableName, "user_name")
	_commit.RepositoryID = field.NewString(tableName, "repository_id")
	_commit.RepositoryName = field.NewString(tableName, "repository_name")
	_commit.CommitID = field.NewString(tableName, "commit_id")
	_commit.CommitName = field.NewString(tableName, "commit_name")
	_commit.DraftName = field.NewString(tableName, "draft_name")
	_commit.CreatedTime = field.NewTime(tableName, "created_time")
	_commit.ManifestDigest = field.NewString(tableName, "manifest_digest")
	_commit.BufManConfigDigest = field.NewString(tableName, "buf_man_config_digest")
	_commit.DocumentDigest = field.NewString(tableName, "document_digest")
	_commit.LicenseDigest = field.NewString(tableName, "license_digest")
	_commit.SequenceID = field.NewInt64(tableName, "sequence_id")

	_commit.fillFieldMap()

	return _commit
}

type commit struct {
	commitDo

	ALL                field.Asterisk
	ID                 field.Int64
	UserID             field.String
	UserName           field.String
	RepositoryID       field.String
	RepositoryName     field.String
	CommitID           field.String
	CommitName         field.String
	DraftName          field.String
	CreatedTime        field.Time
	ManifestDigest     field.String
	BufManConfigDigest field.String
	DocumentDigest     field.String
	LicenseDigest      field.String
	SequenceID         field.Int64

	fieldMap map[string]field.Expr
}

func (c commit) Table(newTableName string) *commit {
	c.commitDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commit) As(alias string) *commit {
	c.commitDo.DO = *(c.commitDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commit) updateTableName(table string) *commit {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UserID = field.NewString(table, "user_id")
	c.UserName = field.NewString(table, "user_name")
	c.RepositoryID = field.NewString(table, "repository_id")
	c.RepositoryName = field.NewString(table, "repository_name")
	c.CommitID = field.NewString(table, "commit_id")
	c.CommitName = field.NewString(table, "commit_name")
	c.DraftName = field.NewString(table, "draft_name")
	c.CreatedTime = field.NewTime(table, "created_time")
	c.ManifestDigest = field.NewString(table, "manifest_digest")
	c.BufManConfigDigest = field.NewString(table, "buf_man_config_digest")
	c.DocumentDigest = field.NewString(table, "document_digest")
	c.LicenseDigest = field.NewString(table, "license_digest")
	c.SequenceID = field.NewInt64(table, "sequence_id")

	c.fillFieldMap()

	return c
}

func (c *commit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commit) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["user_name"] = c.UserName
	c.fieldMap["repository_id"] = c.RepositoryID
	c.fieldMap["repository_name"] = c.RepositoryName
	c.fieldMap["commit_id"] = c.CommitID
	c.fieldMap["commit_name"] = c.CommitName
	c.fieldMap["draft_name"] = c.DraftName
	c.fieldMap["created_time"] = c.CreatedTime
	c.fieldMap["manifest_digest"] = c.ManifestDigest
	c.fieldMap["buf_man_config_digest"] = c.BufManConfigDigest
	c.fieldMap["document_digest"] = c.DocumentDigest
	c.fieldMap["license_digest"] = c.LicenseDigest
	c.fieldMap["sequence_id"] = c.SequenceID
}

func (c commit) clone(db *gorm.DB) commit {
	c.commitDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commit) replaceDB(db *gorm.DB) commit {
	c.commitDo.ReplaceDB(db)
	return c
}

type commitDo struct{ gen.DO }

type ICommitDo interface {
	gen.SubQuery
	Debug() ICommitDo
	WithContext(ctx context.Context) ICommitDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommitDo
	WriteDB() ICommitDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommitDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommitDo
	Not(conds ...gen.Condition) ICommitDo
	Or(conds ...gen.Condition) ICommitDo
	Select(conds ...field.Expr) ICommitDo
	Where(conds ...gen.Condition) ICommitDo
	Order(conds ...field.Expr) ICommitDo
	Distinct(cols ...field.Expr) ICommitDo
	Omit(cols ...field.Expr) ICommitDo
	Join(table schema.Tabler, on ...field.Expr) ICommitDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommitDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommitDo
	Group(cols ...field.Expr) ICommitDo
	Having(conds ...gen.Condition) ICommitDo
	Limit(limit int) ICommitDo
	Offset(offset int) ICommitDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommitDo
	Unscoped() ICommitDo
	Create(values ...*model.Commit) error
	CreateInBatches(values []*model.Commit, batchSize int) error
	Save(values ...*model.Commit) error
	First() (*model.Commit, error)
	Take() (*model.Commit, error)
	Last() (*model.Commit, error)
	Find() ([]*model.Commit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Commit, err error)
	FindInBatches(result *[]*model.Commit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Commit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommitDo
	Assign(attrs ...field.AssignExpr) ICommitDo
	Joins(fields ...field.RelationField) ICommitDo
	Preload(fields ...field.RelationField) ICommitDo
	FirstOrInit() (*model.Commit, error)
	FirstOrCreate() (*model.Commit, error)
	FindByPage(offset int, limit int) (result []*model.Commit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommitDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commitDo) Debug() ICommitDo {
	return c.withDO(c.DO.Debug())
}

func (c commitDo) WithContext(ctx context.Context) ICommitDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commitDo) ReadDB() ICommitDo {
	return c.Clauses(dbresolver.Read)
}

func (c commitDo) WriteDB() ICommitDo {
	return c.Clauses(dbresolver.Write)
}

func (c commitDo) Session(config *gorm.Session) ICommitDo {
	return c.withDO(c.DO.Session(config))
}

func (c commitDo) Clauses(conds ...clause.Expression) ICommitDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commitDo) Returning(value interface{}, columns ...string) ICommitDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commitDo) Not(conds ...gen.Condition) ICommitDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commitDo) Or(conds ...gen.Condition) ICommitDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commitDo) Select(conds ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commitDo) Where(conds ...gen.Condition) ICommitDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commitDo) Order(conds ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commitDo) Distinct(cols ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commitDo) Omit(cols ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commitDo) Join(table schema.Tabler, on ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commitDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommitDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commitDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommitDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commitDo) Group(cols ...field.Expr) ICommitDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commitDo) Having(conds ...gen.Condition) ICommitDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commitDo) Limit(limit int) ICommitDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commitDo) Offset(offset int) ICommitDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommitDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commitDo) Unscoped() ICommitDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commitDo) Create(values ...*model.Commit) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commitDo) CreateInBatches(values []*model.Commit, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commitDo) Save(values ...*model.Commit) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commitDo) First() (*model.Commit, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Commit), nil
	}
}

func (c commitDo) Take() (*model.Commit, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Commit), nil
	}
}

func (c commitDo) Last() (*model.Commit, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Commit), nil
	}
}

func (c commitDo) Find() ([]*model.Commit, error) {
	result, err := c.DO.Find()
	return result.([]*model.Commit), err
}

func (c commitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Commit, err error) {
	buf := make([]*model.Commit, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commitDo) FindInBatches(result *[]*model.Commit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commitDo) Attrs(attrs ...field.AssignExpr) ICommitDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commitDo) Assign(attrs ...field.AssignExpr) ICommitDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commitDo) Joins(fields ...field.RelationField) ICommitDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commitDo) Preload(fields ...field.RelationField) ICommitDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commitDo) FirstOrInit() (*model.Commit, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Commit), nil
	}
}

func (c commitDo) FirstOrCreate() (*model.Commit, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Commit), nil
	}
}

func (c commitDo) FindByPage(offset int, limit int) (result []*model.Commit, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commitDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commitDo) Delete(models ...*model.Commit) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commitDo) withDO(do gen.Dao) *commitDo {
	c.DO = *do.(*gen.DO)
	return c
}
