package lib

import (
	"net/http"
	"image"
	"image/png"
	"log"
)

var SupportedMimeTypes = map[string]bool { "image/png": true, "image/jpeg": true }

func ValidateImageMimeType(url string) error {
	response, err := http.Head(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	contentTypeArray := response.Header["Content-Type"]

	for _, contentType := range contentTypeArray {
		if (SupportedMimeTypes[contentType]) {
			return nil
		}
	}

	return NewUserError("Url does not contain supported MimeType. Supported mimetypes are: image/png and image/jpeg")
}

func DownloadImage(url string) (img *image.NRGBA, e error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pngImage, err := png.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	nrgbaImage := pngImage.(*image.NRGBA)
	return nrgbaImage, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
