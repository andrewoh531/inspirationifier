package controllers

import (
	"spaceship/lib"
	"image"
	"bytes"
	"image/png"
)

func CreateInspiration(url string, text string) ([]byte, error) {
	err := lib.ValidateImageMimeType(url)
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
