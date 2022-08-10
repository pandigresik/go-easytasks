package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tasks/app/models"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func Index (c echo.Context) error {	
	fmt.Println(models.GetTasks())
	return c.Render(http.StatusOK, "index", M{
		"data" : models.GetTasks(),
	})	
}

func Form (c echo.Context) error {	
	return c.Render(http.StatusOK, "form", M{})	
}

func FormEdit (c echo.Context) error {	
	taskId, _ := strconv.Atoi(c.Param("id"))	
	fmt.Println(taskId)
	return c.Render(http.StatusOK, "formedit", M{
		"data" : models.GetTask(taskId),
	})	
}

func CreateTask (c echo.Context) error {					
	status, _ := strconv.Atoi(c.FormValue("status"))
	_, err := models.PutTask(c.FormValue("name"), c.FormValue("pic"), c.FormValue("deadline"),  status)
	
	if err == nil {
		return c.Redirect(http.StatusSeeOther, "/index")
	} else {
		return c.HTML(301, err.Error())
	}	
}

func UpdateTask (c echo.Context) error {					
	status, _ := strconv.Atoi(c.FormValue("status"))
	id, _ := strconv.Atoi(c.FormValue("id"))
	_, err := models.EditTask(id, c.FormValue("name"), c.FormValue("pic"), c.FormValue("deadline"),  status)
	
	if err == nil {
		return c.Redirect(http.StatusSeeOther, "/index")
	} else {
		return c.HTML(301, err.Error())
	}	
}

func DeleteTask(c echo.Context) error {	
	taskId, _ := strconv.Atoi(c.Param("id"))
	
	models.DeleteTask(taskId)		

	return c.Redirect(http.StatusSeeOther, "/")
}