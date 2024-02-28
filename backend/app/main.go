package main

import (
	"github.com/NogueiraMat/app/app/controller"
	"github.com/NogueiraMat/app/app/database"
	"github.com/NogueiraMat/app/app/models"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Anime{})

	controller.SetRouter()
}
