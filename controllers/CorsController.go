package controllers

import "github.com/linxlib/fw"

// CorsController
// @Controller
// @Route /cors
type CorsController struct {
}

// CorsCheck
// @POST /corsCheck
func (c *CorsController) CorsCheck(ctx *fw.Context) {
	ctx.String(200, "content is protected by cors: %s", ctx.GetFastContext().Host())
}
