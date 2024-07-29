package base

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/models"
	"gorm.io/gorm"
	"net/http"
)

/**
CrudController
*/

func NewCrudController[T models.IBaseLong[T]](db *gorm.DB) *CrudController[T] {
	c := &CrudController[T]{
		model: models.NewBaseLong[T](db),
	}
	return c
}

// CrudController
// @Controller
// @BaseCrudController
type CrudController[T models.IBaseLong[T]] struct {
	model models.IBaseLong[T]
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
