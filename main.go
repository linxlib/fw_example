package main

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw/middlewares"
	"github.com/linxlib/fw_example/controllers"
)

var _ controllers.HelloController

func main() {
	s := fw.New()
	s.Use(middlewares.NewServerDownMiddleware("111"))
	s.Use(middlewares.NewWebsocketMiddleware())
	s.Use(middlewares.NewWebsocketHubMiddleware())
	s.Use(middlewares.NewLoggerMiddleware(fw.Logger{}))

	s.RegisterRoute(new(controllers.HelloController))
	s.RegisterRoute(new(controllers.WebsocketHubController))
	s.Run()
}
