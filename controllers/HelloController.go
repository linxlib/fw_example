package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/linxlib/fw"
	"io"
	"net/http"
	"time"
)

// HelloController 哈罗啊
// @Controller
// @Route /api
// @Log
type HelloController struct {
}

// MyBody
// @Body
type MyBody struct {
	A string `json:"a"`
	B int    `json:"b"`
}

// MyPath
// @Path
type MyPath struct {
	Name string `path:"name"`
}

// MyQuery
// @Query
type MyQuery struct {
	A []string `query:"a"`
	B string   `query:"b"`
	C int      `query:"c"`
}

// @Cookie
type MyCookie struct {
	GfsessionId string `cookie:"gfsessionid"`
}

// @Plain
type MyPlain string

func TestPanic() {
	panic("jhhhh")
}

// Get hhh
// @Logger
// @POST /v1/user/{name}
func (this *HelloController) Get(ctx *fw.Context, body1 *MyBody, path1 *MyPath, p *MyQuery, p1 *MyCookie) {
	fmt.Println("pathname", path1.Name)
	fmt.Println(*p)
	fmt.Println(ctx.RemoteIP())
	//ctx.GetFastContext().Response.SetStatusCode(500)
	if p.C != 88 {
		TestPanic()
	}

	ctx.JSON(200, map[string]interface{}{
		"pathname": path1.Name,
		"body1":    body1,
		"path":     path1,
		"query":    path1,
		"cookie":   p1,
		"ip":       ctx.RemoteIP(),
	})

}

// Websocket
// @Logger
// @Websocket
// @WS /echo
func (this *HelloController) Websocket(ctx *fw.Context, msg []byte) {
	fmt.Println(string(msg))
	//time.Sleep(time.Second)
	ctx.Set("fw_err", errors.New("test errors"))
}

// GetIndex
// @GET /index
func (this *HelloController) GetIndex(ctx *fw.Context) {
	resp, _ := http.Get("https://shuye.dev/maintenance-page/")
	bs, _ := io.ReadAll(resp.Body)
	ctx.Data(200, "text/html", bs)
}

// GetImage
// @GET /Maintenance.png
func (this *HelloController) GetImage(ctx *fw.Context) {
	ctx.File("E:\\repos\\htmldemo\\Maintenance.png")
}

// MyGenericType
// @Body
type MyGenericType[T My] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// TestTypeParam
// @POST /testGenericType
func (this *HelloController) TestTypeParam(ctx *fw.Context, body *MyGenericType[int]) {
	ctx.PureJSON(200, body)
}

// TestStream
// @GET /testStream
func (this *HelloController) TestStream(ctx *fw.Context) {
	ctx.Stream(func(w *bufio.Writer) {
		w.WriteString("1")
		time.Sleep(time.Second)
		w.WriteString("2")
		time.Sleep(time.Second)
		w.WriteString("3")
		time.Sleep(time.Second)
		w.WriteString("4")
		time.Sleep(time.Second)
		w.WriteString("5")
		time.Sleep(time.Second)
		w.WriteString("end. nothing")

	})
}

type My interface {
	~int | ~string | byte
}

func Higher[T My](a, b T) bool {
	return a > b
}
