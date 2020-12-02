// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// HashicorpCloudConsulamaAmaClusterMode ClusterMode defines what type of cluster is going to be deployed.
// Some modes will have different billing and options than others.
// PRODUCTION is the default value for this property, that's why it's
// set to be value 0 of the enum.
// swagger:model hashicorp.cloud.consulama.ama.ClusterMode
type HashicorpCloudConsulamaAmaClusterMode string

const (

	// HashicorpCloudConsulamaAmaClusterModePRODUCTION captures enum value "PRODUCTION"
	HashicorpCloudConsulamaAmaClusterModePRODUCTION HashicorpCloudConsulamaAmaClusterMode = "PRODUCTION"

	// HashicorpCloudConsulamaAmaClusterModeDEVELOPMENT captures enum value "DEVELOPMENT"
	HashicorpCloudConsulamaAmaClusterModeDEVELOPMENT HashicorpCloudConsulamaAmaClusterMode = "DEVELOPMENT"
)

// for schema
var hashicorpCloudConsulamaAmaClusterModeEnum []interface{}

func init() {
	var res []HashicorpCloudConsulamaAmaClusterMode
	if err := json.Unmarshal([]byte(`["PRODUCTION","DEVELOPMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudConsulamaAmaClusterModeEnum = append(hashicorpCloudConsulamaAmaClusterModeEnum, v)
	}
}

func (m HashicorpCloudConsulamaAmaClusterMode) validateHashicorpCloudConsulamaAmaClusterModeEnum(path, location string, value HashicorpCloudConsulamaAmaClusterMode) error {
	if err := validate.Enum(path, location, value, hashicorpCloudConsulamaAmaClusterModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud consulama ama cluster mode
func (m HashicorpCloudConsulamaAmaClusterMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudConsulamaAmaClusterModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}