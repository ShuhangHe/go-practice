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
	engine.GET("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("welcome aligads"))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
	})

	engine.GET("/hello", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("now url: " + req.URL.Path))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
