package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("couldn't start web app", err)
	}
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte(GetMessage()))
	if err != nil {
		log.Println("error in indexHandler", err)
	}
}

func GetMessage() string {
	return "Hello World!"
}