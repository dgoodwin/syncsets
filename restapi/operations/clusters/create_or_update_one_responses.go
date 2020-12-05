// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dgoodwin/syncsets/models"
)

// CreateOrUpdateOneOKCode is the HTTP code returned for type CreateOrUpdateOneOK
const CreateOrUpdateOneOKCode int = 200

/*CreateOrUpdateOneOK OK

swagger:response createOrUpdateOneOK
*/
type CreateOrUpdateOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewCreateOrUpdateOneOK creates CreateOrUpdateOneOK with default headers values
func NewCreateOrUpdateOneOK() *CreateOrUpdateOneOK {

	return &CreateOrUpdateOneOK{}
}

// WithPayload adds the payload to the create or update one o k response
func (o *CreateOrUpdateOneOK) WithPayload(payload *models.Cluster) *CreateOrUpdateOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update one o k response
func (o *CreateOrUpdateOneOK) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateOrUpdateOneDefault error

swagger:response createOrUpdateOneDefault
*/
type CreateOrUpdateOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrUpdateOneDefault creates CreateOrUpdateOneDefault with default headers values
func NewCreateOrUpdateOneDefault(code int) *CreateOrUpdateOneDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateOrUpdateOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create or update one default response
func (o *CreateOrUpdateOneDefault) WithStatusCode(code int) *CreateOrUpdateOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create or update one default response
func (o *CreateOrUpdateOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create or update one default response
func (o *CreateOrUpdateOneDefault) WithPayload(payload *models.Error) *CreateOrUpdateOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create or update one default response
func (o *CreateOrUpdateOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrUpdateOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
