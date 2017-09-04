package main

import (
	"log"
	"net/http"
	"fmt"
	"GOLEARN/main/frame"
)

func main() {
	fmt.Println("-------------------------------- Server Start --------------------------------")
	router := frame.NewRouter()
	log.Fatal(http.ListenAndServe("127.0.0.1:2017", router))
}