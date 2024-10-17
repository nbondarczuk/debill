// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAddressHandlerFunc turns a function with the right signature into a get address handler
type GetAddressHandlerFunc func(GetAddressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAddressHandlerFunc) Handle(params GetAddressParams) middleware.Responder {
	return fn(params)
}

// GetAddressHandler interface for that can handle valid get address params
type GetAddressHandler interface {
	Handle(GetAddressParams) middleware.Responder
}

// NewGetAddress creates a new http.Handler for the get address operation
func NewGetAddress(ctx *middleware.Context, handler GetAddressHandler) *GetAddress {
	return &GetAddress{Context: ctx, Handler: handler}
}

/*
	GetAddress swagger:route GET /api/address address getAddress

Get address
*/
type GetAddress struct {
	Context *middleware.Context
	Handler GetAddressHandler
}

func (o *GetAddress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAddressParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
