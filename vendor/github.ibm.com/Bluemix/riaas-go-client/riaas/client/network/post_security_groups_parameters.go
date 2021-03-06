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

// NewPostSecurityGroupsParams creates a new PostSecurityGroupsParams object
// with the default values initialized.
func NewPostSecurityGroupsParams() *PostSecurityGroupsParams {
	var ()
	return &PostSecurityGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSecurityGroupsParamsWithTimeout creates a new PostSecurityGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSecurityGroupsParamsWithTimeout(timeout time.Duration) *PostSecurityGroupsParams {
	var ()
	return &PostSecurityGroupsParams{

		timeout: timeout,
	}
}

// NewPostSecurityGroupsParamsWithContext creates a new PostSecurityGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSecurityGroupsParamsWithContext(ctx context.Context) *PostSecurityGroupsParams {
	var ()
	return &PostSecurityGroupsParams{

		Context: ctx,
	}
}

// NewPostSecurityGroupsParamsWithHTTPClient creates a new PostSecurityGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSecurityGroupsParamsWithHTTPClient(client *http.Client) *PostSecurityGroupsParams {
	var ()
	return &PostSecurityGroupsParams{
		HTTPClient: client,
	}
}

/*PostSecurityGroupsParams contains all the parameters to send to the API endpoint
for the post security groups operation typically these are written to a http.Request
*/
type PostSecurityGroupsParams struct {

	/*Body*/
	Body PostSecurityGroupsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post security groups params
func (o *PostSecurityGroupsParams) WithTimeout(timeout time.Duration) *PostSecurityGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post security groups params
func (o *PostSecurityGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post security groups params
func (o *PostSecurityGroupsParams) WithContext(ctx context.Context) *PostSecurityGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post security groups params
func (o *PostSecurityGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post security groups params
func (o *PostSecurityGroupsParams) WithHTTPClient(client *http.Client) *PostSecurityGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post security groups params
func (o *PostSecurityGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post security groups params
func (o *PostSecurityGroupsParams) WithBody(body PostSecurityGroupsBody) *PostSecurityGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post security groups params
func (o *PostSecurityGroupsParams) SetBody(body PostSecurityGroupsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post security groups params
func (o *PostSecurityGroupsParams) WithGeneration(generation int64) *PostSecurityGroupsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post security groups params
func (o *PostSecurityGroupsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the post security groups params
func (o *PostSecurityGroupsParams) WithVersion(version string) *PostSecurityGroupsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post security groups params
func (o *PostSecurityGroupsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostSecurityGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
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
