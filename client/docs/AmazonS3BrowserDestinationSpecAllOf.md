# AmazonS3BrowserDestinationSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**RegionId** | **string** | AWS region where the Amazon S3 bucket is located. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 

## Methods

### NewAmazonS3BrowserDestinationSpecAllOf

`func NewAmazonS3BrowserDestinationSpecAllOf(regionType EAmazonRegionType, regionId string, bucketName string, ) *AmazonS3BrowserDestinationSpecAllOf`

NewAmazonS3BrowserDestinationSpecAllOf instantiates a new AmazonS3BrowserDestinationSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3BrowserDestinationSpecAllOfWithDefaults

`func NewAmazonS3BrowserDestinationSpecAllOfWithDefaults() *AmazonS3BrowserDestinationSpecAllOf`

NewAmazonS3BrowserDestinationSpecAllOfWithDefaults instantiates a new AmazonS3BrowserDestinationSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonS3BrowserDestinationSpecAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AmazonS3BrowserDestinationSpecAllOf) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionType

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3BrowserDestinationSpecAllOf) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetRegionId

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonS3BrowserDestinationSpecAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonS3BrowserDestinationSpecAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderType

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *AmazonS3BrowserDestinationSpecAllOf) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *AmazonS3BrowserDestinationSpecAllOf) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *AmazonS3BrowserDestinationSpecAllOf) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

