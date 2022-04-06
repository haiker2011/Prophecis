// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// GetModelVersionByNameAndVersionOKCode is the HTTP code returned for type GetModelVersionByNameAndVersionOK
const GetModelVersionByNameAndVersionOKCode int = 200

/*GetModelVersionByNameAndVersionOK OK

swagger:response getModelVersionByNameAndVersionOK
*/
type GetModelVersionByNameAndVersionOK struct {

	/*
	  In: Body
	*/
	Payload models.GetModelVersionResp `json:"body,omitempty"`
}

// NewGetModelVersionByNameAndVersionOK creates GetModelVersionByNameAndVersionOK with default headers values
func NewGetModelVersionByNameAndVersionOK() *GetModelVersionByNameAndVersionOK {

	return &GetModelVersionByNameAndVersionOK{}
}

// WithPayload adds the payload to the get model version by name and version o k response
func (o *GetModelVersionByNameAndVersionOK) WithPayload(payload models.GetModelVersionResp) *GetModelVersionByNameAndVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model version by name and version o k response
func (o *GetModelVersionByNameAndVersionOK) SetPayload(payload models.GetModelVersionResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelVersionByNameAndVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.GetModelVersionResp{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetModelVersionByNameAndVersionUnauthorizedCode is the HTTP code returned for type GetModelVersionByNameAndVersionUnauthorized
const GetModelVersionByNameAndVersionUnauthorizedCode int = 401

/*GetModelVersionByNameAndVersionUnauthorized Unauthorized

swagger:response getModelVersionByNameAndVersionUnauthorized
*/
type GetModelVersionByNameAndVersionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelVersionByNameAndVersionUnauthorized creates GetModelVersionByNameAndVersionUnauthorized with default headers values
func NewGetModelVersionByNameAndVersionUnauthorized() *GetModelVersionByNameAndVersionUnauthorized {

	return &GetModelVersionByNameAndVersionUnauthorized{}
}

// WithPayload adds the payload to the get model version by name and version unauthorized response
func (o *GetModelVersionByNameAndVersionUnauthorized) WithPayload(payload *models.Error) *GetModelVersionByNameAndVersionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model version by name and version unauthorized response
func (o *GetModelVersionByNameAndVersionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelVersionByNameAndVersionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetModelVersionByNameAndVersionNotFoundCode is the HTTP code returned for type GetModelVersionByNameAndVersionNotFound
const GetModelVersionByNameAndVersionNotFoundCode int = 404

/*GetModelVersionByNameAndVersionNotFound The Models cannot be found

swagger:response getModelVersionByNameAndVersionNotFound
*/
type GetModelVersionByNameAndVersionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetModelVersionByNameAndVersionNotFound creates GetModelVersionByNameAndVersionNotFound with default headers values
func NewGetModelVersionByNameAndVersionNotFound() *GetModelVersionByNameAndVersionNotFound {

	return &GetModelVersionByNameAndVersionNotFound{}
}

// WithPayload adds the payload to the get model version by name and version not found response
func (o *GetModelVersionByNameAndVersionNotFound) WithPayload(payload *models.Error) *GetModelVersionByNameAndVersionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get model version by name and version not found response
func (o *GetModelVersionByNameAndVersionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetModelVersionByNameAndVersionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}