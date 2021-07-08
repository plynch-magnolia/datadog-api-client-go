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

// DashboardListAddItemsResponse Response containing a list of added dashboards.
type DashboardListAddItemsResponse struct {
	// List of dashboards added to the dashboard list.
	AddedDashboardsToList *[]DashboardListItemResponse `json:"added_dashboards_to_list,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewDashboardListAddItemsResponse instantiates a new DashboardListAddItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardListAddItemsResponse() *DashboardListAddItemsResponse {
	this := DashboardListAddItemsResponse{}
	return &this
}

// NewDashboardListAddItemsResponseWithDefaults instantiates a new DashboardListAddItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardListAddItemsResponseWithDefaults() *DashboardListAddItemsResponse {
	this := DashboardListAddItemsResponse{}
	return &this
}

// GetAddedDashboardsToList returns the AddedDashboardsToList field value if set, zero value otherwise.
func (o *DashboardListAddItemsResponse) GetAddedDashboardsToList() []DashboardListItemResponse {
	if o == nil || o.AddedDashboardsToList == nil {
		var ret []DashboardListItemResponse
		return ret
	}
	return *o.AddedDashboardsToList
}

// GetAddedDashboardsToListOk returns a tuple with the AddedDashboardsToList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardListAddItemsResponse) GetAddedDashboardsToListOk() (*[]DashboardListItemResponse, bool) {
	if o == nil || o.AddedDashboardsToList == nil {
		return nil, false
	}
	return o.AddedDashboardsToList, true
}

// HasAddedDashboardsToList returns a boolean if a field has been set.
func (o *DashboardListAddItemsResponse) HasAddedDashboardsToList() bool {
	if o != nil && o.AddedDashboardsToList != nil {
		return true
	}

	return false
}

// SetAddedDashboardsToList gets a reference to the given []DashboardListItemResponse and assigns it to the AddedDashboardsToList field.
func (o *DashboardListAddItemsResponse) SetAddedDashboardsToList(v []DashboardListItemResponse) {
	o.AddedDashboardsToList = &v
}

func (o DashboardListAddItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AddedDashboardsToList != nil {
		toSerialize["added_dashboards_to_list"] = o.AddedDashboardsToList
	}
	return json.Marshal(toSerialize)
}

func (o *DashboardListAddItemsResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AddedDashboardsToList *[]DashboardListItemResponse `json:"added_dashboards_to_list,omitempty"`
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
	o.AddedDashboardsToList = all.AddedDashboardsToList
	return nil
}

type NullableDashboardListAddItemsResponse struct {
	value *DashboardListAddItemsResponse
	isSet bool
}

func (v NullableDashboardListAddItemsResponse) Get() *DashboardListAddItemsResponse {
	return v.value
}

func (v *NullableDashboardListAddItemsResponse) Set(val *DashboardListAddItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardListAddItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardListAddItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardListAddItemsResponse(val *DashboardListAddItemsResponse) *NullableDashboardListAddItemsResponse {
	return &NullableDashboardListAddItemsResponse{value: val, isSet: true}
}

func (v NullableDashboardListAddItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardListAddItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
