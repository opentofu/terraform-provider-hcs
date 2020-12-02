// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaResourceNotificationRequestPlan Plan represents an Azure Marketplace plan.
// swagger:model hashicorp.cloud.consulama.ama.ResourceNotificationRequest.Plan
type HashicorpCloudConsulamaAmaResourceNotificationRequestPlan struct {

	// name is the name of the Marketplace plan.
	Name string `json:"name,omitempty"`

	// product is the name of the Marketplace offer.
	Product string `json:"product,omitempty"`

	// publisher is the plan's publisher organization.
	Publisher string `json:"publisher,omitempty"`

	// version is the version of the Marketplace plan.
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama resource notification request plan
func (m *HashicorpCloudConsulamaAmaResourceNotificationRequestPlan) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaResourceNotificationRequestPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaResourceNotificationRequestPlan) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaResourceNotificationRequestPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}