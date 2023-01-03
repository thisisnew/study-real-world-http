package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("C://Users/dhkim/GolandProjects/study-real-world-http/ch03/test.txt")

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)

	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
