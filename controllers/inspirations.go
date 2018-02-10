package controllers

import (
	"spaceship/lib"
	"image"
	"bytes"
	"image/png"
	"strings"
)

func CreateInspiration(url string, text string) ([]byte, error) {
	err := validate(url)
	if err != nil {
		return nil, err
	}

	image, err := lib.DownloadImage(url)
	if err != nil {
		return nil, err
	}

	lib.AddTextToImage(image, text)
	return convertImageNrgbaToBytes(image), nil
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
