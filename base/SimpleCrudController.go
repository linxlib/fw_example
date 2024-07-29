package base

import (
	"errors"
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/models"
	"gorm.io/gorm"
	"math"
	"net/http"
)

func NewSimpleCrudController[T models.IBaseModel](db *gorm.DB) *SimpleCrudController[T] {
	c := &SimpleCrudController[T]{
		db: db,
	}
	return c
}

// SimpleCrudController
// @Controller
type SimpleCrudController[T models.IBaseModel] struct {
	db *gorm.DB
}

// IDQuery2
// @Query
type IDQuery2 struct {
	ID uint `path:"id"`
}

// GetByID 根据ID获取
// @GET /{id}
func (c *SimpleCrudController[T]) GetByID(ctx *fw.Context, q *IDQuery2) {
	v := new(T)
	err := c.db.First(v, q.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, map[string]interface{}{
				"code":    404,
				"message": err.Error(),
				"data":    nil,
			})
		} else {
			ctx.JSON(500, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
				"data":    nil,
			})
		}
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data":    v,
	})
}

// PageSize2
// @Query
type PageSize2 struct {
	Page   int    `query:"page" default:"1"`  //页码
	Size   int    `query:"size" default:"20"` //每页数量
	Search string `query:"search" default:""` //搜索名称
}

// GetPageList 获取分页
// @GET /list
func (c *SimpleCrudController[T]) GetPageList(ctx *fw.Context, q *PageSize2) {
	vs := make([]*T, 0)
	var count int64
	d := c.db
	if q.Search != "" {
		d = d.Where("name like ?", "%"+q.Search+"%")
	}
	d.Limit(q.Size).Offset((q.Page - 1) * q.Size).Count(&count).Find(&vs)

	ctx.JSON(http.StatusOK, map[string]interface{}{
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
func (c *SimpleCrudController[T]) InsertOrUpdate(ctx *fw.Context, body *T) {
	var err error
	if _, ok := (*body).GetID(); ok {
		err = c.db.Save(body).Error
	} else {
		err = c.db.Create(body).Error
	}
	if err != nil {
		ctx.JSON(500, map[string]interface{}{
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

// DeleteBody
// @Body
type DeleteBody struct {
	IDS []uint `json:"ids"`
}

// Delete
// @POST /delete
func (c *SimpleCrudController[T]) Delete(ctx *fw.Context, d *DeleteBody) {
	var tmp = new(T)
	err := c.db.Delete(tmp, d.IDS).Error
	if err != nil {
		ctx.JSON(500, map[string]interface{}{
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
