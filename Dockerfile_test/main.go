package main

import (
	"Dockerfile_test/web/controlleres"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"

	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Use(logger.New())
	//创建mvc
	mvc.New(app).Handle(controlleres.NewMovieController())

	app.Run(iris.Addr(":8080"))
}
