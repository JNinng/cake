package main

import (
	"cake"
	"fmt"
	"net/http"
)

func main() {
	web := cake.New()
	web.GET("/", func(writer http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(writer, "hello world")
	})
	err := web.Run(":8080")
	if err != nil {
		return
	}
}
