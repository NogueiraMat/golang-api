package controller

import (
	"fmt"
	"net/http"

	"github.com/NogueiraMat/app/app/database"
	"github.com/NogueiraMat/app/app/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertAnime(ctx *gin.Context) {
	var anime models.Anime

	if err := ctx.ShouldBindJSON(&anime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		return
	}

	if anime.Name == nil || anime.Gender == nil || anime.ReleaseDate == nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "all de fields name, gender and release_date must be filled!"})

		return
	}

	*anime.ReleaseDate = anime.ReleaseDate.UTC()
	anime.ID = uuid.New()

	result := database.DB.Create(&anime)
	if result.Error != nil {
		fmt.Println("Erro ao criar o anime:", result.Error)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data..."})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{"data": &anime})
}

func FetchAllAnimes(ctx *gin.Context) {
	var animes []models.Anime

	result := database.DB.Find(&animes)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data..."})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": &animes})
}

func FetchAnime(ctx *gin.Context) {
	var anime models.Anime

	animeID := ctx.Param("id")

	uuidAnimeID, err := uuid.Parse(animeID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format..."})
	}

	result := database.DB.First(&anime, uuidAnimeID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Anime not found..."})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch anime..."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": anime})
}
