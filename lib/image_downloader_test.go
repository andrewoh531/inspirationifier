package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TODO think about remove dependency to an external source for this integration test
const ValidPngImageUrl = "https://www.printyourbrackets.com/nfl-logos/seattle-seahawks-logo.png"
const ValidSvgImageUrl = "https://upload.wikimedia.org/wikipedia/commons/f/fa/Apple_logo_black.svg"
const UnsupportedMimeTypeUrl = "https://www.nba.com"

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
	assert.Nil(t, ValidateImageMimeType(ValidPngImageUrl))
}

func TestValidateImageMimeTypeShouldReturnUserErrorIfMimeTypeNotSupported(t *testing.T) {
	actualError := ValidateImageMimeType(UnsupportedMimeTypeUrl)
	expectedError := NewUserError("Url does not contain supported MimeType. Supported mimetypes are: image/png and image/jpeg")
	assert.Equal(t, expectedError, actualError)
}
