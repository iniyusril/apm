package main

import (
	"apm/controller"
	"apm/db"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4/v2"
)

func main() {

	db.Init()
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/panic", func(c echo.Context) error {
		return errors.New("hahahaha")
	})

	g := e.Group("/api/v1/blog")

	g.POST("", controller.InsertBlog)

	e.Logger.Fatal(e.Start(":1323"))
}
