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

// LogsMetricCreateRequest The new log-based metric body.
type LogsMetricCreateRequest struct {
	Data LogsMetricCreateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewLogsMetricCreateRequest instantiates a new LogsMetricCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsMetricCreateRequest(data LogsMetricCreateData) *LogsMetricCreateRequest {
	this := LogsMetricCreateRequest{}
	this.Data = data
	return &this
}

// NewLogsMetricCreateRequestWithDefaults instantiates a new LogsMetricCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsMetricCreateRequestWithDefaults() *LogsMetricCreateRequest {
	this := LogsMetricCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *LogsMetricCreateRequest) GetData() LogsMetricCreateData {
	if o == nil {
		var ret LogsMetricCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LogsMetricCreateRequest) GetDataOk() (*LogsMetricCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LogsMetricCreateRequest) SetData(v LogsMetricCreateData) {
	o.Data = v
}

func (o LogsMetricCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *LogsMetricCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Data *LogsMetricCreateData `json:"data"`
	}{}
	all := struct {
		Data LogsMetricCreateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		o.UnparsedObject = raw
	}
	if _, ok := o.UnparsedObject["data"]; required.Data == nil && !ok {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	o.Data = all.Data
	return nil
}

type NullableLogsMetricCreateRequest struct {
	value *LogsMetricCreateRequest
	isSet bool
}

func (v NullableLogsMetricCreateRequest) Get() *LogsMetricCreateRequest {
	return v.value
}

func (v *NullableLogsMetricCreateRequest) Set(val *LogsMetricCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsMetricCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsMetricCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsMetricCreateRequest(val *LogsMetricCreateRequest) *NullableLogsMetricCreateRequest {
	return &NullableLogsMetricCreateRequest{value: val, isSet: true}
}

func (v NullableLogsMetricCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsMetricCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
