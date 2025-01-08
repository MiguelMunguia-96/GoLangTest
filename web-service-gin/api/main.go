package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	S "web-service-gin/internal/server"
)

func main() {

	fmt.Println("Staring the server...")
	router := gin.Default()
	router.GET("/albums", S.GetAlbum)
	router.POST("/albums", S.PostAlbum)
	router.GET("/albums/:id", S.GetAlbumByID)
	router.PUT("/albums/:id", S.UpdateAlbum)
	router.DELETE("/albums/:id", S.DeleteAlbum)
	router.Run("localhost:8080")
}
