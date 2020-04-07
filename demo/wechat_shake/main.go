package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

func newApp() *iris.Application {
	app := iris.Default()
	//使用中间件
	app.Use(middleware)
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	
	return app
}
func middleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
func main() {
	app := newApp()
	InitLog()
	InitGift()
	app.Get("/ping", func(context iris.Context) {
		_, _ = context.JSON(iris.Map{"message": "pong"})
	})
	_ = app.Run(iris.Server(&http.Server{Addr: ":8081"}))
}
