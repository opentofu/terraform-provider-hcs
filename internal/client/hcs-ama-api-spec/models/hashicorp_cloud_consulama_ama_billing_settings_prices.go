// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaBillingSettingsPrices Prices represents prices of a billing plan.
// swagger:model hashicorp.cloud.consulama.ama.BillingSettings.Prices
type HashicorpCloudConsulamaAmaBillingSettingsPrices struct {

	// dimensions contains all dimensions and their prices that are applied
	// to apps on this billing plan.
	Dimensions []*HashicorpCloudConsulamaAmaBillingSettingsPricesPriceDimension `json:"dimensions"`

	// hourly is the flat hourly fee billed for the usage of this plan
	// independent of other billing dimensions.
	Hourly float64 `json:"hourly,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama billing settings prices
func (m *HashicorpCloudConsulamaAmaBillingSettingsPrices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaBillingSettingsPrices) validateDimensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Dimensions); i++ {
		if swag.IsZero(m.Dimensions[i]) { // not required
			continue
		}

		if m.Dimensions[i] != nil {
			if err := m.Dimensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dimensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBillingSettingsPrices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBillingSettingsPrices) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaBillingSettingsPrices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
