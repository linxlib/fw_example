package base

import (
	"errors"
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/models"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"math"
)

func NewCrudController[T models.IBase[uint]](db *gorm.DB) *CrudController[T] {
	c := &CrudController[T]{
		SimpleCrudController: NewSimpleCrudController[uint, T](db),
	}
	return c
}

// CrudController
// a base crud controller for models base on gorm.Model
type CrudController[T models.IBase[uint]] struct {
	*SimpleCrudController[uint, T]
}

// GetByID 根据ID获取
// @GET /{id}
func (c *CrudController[T]) GetByID(ctx *fw.Context, q *IDQuery2) {
	v := new(T)
	err := c.db.Debug().First(v, q.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(fasthttp.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": err.Error(),
				"data":    nil,
			})
		} else {
			ctx.JSON(fasthttp.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
				"data":    nil,
			})
		}
		return
	}
	ctx.JSON(fasthttp.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data":    v,
	})
}

// GetPageList 获取分页
// @GET /list
func (c *CrudController[T]) GetPageList(ctx *fw.Context, q *PageSize2) {
	var vs []T
	var count int64
	d := c.db
	if q.Search != "" {
		d = d.Where("name like ?", "%"+q.Search+"%")
	}
	d1 := d.Model(new(T))
	d.Count(&count)
	d1.Offset((q.Page - 1) * q.Size).Limit(q.Size).Find(&vs)

	//

	ctx.JSON(fasthttp.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data": map[string]interface{}{
			"list":      vs,
			"total":     count,
			"totalPage": math.Ceil(float64(count) / float64(q.Size)),
		},
	})
}

// InsertOrUpdate 插入或修改
// @POST /
func (c *CrudController[T]) InsertOrUpdate(ctx *fw.Context, body *T) {
	var err error
	if _, ok := (*body).GetID(); ok {
		err = c.db.Save(body).Error
	} else {
		err = c.db.Create(body).Error
	}
	if err != nil {
		ctx.JSON(fasthttp.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "ok",
	})

}

// Delete
// @POST /delete
func (c *CrudController[T]) Delete(ctx *fw.Context, d *DeleteBody) {
	var tmp = new(T)
	err := c.db.Delete(tmp, d.IDS).Error
	if err != nil {
		ctx.JSON(fasthttp.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(fasthttp.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "ok",
	})
}
