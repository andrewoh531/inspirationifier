package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const ValidImageUrl = "https://upload.wikimedia.org/wikipedia/en/thumb/0/01/Golden_State_Warriors_logo.svg/1200px-Golden_State_Warriors_logo.svg.png"

func TestInvalidPathShouldReturn404(t *testing.T) {
	r := gofight.New()

	r.GET("/api/v1/createInspiration").
		SetDebug(true).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestMissingRequiredParameterShouldReturn400(t *testing.T) {
	r := gofight.New()

	r.POST("/api/v1/createInspiration").
		SetJSON(gofight.D{"url": ValidImageUrl}).
		SetDebug(true).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// TODO assert error response body
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}


func TestValidParametersProvidedRequestShouldReturn200(t *testing.T) {
	payload := gofight.D{"url": ValidImageUrl, "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		SetDebug(true).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
		})
}