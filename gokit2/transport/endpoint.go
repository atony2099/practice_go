package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//  like controller

// 1. how to init  service;

// 2.

func GetOrderEndPoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}
