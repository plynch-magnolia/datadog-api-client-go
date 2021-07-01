# ListStreamQuery

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DataSource** | [**ListStreamWidgetDefinitionSource**](ListStreamWidgetDefinitionSource.md) |  | [default to LISTSTREAMWIDGETDEFINITIONSOURCE_ISSUE_STREAM]
**QueryString** | **string** | Widget query. | 

## Methods

### NewListStreamQuery

`func NewListStreamQuery(dataSource ListStreamWidgetDefinitionSource, queryString string, ) *ListStreamQuery`

NewListStreamQuery instantiates a new ListStreamQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStreamQueryWithDefaults

`func NewListStreamQueryWithDefaults() *ListStreamQuery`

NewListStreamQueryWithDefaults instantiates a new ListStreamQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *ListStreamQuery) GetDataSource() ListStreamWidgetDefinitionSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ListStreamQuery) GetDataSourceOk() (*ListStreamWidgetDefinitionSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ListStreamQuery) SetDataSource(v ListStreamWidgetDefinitionSource)`

SetDataSource sets DataSource field to given value.


### GetQueryString

`func (o *ListStreamQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ListStreamQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ListStreamQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


