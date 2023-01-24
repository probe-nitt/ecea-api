package utils

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/fatih/color"
)

var DocumentPath = "static/documents/"

func UploadDocument(file *multipart.FileHeader) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	timeStamp := time.Now()

	if _, err := os.Stat(DocumentPath); os.IsNotExist(err) {
		_ = os.Mkdir(DocumentPath, 0777)
	}
	fileName := timeStamp.Format(time.RFC3339) + ".pdf"
	path := DocumentPath + fileName
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

func DeleteDocument(fileName string) error {
	path := DocumentPath + fileName
	log.Println(color.MagentaString("Deleting: " + path))
	e := os.Remove(path)
	if e != nil {
		return e
	}
	return nil
}
