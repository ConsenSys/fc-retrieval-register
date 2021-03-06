// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProviderRegister Register entry
//
// swagger:model ProviderRegister
type ProviderRegister struct {

	// Filecoin Account to be used with payment channels.
	Address string `json:"address,omitempty"`

	// Admin network addressing information.
	NetworkInfoAdmin string `json:"networkInfoAdmin,omitempty"`

	// Client network addressing information.
	NetworkInfoClient string `json:"networkInfoClient,omitempty"`

	// Gateway network addressing information.
	NetworkInfoGateway string `json:"networkInfoGateway,omitempty"`

	// Node ID.
	NodeID string `json:"nodeId,omitempty"`

	// Region Code.
	RegionCode string `json:"regionCode,omitempty"`

	// Retrieval provider Root Signing Public Key.
	RootSigningKey string `json:"rootSigningKey,omitempty"`

	// Used for signing CID Group Offers and Single CID Offers.
	SigningKey string `json:"signingKey,omitempty"`
}

// Validate validates this provider register
func (m *ProviderRegister) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this provider register based on context it is used
func (m *ProviderRegister) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderRegister) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderRegister) UnmarshalBinary(b []byte) error {
	var res ProviderRegister
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
