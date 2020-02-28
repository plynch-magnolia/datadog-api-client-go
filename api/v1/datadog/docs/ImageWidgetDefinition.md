# ImageWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margin** | Pointer to [**WidgetMargin**](WidgetMargin.md) |  | [optional] 
**Sizing** | Pointer to [**WidgetImageSizing**](WidgetImageSizing.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "image"]
**Url** | Pointer to **string** | URL of the image | 

## Methods

### NewImageWidgetDefinition

`func NewImageWidgetDefinition(type_ string, url string, ) *ImageWidgetDefinition`

NewImageWidgetDefinition instantiates a new ImageWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWidgetDefinitionWithDefaults

`func NewImageWidgetDefinitionWithDefaults() *ImageWidgetDefinition`

NewImageWidgetDefinitionWithDefaults instantiates a new ImageWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMargin

`func (o *ImageWidgetDefinition) GetMargin() WidgetMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ImageWidgetDefinition) GetMarginOk() (WidgetMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMargin

`func (o *ImageWidgetDefinition) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### SetMargin

`func (o *ImageWidgetDefinition) SetMargin(v WidgetMargin)`

SetMargin gets a reference to the given WidgetMargin and assigns it to the Margin field.

### GetSizing

`func (o *ImageWidgetDefinition) GetSizing() WidgetImageSizing`

GetSizing returns the Sizing field if non-nil, zero value otherwise.

### GetSizingOk

`func (o *ImageWidgetDefinition) GetSizingOk() (WidgetImageSizing, bool)`

GetSizingOk returns a tuple with the Sizing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSizing

`func (o *ImageWidgetDefinition) HasSizing() bool`

HasSizing returns a boolean if a field has been set.

### SetSizing

`func (o *ImageWidgetDefinition) SetSizing(v WidgetImageSizing)`

SetSizing gets a reference to the given WidgetImageSizing and assigns it to the Sizing field.

### GetType

`func (o *ImageWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *ImageWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *ImageWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetUrl

`func (o *ImageWidgetDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageWidgetDefinition) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *ImageWidgetDefinition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *ImageWidgetDefinition) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.


### AsWidgetDefinition

`func (s *ImageWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of ImageWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

