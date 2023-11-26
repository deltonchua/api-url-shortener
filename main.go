package main

import (
	"log"

	"github.com/deltonchua/api-url-shortener/api"
)

func main() {
	log.Fatal(api.Run())
}
