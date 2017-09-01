package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// define business logic in an interface
type EdgecastService interface {
	GetData(string) (string, error)
}

type edgecastService struct{}

func (edgecastService) GetData(code string) (string, error) {
	if code == "" {
		return "", ErrNoCode
	}
	resp, err := http.Get(fmt.Sprintf("http://services.faa.gov/airport/status/%s?format=json", code))
	if err != nil {
		return "", ErrReq
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println("INFO:", bodyString)
		return string(bodyString), nil
	}
	return "", ErrNotFound

}

var ErrNoCode = errors.New("no airportcode")
var ErrNotFound = errors.New("not found")
var ErrReq = errors.New("request error")
