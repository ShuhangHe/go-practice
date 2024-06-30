package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动错误: %v", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("hello hanlder, %q\n", req.URL.Path)
}

// index 处理方法
func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		log.Fatalf("write failed: %v", err)
		return
	}
}
