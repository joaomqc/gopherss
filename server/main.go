package main

import (
	"fmt"
	"gopherss/ctrl"
	"gopherss/db"
	"log"
	"os"

	docs "gopherss/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const DEFAULT_PORT = "8080"

//	@title			Gopherss API
//	@version		1.0
//	@description	RSS feed management app
//	@BasePath		/api

var ctrls = []ctrl.Controller{
	ctrl.EntriesController{},
	ctrl.CategoriesController{},
	ctrl.FeedsController{},
}

func main() {
	sqlDb, err := db.GetDb()
	if err != nil {
		log.Fatalln(err)
	}

	err = sqlDb.Initialize()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	apiGroup := r.Group("/api")
	for _, controller := range ctrls {
		controller.Register(apiGroup)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("GOPHERSS_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic("[ERROR] failed to start server: " + err.Error())
	}
	sqlDb.Close()
}
