package utils

import (
	"io"
	"mime/multipart"
	"os"
	"time"
)

var ImagePath = "static/images/"

func UploadImage(file *multipart.FileHeader) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	timeStamp := time.Now()

	if _, err := os.Stat(ImagePath); os.IsNotExist(err) {
		_ = os.Mkdir(ImagePath, 0777)
	}
	fileName := timeStamp.Format(time.RFC3339) + ".jpg"
	path := ImagePath + fileName
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return fileName, nil
}
