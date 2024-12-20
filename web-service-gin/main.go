package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Forever", Artist: "Mystery Skulls", Price: 18.66},
	{ID: "2", Title: "DDC", Artist: "Nothing but thieves", Price: 20.96},
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context){
	var newAlbum album
	
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.POST("/albums", postAlbum)
    router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

