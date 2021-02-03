// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GatewayRegister Register entry
//
// swagger:model GatewayRegister
type GatewayRegister struct {

	// Filecoin Account to be used with payment channels.
	Address string `json:"address,omitempty"`

	// Admin network addressing information.
	NetworkAdminInfo string `json:"networkAdminInfo,omitempty"`

	// Client network addressing information.
	NetworkClientInfo string `json:"networkClientInfo,omitempty"`

	// Gateway network addressing information.
	NetworkGatewayInfo string `json:"networkGatewayInfo,omitempty"`

	// Provider network addressing information.
	NetworkProviderInfo string `json:"networkProviderInfo,omitempty"`

	// Node ID.
	NodeID string `json:"nodeId,omitempty"`

	// Region Code.
	RegionCode string `json:"regionCode,omitempty"`

	// Retrieval provider Root Signing Public Key.
	RootSigningKey string `json:"rootSigningKey,omitempty"`

	// Used for signing CID Group Offers and Single CID Offers.
	SigingKey string `json:"sigingKey,omitempty"`
}

// Validate validates this gateway register
func (m *GatewayRegister) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gateway register based on context it is used
func (m *GatewayRegister) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GatewayRegister) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayRegister) UnmarshalBinary(b []byte) error {
	var res GatewayRegister
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}