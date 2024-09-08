package main

import (
	"gopherss/ctrl"
	"gopherss/db"
	"log"

	"github.com/gin-gonic/gin"
)

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

	for _, controller := range ctrls {
		controller.Register(r)
	}

	r.Run()
	sqlDb.Close()
}
