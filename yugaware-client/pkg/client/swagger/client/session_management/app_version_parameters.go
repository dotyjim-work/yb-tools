// Code generated by go-swagger; DO NOT EDIT.

package session_management

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
)

// NewAppVersionParams creates a new AppVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppVersionParams() *AppVersionParams {
	return &AppVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppVersionParamsWithTimeout creates a new AppVersionParams object
// with the ability to set a timeout on a request.
func NewAppVersionParamsWithTimeout(timeout time.Duration) *AppVersionParams {
	return &AppVersionParams{
		timeout: timeout,
	}
}

// NewAppVersionParamsWithContext creates a new AppVersionParams object
// with the ability to set a context for a request.
func NewAppVersionParamsWithContext(ctx context.Context) *AppVersionParams {
	return &AppVersionParams{
		Context: ctx,
	}
}

// NewAppVersionParamsWithHTTPClient creates a new AppVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppVersionParamsWithHTTPClient(client *http.Client) *AppVersionParams {
	return &AppVersionParams{
		HTTPClient: client,
	}
}

/* AppVersionParams contains all the parameters to send to the API endpoint
   for the app version operation.

   Typically these are written to a http.Request.
*/
type AppVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the app version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppVersionParams) WithDefaults() *AppVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the app version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the app version params
func (o *AppVersionParams) WithTimeout(timeout time.Duration) *AppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app version params
func (o *AppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app version params
func (o *AppVersionParams) WithContext(ctx context.Context) *AppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app version params
func (o *AppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app version params
func (o *AppVersionParams) WithHTTPClient(client *http.Client) *AppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app version params
func (o *AppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}