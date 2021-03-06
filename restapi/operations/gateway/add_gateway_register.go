// Code generated by go-swagger; DO NOT EDIT.

package gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddGatewayRegisterHandlerFunc turns a function with the right signature into a add gateway register handler
type AddGatewayRegisterHandlerFunc func(AddGatewayRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddGatewayRegisterHandlerFunc) Handle(params AddGatewayRegisterParams) middleware.Responder {
	return fn(params)
}

// AddGatewayRegisterHandler interface for that can handle valid add gateway register params
type AddGatewayRegisterHandler interface {
	Handle(AddGatewayRegisterParams) middleware.Responder
}

// NewAddGatewayRegister creates a new http.Handler for the add gateway register operation
func NewAddGatewayRegister(ctx *middleware.Context, handler AddGatewayRegisterHandler) *AddGatewayRegister {
	return &AddGatewayRegister{Context: ctx, Handler: handler}
}

/* AddGatewayRegister swagger:route POST /registers/gateway Gateway addGatewayRegister

Add a Gateway register

<b>Add a Gateway register</b>

*/
type AddGatewayRegister struct {
	Context *middleware.Context
	Handler AddGatewayRegisterHandler
}

func (o *AddGatewayRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddGatewayRegisterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
