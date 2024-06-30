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

	engine.GET("/go", func(context *gee.Context) {
		context.Json(http.StatusOK, gee.MyMap{"name": "cxx", "age": 123})
	})

	engine.GET("/data", func(context *gee.Context) {
		err := context.Data(http.StatusOK, []byte("DATA HELLO"))
		if err != nil {
			context.Error(http.StatusInternalServerError, err)
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
