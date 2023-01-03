package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")

	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("C://Users/dhkim/GolandProjects/study-real-world-http/ch03/michael-jackson-gettyimages-89446104.jpg")

	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)

	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
