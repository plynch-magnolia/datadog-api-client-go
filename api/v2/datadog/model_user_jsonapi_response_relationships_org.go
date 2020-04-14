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

// UserJSONAPIResponseRelationshipsOrg The Organization the user is associated with.
type UserJSONAPIResponseRelationshipsOrg struct {
	Data *OrganizationJSONAPIID `json:"data,omitempty"`
}

// NewUserJSONAPIResponseRelationshipsOrg instantiates a new UserJSONAPIResponseRelationshipsOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserJSONAPIResponseRelationshipsOrg() *UserJSONAPIResponseRelationshipsOrg {
	this := UserJSONAPIResponseRelationshipsOrg{}
	return &this
}

// NewUserJSONAPIResponseRelationshipsOrgWithDefaults instantiates a new UserJSONAPIResponseRelationshipsOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserJSONAPIResponseRelationshipsOrgWithDefaults() *UserJSONAPIResponseRelationshipsOrg {
	this := UserJSONAPIResponseRelationshipsOrg{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserJSONAPIResponseRelationshipsOrg) GetData() OrganizationJSONAPIID {
	if o == nil || o.Data == nil {
		var ret OrganizationJSONAPIID
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJSONAPIResponseRelationshipsOrg) GetDataOk() (*OrganizationJSONAPIID, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserJSONAPIResponseRelationshipsOrg) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given OrganizationJSONAPIID and assigns it to the Data field.
func (o *UserJSONAPIResponseRelationshipsOrg) SetData(v OrganizationJSONAPIID) {
	o.Data = &v
}

func (o UserJSONAPIResponseRelationshipsOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUserJSONAPIResponseRelationshipsOrg struct {
	value *UserJSONAPIResponseRelationshipsOrg
	isSet bool
}

func (v NullableUserJSONAPIResponseRelationshipsOrg) Get() *UserJSONAPIResponseRelationshipsOrg {
	return v.value
}

func (v *NullableUserJSONAPIResponseRelationshipsOrg) Set(val *UserJSONAPIResponseRelationshipsOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableUserJSONAPIResponseRelationshipsOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableUserJSONAPIResponseRelationshipsOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserJSONAPIResponseRelationshipsOrg(val *UserJSONAPIResponseRelationshipsOrg) *NullableUserJSONAPIResponseRelationshipsOrg {
	return &NullableUserJSONAPIResponseRelationshipsOrg{value: val, isSet: true}
}

func (v NullableUserJSONAPIResponseRelationshipsOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserJSONAPIResponseRelationshipsOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
