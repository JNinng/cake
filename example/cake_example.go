package main

import (
	"cake"
	"log"
	"net/http"
	"time"
)

func main() {
	web := cake.Default()
	web.LoadHTMLGlob("./example/templates/*")
	web.Static("/assets", "./example/static")
	web.GET("/", func(c *cake.Context) {
		c.HTML(http.StatusOK, "home.tmpl", cake.H{
			"title": "Home",
			"date":  time.Now().Format(time.DateTime),
		})
	})
	web.GET("/hello", func(c *cake.Context) {
		// /hello?name=admin
		c.String(http.StatusOK, "hello: %s\n", c.Query("name"))
	})
	web.GET("/user/:id", func(c *cake.Context) {
		c.String(http.StatusOK, "user id: %s\n", c.Param("id"))
	})
	web.GET("/file/user/123/*.png", func(c *cake.Context) {
		c.String(http.StatusOK, "file: %s\n", c.Param(".png"))
	})
	web.GET("/out", func(c *cake.Context) {
		out := []string{""}
		c.JSON(http.StatusOK, cake.H{
			"out": out[len(out)+1],
		})
	})
	web.POST("/login", func(c *cake.Context) {
		// /login?name=admin&password=pwd
		c.JSON(http.StatusOK, cake.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	v1 := web.Group("/api/v1")
	v1.Use(V1Logger())
	{
		v1.GET("/user/:id", func(c *cake.Context) {
			c.JSON(http.StatusOK, cake.H{
				"id":      c.Param("id"),
				"version": "v1",
			})
		})
	}
	v2 := web.Group("/api/v2")
	v2.Use(V2Logger())
	{
		v2.GET("/user/:id", func(c *cake.Context) {
			c.JSON(http.StatusOK, cake.H{
				"id":      c.Param("id"),
				"version": "v2",
			})
		})
	}
	err := web.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}

func V1Logger() cake.HandlerFunc {
	return func(c *cake.Context) {
		log.Printf("[v1] %s\n", c.Path)
		c.Next()
	}
}

func V2Logger() cake.HandlerFunc {
	return func(c *cake.Context) {
		log.Printf("[v2] %s\n", c.Path)
		c.Next()
	}
}
