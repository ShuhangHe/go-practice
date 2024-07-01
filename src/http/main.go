package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	// 初始化engine
	engine := gee.New()
	// 添加一个get方法
	engine.GET("/", func(context *gee.Context) {
		err := context.Html(http.StatusOK, "<h1>Hello Aligads</h1>")
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	})

	engine.GET("/hello", func(context *gee.Context) {
		err := context.String(http.StatusOK, "Hello Aligads", "Hello World!")
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	})

	engine.GET("/hello/:name", func(context *gee.Context) {
		context.Json(http.StatusOK, gee.MyMap{"name": context.Params["name"], "age": 123})
	})

	engine.GET("/assets/*filepath", func(context *gee.Context) {
		err := context.Data(http.StatusOK, []byte(context.Params["filepath"]))
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
