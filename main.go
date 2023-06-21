package main

import (
	"github.com/aarnavpant/StudentAPIs.git/src"
	"fmt"
	"net/http"
	"log"
)

func main() {
	muxNew := src.CreateNewMux()
	fmt.Println("server started")
	log.Fatal(http.ListenAndServe(":8081", muxNew))
}