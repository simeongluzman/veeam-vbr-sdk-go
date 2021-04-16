// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// CreateRepositoryReader is a Reader for the CreateRepository structure.
type CreateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepositoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepositoryCreated creates a CreateRepositoryCreated with default headers values
func NewCreateRepositoryCreated() *CreateRepositoryCreated {
	return &CreateRepositoryCreated{}
}

/* CreateRepositoryCreated describes a response with status code 201, with default header values.

Infrastructure session has been created to add the repository. To check the progress, track the session `state`.
*/
type CreateRepositoryCreated struct {
	Payload *models.SessionModel
}

func (o *CreateRepositoryCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/repositories][%d] createRepositoryCreated  %+v", 201, o.Payload)
}
func (o *CreateRepositoryCreated) GetPayload() *models.SessionModel {
	return o.Payload
}

func (o *CreateRepositoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRepositoryBadRequest creates a CreateRepositoryBadRequest with default headers values
func NewCreateRepositoryBadRequest() *CreateRepositoryBadRequest {
	return &CreateRepositoryBadRequest{}
}

/* CreateRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type CreateRepositoryBadRequest struct {
	Payload *models.Error
}

func (o *CreateRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/repositories][%d] createRepositoryBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRepositoryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRepositoryUnauthorized creates a CreateRepositoryUnauthorized with default headers values
func NewCreateRepositoryUnauthorized() *CreateRepositoryUnauthorized {
	return &CreateRepositoryUnauthorized{}
}

/* CreateRepositoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type CreateRepositoryUnauthorized struct {
	Payload *models.Error
}

func (o *CreateRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/repositories][%d] createRepositoryUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateRepositoryUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRepositoryForbidden creates a CreateRepositoryForbidden with default headers values
func NewCreateRepositoryForbidden() *CreateRepositoryForbidden {
	return &CreateRepositoryForbidden{}
}

/* CreateRepositoryForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type CreateRepositoryForbidden struct {
	Payload *models.Error
}

func (o *CreateRepositoryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/repositories][%d] createRepositoryForbidden  %+v", 403, o.Payload)
}
func (o *CreateRepositoryForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRepositoryInternalServerError creates a CreateRepositoryInternalServerError with default headers values
func NewCreateRepositoryInternalServerError() *CreateRepositoryInternalServerError {
	return &CreateRepositoryInternalServerError{}
}

/* CreateRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type CreateRepositoryInternalServerError struct {
	Payload *models.Error
}

func (o *CreateRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/backupInfrastructure/repositories][%d] createRepositoryInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRepositoryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}