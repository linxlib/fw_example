package controllers

import (
	"github.com/linxlib/fw/middlewares"
)

// WebsocketHubController
// @Controller
// @WebsocketHub route=/name
// @Route /hub
type WebsocketHubController struct {
}

// MyDataBody
// @Body
type MyDataBody struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// Broadcast
// @POST /broadcast
func (c *WebsocketHubController) Broadcast(data *MyDataBody, hub *middlewares.Hub) {
	hub.Broadcast([]byte(data.Data))
}

// SendByClient
// @POST /send
func (c *WebsocketHubController) SendByClient(data *MyDataBody, hub *middlewares.Hub) {
	hub.SendByClient(data.ID, []byte(data.Data))
}
