// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SignInHandlerFunc turns a function with the right signature into a sign in handler
type SignInHandlerFunc func(SignInParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SignInHandlerFunc) Handle(params SignInParams) middleware.Responder {
	return fn(params)
}

// SignInHandler interface for that can handle valid sign in params
type SignInHandler interface {
	Handle(SignInParams) middleware.Responder
}

// NewSignIn creates a new http.Handler for the sign in operation
func NewSignIn(ctx *middleware.Context, handler SignInHandler) *SignIn {
	return &SignIn{Context: ctx, Handler: handler}
}

/*
	SignIn swagger:route POST /sign-in auth signIn

sign-in

При запросе на этот эндпоинт произойдет отправка пароля и логина, также в куки будет положен отпечаток браузера (x-request-id) и токен обновления (httpOnly)
*/
type SignIn struct {
	Context *middleware.Context
	Handler SignInHandler
}

func (o *SignIn) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSignInParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}