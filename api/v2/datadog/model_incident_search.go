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

// IncidentSearch Represents an Incident Search.
type IncidentSearch struct {
	Attributes *IncidentSearchAttributes `json:"attributes,omitempty"`
	// JSONAPI Model Type
	Type *string `json:"type,omitempty"`
}

// NewIncidentSearch instantiates a new IncidentSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentSearch() *IncidentSearch {
	this := IncidentSearch{}
	var type_ string = "incidents_search_results"
	this.Type = &type_
	return &this
}

// NewIncidentSearchWithDefaults instantiates a new IncidentSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentSearchWithDefaults() *IncidentSearch {
	this := IncidentSearch{}
	var type_ string = "incidents_search_results"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentSearch) GetAttributes() IncidentSearchAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentSearchAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearch) GetAttributesOk() (*IncidentSearchAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentSearch) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given IncidentSearchAttributes and assigns it to the Attributes field.
func (o *IncidentSearch) SetAttributes(v IncidentSearchAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentSearch) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearch) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentSearch) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IncidentSearch) SetType(v string) {
	o.Type = &v
}

func (o IncidentSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentSearch struct {
	value *IncidentSearch
	isSet bool
}

func (v NullableIncidentSearch) Get() *IncidentSearch {
	return v.value
}

func (v *NullableIncidentSearch) Set(val *IncidentSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentSearch(val *IncidentSearch) *NullableIncidentSearch {
	return &NullableIncidentSearch{value: val, isSet: true}
}

func (v NullableIncidentSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
