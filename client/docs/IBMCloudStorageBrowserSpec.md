# IBMCloudStorageBrowserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | Region of the IBM Cloud object storage. | 
**ServicePoint** | **string** | Endpoint address and port number of the IBM Cloud object storage. | 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 

## Methods

### NewIBMCloudStorageBrowserSpec

`func NewIBMCloudStorageBrowserSpec(regionId string, servicePoint string, ) *IBMCloudStorageBrowserSpec`

NewIBMCloudStorageBrowserSpec instantiates a new IBMCloudStorageBrowserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageBrowserSpecWithDefaults

`func NewIBMCloudStorageBrowserSpecWithDefaults() *IBMCloudStorageBrowserSpec`

NewIBMCloudStorageBrowserSpecWithDefaults instantiates a new IBMCloudStorageBrowserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *IBMCloudStorageBrowserSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *IBMCloudStorageBrowserSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *IBMCloudStorageBrowserSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetServicePoint

`func (o *IBMCloudStorageBrowserSpec) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *IBMCloudStorageBrowserSpec) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *IBMCloudStorageBrowserSpec) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetGatewayServerId

`func (o *IBMCloudStorageBrowserSpec) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *IBMCloudStorageBrowserSpec) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *IBMCloudStorageBrowserSpec) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *IBMCloudStorageBrowserSpec) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

