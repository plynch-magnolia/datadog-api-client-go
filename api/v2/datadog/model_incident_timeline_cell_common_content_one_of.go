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

// IncidentTimelineCellCommonContentOneOf struct for IncidentTimelineCellCommonContentOneOf
type IncidentTimelineCellCommonContentOneOf struct {
	IncidentTimelineCellCommonContentOneOfInterface interface{ GetCellType() string }
}

func (s IncidentTimelineCellCommonContentOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.IncidentTimelineCellCommonContentOneOfInterface)
}

func (s *IncidentTimelineCellCommonContentOneOf) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["cell_type"]; ok {
		switch v {
		case "choice_change":
			var result *IncidentTimelineCellChoiceChangeContent = &IncidentTimelineCellChoiceChangeContent{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentTimelineCellCommonContentOneOfInterface = result
			return nil
		case "markdown":
			var result *IncidentTimelineCellMarkdownContent = &IncidentTimelineCellMarkdownContent{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentTimelineCellCommonContentOneOfInterface = result
			return nil
		case "status_change":
			var result *IncidentTimelineCellStatusChangeContent = &IncidentTimelineCellStatusChangeContent{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentTimelineCellCommonContentOneOfInterface = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'cell_type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'cell_type' not found in unmarshaled payload: %+v", unmarshaled)
	}
}

type NullableIncidentTimelineCellCommonContentOneOf struct {
	value *IncidentTimelineCellCommonContentOneOf
	isSet bool
}

func (v NullableIncidentTimelineCellCommonContentOneOf) Get() *IncidentTimelineCellCommonContentOneOf {
	return v.value
}

func (v *NullableIncidentTimelineCellCommonContentOneOf) Set(val *IncidentTimelineCellCommonContentOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTimelineCellCommonContentOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTimelineCellCommonContentOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTimelineCellCommonContentOneOf(val *IncidentTimelineCellCommonContentOneOf) *NullableIncidentTimelineCellCommonContentOneOf {
	return &NullableIncidentTimelineCellCommonContentOneOf{value: val, isSet: true}
}

func (v NullableIncidentTimelineCellCommonContentOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTimelineCellCommonContentOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
