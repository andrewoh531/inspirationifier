package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const InvalidUrl = "https://someunknowndomain.com"
const ValidPngUrl = "https://upload.wikimedia.org/wikipedia/en/thumb/0/01/Golden_State_Warriors_logo.svg/1200px-Golden_State_Warriors_logo.svg.png"
const ValidJpegUrl = "https://occ-0-2433-448.1.nflxso.net/art/b0b65/b5bd8174a48f773646aefb4bf8e0f1a4b12b0b65.jpg"
const InvalidMimeTypeUrl = "https://www.nba.com"

func TestInvalidPathShouldReturn404(t *testing.T) {
	r := gofight.New()

	r.GET("/api/v1/createInspiration").
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestMissingRequiredParameterShouldReturn400(t *testing.T) {
	r := gofight.New()

	r.POST("/api/v1/createInspiration").
		SetJSON(gofight.D{"url": ValidPngUrl}).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "JSON payload requires both a 'text' and 'url' property", r.Body.String())
		})
}


func TestValidParametersProvidedRequestShouldReturn201(t *testing.T) {
	payload := gofight.D{"url": ValidPngUrl, "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
		})
}

func TestValidJpegRequestShouldReturn201(t *testing.T) {
	payload := gofight.D{"url": ValidJpegUrl, "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusCreated, r.Code)
	})
}

func TestInvalidMimeTypeShouldReturnBadRequest(t *testing.T) {
	payload := gofight.D{"url": InvalidMimeTypeUrl, "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Equal(t, "Url does not contain supported MIME type. Supported MIME types are: image/png and image/jpeg", r.Body.String())
	})
}

func TestInvalidUrlShouldReturnBadRequest(t *testing.T) {
	payload := gofight.D{"url": InvalidUrl, "text": "sample text"}

	r := gofight.New()
	r.POST("/api/v1/createInspiration").
		SetJSON(payload).
		Run(setupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Equal(t, "Invalid URL provided - no such host", r.Body.String())
	})
}