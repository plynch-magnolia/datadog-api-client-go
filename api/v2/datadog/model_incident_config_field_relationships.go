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

// IncidentConfigFieldRelationships The User relationships.
type IncidentConfigFieldRelationships struct {
	CreatedBy      *UserJSONAPIID `json:"created_by,omitempty"`
	LastModifiedBy *UserJSONAPIID `json:"last_modified_by,omitempty"`
}

// NewIncidentConfigFieldRelationships instantiates a new IncidentConfigFieldRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentConfigFieldRelationships() *IncidentConfigFieldRelationships {
	this := IncidentConfigFieldRelationships{}
	return &this
}

// NewIncidentConfigFieldRelationshipsWithDefaults instantiates a new IncidentConfigFieldRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentConfigFieldRelationshipsWithDefaults() *IncidentConfigFieldRelationships {
	this := IncidentConfigFieldRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentConfigFieldRelationships) GetCreatedBy() UserJSONAPIID {
	if o == nil || o.CreatedBy == nil {
		var ret UserJSONAPIID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldRelationships) GetCreatedByOk() (*UserJSONAPIID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentConfigFieldRelationships) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserJSONAPIID and assigns it to the CreatedBy field.
func (o *IncidentConfigFieldRelationships) SetCreatedBy(v UserJSONAPIID) {
	o.CreatedBy = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentConfigFieldRelationships) GetLastModifiedBy() UserJSONAPIID {
	if o == nil || o.LastModifiedBy == nil {
		var ret UserJSONAPIID
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldRelationships) GetLastModifiedByOk() (*UserJSONAPIID, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentConfigFieldRelationships) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given UserJSONAPIID and assigns it to the LastModifiedBy field.
func (o *IncidentConfigFieldRelationships) SetLastModifiedBy(v UserJSONAPIID) {
	o.LastModifiedBy = &v
}

func (o IncidentConfigFieldRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LastModifiedBy != nil {
		toSerialize["last_modified_by"] = o.LastModifiedBy
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentConfigFieldRelationships struct {
	value *IncidentConfigFieldRelationships
	isSet bool
}

func (v NullableIncidentConfigFieldRelationships) Get() *IncidentConfigFieldRelationships {
	return v.value
}

func (v *NullableIncidentConfigFieldRelationships) Set(val *IncidentConfigFieldRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentConfigFieldRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentConfigFieldRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentConfigFieldRelationships(val *IncidentConfigFieldRelationships) *NullableIncidentConfigFieldRelationships {
	return &NullableIncidentConfigFieldRelationships{value: val, isSet: true}
}

func (v NullableIncidentConfigFieldRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentConfigFieldRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
