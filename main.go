package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"spaceship/lib"
	"bytes"
	"image"
	"image/png"
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

		rawImage := lib.CreateInspiration(request.Url, request.Text)
		var img image.Image
		img = rawImage

		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		checkError(err)

		c.Data(http.StatusCreated, "image/png", buf.Bytes())
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hmmm bad request yo"}) // TODO Fix error message returned
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}