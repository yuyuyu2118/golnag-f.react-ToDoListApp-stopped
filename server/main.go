package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 60180...")

	log.Fatal(http.ListenAndServe(":60180", r))
}
