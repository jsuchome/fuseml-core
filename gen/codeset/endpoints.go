// Code generated by goa v3.3.1, DO NOT EDIT.
//
// codeset endpoints
//
// Command:
// $ goa gen github.com/fuseml/fuseml-core/design

package codeset

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "codeset" service endpoints.
type Endpoints struct {
	List     goa.Endpoint
	Register goa.Endpoint
	Get      goa.Endpoint
}

// NewEndpoints wraps the methods of the "codeset" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:     NewListEndpoint(s),
		Register: NewRegisterEndpoint(s),
		Get:      NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "codeset" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Register = m(e.Register)
	e.Get = m(e.Get)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "codeset".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		return s.List(ctx, p)
	}
}

// NewRegisterEndpoint returns an endpoint function that calls the method
// "register" of service "codeset".
func NewRegisterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterPayload)
		return s.Register(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "codeset".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
		return s.Get(ctx, p)
	}
}
