package base

import (
	"errors"
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/models"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"math"
)

func NewSimpleCrudController[E models.PrimaryKey, T models.IBase[E]](db *gorm.DB) *SimpleCrudController[E, T] {
	c := &SimpleCrudController[E, T]{
		db: db,
	}
	return c
}

// SimpleCrudController
// a base crud controller for models which has one primary field
type SimpleCrudController[E models.PrimaryKey, T models.IBase[E]] struct {
	db *gorm.DB
}

// IDQuery2
// @Path
type IDQuery2 struct {
	ID uint `path:"id"`
}

// GetByID 根据ID获取
// @GET /{id}
func (c *SimpleCrudController[E, T]) GetByID(ctx *fw.Context, q *IDQuery2) {
	v := new(T)
	err := c.db.Debug().First(v, q.ID).Error
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
	ctx.JSON(fasthttp.StatusOK, Resp(200, "ok", v))
}

type IData interface {
	any | ListDataBase[any]
}

type RespBase[E int | string, T IData] struct {
	Code    E      `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func RespInt[T IData](code int, message string, data T) RespBase[int, T] {
	return RespBase[int, T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

}
func Resp[E int | string, T IData](code E, message string, data T) RespBase[E, T] {
	return RespBase[E, T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

}
func RespString[T IData](code string, message string, data T) RespBase[string, T] {
	return RespBase[string, T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

}

type PageSizeBase struct {
	Page int `query:"page" default:"1"`  //页码
	Size int `query:"size" default:"20"` //每页数量
}

func (bps PageSizeBase) Offset() int {
	return (bps.Page - 1) * bps.Size
}

// PageSize2
// @Query
type PageSize2 struct {
	PageSizeBase
	Search string `query:"search" default:""` //搜索名称
}

type ListDataBase[T any] struct {
	List      []T   `json:"list"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"totalPage"`
}

func ListData[T any](list []T, total int64, size int) ListDataBase[T] {
	return ListDataBase[T]{
		List:      list,
		Total:     total,
		TotalPage: int(math.Ceil(float64(total) / float64(size))),
	}
}

// GetPageList 获取分页
// @GET /list
func (c *SimpleCrudController[E, T]) GetPageList(ctx *fw.Context, q *PageSize2) {
	var vs []T
	var count int64
	d := c.db
	if q.Search != "" {
		d = d.Where("name like ?", "%"+q.Search+"%")
	}
	d1 := d.Debug().Model(new(T))
	d1.Count(&count)
	d1.Offset(q.Offset()).Limit(q.Size).Find(&vs)
	ctx.JSON(fasthttp.StatusOK, Resp(200, "ok", ListData(vs, count, q.Size)))
}

// InsertOrUpdate 插入或修改
// @POST /
func (c *SimpleCrudController[E, T]) InsertOrUpdate(ctx *fw.Context, body *T) {
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
func (c *SimpleCrudController[E, T]) Delete(ctx *fw.Context, d *DeleteBody) {
	err := c.db.Delete(new(T), d.IDS).Error
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
