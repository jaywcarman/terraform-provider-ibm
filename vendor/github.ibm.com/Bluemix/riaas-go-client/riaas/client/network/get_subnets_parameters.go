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

// NewGetSubnetsParams creates a new GetSubnetsParams object
// with the default values initialized.
func NewGetSubnetsParams() *GetSubnetsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetSubnetsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubnetsParamsWithTimeout creates a new GetSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubnetsParamsWithTimeout(timeout time.Duration) *GetSubnetsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetSubnetsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetSubnetsParamsWithContext creates a new GetSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubnetsParamsWithContext(ctx context.Context) *GetSubnetsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetSubnetsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetSubnetsParamsWithHTTPClient creates a new GetSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubnetsParamsWithHTTPClient(client *http.Client) *GetSubnetsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetSubnetsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetSubnetsParams contains all the parameters to send to the API endpoint
for the get subnets operation typically these are written to a http.Request
*/
type GetSubnetsParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*NetworkACLCrn
	  Filters the collection to subnets with the network ACL of the specified CRN

	*/
	NetworkACLCrn *string
	/*NetworkACLID
	  Filters the collection to subnets with the network ACL of the specified identifier

	*/
	NetworkACLID *string
	/*NetworkACLName
	  Filters the collection to subnets with the network ACL of the specified name

	*/
	NetworkACLName *string
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcCrn
	  Filters the collection to resources within the VPC of the specified CRN

	*/
	VpcCrn *string
	/*VpcID
	  Filters the collection to resources within the VPC of the specified identifier

	*/
	VpcID *string
	/*VpcName
	  Filters the collection to resources within the VPC of the specified name

	*/
	VpcName *string
	/*ZoneName
	  Filters the collection to resources within the specified zone

	*/
	ZoneName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subnets params
