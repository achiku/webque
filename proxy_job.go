package webque

import (
	"log"

	"github.com/bgentry/que-go"
)

// HelloJob says hello
func HelloJob(j *que.Job) error {
	log.Println("hello, world!")
	return nil
}
