package models

import "gorm.io/gorm"

type IBaseModel interface {
	GetID() (uint, bool)
}

var _ IBaseModel = (*BaseModel)(nil)

type BaseModel struct {
	gorm.Model
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
