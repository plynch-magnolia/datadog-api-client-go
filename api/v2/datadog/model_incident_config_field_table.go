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

// IncidentConfigFieldTable ID representing the type of object for the field
type IncidentConfigFieldTable int32

// List of IncidentConfigFieldTable
const (
	INCIDENTCONFIGFIELDTABLE_INCIDENT   IncidentConfigFieldTable = 0
	INCIDENTCONFIGFIELDTABLE_POSTMORTEM IncidentConfigFieldTable = 1
)

// Ptr returns reference to IncidentConfigFieldTable value
func (v IncidentConfigFieldTable) Ptr() *IncidentConfigFieldTable {
	return &v
}

type NullableIncidentConfigFieldTable struct {
	value *IncidentConfigFieldTable
	isSet bool
}

func (v NullableIncidentConfigFieldTable) Get() *IncidentConfigFieldTable {
	return v.value
}

func (v *NullableIncidentConfigFieldTable) Set(val *IncidentConfigFieldTable) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentConfigFieldTable) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentConfigFieldTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentConfigFieldTable(val *IncidentConfigFieldTable) *NullableIncidentConfigFieldTable {
	return &NullableIncidentConfigFieldTable{value: val, isSet: true}
}

func (v NullableIncidentConfigFieldTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentConfigFieldTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
