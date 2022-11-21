// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swg/models"
)

// SignRefreshOKCode is the HTTP code returned for type SignRefreshOK
const SignRefreshOKCode int = 200

/*
SignRefreshOK Токены обновились

swagger:response signRefreshOK
*/
type SignRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *models.SignInRes `json:"body,omitempty"`
}

// NewSignRefreshOK creates SignRefreshOK with default headers values
func NewSignRefreshOK() *SignRefreshOK {

	return &SignRefreshOK{}
}

// WithPayload adds the payload to the sign refresh o k response
func (o *SignRefreshOK) WithPayload(payload *models.SignInRes) *SignRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sign refresh o k response
func (o *SignRefreshOK) SetPayload(payload *models.SignInRes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SignRefreshInternalServerErrorCode is the HTTP code returned for type SignRefreshInternalServerError
const SignRefreshInternalServerErrorCode int = 500

/*
SignRefreshInternalServerError Ошибка сервера

swagger:response signRefreshInternalServerError
*/
type SignRefreshInternalServerError struct {
}

// NewSignRefreshInternalServerError creates SignRefreshInternalServerError with default headers values
func NewSignRefreshInternalServerError() *SignRefreshInternalServerError {

	return &SignRefreshInternalServerError{}
}

// WriteResponse to the client
func (o *SignRefreshInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

/*
SignRefreshDefault Ошибка

swagger:response signRefreshDefault
*/
type SignRefreshDefault struct {
	_statusCode int
}

// NewSignRefreshDefault creates SignRefreshDefault with default headers values
func NewSignRefreshDefault(code int) *SignRefreshDefault {
	if code <= 0 {
		code = 500
	}

	return &SignRefreshDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sign refresh default response
func (o *SignRefreshDefault) WithStatusCode(code int) *SignRefreshDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sign refresh default response
func (o *SignRefreshDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *SignRefreshDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}