// Code generated by go-swagger; DO NOT EDIT.

package universe_information

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

// NewGetUniverseCostForAllParams creates a new GetUniverseCostForAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseCostForAllParams() *GetUniverseCostForAllParams {
	return &GetUniverseCostForAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseCostForAllParamsWithTimeout creates a new GetUniverseCostForAllParams object
// with the ability to set a timeout on a request.
func NewGetUniverseCostForAllParamsWithTimeout(timeout time.Duration) *GetUniverseCostForAllParams {
	return &GetUniverseCostForAllParams{
		timeout: timeout,
	}
}

// NewGetUniverseCostForAllParamsWithContext creates a new GetUniverseCostForAllParams object
// with the ability to set a context for a request.
func NewGetUniverseCostForAllParamsWithContext(ctx context.Context) *GetUniverseCostForAllParams {
	return &GetUniverseCostForAllParams{
		Context: ctx,
	}
}

// NewGetUniverseCostForAllParamsWithHTTPClient creates a new GetUniverseCostForAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseCostForAllParamsWithHTTPClient(client *http.Client) *GetUniverseCostForAllParams {
	return &GetUniverseCostForAllParams{
		HTTPClient: client,
	}
}

/* GetUniverseCostForAllParams contains all the parameters to send to the API endpoint
   for the get universe cost for all operation.

   Typically these are written to a http.Request.
*/
type GetUniverseCostForAllParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe cost for all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseCostForAllParams) WithDefaults() *GetUniverseCostForAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe cost for all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseCostForAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get universe cost for all params
func (o *GetUniverseCostForAllParams) WithTimeout(timeout time.Duration) *GetUniverseCostForAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe cost for all params
func (o *GetUniverseCostForAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe cost for all params
func (o *GetUniverseCostForAllParams) WithContext(ctx context.Context) *GetUniverseCostForAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe cost for all params
func (o *GetUniverseCostForAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe cost for all params
func (o *GetUniverseCostForAllParams) WithHTTPClient(client *http.Client) *GetUniverseCostForAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe cost for all params
func (o *GetUniverseCostForAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get universe cost for all params
func (o *GetUniverseCostForAllParams) WithCUUID(cUUID strfmt.UUID) *GetUniverseCostForAllParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get universe cost for all params
func (o *GetUniverseCostForAllParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseCostForAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}