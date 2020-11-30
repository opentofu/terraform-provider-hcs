// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest See ConsulAMAService.ListUpdateVersion
// swagger:model hashicorp.cloud.consulama.ama.ListConsulUpgradeVersionsRequest
type HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest struct {

	// resource_group is the resource group in which the Consul cluster is
	// running. This is the AMA instance's managed resource group.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// subscription_id is the ID of the Azure subscription the Consul cluster
	// exists in. This is the customer's subscription ID.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama list consul upgrade versions request
func (m *HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaListConsulUpgradeVersionsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
