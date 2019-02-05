package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func collectLocations() {
	client := &http.Client{}
	URL := "https://api.ibm.com/scinsights/run/api/locations/"

	req, err := http.NewRequest("GET", URL, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-IBM-Client-Id", os.Getenv("XIBMClientId"))
	req.Header.Add("X-IBM-Client-Secret", os.Getenv("XIBMClientSecret"))
	req.Header.Add("X-IBM-User-Secret", os.Getenv("XIBMUserSecret"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	header := resp.Header
	content := header.Get("Content-Range")
	count := (strings.Split(content, "/"))[1]
	log.Printf("Number of location records is %v\n", count)
}

func main() {
	log.Printf("SCI tool")
	collectLocations()
}
