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

func newCommitFile(db *gorm.DB, opts ...gen.DOOption) commitFile {
	_commitFile := commitFile{}

	_commitFile.commitFileDo.UseDB(db, opts...)
	_commitFile.commitFileDo.UseModel(&model.CommitFile{})

	tableName := _commitFile.commitFileDo.TableName()
	_commitFile.ALL = field.NewAsterisk(tableName)
	_commitFile.ID = field.NewInt64(tableName, "id")
	_commitFile.Digest = field.NewString(tableName, "digest")
	_commitFile.CommitID = field.NewString(tableName, "commit_id")
	_commitFile.FileName = field.NewString(tableName, "file_name")

	_commitFile.fillFieldMap()

	return _commitFile
}

type commitFile struct {
	commitFileDo

	ALL      field.Asterisk
	ID       field.Int64
	Digest   field.String
	CommitID field.String
	FileName field.String

	fieldMap map[string]field.Expr
}

func (c commitFile) Table(newTableName string) *commitFile {
	c.commitFileDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commitFile) As(alias string) *commitFile {
	c.commitFileDo.DO = *(c.commitFileDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commitFile) updateTableName(table string) *commitFile {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Digest = field.NewString(table, "digest")
	c.CommitID = field.NewString(table, "commit_id")
	c.FileName = field.NewString(table, "file_name")

	c.fillFieldMap()

	return c
}

func (c *commitFile) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commitFile) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["id"] = c.ID
	c.fieldMap["digest"] = c.Digest
	c.fieldMap["commit_id"] = c.CommitID
	c.fieldMap["file_name"] = c.FileName
}

func (c commitFile) clone(db *gorm.DB) commitFile {
	c.commitFileDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commitFile) replaceDB(db *gorm.DB) commitFile {
	c.commitFileDo.ReplaceDB(db)
	return c
}

type commitFileDo struct{ gen.DO }

type ICommitFileDo interface {
	gen.SubQuery
	Debug() ICommitFileDo
	WithContext(ctx context.Context) ICommitFileDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommitFileDo
	WriteDB() ICommitFileDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommitFileDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommitFileDo
	Not(conds ...gen.Condition) ICommitFileDo
	Or(conds ...gen.Condition) ICommitFileDo
	Select(conds ...field.Expr) ICommitFileDo
	Where(conds ...gen.Condition) ICommitFileDo
	Order(conds ...field.Expr) ICommitFileDo
	Distinct(cols ...field.Expr) ICommitFileDo
	Omit(cols ...field.Expr) ICommitFileDo
	Join(table schema.Tabler, on ...field.Expr) ICommitFileDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommitFileDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommitFileDo
	Group(cols ...field.Expr) ICommitFileDo
	Having(conds ...gen.Condition) ICommitFileDo
	Limit(limit int) ICommitFileDo
	Offset(offset int) ICommitFileDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommitFileDo
	Unscoped() ICommitFileDo
	Create(values ...*model.CommitFile) error
	CreateInBatches(values []*model.CommitFile, batchSize int) error
	Save(values ...*model.CommitFile) error
	First() (*model.CommitFile, error)
	Take() (*model.CommitFile, error)
	Last() (*model.CommitFile, error)
	Find() ([]*model.CommitFile, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommitFile, err error)
	FindInBatches(result *[]*model.CommitFile, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CommitFile) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommitFileDo
	Assign(attrs ...field.AssignExpr) ICommitFileDo
	Joins(fields ...field.RelationField) ICommitFileDo
	Preload(fields ...field.RelationField) ICommitFileDo
	FirstOrInit() (*model.CommitFile, error)
	FirstOrCreate() (*model.CommitFile, error)
	FindByPage(offset int, limit int) (result []*model.CommitFile, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommitFileDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commitFileDo) Debug() ICommitFileDo {
	return c.withDO(c.DO.Debug())
}

func (c commitFileDo) WithContext(ctx context.Context) ICommitFileDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commitFileDo) ReadDB() ICommitFileDo {
	return c.Clauses(dbresolver.Read)
}

func (c commitFileDo) WriteDB() ICommitFileDo {
	return c.Clauses(dbresolver.Write)
}

func (c commitFileDo) Session(config *gorm.Session) ICommitFileDo {
	return c.withDO(c.DO.Session(config))
}

func (c commitFileDo) Clauses(conds ...clause.Expression) ICommitFileDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commitFileDo) Returning(value interface{}, columns ...string) ICommitFileDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commitFileDo) Not(conds ...gen.Condition) ICommitFileDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commitFileDo) Or(conds ...gen.Condition) ICommitFileDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commitFileDo) Select(conds ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commitFileDo) Where(conds ...gen.Condition) ICommitFileDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commitFileDo) Order(conds ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commitFileDo) Distinct(cols ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commitFileDo) Omit(cols ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commitFileDo) Join(table schema.Tabler, on ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commitFileDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commitFileDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commitFileDo) Group(cols ...field.Expr) ICommitFileDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commitFileDo) Having(conds ...gen.Condition) ICommitFileDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commitFileDo) Limit(limit int) ICommitFileDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commitFileDo) Offset(offset int) ICommitFileDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commitFileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommitFileDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commitFileDo) Unscoped() ICommitFileDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commitFileDo) Create(values ...*model.CommitFile) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commitFileDo) CreateInBatches(values []*model.CommitFile, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commitFileDo) Save(values ...*model.CommitFile) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commitFileDo) First() (*model.CommitFile, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitFile), nil
	}
}

func (c commitFileDo) Take() (*model.CommitFile, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitFile), nil
	}
}

func (c commitFileDo) Last() (*model.CommitFile, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitFile), nil
	}
}

func (c commitFileDo) Find() ([]*model.CommitFile, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommitFile), err
}

func (c commitFileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommitFile, err error) {
	buf := make([]*model.CommitFile, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commitFileDo) FindInBatches(result *[]*model.CommitFile, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commitFileDo) Attrs(attrs ...field.AssignExpr) ICommitFileDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commitFileDo) Assign(attrs ...field.AssignExpr) ICommitFileDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commitFileDo) Joins(fields ...field.RelationField) ICommitFileDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commitFileDo) Preload(fields ...field.RelationField) ICommitFileDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commitFileDo) FirstOrInit() (*model.CommitFile, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitFile), nil
	}
}

func (c commitFileDo) FirstOrCreate() (*model.CommitFile, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitFile), nil
	}
}

func (c commitFileDo) FindByPage(offset int, limit int) (result []*model.CommitFile, count int64, err error) {
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

func (c commitFileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commitFileDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commitFileDo) Delete(models ...*model.CommitFile) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commitFileDo) withDO(do gen.Dao) *commitFileDo {
	c.DO = *do.(*gen.DO)
	return c
}
