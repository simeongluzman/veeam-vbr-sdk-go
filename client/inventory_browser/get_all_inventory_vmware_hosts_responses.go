// Code generated by go-swagger; DO NOT EDIT.

package inventory_browser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetAllInventoryVmwareHostsReader is a Reader for the GetAllInventoryVmwareHosts structure.
type GetAllInventoryVmwareHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllInventoryVmwareHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllInventoryVmwareHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllInventoryVmwareHostsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllInventoryVmwareHostsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllInventoryVmwareHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllInventoryVmwareHostsOK creates a GetAllInventoryVmwareHostsOK with default headers values
func NewGetAllInventoryVmwareHostsOK() *GetAllInventoryVmwareHostsOK {
	return &GetAllInventoryVmwareHostsOK{}
}

/* GetAllInventoryVmwareHostsOK describes a response with status code 200, with default header values.

OK
*/
type GetAllInventoryVmwareHostsOK struct {
	Payload *models.ViRootsResult
}

func (o *GetAllInventoryVmwareHostsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts][%d] getAllInventoryVmwareHostsOK  %+v", 200, o.Payload)
}
func (o *GetAllInventoryVmwareHostsOK) GetPayload() *models.ViRootsResult {
	return o.Payload
}

func (o *GetAllInventoryVmwareHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ViRootsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllInventoryVmwareHostsUnauthorized creates a GetAllInventoryVmwareHostsUnauthorized with default headers values
func NewGetAllInventoryVmwareHostsUnauthorized() *GetAllInventoryVmwareHostsUnauthorized {
	return &GetAllInventoryVmwareHostsUnauthorized{}
}

/* GetAllInventoryVmwareHostsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetAllInventoryVmwareHostsUnauthorized struct {
	Payload *models.Error
}

func (o *GetAllInventoryVmwareHostsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts][%d] getAllInventoryVmwareHostsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAllInventoryVmwareHostsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllInventoryVmwareHostsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllInventoryVmwareHostsForbidden creates a GetAllInventoryVmwareHostsForbidden with default headers values
func NewGetAllInventoryVmwareHostsForbidden() *GetAllInventoryVmwareHostsForbidden {
	return &GetAllInventoryVmwareHostsForbidden{}
}

/* GetAllInventoryVmwareHostsForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetAllInventoryVmwareHostsForbidden struct {
	Payload *models.Error
}

func (o *GetAllInventoryVmwareHostsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts][%d] getAllInventoryVmwareHostsForbidden  %+v", 403, o.Payload)
}
func (o *GetAllInventoryVmwareHostsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllInventoryVmwareHostsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllInventoryVmwareHostsInternalServerError creates a GetAllInventoryVmwareHostsInternalServerError with default headers values
func NewGetAllInventoryVmwareHostsInternalServerError() *GetAllInventoryVmwareHostsInternalServerError {
	return &GetAllInventoryVmwareHostsInternalServerError{}
}

/* GetAllInventoryVmwareHostsInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetAllInventoryVmwareHostsInternalServerError struct {
	Payload *models.Error
}

func (o *GetAllInventoryVmwareHostsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/inventory/vmware/hosts][%d] getAllInventoryVmwareHostsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAllInventoryVmwareHostsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllInventoryVmwareHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}