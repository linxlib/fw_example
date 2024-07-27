package main

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw/middlewares"
	"github.com/linxlib/fw_example/controllers"
	"time"
)

var _ controllers.HelloController

func main() {
	s := fw.New()
	s.Use(middlewares.NewServerDownMiddleware("111"))
	s.Use(middlewares.NewWebsocketMiddleware())
	s.Use(middlewares.NewWebsocketHubMiddleware())
	s.Use(middlewares.NewLogMiddleware(fw.Logger{}))
	//s.Use(middlewares.NewResponseRewriteMiddleware())
	s.Use(middlewares.NewRecoveryMiddleware())
	s.Use(middlewares.NewBasicAuthMiddleware())
	s.Use(middlewares.NewRedisMiddlewareWithUrl("redis://10.10.0.16:6379/1"))
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
	s.Run()
}
