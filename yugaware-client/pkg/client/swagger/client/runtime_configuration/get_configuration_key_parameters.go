// Code generated by go-swagger; DO NOT EDIT.

package runtime_configuration

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

// NewGetConfigurationKeyParams creates a new GetConfigurationKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigurationKeyParams() *GetConfigurationKeyParams {
	return &GetConfigurationKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationKeyParamsWithTimeout creates a new GetConfigurationKeyParams object
// with the ability to set a timeout on a request.
func NewGetConfigurationKeyParamsWithTimeout(timeout time.Duration) *GetConfigurationKeyParams {
	return &GetConfigurationKeyParams{
		timeout: timeout,
	}
}

// NewGetConfigurationKeyParamsWithContext creates a new GetConfigurationKeyParams object
// with the ability to set a context for a request.
func NewGetConfigurationKeyParamsWithContext(ctx context.Context) *GetConfigurationKeyParams {
	return &GetConfigurationKeyParams{
		Context: ctx,
	}
}

// NewGetConfigurationKeyParamsWithHTTPClient creates a new GetConfigurationKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigurationKeyParamsWithHTTPClient(client *http.Client) *GetConfigurationKeyParams {
	return &GetConfigurationKeyParams{
		HTTPClient: client,
	}
}

/* GetConfigurationKeyParams contains all the parameters to send to the API endpoint
   for the get configuration key operation.

   Typically these are written to a http.Request.
*/
type GetConfigurationKeyParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// Key.
	Key string

	// Scope.
	//
	// Format: uuid
	Scope strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get configuration key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationKeyParams) WithDefaults() *GetConfigurationKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get configuration key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get configuration key params
func (o *GetConfigurationKeyParams) WithTimeout(timeout time.Duration) *GetConfigurationKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration key params
func (o *GetConfigurationKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration key params
func (o *GetConfigurationKeyParams) WithContext(ctx context.Context) *GetConfigurationKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration key params
func (o *GetConfigurationKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration key params
func (o *GetConfigurationKeyParams) WithHTTPClient(client *http.Client) *GetConfigurationKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration key params
func (o *GetConfigurationKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get configuration key params
func (o *GetConfigurationKeyParams) WithCUUID(cUUID strfmt.UUID) *GetConfigurationKeyParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get configuration key params
func (o *GetConfigurationKeyParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithKey adds the key to the get configuration key params
func (o *GetConfigurationKeyParams) WithKey(key string) *GetConfigurationKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get configuration key params
func (o *GetConfigurationKeyParams) SetKey(key string) {
	o.Key = key
}

// WithScope adds the scope to the get configuration key params
func (o *GetConfigurationKeyParams) WithScope(scope strfmt.UUID) *GetConfigurationKeyParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get configuration key params
func (o *GetConfigurationKeyParams) SetScope(scope strfmt.UUID) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param scope
	if err := r.SetPathParam("scope", o.Scope.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}