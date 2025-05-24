package handler

import (
	"net/http"
	"vinyl-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(ctx *gin.Context, db *gorm.DB) {
	var albums []models.Album
	db.Find(&albums)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(ctx *gin.Context, db *gorm.DB) {
	var newAlbum models.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumsByID(ctx *gin.Context, db *gorm.DB) {
	id := ctx.Param("id")

	var album models.Album
	result := db.First(&album, "id=?", id)
	if result.Error != nil {
		// si el usuario no encuentra el album
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album no encontrado"})
		return
	}

	// si el usuario encuentra el album
	ctx.IndentedJSON(http.StatusOK, album)
}

func PutAlbumByID(ctx *gin.Context, db *gorm.DB) {
	id := ctx.Param("id")
	var modifyAlbum models.Album // datos ing por el usuario en el body del put

	var album models.Album // datos obtenidos de la BD

	if err := ctx.BindJSON(&modifyAlbum); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Datos incorrectos"})
		return
	}

	result := db.First(&album, "id=?", id) // obtengo los datos de la BD
	if result.Error != nil {
		// si el usuario no encuentra el album
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album no encontrado"})
		return
	}

	// Si el dni se encontró y es existente, debemos sobrescribir
	album.Title = modifyAlbum.Title
	album.Artist = modifyAlbum.Artist
	album.Year = modifyAlbum.Year
	album.Price = modifyAlbum.Price

	db.Save(&album)
	ctx.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbumByID(ctx *gin.Context, db *gorm.DB) {
	id := ctx.Param("id")

	result := db.Delete(&models.Album{}, "id=?", id)
	if result.RowsAffected == 0 {
		// si el usuario no encuentra el album
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album no encontrado"})
		return
	}

	// si el usuario encuentra el album
	ctx.IndentedJSON(http.StatusOK, gin.H{"Message": "Album eliminado"})
}
