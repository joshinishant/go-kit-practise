package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	svcEndpoints RestEndpoint, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	// HTTP Post - /orders
	r.Methods("POST").Path("/orders").Handler(kithttp.NewServer(
		svcEndpoints.CreateEndpoint,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("GET").Path("/orders/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByIdEndpoint,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/status
	r.Methods("POST").Path("/orders/status").Handler(kithttp.NewServer(
		svcEndpoints.ChangeStatusEndpoint,
		decodeChangeStausRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Order); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetByIDRequest{ID: id}, nil
}

func decodeChangeStausRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req ChangeStatusRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
