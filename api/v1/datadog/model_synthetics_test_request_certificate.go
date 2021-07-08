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

// SyntheticsTestRequestCertificate Client certificate to use when performing the test request.
type SyntheticsTestRequestCertificate struct {
	Cert *SyntheticsTestRequestCertificateItem `json:"cert,omitempty"`
	Key  *SyntheticsTestRequestCertificateItem `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSyntheticsTestRequestCertificate instantiates a new SyntheticsTestRequestCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestRequestCertificate() *SyntheticsTestRequestCertificate {
	this := SyntheticsTestRequestCertificate{}
	return &this
}

// NewSyntheticsTestRequestCertificateWithDefaults instantiates a new SyntheticsTestRequestCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestRequestCertificateWithDefaults() *SyntheticsTestRequestCertificate {
	this := SyntheticsTestRequestCertificate{}
	return &this
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *SyntheticsTestRequestCertificate) GetCert() SyntheticsTestRequestCertificateItem {
	if o == nil || o.Cert == nil {
		var ret SyntheticsTestRequestCertificateItem
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestCertificate) GetCertOk() (*SyntheticsTestRequestCertificateItem, bool) {
	if o == nil || o.Cert == nil {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *SyntheticsTestRequestCertificate) HasCert() bool {
	if o != nil && o.Cert != nil {
		return true
	}

	return false
}

// SetCert gets a reference to the given SyntheticsTestRequestCertificateItem and assigns it to the Cert field.
func (o *SyntheticsTestRequestCertificate) SetCert(v SyntheticsTestRequestCertificateItem) {
	o.Cert = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SyntheticsTestRequestCertificate) GetKey() SyntheticsTestRequestCertificateItem {
	if o == nil || o.Key == nil {
		var ret SyntheticsTestRequestCertificateItem
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestCertificate) GetKeyOk() (*SyntheticsTestRequestCertificateItem, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SyntheticsTestRequestCertificate) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given SyntheticsTestRequestCertificateItem and assigns it to the Key field.
func (o *SyntheticsTestRequestCertificate) SetKey(v SyntheticsTestRequestCertificateItem) {
	o.Key = &v
}

func (o SyntheticsTestRequestCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Cert != nil {
		toSerialize["cert"] = o.Cert
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsTestRequestCertificate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Cert *SyntheticsTestRequestCertificateItem `json:"cert,omitempty"`
		Key  *SyntheticsTestRequestCertificateItem `json:"key,omitempty"`
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
	o.Cert = all.Cert
	o.Key = all.Key
	return nil
}

type NullableSyntheticsTestRequestCertificate struct {
	value *SyntheticsTestRequestCertificate
	isSet bool
}

func (v NullableSyntheticsTestRequestCertificate) Get() *SyntheticsTestRequestCertificate {
	return v.value
}

func (v *NullableSyntheticsTestRequestCertificate) Set(val *SyntheticsTestRequestCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestRequestCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestRequestCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestRequestCertificate(val *SyntheticsTestRequestCertificate) *NullableSyntheticsTestRequestCertificate {
	return &NullableSyntheticsTestRequestCertificate{value: val, isSet: true}
}

func (v NullableSyntheticsTestRequestCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestRequestCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
