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

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized.
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams() *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithTimeout creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		timeout: timeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithContext creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		Context: ctx,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithHTTPClient creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{
		HTTPClient: client,
	}
}

/*GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams contains all the parameters to send to the API endpoint
for the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length operation typically these are written to a http.Request
*/
type GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams struct {

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

// WithTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithGeneration(generation int64) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithID(id string) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetID(id string) {
	o.ID = id
}

// WithPrefixAddress adds the prefixAddress to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithPrefixAddress(prefixAddress string) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixAddress(prefixAddress)
	return o
}

// SetPrefixAddress adds the prefixAddress to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetPrefixAddress(prefixAddress string) {
	o.PrefixAddress = prefixAddress
}

// WithPrefixLength adds the prefixLength to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithPrefixLength(prefixLength string) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixLength(prefixLength)
	return o
}

// SetPrefixLength adds the prefixLength to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetPrefixLength(prefixLength string) {
	o.PrefixLength = prefixLength
}

// WithVersion adds the version to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithVersion(version string) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithVpnGatewayID(vpnGatewayID string) *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the get vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
