package gimme

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// GET request a url and returns the body as a string. Defaults to http unless specificed with https://
func URL(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	return string(body)
}

// Opens a file and returns the contents as a byte array
func File(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panicln(err)
	}
	return file
}
