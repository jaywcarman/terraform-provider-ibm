// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized.
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams() *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithTimeout creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		timeout: timeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithContext creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		Context: ctx,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithHTTPClient creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{
		HTTPClient: client,
	}
}

/*GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams contains all the parameters to send to the API endpoint
for the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length operation typically these are written to a http.Request
*/
type GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The VPN connection identifier

	*/
	ID string
	/*PrefixAddress
	  The prefix address part of the CIDR

	*/
	PrefixAddress string
	/*PrefixLength
	  The prefix length part of the CIDR

	*/
	PrefixLength string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpnGatewayID
	  The VPN gateway identifier

	*/
	VpnGatewayID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithGeneration(generation int64) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithID(id string) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetID(id string) {
	o.ID = id
}

// WithPrefixAddress adds the prefixAddress to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithPrefixAddress(prefixAddress string) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixAddress(prefixAddress)
	return o
}

// SetPrefixAddress adds the prefixAddress to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetPrefixAddress(prefixAddress string) {
	o.PrefixAddress = prefixAddress
}

// WithPrefixLength adds the prefixLength to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithPrefixLength(prefixLength string) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixLength(prefixLength)
	return o
}

// SetPrefixLength adds the prefixLength to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetPrefixLength(prefixLength string) {
	o.PrefixLength = prefixLength
}

// WithVersion adds the version to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithVersion(version string) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithVpnGatewayID(vpnGatewayID string) *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the get vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param prefix_address
	if err := r.SetPathParam("prefix_address", o.PrefixAddress); err != nil {
		return err
	}

	// path param prefix_length
	if err := r.SetPathParam("prefix_length", o.PrefixLength); err != nil {
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

	// path param vpn_gateway_id
	if err := r.SetPathParam("vpn_gateway_id", o.VpnGatewayID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
