package main

import (
	"context" // each request to a server starts a new goroutine with a specific context (deadline, user authorization,
	// passed value, etc.) => each child goroutine is passed this context and if that context times out
	// or gets cancelled, all child routines get cancelled as well
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// create an endpoint for the getData() function of EdgecastService
func makeGetDataEndpoint(svc EdgecastService) endpoint.Endpoint {
	// empty interface{} has no methods and so this function accepts any value => will be converted to type interface{}
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDataRequest)
		v, err := svc.GetData(req.AirportCode)
		if err != nil {
			return getDataResponse{v, err.Error()}, nil
		}
		return getDataResponse{v, ""}, nil
	}
}

// Encode/Decode request/response data in JSON format
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

// define request and response structure
type getDataRequest struct {
	AirportCode string `json:"airportcode"`
}

type getDataResponse struct {
	Data string `json:"data"`
	Err  string `json:"err, omitempty"`
}
