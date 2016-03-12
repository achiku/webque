package webque

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/bgentry/que-go"
)

// NewHTTPClient creates http client
func NewHTTPClient() *http.Client {
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
	}
	return client
}

// HelloJob says hello
func HelloJob(j *que.Job) error {
	log.Println("hello, world!")
	return nil
}

// UpdateDepositJob updates deposit in backend service
func UpdateDepositJob(j *que.Job) error {
	c := NewHTTPClient()
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	url := "http://restapi3.apiary.io/notes"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
