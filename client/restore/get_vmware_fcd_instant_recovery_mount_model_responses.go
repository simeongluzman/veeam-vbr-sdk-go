// Code generated by go-swagger; DO NOT EDIT.

package restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// GetVmwareFcdInstantRecoveryMountModelReader is a Reader for the GetVmwareFcdInstantRecoveryMountModel structure.
type GetVmwareFcdInstantRecoveryMountModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmwareFcdInstantRecoveryMountModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVmwareFcdInstantRecoveryMountModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVmwareFcdInstantRecoveryMountModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVmwareFcdInstantRecoveryMountModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVmwareFcdInstantRecoveryMountModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVmwareFcdInstantRecoveryMountModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVmwareFcdInstantRecoveryMountModelOK creates a GetVmwareFcdInstantRecoveryMountModelOK with default headers values
func NewGetVmwareFcdInstantRecoveryMountModelOK() *GetVmwareFcdInstantRecoveryMountModelOK {
	return &GetVmwareFcdInstantRecoveryMountModelOK{}
}

/* GetVmwareFcdInstantRecoveryMountModelOK describes a response with status code 200, with default header values.

OK
*/
type GetVmwareFcdInstantRecoveryMountModelOK struct {
	Payload *models.VmwareFcdInstantRecoveryMount
}

func (o *GetVmwareFcdInstantRecoveryMountModelOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/restore/instantRecovery/vmware/fcd/{mountId}][%d] getVmwareFcdInstantRecoveryMountModelOK  %+v", 200, o.Payload)
}
func (o *GetVmwareFcdInstantRecoveryMountModelOK) GetPayload() *models.VmwareFcdInstantRecoveryMount {
	return o.Payload
}

func (o *GetVmwareFcdInstantRecoveryMountModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VmwareFcdInstantRecoveryMount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareFcdInstantRecoveryMountModelUnauthorized creates a GetVmwareFcdInstantRecoveryMountModelUnauthorized with default headers values
func NewGetVmwareFcdInstantRecoveryMountModelUnauthorized() *GetVmwareFcdInstantRecoveryMountModelUnauthorized {
	return &GetVmwareFcdInstantRecoveryMountModelUnauthorized{}
}

/* GetVmwareFcdInstantRecoveryMountModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type GetVmwareFcdInstantRecoveryMountModelUnauthorized struct {
	Payload *models.Error
}

func (o *GetVmwareFcdInstantRecoveryMountModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/restore/instantRecovery/vmware/fcd/{mountId}][%d] getVmwareFcdInstantRecoveryMountModelUnauthorized  %+v", 401, o.Payload)
}
func (o *GetVmwareFcdInstantRecoveryMountModelUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareFcdInstantRecoveryMountModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareFcdInstantRecoveryMountModelForbidden creates a GetVmwareFcdInstantRecoveryMountModelForbidden with default headers values
func NewGetVmwareFcdInstantRecoveryMountModelForbidden() *GetVmwareFcdInstantRecoveryMountModelForbidden {
	return &GetVmwareFcdInstantRecoveryMountModelForbidden{}
}

/* GetVmwareFcdInstantRecoveryMountModelForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type GetVmwareFcdInstantRecoveryMountModelForbidden struct {
	Payload *models.Error
}

func (o *GetVmwareFcdInstantRecoveryMountModelForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/restore/instantRecovery/vmware/fcd/{mountId}][%d] getVmwareFcdInstantRecoveryMountModelForbidden  %+v", 403, o.Payload)
}
func (o *GetVmwareFcdInstantRecoveryMountModelForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareFcdInstantRecoveryMountModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareFcdInstantRecoveryMountModelNotFound creates a GetVmwareFcdInstantRecoveryMountModelNotFound with default headers values
func NewGetVmwareFcdInstantRecoveryMountModelNotFound() *GetVmwareFcdInstantRecoveryMountModelNotFound {
	return &GetVmwareFcdInstantRecoveryMountModelNotFound{}
}

/* GetVmwareFcdInstantRecoveryMountModelNotFound describes a response with status code 404, with default header values.

Not found. No object was found with the path parameter specified in the request.
*/
type GetVmwareFcdInstantRecoveryMountModelNotFound struct {
	Payload *models.Error
}

func (o *GetVmwareFcdInstantRecoveryMountModelNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/restore/instantRecovery/vmware/fcd/{mountId}][%d] getVmwareFcdInstantRecoveryMountModelNotFound  %+v", 404, o.Payload)
}
func (o *GetVmwareFcdInstantRecoveryMountModelNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareFcdInstantRecoveryMountModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmwareFcdInstantRecoveryMountModelInternalServerError creates a GetVmwareFcdInstantRecoveryMountModelInternalServerError with default headers values
func NewGetVmwareFcdInstantRecoveryMountModelInternalServerError() *GetVmwareFcdInstantRecoveryMountModelInternalServerError {
	return &GetVmwareFcdInstantRecoveryMountModelInternalServerError{}
}

/* GetVmwareFcdInstantRecoveryMountModelInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type GetVmwareFcdInstantRecoveryMountModelInternalServerError struct {
	Payload *models.Error
}

func (o *GetVmwareFcdInstantRecoveryMountModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/restore/instantRecovery/vmware/fcd/{mountId}][%d] getVmwareFcdInstantRecoveryMountModelInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVmwareFcdInstantRecoveryMountModelInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVmwareFcdInstantRecoveryMountModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}