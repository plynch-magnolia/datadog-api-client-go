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

// IncidentUpdateRelationships The incident's relationships for an update request.
type IncidentUpdateRelationships struct {
	CommanderUser      *RelationshipToUser                         `json:"commander_user,omitempty"`
	CreatedByUser      *RelationshipToUser                         `json:"created_by_user,omitempty"`
	Integrations       *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
	LastModifiedByUser *RelationshipToUser                         `json:"last_modified_by_user,omitempty"`
	Postmortem         *RelationshipToIncidentPostmortem           `json:"postmortem,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentUpdateRelationships instantiates a new IncidentUpdateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentUpdateRelationships() *IncidentUpdateRelationships {
	this := IncidentUpdateRelationships{}
	return &this
}

// NewIncidentUpdateRelationshipsWithDefaults instantiates a new IncidentUpdateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentUpdateRelationshipsWithDefaults() *IncidentUpdateRelationships {
	this := IncidentUpdateRelationships{}
	return &this
}

// GetCommanderUser returns the CommanderUser field value if set, zero value otherwise.
func (o *IncidentUpdateRelationships) GetCommanderUser() RelationshipToUser {
	if o == nil || o.CommanderUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CommanderUser
}

// GetCommanderUserOk returns a tuple with the CommanderUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRelationships) GetCommanderUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CommanderUser == nil {
		return nil, false
	}
	return o.CommanderUser, true
}

// HasCommanderUser returns a boolean if a field has been set.
func (o *IncidentUpdateRelationships) HasCommanderUser() bool {
	if o != nil && o.CommanderUser != nil {
		return true
	}

	return false
}

// SetCommanderUser gets a reference to the given RelationshipToUser and assigns it to the CommanderUser field.
func (o *IncidentUpdateRelationships) SetCommanderUser(v RelationshipToUser) {
	o.CommanderUser = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *IncidentUpdateRelationships) GetCreatedByUser() RelationshipToUser {
	if o == nil || o.CreatedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *IncidentUpdateRelationships) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given RelationshipToUser and assigns it to the CreatedByUser field.
func (o *IncidentUpdateRelationships) SetCreatedByUser(v RelationshipToUser) {
	o.CreatedByUser = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *IncidentUpdateRelationships) GetIntegrations() RelationshipToIncidentIntegrationMetadatas {
	if o == nil || o.Integrations == nil {
		var ret RelationshipToIncidentIntegrationMetadatas
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRelationships) GetIntegrationsOk() (*RelationshipToIncidentIntegrationMetadatas, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *IncidentUpdateRelationships) HasIntegrations() bool {
	if o != nil && o.Integrations != nil {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given RelationshipToIncidentIntegrationMetadatas and assigns it to the Integrations field.
func (o *IncidentUpdateRelationships) SetIntegrations(v RelationshipToIncidentIntegrationMetadatas) {
	o.Integrations = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentUpdateRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentUpdateRelationships) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentUpdateRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetPostmortem returns the Postmortem field value if set, zero value otherwise.
func (o *IncidentUpdateRelationships) GetPostmortem() RelationshipToIncidentPostmortem {
	if o == nil || o.Postmortem == nil {
		var ret RelationshipToIncidentPostmortem
		return ret
	}
	return *o.Postmortem
}

// GetPostmortemOk returns a tuple with the Postmortem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRelationships) GetPostmortemOk() (*RelationshipToIncidentPostmortem, bool) {
	if o == nil || o.Postmortem == nil {
		return nil, false
	}
	return o.Postmortem, true
}

// HasPostmortem returns a boolean if a field has been set.
func (o *IncidentUpdateRelationships) HasPostmortem() bool {
	if o != nil && o.Postmortem != nil {
		return true
	}

	return false
}

// SetPostmortem gets a reference to the given RelationshipToIncidentPostmortem and assigns it to the Postmortem field.
func (o *IncidentUpdateRelationships) SetPostmortem(v RelationshipToIncidentPostmortem) {
	o.Postmortem = &v
}

func (o IncidentUpdateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CommanderUser != nil {
		toSerialize["commander_user"] = o.CommanderUser
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.Postmortem != nil {
		toSerialize["postmortem"] = o.Postmortem
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentUpdateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CommanderUser      *RelationshipToUser                         `json:"commander_user,omitempty"`
		CreatedByUser      *RelationshipToUser                         `json:"created_by_user,omitempty"`
		Integrations       *RelationshipToIncidentIntegrationMetadatas `json:"integrations,omitempty"`
		LastModifiedByUser *RelationshipToUser                         `json:"last_modified_by_user,omitempty"`
		Postmortem         *RelationshipToIncidentPostmortem           `json:"postmortem,omitempty"`
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
	o.CommanderUser = all.CommanderUser
	o.CreatedByUser = all.CreatedByUser
	o.Integrations = all.Integrations
	o.LastModifiedByUser = all.LastModifiedByUser
	o.Postmortem = all.Postmortem
	return nil
}

type NullableIncidentUpdateRelationships struct {
	value *IncidentUpdateRelationships
	isSet bool
}

func (v NullableIncidentUpdateRelationships) Get() *IncidentUpdateRelationships {
	return v.value
}

func (v *NullableIncidentUpdateRelationships) Set(val *IncidentUpdateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentUpdateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentUpdateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentUpdateRelationships(val *IncidentUpdateRelationships) *NullableIncidentUpdateRelationships {
	return &NullableIncidentUpdateRelationships{value: val, isSet: true}
}

func (v NullableIncidentUpdateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentUpdateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
