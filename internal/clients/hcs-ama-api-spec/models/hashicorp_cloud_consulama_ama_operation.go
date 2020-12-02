// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HashicorpCloudConsulamaAmaOperation Operation represents a single operation.
// swagger:model hashicorp.cloud.consulama.ama.Operation
type HashicorpCloudConsulamaAmaOperation struct {

	// error is the error that occurred in the operation.
	Error *GoogleRPCStatus `json:"error,omitempty"`

	// id is the unique ID for this operation. This is only unique within the
	// context of a single managed app.
	ID string `json:"id,omitempty"`

	// response is the result of the operation. See the documentation for the API
	// call creating the operation to understand what the value of this is.
	Response *GoogleProtobufAny `json:"response,omitempty"`

	// state is the current state of the operation. This is a simple tri-state:
	// PENDING means the operation is created but not yet started, RUNNING means
	// the operation is currently running (though it may be very long-running),
	// and DONE means the operation is complete whether successfully or not.
	State HashicorpCloudConsulamaAmaOperationState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud consulama ama operation
func (m *HashicorpCloudConsulamaAmaOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudConsulamaAmaOperation) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulamaAmaOperation) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudConsulamaAmaOperation) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudConsulamaAmaOperation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudConsulamaAmaOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}