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

// APIErrorResponse API Error response.
type APIErrorResponse struct {
	// An array of error messages.
	Errors []string `json:"errors"`
}

// NewAPIErrorResponse instantiates a new APIErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIErrorResponse(errors []string) *APIErrorResponse {
	this := APIErrorResponse{}
	this.Errors = errors
	return &this
}

// NewAPIErrorResponseWithDefaults instantiates a new APIErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIErrorResponseWithDefaults() *APIErrorResponse {
	this := APIErrorResponse{}
	return &this
}

// GetErrors returns the Errors field value
func (o *APIErrorResponse) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *APIErrorResponse) GetErrorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *APIErrorResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o APIErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableAPIErrorResponse struct {
	value *APIErrorResponse
	isSet bool
}

func (v NullableAPIErrorResponse) Get() *APIErrorResponse {
	return v.value
}

func (v *NullableAPIErrorResponse) Set(val *APIErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIErrorResponse(val *APIErrorResponse) *NullableAPIErrorResponse {
	return &NullableAPIErrorResponse{value: val, isSet: true}
}

func (v NullableAPIErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
