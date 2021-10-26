package utils

import (
	"errors"
	"fmt"
	"io"
	"os"

	uuid "github.com/satori/go.uuid"
)

const (
	VIDEO_QUALITY = "medium"
)

var AllowedExtensions = map[string]string {
	"mp4": ".mp4",
}

func CreateMedia(fileExtension string, stream io.ReadCloser) error {
	path := os.Getenv("DOCKER_CONTAINER_PATH")
	if len(path) <= 0 {
		return errors.New("Path not found")
	}
	if len(AllowedExtensions[fileExtension]) <= 0 {
		return errors.New("Extension not allowed")
	}
	file, err := os.Create(fmt.Sprintf("%s%s%s", path, uuid.NewV4().String(), fileExtension))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}
