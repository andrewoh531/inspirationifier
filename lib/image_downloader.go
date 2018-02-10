package lib

// TODO change package name
// TODO rethink project layout

import (
	"log"
	"net/http"
	//"os"
	//"io"
	"fmt"
	"io/ioutil"
)



func DownloadImageAsBytes(url string) []byte {

	validateImageMimeType(url)
	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return contents
	//
	//
	////open a file for writing
	//file, err := os.Create("/tmp/asdf.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// Use io.Copy to just dump the response body to the file. This supports huge files
	//_, err = io.Copy(file, response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//file.Close()
	//fmt.Println("Success!")
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

func addTextToImage(rawImage []byte, text string) []byte {



	return nil
}


