package middlewares

import (
	"github.com/linxlib/fw"
	"gorm.io/gorm"
)

type MySQLMiddleware struct {
	*fw.MiddlewareGlobal
	db *gorm.DB
}

func NewMySQLMiddleware(db *gorm.DB) *MySQLMiddleware {
	m := &MySQLMiddleware{
		MiddlewareGlobal: fw.NewMiddlewareGlobal("gorm_mysql"),
		db:               db}

	return m
}
func (m *MySQLMiddleware) GetDB() *gorm.DB {
	return m.db
}

func (m *MySQLMiddleware) CloneAsMethod() fw.IMiddlewareMethod {
	return m.CloneAsCtl()
}

func (m *MySQLMiddleware) HandlerMethod(h fw.HandlerFunc) fw.HandlerFunc {

	return func(context *fw.Context) {
		context.Map(m.db)
		h(context)
	}
}

func (m *MySQLMiddleware) CloneAsCtl() fw.IMiddlewareCtl {
	return NewMySQLMiddleware(m.db)
}

func (m *MySQLMiddleware) HandlerController(string) []*fw.RouteItem {
	return fw.EmptyRouteItem(m)
}
