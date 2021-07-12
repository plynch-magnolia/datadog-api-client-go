# FastlyService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Your Fastly service ID. | 
**ServiceTags** | **string** | A list of tags(key:value), seperated by commas, to assign to each Fastly metric. | 
**ServiceAccount** | **string** | Your service account associated with this Fastly integration. | 

## Methods

### NewFastlyService

`func NewFastlyService(serviceId string, serviceTags string, serviceAccount string, ) *FastlyService`

NewFastlyService instantiates a new FastlyService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastlyServiceWithDefaults

`func NewFastlyServiceWithDefaults() *FastlyService`

NewFastlyServiceWithDefaults instantiates a new FastlyService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *FastlyService) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FastlyService) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FastlyService) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceTags

`func (o *FastlyService) GetServiceTags() string`

GetServiceTags returns the ServiceTags field if non-nil, zero value otherwise.

### GetServiceTagsOk

`func (o *FastlyService) GetServiceTagsOk() (*string, bool)`

GetServiceTagsOk returns a tuple with the ServiceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTags

`func (o *FastlyService) SetServiceTags(v string)`

SetServiceTags sets ServiceTags field to given value.


### GetServiceAccount

`func (o *FastlyService) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *FastlyService) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *FastlyService) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


