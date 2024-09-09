package main

import (
	"gopherss/ctrl"
	"gopherss/db"
	"log"

	_ "gopherss/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Gopherss API
//	@version		1.0
//	@description	RSS feed management app
//	@host			localhost:8889
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
	r.Run(":8889")
	sqlDb.Close()
}
