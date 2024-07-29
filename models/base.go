package models

import (
	"go/constant"
	"go/token"
	"gorm.io/gorm"
)

type IBaseModel interface {
	GetID() (uint, bool)
}

type PrimaryKey interface {
	comparable
}
type IBase[T PrimaryKey] interface {
	GetID() (T, bool)
	GetSearchColumns() map[string]string
}
type Base[T PrimaryKey] struct {
	ID T `gorm:"primaryKey" json:"id"`
}

func (b *Base[T]) GetID() (T, bool) {
	v := constant.Make(b.ID)
	switch v.Kind() {
	case constant.Int:
		if constant.Compare(v, token.EQL, constant.Make(0)) {
			return b.ID, true
		}
		return b.ID, false
	case constant.String:
		if constant.Compare(v, token.EQL, constant.MakeString("")) {
			return b.ID, true
		}
		return b.ID, false
	default:
		return b.ID, false
	}
}

var _ IBaseModel = (*BaseModel)(nil)

type BaseModel struct {
	gorm.Model
}

type BaseModel2 struct {
	ID uint `gorm:"primarykey" json:"id"`
}

func (b *BaseModel2) GetID() (uint, bool) {
	if b.ID == 0 {
		return 0, false
	}
	return b.ID, true
}

func (b *BaseModel) GetID() (uint, bool) {
	if b.ID == 0 {
		return 0, false
	}
	return b.ID, true
}

func NewBaseLong[T any](db *gorm.DB) IBaseLong[T] {
	return &BaseLong[T]{
		db: db,
	}
}

type IBaseLong[T any] interface {
	SetDB(*gorm.DB)
	GetID() int64
	GetByID(id int) *T
	GetPageList(page int, pageSize int) []*T
	Create(obj *T)
}

var _ IBaseLong[AdminUser] = (*BaseLong[AdminUser])(nil)

type BaseLong[T any] struct {
	db *gorm.DB
	ID int64 `gorm:"primaryKey" json:"id"`
}

func (b *BaseLong[T]) SetDB(db *gorm.DB) {
	b.db = db
}
func (b *BaseLong[T]) Create(obj *T) {
	b.db.Create(&obj)
}

func (b *BaseLong[T]) GetPageList(page int, pageSize int) []*T {
	r := make([]*T, 0)
	b.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&r)
	return r
}

func (b *BaseLong[T]) GetByID(id int) *T {
	r := new(T)
	b.db.Where("id=?", id).First(r)
	return r
}

func (b *BaseLong[T]) GetID() int64 {
	return b.ID
}