func (o *GetSubnetsParams) WithTimeout(timeout time.Duration) *GetSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subnets params
func (o *GetSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subnets params
func (o *GetSubnetsParams) WithContext(ctx context.Context) *GetSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subnets params
func (o *GetSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subnets params
func (o *GetSubnetsParams) WithHTTPClient(client *http.Client) *GetSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subnets params
func (o *GetSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get subnets params
func (o *GetSubnetsParams) WithGeneration(generation int64) *GetSubnetsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get subnets params
func (o *GetSubnetsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get subnets params
func (o *GetSubnetsParams) WithLimit(limit *int32) *GetSubnetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get subnets params
func (o *GetSubnetsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNetworkACLCrn adds the networkACLCrn to the get subnets params
func (o *GetSubnetsParams) WithNetworkACLCrn(networkACLCrn *string) *GetSubnetsParams {
	o.SetNetworkACLCrn(networkACLCrn)
	return o
}

// SetNetworkACLCrn adds the networkAclCrn to the get subnets params
func (o *GetSubnetsParams) SetNetworkACLCrn(networkACLCrn *string) {
	o.NetworkACLCrn = networkACLCrn
}

// WithNetworkACLID adds the networkACLID to the get subnets params
func (o *GetSubnetsParams) WithNetworkACLID(networkACLID *string) *GetSubnetsParams {
	o.SetNetworkACLID(networkACLID)
	return o
}

// SetNetworkACLID adds the networkAclId to the get subnets params
func (o *GetSubnetsParams) SetNetworkACLID(networkACLID *string) {
	o.NetworkACLID = networkACLID
}

// WithNetworkACLName adds the networkACLName to the get subnets params
func (o *GetSubnetsParams) WithNetworkACLName(networkACLName *string) *GetSubnetsParams {
	o.SetNetworkACLName(networkACLName)
	return o
}

// SetNetworkACLName adds the networkAclName to the get subnets params
func (o *GetSubnetsParams) SetNetworkACLName(networkACLName *string) {
	o.NetworkACLName = networkACLName
}

// WithResourceGroupID adds the resourceGroupID to the get subnets params
func (o *GetSubnetsParams) WithResourceGroupID(resourceGroupID *string) *GetSubnetsParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get subnets params
func (o *GetSubnetsParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get subnets params
func (o *GetSubnetsParams) WithStart(start *string) *GetSubnetsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get subnets params
func (o *GetSubnetsParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get subnets params
func (o *GetSubnetsParams) WithVersion(version string) *GetSubnetsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get subnets params
func (o *GetSubnetsParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcCrn adds the vpcCrn to the get subnets params
func (o *GetSubnetsParams) WithVpcCrn(vpcCrn *string) *GetSubnetsParams {
	o.SetVpcCrn(vpcCrn)
	return o
}

// SetVpcCrn adds the vpcCrn to the get subnets params
func (o *GetSubnetsParams) SetVpcCrn(vpcCrn *string) {
	o.VpcCrn = vpcCrn
}

// WithVpcID adds the vpcID to the get subnets params
func (o *GetSubnetsParams) WithVpcID(vpcID *string) *GetSubnetsParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get subnets params
func (o *GetSubnetsParams) SetVpcID(vpcID *string) {
	o.VpcID = vpcID
}

// WithVpcName adds the vpcName to the get subnets params
func (o *GetSubnetsParams) WithVpcName(vpcName *string) *GetSubnetsParams {
	o.SetVpcName(vpcName)
	return o
}

// SetVpcName adds the vpcName to the get subnets params
func (o *GetSubnetsParams) SetVpcName(vpcName *string) {
	o.VpcName = vpcName
}

// WithZoneName adds the zoneName to the get subnets params
func (o *GetSubnetsParams) WithZoneName(zoneName *string) *GetSubnetsParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the get subnets params
func (o *GetSubnetsParams) SetZoneName(zoneName *string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.NetworkACLCrn != nil {

		// query param network_acl.crn
		var qrNetworkACLCrn string
		if o.NetworkACLCrn != nil {
			qrNetworkACLCrn = *o.NetworkACLCrn
		}
		qNetworkACLCrn := qrNetworkACLCrn
		if qNetworkACLCrn != "" {
			if err := r.SetQueryParam("network_acl.crn", qNetworkACLCrn); err != nil {
				return err
			}
		}

	}

	if o.NetworkACLID != nil {

		// query param network_acl.id
		var qrNetworkACLID string
		if o.NetworkACLID != nil {
			qrNetworkACLID = *o.NetworkACLID
		}
		qNetworkACLID := qrNetworkACLID
		if qNetworkACLID != "" {
			if err := r.SetQueryParam("network_acl.id", qNetworkACLID); err != nil {
				return err
			}
		}

	}

	if o.NetworkACLName != nil {

		// query param network_acl.name
		var qrNetworkACLName string
		if o.NetworkACLName != nil {
			qrNetworkACLName = *o.NetworkACLName
		}
		qNetworkACLName := qrNetworkACLName
		if qNetworkACLName != "" {
			if err := r.SetQueryParam("network_acl.name", qNetworkACLName); err != nil {
				return err
			}
		}

	}

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
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

	if o.VpcCrn != nil {

		// query param vpc.crn
		var qrVpcCrn string
		if o.VpcCrn != nil {
			qrVpcCrn = *o.VpcCrn
		}
		qVpcCrn := qrVpcCrn
		if qVpcCrn != "" {
			if err := r.SetQueryParam("vpc.crn", qVpcCrn); err != nil {
				return err
			}
		}

	}

	if o.VpcID != nil {

		// query param vpc.id
		var qrVpcID string
		if o.VpcID != nil {
			qrVpcID = *o.VpcID
		}
		qVpcID := qrVpcID
		if qVpcID != "" {
			if err := r.SetQueryParam("vpc.id", qVpcID); err != nil {
				return err
			}
		}

	}

	if o.VpcName != nil {

		// query param vpc.name
		var qrVpcName string
		if o.VpcName != nil {
			qrVpcName = *o.VpcName
		}
		qVpcName := qrVpcName
		if qVpcName != "" {
			if err := r.SetQueryParam("vpc.name", qVpcName); err != nil {
				return err
			}
		}

	}

	if o.ZoneName != nil {

		// query param zone.name
		var qrZoneName string
		if o.ZoneName != nil {
			qrZoneName = *o.ZoneName
		}
		qZoneName := qrZoneName
		if qZoneName != "" {
			if err := r.SetQueryParam("zone.name", qZoneName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
