package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/h2non/gock"
	"fmt"
	"net/url"
	//"io/ioutil"
	"os"
	"errors"
)

const UrlBase = "http://foo.com"
const UrlPath = "/bar"
const DummyImagePath = "../testUtilities/seattle-seahawks-logo.png"
var DummyError = errors.New("Dummy error")

func TestDownloadImageShouldReturnImageNrgba(t *testing.T) {
	// Given
	dummyImage, err := os.Open(DummyImagePath)
	if err != nil {
		assert.Fail(t, "Error retrieving dummy image from " +DummyImagePath)
	}

	defer gock.Off()
	gock.New(UrlBase).
		Get(UrlPath).
		Reply(200).
			Body(dummyImage)

	// When
	actualImage, err := DownloadImage(UrlBase + UrlPath)

	// Then
	assert.Equal(t, "*image.NRGBA", fmt.Sprintf("%T", actualImage))
}

func TestDownloadImageShouldReturnErrorWhenGetFails(t *testing.T) {
	// Given
	expectedError := errors.New("Dummy error")
	defer gock.Off()
	gock.New(UrlBase).
		Get(UrlPath).
		ReplyError(expectedError)

	// When
	_, actualError := DownloadImage(UrlBase + UrlPath)

	// Then
	assert.Equal(t, "*url.Error", fmt.Sprintf("%T", actualError))
	assert.Equal(t, DummyError, actualError.(*url.Error).Err)
}

func TestValidateImageMimeTypeShouldNotReturnErrorIfValidMimeType(t *testing.T) {
	defer gock.Off()
	gock.New(UrlBase).
		Head(UrlPath).
		Reply(200).
		SetHeader("Content-Type", "image/jpeg")

	assert.Nil(t, ValidateImageMimeType(UrlBase + UrlPath))
}

func TestValidateImageMimeTypeShouldReturnUserErrorIfMimeTypeNotSupported(t *testing.T) {
	defer gock.Off()
	gock.New(UrlBase).
		Head(UrlPath).
		Reply(200).
		SetHeader("Content-Type", "image/gif")

	actualError := ValidateImageMimeType(UrlBase + UrlPath)
	expectedError := NewUserError("Url does not contain supported MIME type. Supported MIME types are: image/png and image/jpeg")
	assert.Equal(t, expectedError, actualError)
}

func TestValidateImageMimeTypeShouldReturnErrFromHeadRequest(t *testing.T) {
	defer gock.Off()

	gock.New(UrlBase).
		Head(UrlPath).
		ReplyError(DummyError)

	actualError := ValidateImageMimeType(UrlBase + UrlPath)
	assert.Equal(t, "*url.Error", fmt.Sprintf("%T", actualError))
	assert.Equal(t, DummyError, actualError.(*url.Error).Err)
}
