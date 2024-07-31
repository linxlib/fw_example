package controllers

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw_example/base"
	"github.com/linxlib/fw_example/models"
	"gorm.io/gorm"
)

// UserCrud2Controller
// @Route /user2
// @Controller
// @Log
type UserCrud2Controller struct {
	*base.SimpleCrudController[int, AdminUser]
}

// Test
// @GET /test
func (u *UserCrud2Controller) Test(ctx *fw.Context) {
	ctx.String(200, "test")
}

// AdminUser
// @Body
type AdminUser struct {
	*models.Base[int] `gorm:"embedded"`
	LoginName         string `json:"login_name"`
	Password          string `json:"password"`
}

func NewUserCrud2Controller(db *gorm.DB) *UserCrud2Controller {
	a := &UserCrud2Controller{
		SimpleCrudController: base.NewSimpleCrudController[int, AdminUser](db),
	}
	return a
}
