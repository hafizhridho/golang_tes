package main

import (
	"berita/configs"
	"berita/controllers"

	"github.com/labstack/echo/v4"
)

func main () {
	configs.Loadenv()
	configs.InitDatabase()
	e := echo.New()
	e.GET("/news", controllers.GetNewsControllers)
	e.POST("/news", controllers.AddNewsController)
	e.DELETE("/news/:id", controllers.DeleteNewsController)
	e.Start(":8000")
}