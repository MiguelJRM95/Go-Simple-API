package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/web-service-gin/config"
)

// albums slice to seed record album data.
//var albumes = []Album{
//    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}
var conf *config.Config

func init() {
    if err := godotenv.Load("local.env"); err != nil {
		log.Print("No .env file found")
    }
}

func main() {
	conf = config.New()
	Conn()

	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbum)
	router.PUT("/albums/:id", UpdateAlbum)
	router.DELETE("/albums/:id", DeleteAlbum)

	router.Run(":8080")
}