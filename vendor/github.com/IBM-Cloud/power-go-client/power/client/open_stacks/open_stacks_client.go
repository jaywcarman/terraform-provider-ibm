// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	"github.com/go-openapi/strfmt"
)

// New creates a new open stacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for open stacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBrokerOpenstacksGet lists all open stack instances being managed
*/
func (a *Client) ServiceBrokerOpenstacksGet(params *ServiceBrokerOpenstacksGetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerOpenstacksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerOpenstacksGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.openstacks.get",
		Method:             "GET",
		PathPattern:        "/broker/v1/openstacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerOpenstacksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBrokerOpenstacksGetOK), nil

}

/*
ServiceBrokerOpenstacksPost creates a new open stack instance to be managed
*/
func (a *Client) ServiceBrokerOpenstacksPost(params *ServiceBrokerOpenstacksPostParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBrokerOpenstacksPostOK, *ServiceBrokerOpenstacksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerOpenstacksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBroker.openstacks.post",
		Method:             "POST",
		PathPattern:        "/broker/v1/openstacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerOpenstacksPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceBrokerOpenstacksPostOK:
		return value, nil, nil
	case *ServiceBrokerOpenstacksPostCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
