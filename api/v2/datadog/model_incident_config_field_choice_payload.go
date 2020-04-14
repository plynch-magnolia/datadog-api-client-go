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

// IncidentConfigFieldChoicePayload Represents the JSON API Payload of an Incident Config Field Choice Item.
type IncidentConfigFieldChoicePayload struct {
	Data IncidentConfigFieldChoice `json:"data"`
	// The User relationships.
	Included *[]UserJSONAPIResponse `json:"included,omitempty"`
}

// NewIncidentConfigFieldChoicePayload instantiates a new IncidentConfigFieldChoicePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentConfigFieldChoicePayload(data IncidentConfigFieldChoice) *IncidentConfigFieldChoicePayload {
	this := IncidentConfigFieldChoicePayload{}
	this.Data = data
	return &this
}

// NewIncidentConfigFieldChoicePayloadWithDefaults instantiates a new IncidentConfigFieldChoicePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentConfigFieldChoicePayloadWithDefaults() *IncidentConfigFieldChoicePayload {
	this := IncidentConfigFieldChoicePayload{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentConfigFieldChoicePayload) GetData() IncidentConfigFieldChoice {
	if o == nil {
		var ret IncidentConfigFieldChoice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldChoicePayload) GetDataOk() (*IncidentConfigFieldChoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentConfigFieldChoicePayload) SetData(v IncidentConfigFieldChoice) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentConfigFieldChoicePayload) GetIncluded() []UserJSONAPIResponse {
	if o == nil || o.Included == nil {
		var ret []UserJSONAPIResponse
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigFieldChoicePayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentConfigFieldChoicePayload) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []UserJSONAPIResponse and assigns it to the Included field.
func (o *IncidentConfigFieldChoicePayload) SetIncluded(v []UserJSONAPIResponse) {
	o.Included = &v
}

func (o IncidentConfigFieldChoicePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentConfigFieldChoicePayload struct {
	value *IncidentConfigFieldChoicePayload
	isSet bool
}

func (v NullableIncidentConfigFieldChoicePayload) Get() *IncidentConfigFieldChoicePayload {
	return v.value
}

func (v *NullableIncidentConfigFieldChoicePayload) Set(val *IncidentConfigFieldChoicePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentConfigFieldChoicePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentConfigFieldChoicePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentConfigFieldChoicePayload(val *IncidentConfigFieldChoicePayload) *NullableIncidentConfigFieldChoicePayload {
	return &NullableIncidentConfigFieldChoicePayload{value: val, isSet: true}
}

func (v NullableIncidentConfigFieldChoicePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentConfigFieldChoicePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
