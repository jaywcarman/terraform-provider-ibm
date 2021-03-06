// Code generated by go-swagger; DO NOT EDIT.

package v_p_cs

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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewUpdateVpcRouteParams creates a new UpdateVpcRouteParams object
// with the default values initialized.
func NewUpdateVpcRouteParams() *UpdateVpcRouteParams {
	var ()
	return &UpdateVpcRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVpcRouteParamsWithTimeout creates a new UpdateVpcRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVpcRouteParamsWithTimeout(timeout time.Duration) *UpdateVpcRouteParams {
	var ()
	return &UpdateVpcRouteParams{

		timeout: timeout,
	}
}

// NewUpdateVpcRouteParamsWithContext creates a new UpdateVpcRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVpcRouteParamsWithContext(ctx context.Context) *UpdateVpcRouteParams {
	var ()
	return &UpdateVpcRouteParams{

		Context: ctx,
	}
}

// NewUpdateVpcRouteParamsWithHTTPClient creates a new UpdateVpcRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVpcRouteParamsWithHTTPClient(client *http.Client) *UpdateVpcRouteParams {
	var ()
	return &UpdateVpcRouteParams{
		HTTPClient: client,
	}
}

/*UpdateVpcRouteParams contains all the parameters to send to the API endpoint
for the update vpc route operation typically these are written to a http.Request
*/
type UpdateVpcRouteParams struct {

	/*RoutePatch
	  The route patch

	*/
	RoutePatch *models.RoutePatch
	/*Generation
	  The infrastructure generation for the request. For the API behavior documented here, use `1` or `2`.

	*/
	Generation int64
	/*ID
	  The route identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update vpc route params
func (o *UpdateVpcRouteParams) WithTimeout(timeout time.Duration) *UpdateVpcRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vpc route params
func (o *UpdateVpcRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vpc route params
func (o *UpdateVpcRouteParams) WithContext(ctx context.Context) *UpdateVpcRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vpc route params
func (o *UpdateVpcRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vpc route params
func (o *UpdateVpcRouteParams) WithHTTPClient(client *http.Client) *UpdateVpcRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vpc route params
func (o *UpdateVpcRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoutePatch adds the routePatch to the update vpc route params
func (o *UpdateVpcRouteParams) WithRoutePatch(routePatch *models.RoutePatch) *UpdateVpcRouteParams {
	o.SetRoutePatch(routePatch)
	return o
}

// SetRoutePatch adds the routePatch to the update vpc route params
func (o *UpdateVpcRouteParams) SetRoutePatch(routePatch *models.RoutePatch) {
	o.RoutePatch = routePatch
}

// WithGeneration adds the generation to the update vpc route params
func (o *UpdateVpcRouteParams) WithGeneration(generation int64) *UpdateVpcRouteParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the update vpc route params
func (o *UpdateVpcRouteParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the update vpc route params
func (o *UpdateVpcRouteParams) WithID(id string) *UpdateVpcRouteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update vpc route params
func (o *UpdateVpcRouteParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the update vpc route params
func (o *UpdateVpcRouteParams) WithVersion(version string) *UpdateVpcRouteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update vpc route params
func (o *UpdateVpcRouteParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the update vpc route params
func (o *UpdateVpcRouteParams) WithVpcID(vpcID string) *UpdateVpcRouteParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the update vpc route params
func (o *UpdateVpcRouteParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVpcRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoutePatch != nil {
		if err := r.SetBodyParam(o.RoutePatch); err != nil {
			return err
		}
	}

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

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
