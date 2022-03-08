// Code generated by go-swagger; DO NOT EDIT.

package asynchronous_replication

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

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// NewEditXClusterConfigParams creates a new EditXClusterConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEditXClusterConfigParams() *EditXClusterConfigParams {
	return &EditXClusterConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEditXClusterConfigParamsWithTimeout creates a new EditXClusterConfigParams object
// with the ability to set a timeout on a request.
func NewEditXClusterConfigParamsWithTimeout(timeout time.Duration) *EditXClusterConfigParams {
	return &EditXClusterConfigParams{
		timeout: timeout,
	}
}

// NewEditXClusterConfigParamsWithContext creates a new EditXClusterConfigParams object
// with the ability to set a context for a request.
func NewEditXClusterConfigParamsWithContext(ctx context.Context) *EditXClusterConfigParams {
	return &EditXClusterConfigParams{
		Context: ctx,
	}
}

// NewEditXClusterConfigParamsWithHTTPClient creates a new EditXClusterConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewEditXClusterConfigParamsWithHTTPClient(client *http.Client) *EditXClusterConfigParams {
	return &EditXClusterConfigParams{
		HTTPClient: client,
	}
}

/* EditXClusterConfigParams contains all the parameters to send to the API endpoint
   for the edit x cluster config operation.

   Typically these are written to a http.Request.
*/
type EditXClusterConfigParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// XccUUID.
	//
	// Format: uuid
	XccUUID strfmt.UUID

	/* XclusterReplicationEditFormData.

	   XCluster Replication Edit Form Data
	*/
	XclusterReplicationEditFormData *models.XClusterConfigEditFormData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edit x cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditXClusterConfigParams) WithDefaults() *EditXClusterConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edit x cluster config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EditXClusterConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edit x cluster config params
func (o *EditXClusterConfigParams) WithTimeout(timeout time.Duration) *EditXClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit x cluster config params
func (o *EditXClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit x cluster config params
func (o *EditXClusterConfigParams) WithContext(ctx context.Context) *EditXClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit x cluster config params
func (o *EditXClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit x cluster config params
func (o *EditXClusterConfigParams) WithHTTPClient(client *http.Client) *EditXClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit x cluster config params
func (o *EditXClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the edit x cluster config params
func (o *EditXClusterConfigParams) WithCUUID(cUUID strfmt.UUID) *EditXClusterConfigParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the edit x cluster config params
func (o *EditXClusterConfigParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithXccUUID adds the xccUUID to the edit x cluster config params
func (o *EditXClusterConfigParams) WithXccUUID(xccUUID strfmt.UUID) *EditXClusterConfigParams {
	o.SetXccUUID(xccUUID)
	return o
}

// SetXccUUID adds the xccUuid to the edit x cluster config params
func (o *EditXClusterConfigParams) SetXccUUID(xccUUID strfmt.UUID) {
	o.XccUUID = xccUUID
}

// WithXclusterReplicationEditFormData adds the xclusterReplicationEditFormData to the edit x cluster config params
func (o *EditXClusterConfigParams) WithXclusterReplicationEditFormData(xclusterReplicationEditFormData *models.XClusterConfigEditFormData) *EditXClusterConfigParams {
	o.SetXclusterReplicationEditFormData(xclusterReplicationEditFormData)
	return o
}

// SetXclusterReplicationEditFormData adds the xclusterReplicationEditFormData to the edit x cluster config params
func (o *EditXClusterConfigParams) SetXclusterReplicationEditFormData(xclusterReplicationEditFormData *models.XClusterConfigEditFormData) {
	o.XclusterReplicationEditFormData = xclusterReplicationEditFormData
}

// WriteToRequest writes these params to a swagger request
func (o *EditXClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param xccUUID
	if err := r.SetPathParam("xccUUID", o.XccUUID.String()); err != nil {
		return err
	}
	if o.XclusterReplicationEditFormData != nil {
		if err := r.SetBodyParam(o.XclusterReplicationEditFormData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}