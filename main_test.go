package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

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
		SetJSON(gofight.D{"url": "www.nba.com"}).
		SetDebug(true).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}


func TestValidParametersProvidedRequestShouldReturn200(t *testing.T) {
	payload := gofight.D{"url": "www.nba.com", "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		SetDebug(true).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusCreated, r.Code)
	})
}