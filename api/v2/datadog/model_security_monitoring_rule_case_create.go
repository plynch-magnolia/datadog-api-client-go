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

// SecurityMonitoringRuleCaseCreate Case when signal is generated.
type SecurityMonitoringRuleCaseCreate struct {
	// A rule case contains logical operations (`>`,`>=`, `&&`, `||`) to determine if a signal should be generated based on the event counts in the previously defined queries.
	Condition *string `json:"condition,omitempty"`
	// Name of the case.
	Name *string `json:"name,omitempty"`
	// Notification targets for each rule case.
	Notifications *[]string                      `json:"notifications,omitempty"`
	Status        SecurityMonitoringRuleSeverity `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSecurityMonitoringRuleCaseCreate instantiates a new SecurityMonitoringRuleCaseCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleCaseCreate(status SecurityMonitoringRuleSeverity) *SecurityMonitoringRuleCaseCreate {
	this := SecurityMonitoringRuleCaseCreate{}
	this.Status = status
	return &this
}

// NewSecurityMonitoringRuleCaseCreateWithDefaults instantiates a new SecurityMonitoringRuleCaseCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleCaseCreateWithDefaults() *SecurityMonitoringRuleCaseCreate {
	this := SecurityMonitoringRuleCaseCreate{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseCreate) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseCreate) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseCreate) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *SecurityMonitoringRuleCaseCreate) SetCondition(v string) {
	o.Condition = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleCaseCreate) SetName(v string) {
	o.Name = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleCaseCreate) GetNotifications() []string {
	if o == nil || o.Notifications == nil {
		var ret []string
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseCreate) GetNotificationsOk() (*[]string, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleCaseCreate) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *SecurityMonitoringRuleCaseCreate) SetNotifications(v []string) {
	o.Notifications = &v
}

// GetStatus returns the Status field value
func (o *SecurityMonitoringRuleCaseCreate) GetStatus() SecurityMonitoringRuleSeverity {
	if o == nil {
		var ret SecurityMonitoringRuleSeverity
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleCaseCreate) GetStatusOk() (*SecurityMonitoringRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SecurityMonitoringRuleCaseCreate) SetStatus(v SecurityMonitoringRuleSeverity) {
	o.Status = v
}

func (o SecurityMonitoringRuleCaseCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleCaseCreate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Status *SecurityMonitoringRuleSeverity `json:"status"`
	}{}
	all := struct {
		Condition     *string                        `json:"condition,omitempty"`
		Name          *string                        `json:"name,omitempty"`
		Notifications *[]string                      `json:"notifications,omitempty"`
		Status        SecurityMonitoringRuleSeverity `json:"status"`
	}{}
	err = json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		o.UnparsedObject = raw
	}
	if _, ok := o.UnparsedObject["status"]; required.Status == nil && !ok {
		return fmt.Errorf("Required field status missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Status; !v.IsValid() {
		o.UnparsedObject = raw
		return nil
	}
	o.Condition = all.Condition
	o.Name = all.Name
	o.Notifications = all.Notifications
	o.Status = all.Status
	return nil
}

type NullableSecurityMonitoringRuleCaseCreate struct {
	value *SecurityMonitoringRuleCaseCreate
	isSet bool
}

func (v NullableSecurityMonitoringRuleCaseCreate) Get() *SecurityMonitoringRuleCaseCreate {
	return v.value
}

func (v *NullableSecurityMonitoringRuleCaseCreate) Set(val *SecurityMonitoringRuleCaseCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleCaseCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleCaseCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleCaseCreate(val *SecurityMonitoringRuleCaseCreate) *NullableSecurityMonitoringRuleCaseCreate {
	return &NullableSecurityMonitoringRuleCaseCreate{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleCaseCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleCaseCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
