package main

import (
	"cake"
	"log"
	"net/http"
)

func main() {
	web := cake.New()
	web.GET("/", func(c *cake.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	web.GET("/hello", func(c *cake.Context) {
		// /hello?name=admin
		c.String(http.StatusOK, "hello: %s\n", c.Query("name"))
	})
	web.POST("/login", func(c *cake.Context) {
		// /login?name=admin&password=pwd
		c.JSON(http.StatusOK, cake.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	err := web.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
