package main

import (
	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
//var albumes = []Album{
//    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}

func main() {
	Conn()
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbum)
	router.PUT("/albums/:id", UpdateAlbum)
	router.DELETE("/albums/:id", DeleteAlbum)

	router.Run("localhost:8080")
}