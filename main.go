package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	tm_eventsURL = "https://app.ticketmaster.com/discovery/v2/events.json"
)

func main() {
	tm_apikey := os.Getenv("TM_APIKEY")

	req, err := http.NewRequest(http.MethodGet, tm_eventsURL, nil)
	if err != nil {
		fmt.Println("New tm request failed")
		return
	}

	q := req.URL.Query()
	q.Add("size", "1")
	q.Add("apikey", tm_apikey)
	q.Add("countryCode", "CA")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Perform tm request failed")
		return
	}
	defer resp.Body.Close()
	results, _ := io.ReadAll(resp.Body)
	fmt.Print(string(results))
}
