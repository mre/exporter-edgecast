package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGetDataEndpoint(svc EdgecastService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDataRequest)
		v, err := svc.GetData(req.AirportCode)
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
	AirportCode string `json:"airportcode"`
}

type getDataResponse struct {
	Data string `json:"data"`
	Err  string `json:"err, omitempty"`
}
