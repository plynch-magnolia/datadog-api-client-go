/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageCloudSecurityPostureManagementHour Cloud Security Posture Management usage for a given organization for a given hour.
type UsageCloudSecurityPostureManagementHour struct {
	// The total number of Cloud Security Posture Management containers during a given hour.
	ContainerCount *int64 `json:"container_count,omitempty"`
	// The total number of Cloud Security Posture Management hosts during a given hour.
	HostCount *int64 `json:"host_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewUsageCloudSecurityPostureManagementHour instantiates a new UsageCloudSecurityPostureManagementHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageCloudSecurityPostureManagementHour() *UsageCloudSecurityPostureManagementHour {
	this := UsageCloudSecurityPostureManagementHour{}
	return &this
}

// NewUsageCloudSecurityPostureManagementHourWithDefaults instantiates a new UsageCloudSecurityPostureManagementHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageCloudSecurityPostureManagementHourWithDefaults() *UsageCloudSecurityPostureManagementHour {
	this := UsageCloudSecurityPostureManagementHour{}
	return &this
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise.
func (o *UsageCloudSecurityPostureManagementHour) GetContainerCount() int64 {
	if o == nil || o.ContainerCount == nil {
		var ret int64
		return ret
	}
	return *o.ContainerCount
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCloudSecurityPostureManagementHour) GetContainerCountOk() (*int64, bool) {
	if o == nil || o.ContainerCount == nil {
		return nil, false
	}
	return o.ContainerCount, true
}

// HasContainerCount returns a boolean if a field has been set.
func (o *UsageCloudSecurityPostureManagementHour) HasContainerCount() bool {
	if o != nil && o.ContainerCount != nil {
		return true
	}

	return false
}

// SetContainerCount gets a reference to the given int64 and assigns it to the ContainerCount field.
func (o *UsageCloudSecurityPostureManagementHour) SetContainerCount(v int64) {
	o.ContainerCount = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *UsageCloudSecurityPostureManagementHour) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCloudSecurityPostureManagementHour) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *UsageCloudSecurityPostureManagementHour) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *UsageCloudSecurityPostureManagementHour) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageCloudSecurityPostureManagementHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCloudSecurityPostureManagementHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageCloudSecurityPostureManagementHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageCloudSecurityPostureManagementHour) SetHour(v time.Time) {
	o.Hour = &v
}

func (o UsageCloudSecurityPostureManagementHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ContainerCount != nil {
		toSerialize["container_count"] = o.ContainerCount
	}
	if o.HostCount != nil {
		toSerialize["host_count"] = o.HostCount
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	return json.Marshal(toSerialize)
}

func (o *UsageCloudSecurityPostureManagementHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ContainerCount *int64     `json:"container_count,omitempty"`
		HostCount      *int64     `json:"host_count,omitempty"`
		Hour           *time.Time `json:"hour,omitempty"`
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
	o.ContainerCount = all.ContainerCount
	o.HostCount = all.HostCount
	o.Hour = all.Hour
	return nil
}

type NullableUsageCloudSecurityPostureManagementHour struct {
	value *UsageCloudSecurityPostureManagementHour
	isSet bool
}

func (v NullableUsageCloudSecurityPostureManagementHour) Get() *UsageCloudSecurityPostureManagementHour {
	return v.value
}

func (v *NullableUsageCloudSecurityPostureManagementHour) Set(val *UsageCloudSecurityPostureManagementHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageCloudSecurityPostureManagementHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageCloudSecurityPostureManagementHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageCloudSecurityPostureManagementHour(val *UsageCloudSecurityPostureManagementHour) *NullableUsageCloudSecurityPostureManagementHour {
	return &NullableUsageCloudSecurityPostureManagementHour{value: val, isSet: true}
}

func (v NullableUsageCloudSecurityPostureManagementHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageCloudSecurityPostureManagementHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
