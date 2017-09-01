package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"log"
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
		return "", err
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

func main() {
	svc := edgecastService{}

	getDataHandler := httptransport.NewServer(
		makeGetDataEndpoint(svc),
		decodeGetDataRequest,
		encodeResponse,
	)

	http.Handle("/getdata", getDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeGetDataEndpoint(svc EdgecastService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDataRequest)
		v, err := svc.GetData(req.S)
		if err != nil {
			return getDataResponse{v, err.Error()}, nil
		}
		return getDataResponse{v, ""}, nil
	}
}

func decodeGetDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getDataRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getDataRequest struct {
	S string `json:"s"`
}

type getDataResponse struct {
	V   string `json:"v"`
	Err string `json:"err, omitempty"`
}

var ErrNoCode = errors.New("no airportcode")
var ErrNotFound = errors.New("not found")
var ErrReq = errors.New("request error")
