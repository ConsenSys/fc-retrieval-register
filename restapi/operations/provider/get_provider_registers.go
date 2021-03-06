// Code generated by go-swagger; DO NOT EDIT.

package provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProviderRegistersHandlerFunc turns a function with the right signature into a get provider registers handler
type GetProviderRegistersHandlerFunc func(GetProviderRegistersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProviderRegistersHandlerFunc) Handle(params GetProviderRegistersParams) middleware.Responder {
	return fn(params)
}

// GetProviderRegistersHandler interface for that can handle valid get provider registers params
type GetProviderRegistersHandler interface {
	Handle(GetProviderRegistersParams) middleware.Responder
}

// NewGetProviderRegisters creates a new http.Handler for the get provider registers operation
func NewGetProviderRegisters(ctx *middleware.Context, handler GetProviderRegistersHandler) *GetProviderRegisters {
	return &GetProviderRegisters{Context: ctx, Handler: handler}
}

/* GetProviderRegisters swagger:route GET /registers/provider Provider getProviderRegisters

Get Provider register list

<b>Get Provider register list</b>

*/
type GetProviderRegisters struct {
	Context *middleware.Context
	Handler GetProviderRegistersHandler
}

func (o *GetProviderRegisters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProviderRegistersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
