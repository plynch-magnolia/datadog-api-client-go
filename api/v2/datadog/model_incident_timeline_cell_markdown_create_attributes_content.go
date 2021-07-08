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

// IncidentTimelineCellMarkdownCreateAttributesContent The Markdown timeline cell contents.
type IncidentTimelineCellMarkdownCreateAttributesContent struct {
	// The Markdown content of the cell.
	Content *string `json:"content,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentTimelineCellMarkdownCreateAttributesContent instantiates a new IncidentTimelineCellMarkdownCreateAttributesContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTimelineCellMarkdownCreateAttributesContent() *IncidentTimelineCellMarkdownCreateAttributesContent {
	this := IncidentTimelineCellMarkdownCreateAttributesContent{}
	return &this
}

// NewIncidentTimelineCellMarkdownCreateAttributesContentWithDefaults instantiates a new IncidentTimelineCellMarkdownCreateAttributesContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTimelineCellMarkdownCreateAttributesContentWithDefaults() *IncidentTimelineCellMarkdownCreateAttributesContent {
	this := IncidentTimelineCellMarkdownCreateAttributesContent{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *IncidentTimelineCellMarkdownCreateAttributesContent) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributesContent) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *IncidentTimelineCellMarkdownCreateAttributesContent) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *IncidentTimelineCellMarkdownCreateAttributesContent) SetContent(v string) {
	o.Content = &v
}

func (o IncidentTimelineCellMarkdownCreateAttributesContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentTimelineCellMarkdownCreateAttributesContent) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Content *string `json:"content,omitempty"`
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
	o.Content = all.Content
	return nil
}

type NullableIncidentTimelineCellMarkdownCreateAttributesContent struct {
	value *IncidentTimelineCellMarkdownCreateAttributesContent
	isSet bool
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributesContent) Get() *IncidentTimelineCellMarkdownCreateAttributesContent {
	return v.value
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributesContent) Set(val *IncidentTimelineCellMarkdownCreateAttributesContent) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributesContent) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributesContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTimelineCellMarkdownCreateAttributesContent(val *IncidentTimelineCellMarkdownCreateAttributesContent) *NullableIncidentTimelineCellMarkdownCreateAttributesContent {
	return &NullableIncidentTimelineCellMarkdownCreateAttributesContent{value: val, isSet: true}
}

func (v NullableIncidentTimelineCellMarkdownCreateAttributesContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTimelineCellMarkdownCreateAttributesContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
