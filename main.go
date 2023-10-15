package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	tm_eventsURL   = "https://app.ticketmaster.com/discovery/v2/events.json"
	sp_apiURL      = "https://accounts.spotify.com/api/token"
	sp_authURL     = "https://accounts.spotify.com/authorize"
	sp_redirectURL = "http://localhost:3000"
)

func main() {

	//TODO need a client running at the redirectURL for auth purposes
	sp_clientID := os.Getenv("SP_CLIENTID")
	req, _ := http.NewRequest(http.MethodGet, sp_authURL, nil)
	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", sp_clientID)
	q.Add("redirect_uri", sp_redirectURL)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Perform tm request failed")

	}
	defer resp.Body.Close()
	results, _ := io.ReadAll(resp.Body)
	fmt.Print(string(results))
}

func sp_getAccessToken() string {
	sp_clientID := os.Getenv("SP_CLIENTID")
	sp_clientSecret := os.Getenv("SP_CLIENTSECRET")
	authorizationHeader := "Basic " + b64.StdEncoding.EncodeToString([]byte(sp_clientID+":"+sp_clientSecret))

	body := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest(http.MethodPost, sp_apiURL, body)

	req.Header.Set("Authorization", authorizationHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print("Perform tm request failed")
	}
	defer resp.Body.Close()
	results, _ := io.ReadAll(resp.Body)

	return string(results)
}

func tm_call() string {
	tm_apikey := os.Getenv("TM_APIKEY")

	req, err := http.NewRequest(http.MethodGet, tm_eventsURL, nil)
	if err != nil {
		return "New tm request failed"
	}

	q := req.URL.Query()
	q.Add("size", "3")
	q.Add("apikey", tm_apikey)
	q.Add("countryCode", "CA")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Perform tm request failed"

	}
	defer resp.Body.Close()
	results, _ := io.ReadAll(resp.Body)

	return string(results)
}
