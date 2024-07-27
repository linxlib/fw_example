package controllers

import (
	"context"
	"github.com/linxlib/fw"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisController
// @Controller
// @Route /api
type RedisController struct {
}

// Get
// @GET /testRedisWrite
func (c *RedisController) Write(ctx *fw.Context, client *redis.Client) {
	client.Set(context.Background(), "test_redis", "Hello world!", time.Hour)
	ctx.String(200, "write to redis success")

}

// Get
// @GET /testRedisRead
func (c *RedisController) Read(ctx *fw.Context, client *redis.Client) {
	var result = client.Get(context.Background(), "test_redis").Val()
	ctx.String(200, "read from redis: %s", result)

}
