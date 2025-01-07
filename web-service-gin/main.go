package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	MetaReady   bool   `json:"metaReady"`
	BackupReady bool   `json:"backupReady"`
	Mp3Ready    bool   `json:"mp3Ready"`
}

var albums = []Album{
	{ID: "1", Title: "Forever", Artist: "Mystery Skulls", MetaReady: true, BackupReady: true, Mp3Ready: true},
	{ID: "2", Title: "DDC", Artist: "Nothing but thieves", MetaReady: true, BackupReady: true, Mp3Ready: true},
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func findAlbumIndex(id string) (int, bool) {
	for index, album := range albums {
		if album.ID == id {
			return index, true
		}
	}
	return -1, false
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	if index, isPresent := findAlbumIndex(id); isPresent {
		c.IndentedJSON(http.StatusOK, albums[index])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}
}

func deleteAlbum(c *gin.Context) {
	requestID := c.Param("id")

	if index, isPresent := findAlbumIndex(requestID); isPresent {
		albums = append(albums[:index], albums[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": albums[index].Title + " album deleted "})
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}
}
func updateAlbum(c *gin.Context) {
	requestID := c.Param("id")

	if index, isPresent := findAlbumIndex(requestID); isPresent {
		var newAlbum Album
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}
		albums[index] = newAlbum
		c.IndentedJSON(http.StatusOK, albums[index])
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}
}

func main() {

	fmt.Println("Staring the server...")
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
	router.Run("localhost:8080")

}
