package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url-shorter/apis"
)

var PORT int = 8080

func main() {
	http.HandleFunc("/shorter/", apis.Shorter)

	fmt.Println("Running server on Port:", PORT)
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
