package main

import (
	"log"

	"github.com/achiku/webque"
)

func main() {
	log.Println("starting backend service...")
	webque.BackendRun()
}
