// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaUpdateClusterRequest See ConsulAMAService.UpdateCluster
// swagger:model hashicorp.cloud.consulama.ama.UpdateClusterRequest
type HashicorpCloudConsulamaAmaUpdateClusterRequest struct {

	// resource_group is the resource group in which the Consul cluster is
	// running. This is the AMA instance's managed resource group.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// subscription_id is the ID of the Azure subscription the Consul cluster
	// exists in. This is the customer's subscription ID.
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// update contains the details of the Consul cluster to be updated.
	Update *HashicorpCloudConsulamaAmaClusterUpdate `json:"update,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama update cluster request
func (m *HashicorpCloudConsulamaAmaUpdateClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaUpdateClusterRequest) validateUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.Update) { // not required
		return nil
	}

	if m.Update != nil {
		if err := m.Update.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaUpdateClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaUpdateClusterRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaUpdateClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
