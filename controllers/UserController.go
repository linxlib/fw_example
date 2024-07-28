package controllers

import (
	"github.com/linxlib/fw"
	"gorm.io/gorm"
)

// UserCrudController
// @Route /user
type UserCrudController struct {
	*CrudController[AdminUser]
}

// Test
// @GET /test
func (u *UserCrudController) Test(ctx *fw.Context) {
	ctx.String(200, "test")
}

func NewUserCrudController(db *gorm.DB) *UserCrudController {
	a := &UserCrudController{
		CrudController: NewCrudController[AdminUser](db),
	}
	return a
}

// AdminUser
// @Body
type AdminUser struct {
	*BaseLong[AdminUser] `gorm:"embedded"`
	LoginName            string `json:"login_name"`
	Password             string `json:"password"`
}
