// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaRestoreSnapshotResponse See ConsulAMAService.RestoreSnapshot
// swagger:model hashicorp.cloud.consulama.ama.RestoreSnapshotResponse
type HashicorpCloudConsulamaAmaRestoreSnapshotResponse struct {

	// operation used to track the progress of the snapshot restore
	Operation *HashicorpCloudConsulamaAmaOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama restore snapshot response
func (m *HashicorpCloudConsulamaAmaRestoreSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaRestoreSnapshotResponse) validateOperation(formats strfmt.Registry) error {

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
func (m *HashicorpCloudConsulamaAmaRestoreSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaRestoreSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaRestoreSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}