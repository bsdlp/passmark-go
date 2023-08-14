package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bsdlp/passmark-go"
)

func main() {
	var sessionCookie *http.Cookie
	resp, err := http.Get("https://www.cpubenchmark.net/CPU_mega_page.html")
	if err != nil {
		log.Fatalf("error getting session cookie: %s", err.Error())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "PHPSESSID" {
			sessionCookie = cookie
		}
	}
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("error discarding body: %s", err.Error())
	}
	resp.Body.Close()

	req, err := http.NewRequest("GET", "https://www.cpubenchmark.net/data/", nil)
	if err != nil {
		log.Fatalf("unable to call data endpoint: %s", err.Error())
	}
	req.AddCookie(sessionCookie)
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error fetching data: %s", err.Error())
	}

	var d responseData
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Fatalf("error decoding response data: %s", err.Error())
	}

	resp.Body.Close()

	err = json.NewEncoder(os.Stdout).Encode(d.Data)
	if err != nil {
		log.Fatalf("error encoding data: %s", err.Error())
	}
}

type responseData struct {
	Data []passmark.CPU `json:"data"`
}
