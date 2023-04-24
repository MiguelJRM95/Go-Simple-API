package main

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []Album
	var err error
	albums, err = GetAll()
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context){
	
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

func PostAlbum(c *gin.Context) {
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

func UpdateAlbum(c *gin.Context) {
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


func DeleteAlbum(c *gin.Context) {
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