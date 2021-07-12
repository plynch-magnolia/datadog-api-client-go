# FastlyServiceTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceTags** | **string** | A list of tags(key:value), seperated by commas, to assign to each Fastly metric. | 

## Methods

### NewFastlyServiceTags

`func NewFastlyServiceTags(serviceTags string, ) *FastlyServiceTags`

NewFastlyServiceTags instantiates a new FastlyServiceTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastlyServiceTagsWithDefaults

`func NewFastlyServiceTagsWithDefaults() *FastlyServiceTags`

NewFastlyServiceTagsWithDefaults instantiates a new FastlyServiceTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceTags

`func (o *FastlyServiceTags) GetServiceTags() string`

GetServiceTags returns the ServiceTags field if non-nil, zero value otherwise.

### GetServiceTagsOk

`func (o *FastlyServiceTags) GetServiceTagsOk() (*string, bool)`

GetServiceTagsOk returns a tuple with the ServiceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTags

`func (o *FastlyServiceTags) SetServiceTags(v string)`

SetServiceTags sets ServiceTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


