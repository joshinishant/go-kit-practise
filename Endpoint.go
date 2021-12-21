package main

import (
	"context"
	"go-kit/orderservice"

	"github.com/go-kit/kit/endpoint"
)

type RestEndpoint struct {
	GetByIdEndpoint      endpoint.Endpoint
	CreateEndpoint       endpoint.Endpoint
	ChangeStatusEndpoint endpoint.Endpoint
}

func MakeEndpoints(s orderservice.OderService) RestEndpoint {
	return RestEndpoint{
		GetByIdEndpoint:      makeGetByIdEndpoint(s),
		CreateEndpoint:       makeCreateEndpoint(s),
		ChangeStatusEndpoint: makeChangeStatusEndpoint(s),
	}
}

func makeGetByIdEndpoint(s orderservice.OderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		order, err := s.GetById(ctx, req.ID)
		return GetByIDResponse{Order: order, Err: err}, nil
	}
}

func makeCreateEndpoint(s orderservice.OderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Order)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeChangeStatusEndpoint(s orderservice.OderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeStatusRequest)
		err := s.ChangeStatus(ctx, req.ID, req.Status)
		return ChangeStatusResponse{Err: err}, nil
	}
}
