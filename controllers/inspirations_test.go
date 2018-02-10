package controllers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"spaceship/lib"
	"spaceship/testUtilities"
)

const ProtocolNotSpecified = "www.nba.com"

func TestWhenProtocolNotSpecifiedUserErrorShouldBeReturned(t *testing.T) {
	imageChannel := make(chan InspirationResult)
	go CreateInspiration(ProtocolNotSpecified, "", imageChannel)

	response := <-imageChannel
	expectedError := lib.NewUserError("Please provide 'http' or 'https' protocol.")
	assert.Equal(t, expectedError, response.Error)
}

func TestWhenValidValuesProvidedImageShouldBeReturned(t *testing.T) {
	imageChannel := make(chan InspirationResult)
	go CreateInspiration(testUtilities.ValidPngImageUrl, "Sample text", imageChannel)

	response := <-imageChannel
	assert.Nil(t, response.Error)
	assert.True(t, len(response.ImageBytes) > 0)
}