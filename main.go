package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"spaceship/controllers"
	"spaceship/lib"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/healthcheck", healthCheck)
		v1.POST("/createInspiration", createInspiration)
	}

	return router
}

func main() {
	r := setupRouter()
	r.Run()
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
		imageChannel := make(chan controllers.InspirationResult)
		go controllers.CreateInspiration(request.Url, request.Text, imageChannel)

		response := <-imageChannel
		if response.Error != nil {
			handleError(c, response.Error)
		} else {
			c.Data(http.StatusCreated, "image/png", response.ImageBytes)
		}

	} else {
		c.String(http.StatusBadRequest, "JSON payload requires both a 'text' and 'url' property")
	}
}

func handleError(c *gin.Context, err error) {

	switch err.(type) {
		case *lib.UserError:
			c.String(http.StatusBadRequest, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
	}

}
