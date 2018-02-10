package lib

// TODO rethink project layout

import (
	"log"
	"net/http"
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
)


func CreateInspiration(url string, text string) *image.NRGBA {
	image := downloadImageAsBytes(url)
	addText(&image, text)
	saveImageToDisk(&image)
	return &image
}

func downloadImageAsBytes(url string) image.NRGBA {

	validateImageMimeType(url)
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	pngImage, err := png.Decode(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	nrgbaImage := pngImage.(*image.NRGBA)

	return *nrgbaImage
}

func saveImageToDisk(pngImage *image.NRGBA) {
	file, err := os.Create("/tmp/image.png")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, pngImage)

	if err != nil {
		log.Fatal(err)
	}
}

func validateImageMimeType(url string) {
	response, e := http.Head(url)
	defer response.Body.Close()

	if e != nil {
		// TODO handle error
		log.Fatal(e)
	}

	contentType := response.Header["Content-Type"]
	fmt.Println("Content type: ", contentType)
}

func addText(img *image.NRGBA, text string) {
	addLabelHelper(img, 20, 30, text)
}

func addLabelHelper(img *image.NRGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
