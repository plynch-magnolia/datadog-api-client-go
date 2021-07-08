/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsListGlobalVariablesResponse Object containing an array of Synthetic global variables.
type SyntheticsListGlobalVariablesResponse struct {
	// Array of Synthetic global variables.
	Variables *[]SyntheticsGlobalVariable `json:"variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsListGlobalVariablesResponse instantiates a new SyntheticsListGlobalVariablesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsListGlobalVariablesResponse() *SyntheticsListGlobalVariablesResponse {
	this := SyntheticsListGlobalVariablesResponse{}
	return &this
}

// NewSyntheticsListGlobalVariablesResponseWithDefaults instantiates a new SyntheticsListGlobalVariablesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsListGlobalVariablesResponseWithDefaults() *SyntheticsListGlobalVariablesResponse {
	this := SyntheticsListGlobalVariablesResponse{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsListGlobalVariablesResponse) GetVariables() []SyntheticsGlobalVariable {
	if o == nil || o.Variables == nil {
		var ret []SyntheticsGlobalVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsListGlobalVariablesResponse) GetVariablesOk() (*[]SyntheticsGlobalVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsListGlobalVariablesResponse) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []SyntheticsGlobalVariable and assigns it to the Variables field.
func (o *SyntheticsListGlobalVariablesResponse) SetVariables(v []SyntheticsGlobalVariable) {
	o.Variables = &v
}

func (o SyntheticsListGlobalVariablesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsListGlobalVariablesResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Variables *[]SyntheticsGlobalVariable `json:"variables,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	o.Variables = all.Variables
	return nil
}

type NullableSyntheticsListGlobalVariablesResponse struct {
	value *SyntheticsListGlobalVariablesResponse
	isSet bool
}

func (v NullableSyntheticsListGlobalVariablesResponse) Get() *SyntheticsListGlobalVariablesResponse {
	return v.value
}

func (v *NullableSyntheticsListGlobalVariablesResponse) Set(val *SyntheticsListGlobalVariablesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsListGlobalVariablesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsListGlobalVariablesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsListGlobalVariablesResponse(val *SyntheticsListGlobalVariablesResponse) *NullableSyntheticsListGlobalVariablesResponse {
	return &NullableSyntheticsListGlobalVariablesResponse{value: val, isSet: true}
}

func (v NullableSyntheticsListGlobalVariablesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsListGlobalVariablesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
