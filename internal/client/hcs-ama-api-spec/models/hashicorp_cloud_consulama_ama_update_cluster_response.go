// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaUpdateClusterResponse See ConsulAMAService.UpdateCluster
// swagger:model hashicorp.cloud.consulama.ama.UpdateClusterResponse
type HashicorpCloudConsulamaAmaUpdateClusterResponse struct {

	// operation used to track the progress of the cluster update
	Operation *HashicorpCloudConsulamaAmaOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama update cluster response
func (m *HashicorpCloudConsulamaAmaUpdateClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaUpdateClusterResponse) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaUpdateClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaUpdateClusterResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaUpdateClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
