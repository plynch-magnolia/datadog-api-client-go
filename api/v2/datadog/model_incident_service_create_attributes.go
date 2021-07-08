/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// IncidentServiceCreateAttributes The incident service's attributes for a create request.
type IncidentServiceCreateAttributes struct {
	// Name of the incident service.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentServiceCreateAttributes instantiates a new IncidentServiceCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServiceCreateAttributes(name string) *IncidentServiceCreateAttributes {
	this := IncidentServiceCreateAttributes{}
	this.Name = name
	return &this
}

// NewIncidentServiceCreateAttributesWithDefaults instantiates a new IncidentServiceCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServiceCreateAttributesWithDefaults() *IncidentServiceCreateAttributes {
	this := IncidentServiceCreateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *IncidentServiceCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentServiceCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IncidentServiceCreateAttributes) SetName(v string) {
	o.Name = v
}

func (o IncidentServiceCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentServiceCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Name *string `json:"name"`
	}{}
	all := struct {
		Name string `json:"name"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		o.UnparsedObject = raw
	}
	if _, ok := o.UnparsedObject["name"]; required.Name == nil && !ok {
		return fmt.Errorf("Required field name missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	o.Name = all.Name
	return nil
}

type NullableIncidentServiceCreateAttributes struct {
	value *IncidentServiceCreateAttributes
	isSet bool
}

func (v NullableIncidentServiceCreateAttributes) Get() *IncidentServiceCreateAttributes {
	return v.value
}

func (v *NullableIncidentServiceCreateAttributes) Set(val *IncidentServiceCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServiceCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServiceCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServiceCreateAttributes(val *IncidentServiceCreateAttributes) *NullableIncidentServiceCreateAttributes {
	return &NullableIncidentServiceCreateAttributes{value: val, isSet: true}
}

func (v NullableIncidentServiceCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServiceCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
