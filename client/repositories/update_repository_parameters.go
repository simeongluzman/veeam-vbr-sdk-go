// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/veeamhub/veeam-vbr-sdk-go/models"
)

// NewUpdateRepositoryParams creates a new UpdateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepositoryParams() *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepositoryParamsWithTimeout creates a new UpdateRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateRepositoryParamsWithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateRepositoryParamsWithContext creates a new UpdateRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateRepositoryParamsWithContext(ctx context.Context) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateRepositoryParamsWithHTTPClient creates a new UpdateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepositoryParamsWithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		HTTPClient: client,
	}
}

/* UpdateRepositoryParams contains all the parameters to send to the API endpoint
   for the update repository operation.

   Typically these are written to a http.Request.
*/
type UpdateRepositoryParams struct {

	// Body.
	Body *models.RepositoryModel

	/* ID.

	   ID of the backup repository.

	   Format: uuid
	*/
	ID strfmt.UUID

	/* XAPIVersion.

	     Version and revision of the client REST API. Must be in the following
	format: *\<version\>-\<revision\>*.


	     Default: "1.0-rev1"
	*/
	XAPIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepositoryParams) WithDefaults() *UpdateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepositoryParams) SetDefaults() {
	var (
		xAPIVersionDefault = string("1.0-rev1")
	)

	val := UpdateRepositoryParams{
		XAPIVersion: xAPIVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) WithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository params
func (o *UpdateRepositoryParams) WithContext(ctx context.Context) *UpdateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository params
func (o *UpdateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) WithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update repository params
func (o *UpdateRepositoryParams) WithBody(body *models.RepositoryModel) *UpdateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update repository params
func (o *UpdateRepositoryParams) SetBody(body *models.RepositoryModel) {
	o.Body = body
}

// WithID adds the id to the update repository params
func (o *UpdateRepositoryParams) WithID(id strfmt.UUID) *UpdateRepositoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update repository params
func (o *UpdateRepositoryParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithXAPIVersion adds the xAPIVersion to the update repository params
func (o *UpdateRepositoryParams) WithXAPIVersion(xAPIVersion string) *UpdateRepositoryParams {
	o.SetXAPIVersion(xAPIVersion)
	return o
}

// SetXAPIVersion adds the xApiVersion to the update repository params
func (o *UpdateRepositoryParams) SetXAPIVersion(xAPIVersion string) {
	o.XAPIVersion = xAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// header param x-api-version
	if err := r.SetHeaderParam("x-api-version", o.XAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}