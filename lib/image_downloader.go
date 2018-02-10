package lib

import (
	"net/http"
	"fmt"
	"image"
	"image/png"
	"bytes"
)


func CreateInspiration(url string, text string) []byte {
	validateImageMimeType(url)
	image := downloadImage(url)
	addText(image, text)
	return convertImageNrgbaToBytes(image)
}

func validateImageMimeType(url string) {
	response, err := http.Head(url)
	defer response.Body.Close()
	checkError(err)

	contentType := response.Header["Content-Type"]
	fmt.Println("Content type: ", contentType)
}

func downloadImage(url string) *image.NRGBA {
	response, err := http.Get(url)
	defer response.Body.Close()
	checkError(err)

	pngImage, err := png.Decode(response.Body)
	checkError(err)

	nrgbaImage := pngImage.(*image.NRGBA)
	return nrgbaImage
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
