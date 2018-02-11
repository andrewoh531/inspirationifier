package controllers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"spaceship/lib"
	"os"
	"github.com/h2non/gock"
)

const ProtocolNotSpecified = "www.nba.com"
const UrlBase = "http://foo.com"
const UrlPath = "/bar"
const SamplePngImagePath = "../test-resources/sample.png"

func TestWhenProtocolNotSpecifiedUserErrorShouldBeReturned(t *testing.T) {
	imageChannel := make(chan InspirationResult)
	go CreateInspiration(ProtocolNotSpecified, "", imageChannel)

	response := <-imageChannel
	expectedError := lib.NewUserError("Please provide 'http' or 'https' protocol.")
	assert.Equal(t, expectedError, response.Error)
}

func TestWhenValidValuesProvidedImageShouldBeReturned(t *testing.T) {
	// Given
	dummyImage, err := os.Open(SamplePngImagePath)
	if err != nil {
		assert.Fail(t, "Error retrieving dummy image from " + SamplePngImagePath)
	}

	defer gock.Off()
	gock.New(UrlBase).
		Head(UrlPath).
		Reply(200).
		SetHeader("Content-Type", "image/jpeg")

	gock.New(UrlBase).
		Get(UrlPath).
		Reply(200).
		SetHeader("Content-Type", "image/png").
		Body(dummyImage)

	// When
	imageChannel := make(chan InspirationResult)
	go CreateInspiration(UrlBase + UrlPath, "Sample text", imageChannel)

	// Then
	response := <-imageChannel
	assert.Nil(t, response.Error)
	assert.True(t, len(response.ImageBytes) > 0)
}