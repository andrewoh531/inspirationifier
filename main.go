package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"spaceship/lib"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/ping", healthCheck)
		v1.POST("/createInspiration", createInspiration)
	}

	return router
}

func main() {
	r := setupRouter()
	r.Run(":8080") // TODO Pick this up as environment variable
}

type inspirationPayload struct {
	Text string `json:"text" binding:"required"`
	Url string `json:"url" binding:"required"`
}

func healthCheck(c *gin.Context) {
	c.Status(http.StatusOK);
}

func createInspiration(c *gin.Context) {
	var inspirationPayload inspirationPayload

	if err := c.ShouldBindWith(&inspirationPayload, binding.JSON); err == nil {

		rawImage := lib.DownloadImageAsBytes(inspirationPayload.Url)

		// Add text to the image

		// Return image as downloadable response

		message := fmt.Sprintf("Payload received %s, %s", inspirationPayload.Text, inspirationPayload.Url)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": message})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hmmm bad request yo"}) // TODO Fix error message returned
	}
}
