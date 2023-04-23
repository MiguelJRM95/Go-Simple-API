package main

import (
	"strconv"
	"net/http"

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
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	var albums []Album
	var err error
	albums, err = GetAll()
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context){
	
	var album Album
	var err error
	
	id := c.Param("id")

	//for _, a := range albumes{
	//	if a.ID == id {
	//		c.IndentedJSON(http.StatusOK, a)
	//		return
	//	}
	//}

	castedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
    	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not a number"})
		return
	}	

	album, err = GetById(castedId)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func postAlbum(c *gin.Context) {
	var newAlbum NewAlbum
	var savedAlbum Album
	var err error
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failure"})
		return
	}
	//albumes = append(albumes, newAlbum)
	savedAlbum, err = Save(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not save the album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, savedAlbum)
}

func updateAlbum(c *gin.Context) {
	var newAlbum NewAlbum
	var updatedAlbum Album
	var err error

	id := c.Param("id")
	castedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
    	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not a number"})
		return
	}	
	
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failure"})
		return
	}

	updatedAlbum, err = Update(castedId, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not update the album"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}


func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	castedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
    	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not a number"})
		return
	}
	
	_,err = Delete(castedId)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album could not be deleted"})
		return
	}

	c.Status(http.StatusOK)
}