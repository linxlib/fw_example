package controllers

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/base"
	"github.com/linxlib/fw_example/models"
	"gorm.io/gorm"
)

// UserCrudController
// @Route /user
type UserCrudController struct {
	*base.CrudController[models.AdminUser]
}

// Test
// @GET /test
func (u *UserCrudController) Test(ctx *fw.Context) {
	ctx.String(200, "test")
}

func NewUserCrudController(db *gorm.DB) *UserCrudController {
	a := &UserCrudController{
		CrudController: base.NewCrudController[models.AdminUser](db),
	}
	return a
}

// UserCrud2Controller
// @Route /user2
type UserCrud2Controller struct {
	*base.SimpleCrudController[AdminUser]
}

// Test
// @GET /test
func (u *UserCrud2Controller) Test(ctx *fw.Context) {
	ctx.String(200, "test")
}

// AdminUser
// @Body
type AdminUser struct {
	*models.BaseModel `gorm:"embedded"`
	LoginName         string `json:"login_name"`
	Password          string `json:"password"`
}

func NewUserCrud2Controller(db *gorm.DB) *UserCrud2Controller {
	a := &UserCrud2Controller{
		SimpleCrudController: base.NewSimpleCrudController[AdminUser](db),
	}
	return a
}
