package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func downloadFileFromURL(fileUrl string) (string, error) {
	response, err := http.Get(fileUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}

	parsedURL, err := url.Parse(fileUrl)
	if err != nil {
		return "", err
	}

	fileName := path.Base(parsedURL.Path)

	file, err := os.Create("/tmp/" + fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return "/tmp/" + fileName, nil
}
