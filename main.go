package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"spaceship/controllers"
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
	var request inspirationPayload

	if err := c.ShouldBindWith(&request, binding.JSON); err == nil {
		imageBytes := controllers.CreateInspiration(request.Url, request.Text)
		c.Data(http.StatusCreated, "image/png", imageBytes)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hmmm bad request yo"}) // TODO Fix error message returned
	}
}
