package main

import (
	"log"

	"github.com/achiku/webque"
)

func main() {
	log.Println("starting proxy service...")
	webque.ProxyRun()
}
