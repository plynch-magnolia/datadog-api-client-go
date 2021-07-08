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

// IncidentServicesResponseMetaPagination Pagination properties.
type IncidentServicesResponseMetaPagination struct {
	// The index of the first element in the next page of results. Equal to page size added to the current offset.
	NextOffset *int64 `json:"next_offset,omitempty"`
	// The index of the first element in the results.
	Offset *int64 `json:"offset,omitempty"`
	// Maximum size of pages to return.
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentServicesResponseMetaPagination instantiates a new IncidentServicesResponseMetaPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentServicesResponseMetaPagination() *IncidentServicesResponseMetaPagination {
	this := IncidentServicesResponseMetaPagination{}
	return &this
}

// NewIncidentServicesResponseMetaPaginationWithDefaults instantiates a new IncidentServicesResponseMetaPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentServicesResponseMetaPaginationWithDefaults() *IncidentServicesResponseMetaPagination {
	this := IncidentServicesResponseMetaPagination{}
	return &this
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *IncidentServicesResponseMetaPagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponseMetaPagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil || o.NextOffset == nil {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *IncidentServicesResponseMetaPagination) HasNextOffset() bool {
	if o != nil && o.NextOffset != nil {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given int64 and assigns it to the NextOffset field.
func (o *IncidentServicesResponseMetaPagination) SetNextOffset(v int64) {
	o.NextOffset = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *IncidentServicesResponseMetaPagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponseMetaPagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *IncidentServicesResponseMetaPagination) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *IncidentServicesResponseMetaPagination) SetOffset(v int64) {
	o.Offset = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *IncidentServicesResponseMetaPagination) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentServicesResponseMetaPagination) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *IncidentServicesResponseMetaPagination) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *IncidentServicesResponseMetaPagination) SetSize(v int64) {
	o.Size = &v
}

func (o IncidentServicesResponseMetaPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.NextOffset != nil {
		toSerialize["next_offset"] = o.NextOffset
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentServicesResponseMetaPagination) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		NextOffset *int64 `json:"next_offset,omitempty"`
		Offset     *int64 `json:"offset,omitempty"`
		Size       *int64 `json:"size,omitempty"`
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
	o.NextOffset = all.NextOffset
	o.Offset = all.Offset
	o.Size = all.Size
	return nil
}

type NullableIncidentServicesResponseMetaPagination struct {
	value *IncidentServicesResponseMetaPagination
	isSet bool
}

func (v NullableIncidentServicesResponseMetaPagination) Get() *IncidentServicesResponseMetaPagination {
	return v.value
}

func (v *NullableIncidentServicesResponseMetaPagination) Set(val *IncidentServicesResponseMetaPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentServicesResponseMetaPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentServicesResponseMetaPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentServicesResponseMetaPagination(val *IncidentServicesResponseMetaPagination) *NullableIncidentServicesResponseMetaPagination {
	return &NullableIncidentServicesResponseMetaPagination{value: val, isSet: true}
}

func (v NullableIncidentServicesResponseMetaPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentServicesResponseMetaPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
