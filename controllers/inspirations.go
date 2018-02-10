package controllers

import (
	"spaceship/lib"
	"image"
	"bytes"
	"image/png"
	"strings"
)

type InspirationResult struct {
	ImageBytes []byte
	Error      error
}

func CreateInspiration(url string, text string, done chan InspirationResult) {
	err := validate(url)
	if err != nil {
		done <- InspirationResult{nil, err}
		return
	}

	image, err := lib.DownloadImage(url)
	if err != nil {
		done <- InspirationResult{nil, err}
		return
	}

	lib.AddTextToImage(image, text)
	done <- InspirationResult{ convertImageNrgbaToBytes(image), nil }
}

func validate(url string) error {
	trimmed := strings.TrimSpace(url)
	if strings.Index(trimmed, "http") != 0 {
		return lib.NewUserError("Please provide 'http' or 'https' protocol.")
	}
	return lib.ValidateImageMimeType(url)
}

func convertImageNrgbaToBytes(nrgbaImage *image.NRGBA) []byte {
	var img image.Image
	img = nrgbaImage

	buf := new(bytes.Buffer)

	err := png.Encode(buf, img)
	checkError(err)

	return buf.Bytes()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
