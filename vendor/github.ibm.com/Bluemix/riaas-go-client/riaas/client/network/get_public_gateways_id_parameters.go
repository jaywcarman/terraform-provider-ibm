// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicGatewaysIDParams creates a new GetPublicGatewaysIDParams object
// with the default values initialized.
func NewGetPublicGatewaysIDParams() *GetPublicGatewaysIDParams {
	var ()
	return &GetPublicGatewaysIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGatewaysIDParamsWithTimeout creates a new GetPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGatewaysIDParamsWithTimeout(timeout time.Duration) *GetPublicGatewaysIDParams {
	var ()
	return &GetPublicGatewaysIDParams{

		timeout: timeout,
	}
}

// NewGetPublicGatewaysIDParamsWithContext creates a new GetPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGatewaysIDParamsWithContext(ctx context.Context) *GetPublicGatewaysIDParams {
	var ()
	return &GetPublicGatewaysIDParams{

		Context: ctx,
	}
}

// NewGetPublicGatewaysIDParamsWithHTTPClient creates a new GetPublicGatewaysIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGatewaysIDParamsWithHTTPClient(client *http.Client) *GetPublicGatewaysIDParams {
	var ()
	return &GetPublicGatewaysIDParams{
		HTTPClient: client,
	}
}

/*GetPublicGatewaysIDParams contains all the parameters to send to the API endpoint
for the get public gateways ID operation typically these are written to a http.Request
*/
type GetPublicGatewaysIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The public gateway idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithTimeout(timeout time.Duration) *GetPublicGatewaysIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithContext(ctx context.Context) *GetPublicGatewaysIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithHTTPClient(client *http.Client) *GetPublicGatewaysIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithGeneration(generation int64) *GetPublicGatewaysIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithID(id string) *GetPublicGatewaysIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) WithVersion(version string) *GetPublicGatewaysIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get public gateways ID params
func (o *GetPublicGatewaysIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGatewaysIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
