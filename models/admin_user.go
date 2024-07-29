package models

// AdminUser
// @Body
type AdminUser struct {
	*BaseLong[AdminUser] `gorm:"embedded"`
	LoginName            string `json:"login_name"`
	Password             string `json:"password"`
}
