package main

import (
	"log"
	"time"
	database "vinyl-api/db"
	"vinyl-api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // o 3000
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/albums", func(ctx *gin.Context) {
		handler.GetAlbums(ctx, db)
	})

	router.GET("/albums/:id", func(ctx *gin.Context) {
		handler.GetAlbumsByID(ctx, db)
	})

	router.POST("/albums", func(ctx *gin.Context) {
		handler.PostAlbums(ctx, db)
	})

	router.PUT("/albums/:id", func(ctx *gin.Context) {
		handler.PutAlbumByID(ctx, db)
	})

	router.DELETE("/albums/:id", func(ctx *gin.Context) {
		handler.DeleteAlbumByID(ctx, db)
	})

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumsByID)
	// router.POST("/albums", postAlbums)
	// router.PUT("/albums/:id", putAlbumByID) // muy básico, falta hacer más cosas
	// router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
