# AWSDeleteAccountRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AccessKeyId** | Pointer to **string** | Your AWS access key ID. Only required if your AWS account is a GovCloud or China account. | [optional] 
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | [optional] 
**RoleName** | Pointer to **string** | Your Datadog role delegation name. | [optional] 

## Methods

### NewAWSDeleteAccountRequest

`func NewAWSDeleteAccountRequest() *AWSDeleteAccountRequest`

NewAWSDeleteAccountRequest instantiates a new AWSDeleteAccountRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSDeleteAccountRequestWithDefaults

`func NewAWSDeleteAccountRequestWithDefaults() *AWSDeleteAccountRequest`

NewAWSDeleteAccountRequestWithDefaults instantiates a new AWSDeleteAccountRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccessKeyId

`func (o *AWSDeleteAccountRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSDeleteAccountRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSDeleteAccountRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AWSDeleteAccountRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccountId

`func (o *AWSDeleteAccountRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSDeleteAccountRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSDeleteAccountRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSDeleteAccountRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRoleName

`func (o *AWSDeleteAccountRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AWSDeleteAccountRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AWSDeleteAccountRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AWSDeleteAccountRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


