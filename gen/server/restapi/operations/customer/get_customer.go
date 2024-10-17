// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCustomerHandlerFunc turns a function with the right signature into a get customer handler
type GetCustomerHandlerFunc func(GetCustomerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCustomerHandlerFunc) Handle(params GetCustomerParams) middleware.Responder {
	return fn(params)
}

// GetCustomerHandler interface for that can handle valid get customer params
type GetCustomerHandler interface {
	Handle(GetCustomerParams) middleware.Responder
}

// NewGetCustomer creates a new http.Handler for the get customer operation
func NewGetCustomer(ctx *middleware.Context, handler GetCustomerHandler) *GetCustomer {
	return &GetCustomer{Context: ctx, Handler: handler}
}

/*
	GetCustomer swagger:route GET /api/customer customer getCustomer

Get customer
*/
type GetCustomer struct {
	Context *middleware.Context
	Handler GetCustomerHandler
}

func (o *GetCustomer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCustomerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
