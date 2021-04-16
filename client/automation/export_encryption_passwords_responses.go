// Code generated by go-swagger; DO NOT EDIT.

package automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// ExportEncryptionPasswordsReader is a Reader for the ExportEncryptionPasswords structure.
type ExportEncryptionPasswordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportEncryptionPasswordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportEncryptionPasswordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExportEncryptionPasswordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExportEncryptionPasswordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportEncryptionPasswordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExportEncryptionPasswordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportEncryptionPasswordsOK creates a ExportEncryptionPasswordsOK with default headers values
func NewExportEncryptionPasswordsOK() *ExportEncryptionPasswordsOK {
	return &ExportEncryptionPasswordsOK{}
}

/* ExportEncryptionPasswordsOK describes a response with status code 200, with default header values.

OK
*/
type ExportEncryptionPasswordsOK struct {
	Payload *models.EncryptionPasswordImportSpecCollection
}

func (o *ExportEncryptionPasswordsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/automation/encryptionPasswords/export][%d] exportEncryptionPasswordsOK  %+v", 200, o.Payload)
}
func (o *ExportEncryptionPasswordsOK) GetPayload() *models.EncryptionPasswordImportSpecCollection {
	return o.Payload
}

func (o *ExportEncryptionPasswordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionPasswordImportSpecCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionPasswordsBadRequest creates a ExportEncryptionPasswordsBadRequest with default headers values
func NewExportEncryptionPasswordsBadRequest() *ExportEncryptionPasswordsBadRequest {
	return &ExportEncryptionPasswordsBadRequest{}
}

/* ExportEncryptionPasswordsBadRequest describes a response with status code 400, with default header values.

Bad request. This error is related to POST/PUT requests. The request body is malformed, incomplete or otherwise invalid.
*/
type ExportEncryptionPasswordsBadRequest struct {
	Payload *models.Error
}

func (o *ExportEncryptionPasswordsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/automation/encryptionPasswords/export][%d] exportEncryptionPasswordsBadRequest  %+v", 400, o.Payload)
}
func (o *ExportEncryptionPasswordsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportEncryptionPasswordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionPasswordsUnauthorized creates a ExportEncryptionPasswordsUnauthorized with default headers values
func NewExportEncryptionPasswordsUnauthorized() *ExportEncryptionPasswordsUnauthorized {
	return &ExportEncryptionPasswordsUnauthorized{}
}

/* ExportEncryptionPasswordsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The authorization header has been expected but not found (or found but is expired).
*/
type ExportEncryptionPasswordsUnauthorized struct {
	Payload *models.Error
}

func (o *ExportEncryptionPasswordsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/automation/encryptionPasswords/export][%d] exportEncryptionPasswordsUnauthorized  %+v", 401, o.Payload)
}
func (o *ExportEncryptionPasswordsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportEncryptionPasswordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionPasswordsForbidden creates a ExportEncryptionPasswordsForbidden with default headers values
func NewExportEncryptionPasswordsForbidden() *ExportEncryptionPasswordsForbidden {
	return &ExportEncryptionPasswordsForbidden{}
}

/* ExportEncryptionPasswordsForbidden describes a response with status code 403, with default header values.

Forbidden. The user sending the request does not have adequate privileges to access one or more objects specified in the request.
*/
type ExportEncryptionPasswordsForbidden struct {
	Payload *models.Error
}

func (o *ExportEncryptionPasswordsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/automation/encryptionPasswords/export][%d] exportEncryptionPasswordsForbidden  %+v", 403, o.Payload)
}
func (o *ExportEncryptionPasswordsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportEncryptionPasswordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEncryptionPasswordsInternalServerError creates a ExportEncryptionPasswordsInternalServerError with default headers values
func NewExportEncryptionPasswordsInternalServerError() *ExportEncryptionPasswordsInternalServerError {
	return &ExportEncryptionPasswordsInternalServerError{}
}

/* ExportEncryptionPasswordsInternalServerError describes a response with status code 500, with default header values.

Internal server error. The request has been received but could not be completed because of an internal error at the server side.
*/
type ExportEncryptionPasswordsInternalServerError struct {
	Payload *models.Error
}

func (o *ExportEncryptionPasswordsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/automation/encryptionPasswords/export][%d] exportEncryptionPasswordsInternalServerError  %+v", 500, o.Payload)
}
func (o *ExportEncryptionPasswordsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExportEncryptionPasswordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}