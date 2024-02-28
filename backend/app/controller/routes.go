package controller

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/anime", InsertAnime)
		v1.GET("/animes", FetchAllAnimes)
		v1.GET("/animes/:id", FetchAnime)
	}

	router.Run(":3333")
}
