package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
)

func newApp() *iris.Application {
	app := iris.Default()
	
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}
func main() {
	app := newApp()
	err := app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatal(err)
	}
}
