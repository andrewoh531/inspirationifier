package lib

import (
	"net/http"
	"image"
	"image/png"
	"image/jpeg"
	log "github.com/sirupsen/logrus"
	"io"
	"image/draw"
	"net/url"
	"net"
)

const MimeTypePng = "image/png"
const MimeTypeJpeg = "image/jpeg"
var SupportedMimeTypes = map[string]bool { MimeTypePng: true, MimeTypeJpeg: true }

func ValidateImageMimeType(url string) error {
	response, err := http.Head(url)
	if err != nil {

		if isNoSuchHostError(err) {
			log.Info("Invalid URL provided: " + url)
			return NewUserError("Invalid URL provided - no such host")
		}

		log.Error(err)
		return err
	}
	defer response.Body.Close()

	contentTypeArray := response.Header["Content-Type"]

	for _, contentType := range contentTypeArray {
		if (SupportedMimeTypes[contentType]) {
			return nil
		}
	}

	return NewUserError("Url does not contain supported MIME type. Supported MIME types are: image/png and image/jpeg")
}

func isNoSuchHostError(err error) bool {

	if err1, ok1 := err.(*url.Error); ok1 {
		if opErr, ok2 := err1.Err.(*net.OpError); ok2 {
			if dnsErr, ok3 := opErr.Err.(*net.DNSError); ok3 {
				if (dnsErr.Err == "no such host") {
					return true
				}
			}
		}
	}

	return false
}

func DownloadImage(url string) (*image.NRGBA, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Error("Error downloading image: ", err)
		return nil, err
	}
	defer response.Body.Close()

	decodedImg, err := decodeImage(response.Header["Content-Type"], response.Body)
	if err != nil {
		log.Error("Error decoding downloaded image: ", err)
		return nil, err
	}

	nrgbaImage := decodedImg.(*image.NRGBA)
	return nrgbaImage, nil
}

func decodeImage(contentType []string, reader io.ReadCloser) (image.Image, error) {
	var origImg image.Image
	var err error

	if (contains(contentType, MimeTypePng)) {
		origImg, err = png.Decode(reader)
	} else if (contains(contentType, MimeTypeJpeg)) {
		origImg, err = jpeg.Decode(reader)
	}

	if err != nil {
		log.Error("Error decoding image: ", err)
		return nil, err
	}

	bounds := origImg.Bounds()
	rgbaImage := image.NewNRGBA(origImg.Bounds())
	draw.Draw(rgbaImage, rgbaImage.Bounds(), origImg, bounds.Min, draw.Src)

	return rgbaImage, nil
}

func contains(stringSlice []string, searchStr string) bool {
	for _, value := range stringSlice {
		if value == searchStr {
			return true
		}
	}
	return false
}