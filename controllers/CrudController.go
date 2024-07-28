package controllers

import (
	"github.com/linxlib/fw"
	"gorm.io/gorm"
	"net/http"
)

/**
CrudController
*/

func NewCrudController[T IBaseLong[T]](db *gorm.DB) *CrudController[T] {
	c := &CrudController[T]{
		model: NewBaseLong[T](db),
	}
	return c
}

func NewBaseLong[T any](db *gorm.DB) IBaseLong[T] {
	return &BaseLong[T]{
		db: db,
	}
}

// CrudController
// @Controller
// @BaseCrudController
type CrudController[T IBaseLong[T]] struct {
	model IBaseLong[T]
}

// IDQuery
// @Query
type IDQuery struct {
	ID int `query:"id"`
}

// GetOne
// @GET /one
func (c *CrudController[T]) GetOne(ctx *fw.Context, q *IDQuery) {
	v := c.model.GetByID(q.ID)

	ctx.JSON(http.StatusOK, v)
}

// PageSize
// @Query
type PageSize struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

// GetPageList
// @GET /page
func (c *CrudController[T]) GetPageList(ctx *fw.Context, q *PageSize) {
	v := c.model.GetPageList(q.Page, q.Size)
	ctx.JSON(http.StatusOK, v)
}

// Insert
// @POST /insert
func (c *CrudController[T]) Insert(ctx *fw.Context, body *T) {
	c.model.Create(body)
}

type IBaseLong[T any] interface {
	SetDB(*gorm.DB)
	SetBaseLong(*BaseLong[T])
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

func (b *BaseLong[T]) SetBaseLong(b2 *BaseLong[T]) {
	b = b2
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
