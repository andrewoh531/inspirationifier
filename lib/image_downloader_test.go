package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)


// TODO think about remove dependency to an external source for this integration test
const ValidPngImageUrl = "https://www.printyourbrackets.com/nfl-logos/seattle-seahawks-logo.png"
const ValidSvgImageUrl = "https://upload.wikimedia.org/wikipedia/commons/f/fa/Apple_logo_black.svg"

const ExpectedPngImagePath = "./testUtilities/seattle-seahawks-logo.png"

func TestDownloadImageWithValidPngUrl(t *testing.T) {
	downloadedContent := DownloadImageAsBytes(ValidPngImageUrl)

	//file, err := os.Open(ExpectedPngImagePath) // For read access.
	//if err != nil {
	//	assert.Fail(t, "Error retrieving expected image contents from " + ExpectedPngImagePath)
	//}

	expectedContents, err := ioutil.ReadFile(ExpectedPngImagePath)
	if err != nil {
		assert.Fail(t, "Error retrieving expected image contents from " + ExpectedPngImagePath)
	}


	// TODO file comparison does not equate
	assert.Equal(t, expectedContents, downloadedContent)
}


// TODO Add support for multiple image types
//func TestDownloadImageWithValidSvgUrl(t *testing.T) {
//	DownloadImageAsBytes(ValidSvgImageUrl)
//}
