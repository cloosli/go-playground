package main

import "net/http"

import "log"

import "fmt"

var url string = "http://limesapps.ch"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Cannot open url", err)
	}
	fmt.Printf("HTTP Status Code: %v \n", resp.Status)
	if resp.StatusCode < 400 {
		body := resp.Body
		fmt.Printf("Body: %v \n", body)
	}
}
