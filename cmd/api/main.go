package main

import (
	"log"
	"net/http"
	"github.com/silasms/octarine/internal/handler"
)

const port string = ":8080"

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", handler.ImageGetAll)
	http.HandleFunc("/send-img", handler.SendImg)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}