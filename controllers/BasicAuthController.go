package controllers

import (
	"github.com/linxlib/fw"
	"github.com/linxlib/fw/middlewares"
)

// BasicAuthController
// @Controller
// @BasicAuth admin=admin
type BasicAuthController struct {
}

// NeedBasicAuth
// @GET /needAuth
// @Websocket
func (c *BasicAuthController) NeedBasicAuth(ctx *fw.Context) {
	ctx.JSON(200, map[string]interface{}{
		"middleware": "BasicAuthMiddleware",
		"proxy":      false,
		"realm":      "Authorization Required",
		"username":   ctx.GetString(middlewares.AuthUserKey),
	})
}
