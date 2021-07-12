/*
 * Datadog API V1 Collection
 *
 * Collection of all Datadog Public endpoints.
 *
 * API version: 1.0
 * Contact: support@datadoghq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// FastlyServiceId Fastly service ID.
type FastlyServiceId struct {
	// Your service ID from Fastly.
	ServiceId string `json:"service_id"`
}

// NewFastlyServiceId instantiates a new FastlyServiceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFastlyServiceId(serviceId string) *FastlyServiceId {
	this := FastlyServiceId{}
	this.ServiceId = serviceId
	return &this
}

// NewFastlyServiceIdWithDefaults instantiates a new FastlyServiceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFastlyServiceIdWithDefaults() *FastlyServiceId {
	this := FastlyServiceId{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *FastlyServiceId) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *FastlyServiceId) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *FastlyServiceId) SetServiceId(v string) {
	o.ServiceId = v
}

func (o FastlyServiceId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service_id"] = o.ServiceId
	}
	return json.Marshal(toSerialize)
}

type NullableFastlyServiceId struct {
	value *FastlyServiceId
	isSet bool
}

func (v NullableFastlyServiceId) Get() *FastlyServiceId {
	return v.value
}

func (v *NullableFastlyServiceId) Set(val *FastlyServiceId) {
	v.value = val
	v.isSet = true
}

func (v NullableFastlyServiceId) IsSet() bool {
	return v.isSet
}

func (v *NullableFastlyServiceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFastlyServiceId(val *FastlyServiceId) *NullableFastlyServiceId {
	return &NullableFastlyServiceId{value: val, isSet: true}
}

func (v NullableFastlyServiceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFastlyServiceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
