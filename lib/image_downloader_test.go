package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"spaceship/testUtilities"
)

// TODO think about remove dependency to an external source for this integration test
const ExpectedPngImagePath = "./testUtilities/seattle-seahawks-logo.png"

//func TestDownloadImageWithValidPngUrl(t *testing.T) {
//	image := DownloadImage(ValidPngImageUrl)
//
//	expectedContents, err := ioutil.ReadFile(ExpectedPngImagePath)
//	if err != nil {
//		assert.Fail(t, "Error retrieving expected image contents from " + ExpectedPngImagePath)
//	}
//
//	assert.Equal(t, expectedContents, image)
//}


func TestValidateImageMimeTypeShouldNotReturnErrorIfValidMimeType(t *testing.T) {
	assert.Nil(t, ValidateImageMimeType(testUtilities.ValidPngImageUrl))
}

func TestValidateImageMimeTypeShouldReturnUserErrorIfMimeTypeNotSupported(t *testing.T) {
	actualError := ValidateImageMimeType(testUtilities.UnsupportedMimeTypeUrl)
	expectedError := NewUserError("Url does not contain supported MIME type. Supported MIME types are: image/png and image/jpeg")
	assert.Equal(t, expectedError, actualError)
}
