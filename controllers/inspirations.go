package controllers

import (
	"spaceship/lib"
	"image"
	"bytes"
	"image/png"
)

func CreateInspiration(url string, text string) []byte {
	lib.ValidateImageMimeType(url)
	image, _ := lib.DownloadImage(url)
	lib.AddTextToImage(image, text)
	return convertImageNrgbaToBytes(image)
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
