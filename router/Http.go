package router

import (
	"tasks/app/controllers"

	"github.com/labstack/echo/v4"
)

func Http() *echo.Echo{
	e := Config()
	e.GET("/", controllers.Index)
	e.GET("/index", controllers.Index)
	e.GET("/form", controllers.Form)
	e.GET("/editform/:id", controllers.FormEdit)
	e.POST("/task", controllers.CreateTask)
	e.POST("/taskEdit", controllers.UpdateTask)
	e.GET("/deleteTask/:id", controllers.DeleteTask)
	return e
}