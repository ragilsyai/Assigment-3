package main

import (
	"Assigment3/config"
	"Assigment3/controller"
	"Assigment3/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*")

	webService := services.NewWebService()
	webController := controller.NewWebController(webService)

	router.GET("/status", webController.GetStatus)
	//router.Static("/js", "./js")

	router.Run(config.APP_PORT)
}
