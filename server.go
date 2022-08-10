package main

import (
	"tasks/app/models"
	"tasks/router"

	"github.com/labstack/gommon/log"
)


func main() {	
	models.InitDB("storage.db")	
	e := router.Http()	
	e.Debug = true	
	e.Logger.SetLevel(log.INFO)	
	e.Logger.Fatal(e.Start(":8880"))
}
