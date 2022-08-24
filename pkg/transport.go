package pkg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type DemoRequest struct {
	Name string `json:"name"`
}

type DemoResponse struct {
	S   string `json:"resp"`
	Err string `json:"err,omitempty"`
}

func makeDemoEndpoint(srv demoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DemoRequest)
		v, err := srv.Demo(req.Name)
		if err != nil {
			return nil, err
		}

		return DemoResponse{v, ""}, nil
	}
}

func decodeDemoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DemoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
