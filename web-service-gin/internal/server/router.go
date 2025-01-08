package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	A "web-service-gin/internal/model"
)

var albums = []A.Album{
	{ID: "1", Title: "Forever", Artist: "Mystery Skulls", MetaReady: true, BackupReady: true, Mp3Ready: true},
	{ID: "2", Title: "DDC", Artist: "Nothing but thieves", MetaReady: true, BackupReady: true, Mp3Ready: true},
}

func GetAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum A.Album

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

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	if index, isPresent := findAlbumIndex(id); isPresent {
		c.IndentedJSON(http.StatusOK, albums[index])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}
}

func DeleteAlbum(c *gin.Context) {
	requestID := c.Param("id")

	if index, isPresent := findAlbumIndex(requestID); isPresent {
		name := albums[index].Title
		albums = append(albums[:index], albums[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": name + " album deleted "})
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}
}
func UpdateAlbum(c *gin.Context) {
	requestID := c.Param("id")

	if index, isPresent := findAlbumIndex(requestID); isPresent {
		var newAlbum A.Album
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
