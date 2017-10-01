package main

import (
	"errors"
	"fmt"
	ec "github.com/iwilltry42/edgecast"
	"io/ioutil"
	"net/http"
)

/*
 * GLOBAL VARIABLES //TODO: maybe replace with environment-variables
 */
var EDGECAST_ENDPOINT = "https://api.edgecast.com/v2/realtimestats/customers/%s/media/%d/%s"

// define business logic in an interface
type EdgecastService interface {
	GetData(string) (string, error)
}

type edgecastService struct {
	ecs *ec.Edgecast
}

// Contacting the API to fetch the required data using the transmitted code
func (e edgecastService) GetData(code string) (string, error) {

	// catch empty input
	if code == "" {
		return "", ErrNoCode
	}

	// GET request to API-Endpoint
	e.ecs.Bandwidth(2)
	resp, err := http.Get(fmt.Sprintf("http://services.faa.gov/airport/status/%s?format=json", code))
	if err != nil {
		return "", ErrReq
	}
	defer resp.Body.Close() // in any case of return, close before returning

	// OK-Response from server -> extract body data
	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return string(bodyString), nil
	}
	return "", ErrNotFound

}

// Custom Errors
var ErrNoCode = errors.New("no airportcode")
var ErrNotFound = errors.New("not found")
var ErrReq = errors.New("request error")
