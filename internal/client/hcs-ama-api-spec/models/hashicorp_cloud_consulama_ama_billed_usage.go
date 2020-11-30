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

// HashicorpCloudConsulamaAmaBilledUsage BilledUsage describes the usage that was recorded and billed for a specific
// time frame.
// swagger:model hashicorp.cloud.consulama.ama.BilledUsage
type HashicorpCloudConsulamaAmaBilledUsage struct {

	// cost is the total cost billed for the time frame.
	Cost float64 `json:"cost,omitempty"`

	// details contains usage by plan. Because there can be multiple settings
	// applied to one time frame, this needs to be a list.
	Details []*HashicorpCloudConsulamaAmaBilledUsagePlanUsage `json:"details"`

	// end is the end time of the time frame.
	End string `json:"end,omitempty"`

	// start is the start time of the time frame.
	Start string `json:"start,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama billed usage
func (m *HashicorpCloudConsulamaAmaBilledUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaBilledUsage) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {
		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {
			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBilledUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaBilledUsage) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaBilledUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
