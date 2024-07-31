package main

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw/middlewares"
	"github.com/linxlib/fw_example/controllers"
	middlewares2 "github.com/linxlib/fw_example/middlewares"
	"github.com/linxlib/fw_example/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//go:generate go run github.com/linxlib/astp/astpg -o gen.json

var _ controllers.HelloController

func main() {
	//os.Setenv("FW_DEV", "false")
	s := fw.New()
	s.Use(middlewares.NewServerDownMiddleware("111"))
	s.Use(middlewares.NewWebsocketMiddleware())
	s.Use(middlewares.NewWebsocketHubMiddleware())
	var logger = logrus.New()
	s.Provide(logger)
	s.Use(middlewares.NewLoggerMiddleware(logger))
	//s.Use(middlewares.NewResponseRewriteMiddleware())
	s.Use(middlewares.NewRecoveryMiddleware(&middlewares.RecoveryOptions{
		NiceWeb: true,
		Console: true,
		Output:  logger.WriterLevel(logrus.ErrorLevel),
	}))
	s.Use(middlewares.NewBasicAuthMiddleware())
	// connect mysql
	db, err := gorm.Open(mysql.Open("root:root@tcp(10.10.0.16:3306)/wanshengserver?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbmid := middlewares2.NewMySQLMiddleware(db)

	s.Use(dbmid)

	s.Use(middlewares2.NewRedisMiddlewareWithUrl("redis://10.10.0.16:6379/1"))
	s.Use(middlewares.NewCorsMiddleware(middlewares.Config{

		AllowOrigins:     []string{"http://www.example.com:5500"},
		AllowMethods:     []string{"POST", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	s.RegisterRoute(new(controllers.HelloController))
	s.RegisterRoute(new(controllers.WebsocketHubController))
	s.RegisterRoute(new(controllers.BasicAuthController))
	s.RegisterRoute(new(controllers.RedisController))
	s.RegisterRoute(new(controllers.CorsController))
	db.AutoMigrate(new(models.AdminUser))

	s.RegisterRoute(controllers.NewUserCrud2Controller(db))
	s.Map(dbmid.GetDB())
	s.Run()
}
