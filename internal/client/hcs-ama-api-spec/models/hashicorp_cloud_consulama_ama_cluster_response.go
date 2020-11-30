// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaClusterResponse See ConsulAMAService.GetCluster, ConsulAMAService.CreateCluster.
// swagger:model hashicorp.cloud.consulama.ama.ClusterResponse
type HashicorpCloudConsulamaAmaClusterResponse struct {

	// id is the complete resource ID.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name is the resource name as given by Azure in the GetClusterRequest.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// properties are the properties of the cluster.
	Properties *HashicorpCloudConsulamaAmaClusterProperties `json:"properties,omitempty"`

	// type is the resource type as given by Azure in the GetClusterRequest.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama cluster response
func (m *HashicorpCloudConsulamaAmaClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaClusterResponse) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaClusterResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
