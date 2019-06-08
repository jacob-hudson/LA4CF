package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// building the request URL
	var prefix = "https://api.cloudflare.com/"
	var path = "client/v4/zones/"
	var zone = os.Getenv("CF_ZONE")
	var suffix = "/dns_analytics/report/bytime"

	var request_url = prefix + path + zone + suffix

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("X-Auth-Email", "jacob.alan.hudson@gmail.com")
	req.Header.Set("X-Auth-Key", os.Getenv("CF_KEY"))

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}
