package function

import (
	"io"
	"net/http"
	"log"
	"crypto/tls"
)

type CountryResponse struct {
	CountryCode string `json:"countryCode"`
	PostalCode  string `json:"postalCode"`
}

func doRequest(method, apiURL, tokenBearer string, body io.Reader) *http.Response{
	reqNew, _ := http.NewRequest(method, apiURL, body)
	reqNew.Header.Set("Content-Type", "application/json")
	reqNew.Header.Set("Authorization", "Bearer "+tokenBearer)
	log.Println("Doing The request: ")
	log.Println(reqNew)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respNew, _ := client.Do(reqNew)
	log.Println("Getting the next response: ")
	log.Println(respNew)
	return respNew
}
