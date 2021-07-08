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

// MetricTagConfigurationCreateAttributes Object containing the definition of a metric tag configuration to be created.
type MetricTagConfigurationCreateAttributes struct {
	// Toggle to include/exclude percentiles for a distribution metric. Defaults to false. Can only be applied to metrics that have a `metric_type` of `distribution`.
	IncludePercentiles *bool                             `json:"include_percentiles,omitempty"`
	MetricType         MetricTagConfigurationMetricTypes `json:"metric_type"`
	// A list of tag keys that will be queryable for your metric.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewMetricTagConfigurationCreateAttributes instantiates a new MetricTagConfigurationCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationCreateAttributes(metricType MetricTagConfigurationMetricTypes, tags []string) *MetricTagConfigurationCreateAttributes {
	this := MetricTagConfigurationCreateAttributes{}
	var includePercentiles bool = false
	this.IncludePercentiles = &includePercentiles
	this.MetricType = metricType
	this.Tags = tags
	return &this
}

// NewMetricTagConfigurationCreateAttributesWithDefaults instantiates a new MetricTagConfigurationCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationCreateAttributesWithDefaults() *MetricTagConfigurationCreateAttributes {
	this := MetricTagConfigurationCreateAttributes{}
	var includePercentiles bool = false
	this.IncludePercentiles = &includePercentiles
	var metricType MetricTagConfigurationMetricTypes = METRICTAGCONFIGURATIONMETRICTYPES_GAUGE
	this.MetricType = metricType
	return &this
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *MetricTagConfigurationCreateAttributes) HasIncludePercentiles() bool {
	if o != nil && o.IncludePercentiles != nil {
		return true
	}

	return false
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *MetricTagConfigurationCreateAttributes) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetMetricType returns the MetricType field value
func (o *MetricTagConfigurationCreateAttributes) GetMetricType() MetricTagConfigurationMetricTypes {
	if o == nil {
		var ret MetricTagConfigurationMetricTypes
		return ret
	}

	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetMetricTypeOk() (*MetricTagConfigurationMetricTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value
func (o *MetricTagConfigurationCreateAttributes) SetMetricType(v MetricTagConfigurationMetricTypes) {
	o.MetricType = v
}

// GetTags returns the Tags field value
func (o *MetricTagConfigurationCreateAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *MetricTagConfigurationCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

func (o MetricTagConfigurationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	if true {
		toSerialize["metric_type"] = o.MetricType
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

func (o *MetricTagConfigurationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		MetricType *MetricTagConfigurationMetricTypes `json:"metric_type"`
		Tags       *[]string                          `json:"tags"`
	}{}
	all := struct {
		IncludePercentiles *bool                             `json:"include_percentiles,omitempty"`
		MetricType         MetricTagConfigurationMetricTypes `json:"metric_type"`
		Tags               []string                          `json:"tags"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		o.UnparsedObject = raw
	}
	if _, ok := o.UnparsedObject["metric_type"]; required.MetricType == nil && !ok {
		return fmt.Errorf("Required field metric_type missing")
	}
	if _, ok := o.UnparsedObject["tags"]; required.Tags == nil && !ok {
		return fmt.Errorf("Required field tags missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	if v := all.MetricType; !v.IsValid() {
		o.UnparsedObject = raw
		return nil
	}
	o.IncludePercentiles = all.IncludePercentiles
	o.MetricType = all.MetricType
	o.Tags = all.Tags
	return nil
}

type NullableMetricTagConfigurationCreateAttributes struct {
	value *MetricTagConfigurationCreateAttributes
	isSet bool
}

func (v NullableMetricTagConfigurationCreateAttributes) Get() *MetricTagConfigurationCreateAttributes {
	return v.value
}

func (v *NullableMetricTagConfigurationCreateAttributes) Set(val *MetricTagConfigurationCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationCreateAttributes(val *MetricTagConfigurationCreateAttributes) *NullableMetricTagConfigurationCreateAttributes {
	return &NullableMetricTagConfigurationCreateAttributes{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
