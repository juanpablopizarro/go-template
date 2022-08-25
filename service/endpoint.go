package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type HashRequest struct {
	S string `json:"string"`
}

type HashResponse struct {
	H   string `json:"hash"`
	Err string `json:"error,omitempty"`
}

func makeHashEndpoint(srv Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(HashRequest)
		v, err := srv.Hash(req.S)
		if err != nil {
			return nil, ErrInvalidInput
		}
		return HashResponse{v, ""}, nil
	}
}
