package models

// AdminUser
// @Body
type AdminUser struct {
	*Base[int] `gorm:"embedded"`
	LoginName  string `json:"login_name" gorm:"column:login_name"`
	Password   string `json:"password" gorm:"column:password"`
}
