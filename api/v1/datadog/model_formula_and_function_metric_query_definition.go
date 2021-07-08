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

// FormulaAndFunctionMetricQueryDefinition A formula and functions metrics query.
type FormulaAndFunctionMetricQueryDefinition struct {
	Aggregator *FormulaAndFunctionMetricAggregation `json:"aggregator,omitempty"`
	DataSource FormulaAndFunctionMetricDataSource   `json:"data_source"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Metrics query definition.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewFormulaAndFunctionMetricQueryDefinition instantiates a new FormulaAndFunctionMetricQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormulaAndFunctionMetricQueryDefinition(dataSource FormulaAndFunctionMetricDataSource, name string, query string) *FormulaAndFunctionMetricQueryDefinition {
	this := FormulaAndFunctionMetricQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionMetricQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormulaAndFunctionMetricQueryDefinitionWithDefaults() *FormulaAndFunctionMetricQueryDefinition {
	this := FormulaAndFunctionMetricQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation {
	if o == nil || o.Aggregator == nil {
		var ret FormulaAndFunctionMetricAggregation
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) HasAggregator() bool {
	if o != nil && o.Aggregator != nil {
		return true
	}

	return false
}

// SetAggregator gets a reference to the given FormulaAndFunctionMetricAggregation and assigns it to the Aggregator field.
func (o *FormulaAndFunctionMetricQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation) {
	o.Aggregator = &v
}

// GetDataSource returns the DataSource field value
func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSource() FormulaAndFunctionMetricDataSource {
	if o == nil {
		var ret FormulaAndFunctionMetricDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *FormulaAndFunctionMetricQueryDefinition) SetDataSource(v FormulaAndFunctionMetricDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value
func (o *FormulaAndFunctionMetricQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormulaAndFunctionMetricQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value
func (o *FormulaAndFunctionMetricQueryDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *FormulaAndFunctionMetricQueryDefinition) SetQuery(v string) {
	o.Query = v
}

func (o FormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if true {
		toSerialize["data_source"] = o.DataSource
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

func (o *FormulaAndFunctionMetricQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		DataSource *FormulaAndFunctionMetricDataSource `json:"data_source"`
		Name       *string                             `json:"name"`
		Query      *string                             `json:"query"`
	}{}
	all := struct {
		Aggregator *FormulaAndFunctionMetricAggregation `json:"aggregator,omitempty"`
		DataSource FormulaAndFunctionMetricDataSource   `json:"data_source"`
		Name       string                               `json:"name"`
		Query      string                               `json:"query"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		o.UnparsedObject = raw
	}
	if _, ok := o.UnparsedObject["data_source"]; required.DataSource == nil && !ok {
		return fmt.Errorf("Required field data_source missing")
	}
	if _, ok := o.UnparsedObject["name"]; required.Name == nil && !ok {
		return fmt.Errorf("Required field name missing")
	}
	if _, ok := o.UnparsedObject["query"]; required.Query == nil && !ok {
		return fmt.Errorf("Required field query missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Aggregator; v != nil && !v.IsValid() {
		o.UnparsedObject = raw
		return nil
	}
	if v := all.DataSource; !v.IsValid() {
		o.UnparsedObject = raw
		return nil
	}
	o.Aggregator = all.Aggregator
	o.DataSource = all.DataSource
	o.Name = all.Name
	o.Query = all.Query
	return nil
}

type NullableFormulaAndFunctionMetricQueryDefinition struct {
	value *FormulaAndFunctionMetricQueryDefinition
	isSet bool
}

func (v NullableFormulaAndFunctionMetricQueryDefinition) Get() *FormulaAndFunctionMetricQueryDefinition {
	return v.value
}

func (v *NullableFormulaAndFunctionMetricQueryDefinition) Set(val *FormulaAndFunctionMetricQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionMetricQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionMetricQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionMetricQueryDefinition(val *FormulaAndFunctionMetricQueryDefinition) *NullableFormulaAndFunctionMetricQueryDefinition {
	return &NullableFormulaAndFunctionMetricQueryDefinition{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionMetricQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
