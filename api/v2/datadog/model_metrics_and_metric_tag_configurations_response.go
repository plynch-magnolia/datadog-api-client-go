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

// MetricsAndMetricTagConfigurationsResponse Response object that includes metrics and metric tag configurations.
type MetricsAndMetricTagConfigurationsResponse struct {
	// Array of metrics and metric tag configurations.
	Data *[]MetricsAndMetricTagConfigurations `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewMetricsAndMetricTagConfigurationsResponse instantiates a new MetricsAndMetricTagConfigurationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsAndMetricTagConfigurationsResponse() *MetricsAndMetricTagConfigurationsResponse {
	this := MetricsAndMetricTagConfigurationsResponse{}
	return &this
}

// NewMetricsAndMetricTagConfigurationsResponseWithDefaults instantiates a new MetricsAndMetricTagConfigurationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsAndMetricTagConfigurationsResponseWithDefaults() *MetricsAndMetricTagConfigurationsResponse {
	this := MetricsAndMetricTagConfigurationsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MetricsAndMetricTagConfigurationsResponse) GetData() []MetricsAndMetricTagConfigurations {
	if o == nil || o.Data == nil {
		var ret []MetricsAndMetricTagConfigurations
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsAndMetricTagConfigurationsResponse) GetDataOk() (*[]MetricsAndMetricTagConfigurations, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MetricsAndMetricTagConfigurationsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []MetricsAndMetricTagConfigurations and assigns it to the Data field.
func (o *MetricsAndMetricTagConfigurationsResponse) SetData(v []MetricsAndMetricTagConfigurations) {
	o.Data = &v
}

func (o MetricsAndMetricTagConfigurationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *MetricsAndMetricTagConfigurationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *[]MetricsAndMetricTagConfigurations `json:"data,omitempty"`
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
	o.Data = all.Data
	return nil
}

type NullableMetricsAndMetricTagConfigurationsResponse struct {
	value *MetricsAndMetricTagConfigurationsResponse
	isSet bool
}

func (v NullableMetricsAndMetricTagConfigurationsResponse) Get() *MetricsAndMetricTagConfigurationsResponse {
	return v.value
}

func (v *NullableMetricsAndMetricTagConfigurationsResponse) Set(val *MetricsAndMetricTagConfigurationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsAndMetricTagConfigurationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsAndMetricTagConfigurationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsAndMetricTagConfigurationsResponse(val *MetricsAndMetricTagConfigurationsResponse) *NullableMetricsAndMetricTagConfigurationsResponse {
	return &NullableMetricsAndMetricTagConfigurationsResponse{value: val, isSet: true}
}

func (v NullableMetricsAndMetricTagConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsAndMetricTagConfigurationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
