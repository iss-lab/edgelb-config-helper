// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V2CloudProviderAwsElb An AWS ELB configuration.
//
// swagger:model V2CloudProviderAwsElb
type V2CloudProviderAwsElb struct {

	// Location of the access logs of a load balancer, with format `S3Bucket/prefix`, if you don’t specify a prefix, the access logs are stored in the root of the bucket. Access logs are created only if the load balancer has a TLS listener and they contain information only about TLS requests.
	AccessLogS3Location *string `json:"accessLogS3Location,omitempty"`

	// ARN of the ELB to use. If it is specified, it is assumed that there is a pre-created ELB that Edge-LB is not supposed to create or delete.
	// Min Length: 1
	Arn string `json:"arn,omitempty"`

	// If true, the load balancer will not be deleted when the corresponding pool is deleted.
	DeletionProtection bool `json:"deletionProtection,omitempty"`

	// Elastic IP addresses (IPv4 addresses and allocation IDs) that should be associated with subnets, the ELB is attached to. For instance, ["1.1.1.1", "eip-12345678"]
	ElasticIPs []string `json:"elasticIPs"`

	// If true (default), a load balancer is created as internal; otherwise - as internet-facing.
	Internal *bool `json:"internal,omitempty"`

	// listeners
	// Required: true
	// Min Items: 1
	Listeners []*V2CloudProviderAwsElbListener `json:"listeners"`

	// Name of the ELB which should be unique within an Edge-LB pool. The actual ELB name would be derived from this name and would be of form dcos-lb-hash(cluster-id)-hash(edgelb-app-id)-hash(pool-name)-hash(elb-name).
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// IDs of subnets to attach the ELB to. If it is empty, subnets are autodiscovered based on Edge-LB load balancers registered with apiserver.
	Subnets []string `json:"subnets"`

	// tags
	// Max Items: 30
	// Min Items: 1
	Tags []*V2CloudProviderAwsTag `json:"tags"`

	// An AWS ELB type.
	// Required: true
	// Enum: [NLB]
	Type *string `json:"type"`
}

// Validate validates this v2 cloud provider aws elb
func (m *V2CloudProviderAwsElb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListeners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2CloudProviderAwsElb) validateArn(formats strfmt.Registry) error {

	if swag.IsZero(m.Arn) { // not required
		return nil
	}

	if err := validate.MinLength("arn", "body", string(m.Arn), 1); err != nil {
		return err
	}

	return nil
}

func (m *V2CloudProviderAwsElb) validateElasticIPs(formats strfmt.Registry) error {

	if swag.IsZero(m.ElasticIPs) { // not required
		return nil
	}

	for i := 0; i < len(m.ElasticIPs); i++ {

		if err := validate.MinLength("elasticIPs"+"."+strconv.Itoa(i), "body", string(m.ElasticIPs[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V2CloudProviderAwsElb) validateListeners(formats strfmt.Registry) error {

	if err := validate.Required("listeners", "body", m.Listeners); err != nil {
		return err
	}

	iListenersSize := int64(len(m.Listeners))

	if err := validate.MinItems("listeners", "body", iListenersSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Listeners); i++ {
		if swag.IsZero(m.Listeners[i]) { // not required
			continue
		}

		if m.Listeners[i] != nil {
			if err := m.Listeners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listeners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V2CloudProviderAwsElb) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *V2CloudProviderAwsElb) validateSubnets(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {

		if err := validate.MinLength("subnets"+"."+strconv.Itoa(i), "body", string(m.Subnets[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V2CloudProviderAwsElb) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	iTagsSize := int64(len(m.Tags))

	if err := validate.MinItems("tags", "body", iTagsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("tags", "body", iTagsSize, 30); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v2CloudProviderAwsElbTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NLB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2CloudProviderAwsElbTypeTypePropEnum = append(v2CloudProviderAwsElbTypeTypePropEnum, v)
	}
}

const (

	// V2CloudProviderAwsElbTypeNLB captures enum value "NLB"
	V2CloudProviderAwsElbTypeNLB string = "NLB"
)

// prop value enum
func (m *V2CloudProviderAwsElb) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v2CloudProviderAwsElbTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V2CloudProviderAwsElb) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2CloudProviderAwsElb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2CloudProviderAwsElb) UnmarshalBinary(b []byte) error {
	var res V2CloudProviderAwsElb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
