/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RepositoryImportSpecCollection struct for RepositoryImportSpecCollection
type RepositoryImportSpecCollection struct {
	// Array of Microsoft Windows-based repositories.
	WindowsLocalRepositories []WindowsLocalStorageImportSpec `json:"WindowsLocalRepositories,omitempty"`
	// Array of Linux-based repositories.
	LinuxLocalRepositories []LinuxLocalStorageImportSpec `json:"LinuxLocalRepositories,omitempty"`
	// Array of SMB backup repositories.
	SmbRepositories []SmbStorageImportSpec `json:"SmbRepositories,omitempty"`
	// Array of NFS backup repositories.
	NfsRepositories []NfsStorageImportSpec `json:"NfsRepositories,omitempty"`
	// Array of Microsoft Azure Blob storages.
	AzureBlobStorages []AzureBlobStorageImportSpec `json:"AzureBlobStorages,omitempty"`
	// Array of Microsoft Azure Data Box storages.
	AzureDataBoxStorages []AzureDataBoxStorageImportSpec `json:"AzureDataBoxStorages,omitempty"`
	// Array of Amazon S3 storages.
	AmazonS3Storages []AmazonS3StorageImportSpec `json:"AmazonS3Storages,omitempty"`
	// Array of AWS Snowball Edge storages.
	AmazonSnowballEdgeStorages []AmazonSnowballEdgeStorageImportSpec `json:"AmazonSnowballEdgeStorages,omitempty"`
	// Array of S3 compatible storages.
	S3CompatibleStorages []S3CompatibleStorageImportSpec `json:"S3CompatibleStorages,omitempty"`
	// Array of Google Cloud storages.
	GoogleCloudStorages []GoogleCloudStorageImportSpec `json:"GoogleCloudStorages,omitempty"`
	// Array of IBM Cloud storages.
	IBMCloudStorages []IBMCloudStorageImportSpec `json:"IBMCloudStorages,omitempty"`
	// Array of Amazon S3 Glacier storages.
	AmazonS3GlacierStorages []AmazonS3GlacierStorageImportSpec `json:"AmazonS3GlacierStorages,omitempty"`
	// Array of Microsoft Azure Archive storages.
	AzureArchiveStorages []AzureArchiveStorageImportSpec `json:"AzureArchiveStorages,omitempty"`
	// Array of Wasabi Cloud storages.
	WasabiCloudStorages []WasabiCloudStorageImportSpec `json:"WasabiCloudStorages,omitempty"`
	// Array of Linux hardened repositories.
	LinuxHardenedRepositories []LinuxHardenedStorageImportSpec `json:"LinuxHardenedRepositories,omitempty"`
}

// NewRepositoryImportSpecCollection instantiates a new RepositoryImportSpecCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryImportSpecCollection() *RepositoryImportSpecCollection {
	this := RepositoryImportSpecCollection{}
	return &this
}

// NewRepositoryImportSpecCollectionWithDefaults instantiates a new RepositoryImportSpecCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryImportSpecCollectionWithDefaults() *RepositoryImportSpecCollection {
	this := RepositoryImportSpecCollection{}
	return &this
}

// GetWindowsLocalRepositories returns the WindowsLocalRepositories field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetWindowsLocalRepositories() []WindowsLocalStorageImportSpec {
	if o == nil || isNil(o.WindowsLocalRepositories) {
		var ret []WindowsLocalStorageImportSpec
		return ret
	}
	return o.WindowsLocalRepositories
}

// GetWindowsLocalRepositoriesOk returns a tuple with the WindowsLocalRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetWindowsLocalRepositoriesOk() ([]WindowsLocalStorageImportSpec, bool) {
	if o == nil || isNil(o.WindowsLocalRepositories) {
    return nil, false
	}
	return o.WindowsLocalRepositories, true
}

// HasWindowsLocalRepositories returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasWindowsLocalRepositories() bool {
	if o != nil && !isNil(o.WindowsLocalRepositories) {
		return true
	}

	return false
}

// SetWindowsLocalRepositories gets a reference to the given []WindowsLocalStorageImportSpec and assigns it to the WindowsLocalRepositories field.
func (o *RepositoryImportSpecCollection) SetWindowsLocalRepositories(v []WindowsLocalStorageImportSpec) {
	o.WindowsLocalRepositories = v
}

// GetLinuxLocalRepositories returns the LinuxLocalRepositories field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetLinuxLocalRepositories() []LinuxLocalStorageImportSpec {
	if o == nil || isNil(o.LinuxLocalRepositories) {
		var ret []LinuxLocalStorageImportSpec
		return ret
	}
	return o.LinuxLocalRepositories
}

// GetLinuxLocalRepositoriesOk returns a tuple with the LinuxLocalRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetLinuxLocalRepositoriesOk() ([]LinuxLocalStorageImportSpec, bool) {
	if o == nil || isNil(o.LinuxLocalRepositories) {
    return nil, false
	}
	return o.LinuxLocalRepositories, true
}

// HasLinuxLocalRepositories returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasLinuxLocalRepositories() bool {
	if o != nil && !isNil(o.LinuxLocalRepositories) {
		return true
	}

	return false
}

// SetLinuxLocalRepositories gets a reference to the given []LinuxLocalStorageImportSpec and assigns it to the LinuxLocalRepositories field.
func (o *RepositoryImportSpecCollection) SetLinuxLocalRepositories(v []LinuxLocalStorageImportSpec) {
	o.LinuxLocalRepositories = v
}

// GetSmbRepositories returns the SmbRepositories field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetSmbRepositories() []SmbStorageImportSpec {
	if o == nil || isNil(o.SmbRepositories) {
		var ret []SmbStorageImportSpec
		return ret
	}
	return o.SmbRepositories
}

// GetSmbRepositoriesOk returns a tuple with the SmbRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetSmbRepositoriesOk() ([]SmbStorageImportSpec, bool) {
	if o == nil || isNil(o.SmbRepositories) {
    return nil, false
	}
	return o.SmbRepositories, true
}

// HasSmbRepositories returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasSmbRepositories() bool {
	if o != nil && !isNil(o.SmbRepositories) {
		return true
	}

	return false
}

// SetSmbRepositories gets a reference to the given []SmbStorageImportSpec and assigns it to the SmbRepositories field.
func (o *RepositoryImportSpecCollection) SetSmbRepositories(v []SmbStorageImportSpec) {
	o.SmbRepositories = v
}

// GetNfsRepositories returns the NfsRepositories field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetNfsRepositories() []NfsStorageImportSpec {
	if o == nil || isNil(o.NfsRepositories) {
		var ret []NfsStorageImportSpec
		return ret
	}
	return o.NfsRepositories
}

// GetNfsRepositoriesOk returns a tuple with the NfsRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetNfsRepositoriesOk() ([]NfsStorageImportSpec, bool) {
	if o == nil || isNil(o.NfsRepositories) {
    return nil, false
	}
	return o.NfsRepositories, true
}

// HasNfsRepositories returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasNfsRepositories() bool {
	if o != nil && !isNil(o.NfsRepositories) {
		return true
	}

	return false
}

// SetNfsRepositories gets a reference to the given []NfsStorageImportSpec and assigns it to the NfsRepositories field.
func (o *RepositoryImportSpecCollection) SetNfsRepositories(v []NfsStorageImportSpec) {
	o.NfsRepositories = v
}

// GetAzureBlobStorages returns the AzureBlobStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAzureBlobStorages() []AzureBlobStorageImportSpec {
	if o == nil || isNil(o.AzureBlobStorages) {
		var ret []AzureBlobStorageImportSpec
		return ret
	}
	return o.AzureBlobStorages
}

// GetAzureBlobStoragesOk returns a tuple with the AzureBlobStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAzureBlobStoragesOk() ([]AzureBlobStorageImportSpec, bool) {
	if o == nil || isNil(o.AzureBlobStorages) {
    return nil, false
	}
	return o.AzureBlobStorages, true
}

// HasAzureBlobStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAzureBlobStorages() bool {
	if o != nil && !isNil(o.AzureBlobStorages) {
		return true
	}

	return false
}

// SetAzureBlobStorages gets a reference to the given []AzureBlobStorageImportSpec and assigns it to the AzureBlobStorages field.
func (o *RepositoryImportSpecCollection) SetAzureBlobStorages(v []AzureBlobStorageImportSpec) {
	o.AzureBlobStorages = v
}

// GetAzureDataBoxStorages returns the AzureDataBoxStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAzureDataBoxStorages() []AzureDataBoxStorageImportSpec {
	if o == nil || isNil(o.AzureDataBoxStorages) {
		var ret []AzureDataBoxStorageImportSpec
		return ret
	}
	return o.AzureDataBoxStorages
}

// GetAzureDataBoxStoragesOk returns a tuple with the AzureDataBoxStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAzureDataBoxStoragesOk() ([]AzureDataBoxStorageImportSpec, bool) {
	if o == nil || isNil(o.AzureDataBoxStorages) {
    return nil, false
	}
	return o.AzureDataBoxStorages, true
}

// HasAzureDataBoxStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAzureDataBoxStorages() bool {
	if o != nil && !isNil(o.AzureDataBoxStorages) {
		return true
	}

	return false
}

// SetAzureDataBoxStorages gets a reference to the given []AzureDataBoxStorageImportSpec and assigns it to the AzureDataBoxStorages field.
func (o *RepositoryImportSpecCollection) SetAzureDataBoxStorages(v []AzureDataBoxStorageImportSpec) {
	o.AzureDataBoxStorages = v
}

// GetAmazonS3Storages returns the AmazonS3Storages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAmazonS3Storages() []AmazonS3StorageImportSpec {
	if o == nil || isNil(o.AmazonS3Storages) {
		var ret []AmazonS3StorageImportSpec
		return ret
	}
	return o.AmazonS3Storages
}

// GetAmazonS3StoragesOk returns a tuple with the AmazonS3Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAmazonS3StoragesOk() ([]AmazonS3StorageImportSpec, bool) {
	if o == nil || isNil(o.AmazonS3Storages) {
    return nil, false
	}
	return o.AmazonS3Storages, true
}

// HasAmazonS3Storages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAmazonS3Storages() bool {
	if o != nil && !isNil(o.AmazonS3Storages) {
		return true
	}

	return false
}

// SetAmazonS3Storages gets a reference to the given []AmazonS3StorageImportSpec and assigns it to the AmazonS3Storages field.
func (o *RepositoryImportSpecCollection) SetAmazonS3Storages(v []AmazonS3StorageImportSpec) {
	o.AmazonS3Storages = v
}

// GetAmazonSnowballEdgeStorages returns the AmazonSnowballEdgeStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAmazonSnowballEdgeStorages() []AmazonSnowballEdgeStorageImportSpec {
	if o == nil || isNil(o.AmazonSnowballEdgeStorages) {
		var ret []AmazonSnowballEdgeStorageImportSpec
		return ret
	}
	return o.AmazonSnowballEdgeStorages
}

// GetAmazonSnowballEdgeStoragesOk returns a tuple with the AmazonSnowballEdgeStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAmazonSnowballEdgeStoragesOk() ([]AmazonSnowballEdgeStorageImportSpec, bool) {
	if o == nil || isNil(o.AmazonSnowballEdgeStorages) {
    return nil, false
	}
	return o.AmazonSnowballEdgeStorages, true
}

// HasAmazonSnowballEdgeStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAmazonSnowballEdgeStorages() bool {
	if o != nil && !isNil(o.AmazonSnowballEdgeStorages) {
		return true
	}

	return false
}

// SetAmazonSnowballEdgeStorages gets a reference to the given []AmazonSnowballEdgeStorageImportSpec and assigns it to the AmazonSnowballEdgeStorages field.
func (o *RepositoryImportSpecCollection) SetAmazonSnowballEdgeStorages(v []AmazonSnowballEdgeStorageImportSpec) {
	o.AmazonSnowballEdgeStorages = v
}

// GetS3CompatibleStorages returns the S3CompatibleStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetS3CompatibleStorages() []S3CompatibleStorageImportSpec {
	if o == nil || isNil(o.S3CompatibleStorages) {
		var ret []S3CompatibleStorageImportSpec
		return ret
	}
	return o.S3CompatibleStorages
}

// GetS3CompatibleStoragesOk returns a tuple with the S3CompatibleStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetS3CompatibleStoragesOk() ([]S3CompatibleStorageImportSpec, bool) {
	if o == nil || isNil(o.S3CompatibleStorages) {
    return nil, false
	}
	return o.S3CompatibleStorages, true
}

// HasS3CompatibleStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasS3CompatibleStorages() bool {
	if o != nil && !isNil(o.S3CompatibleStorages) {
		return true
	}

	return false
}

// SetS3CompatibleStorages gets a reference to the given []S3CompatibleStorageImportSpec and assigns it to the S3CompatibleStorages field.
func (o *RepositoryImportSpecCollection) SetS3CompatibleStorages(v []S3CompatibleStorageImportSpec) {
	o.S3CompatibleStorages = v
}

// GetGoogleCloudStorages returns the GoogleCloudStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetGoogleCloudStorages() []GoogleCloudStorageImportSpec {
	if o == nil || isNil(o.GoogleCloudStorages) {
		var ret []GoogleCloudStorageImportSpec
		return ret
	}
	return o.GoogleCloudStorages
}

// GetGoogleCloudStoragesOk returns a tuple with the GoogleCloudStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetGoogleCloudStoragesOk() ([]GoogleCloudStorageImportSpec, bool) {
	if o == nil || isNil(o.GoogleCloudStorages) {
    return nil, false
	}
	return o.GoogleCloudStorages, true
}

// HasGoogleCloudStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasGoogleCloudStorages() bool {
	if o != nil && !isNil(o.GoogleCloudStorages) {
		return true
	}

	return false
}

// SetGoogleCloudStorages gets a reference to the given []GoogleCloudStorageImportSpec and assigns it to the GoogleCloudStorages field.
func (o *RepositoryImportSpecCollection) SetGoogleCloudStorages(v []GoogleCloudStorageImportSpec) {
	o.GoogleCloudStorages = v
}

// GetIBMCloudStorages returns the IBMCloudStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetIBMCloudStorages() []IBMCloudStorageImportSpec {
	if o == nil || isNil(o.IBMCloudStorages) {
		var ret []IBMCloudStorageImportSpec
		return ret
	}
	return o.IBMCloudStorages
}

// GetIBMCloudStoragesOk returns a tuple with the IBMCloudStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetIBMCloudStoragesOk() ([]IBMCloudStorageImportSpec, bool) {
	if o == nil || isNil(o.IBMCloudStorages) {
    return nil, false
	}
	return o.IBMCloudStorages, true
}

// HasIBMCloudStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasIBMCloudStorages() bool {
	if o != nil && !isNil(o.IBMCloudStorages) {
		return true
	}

	return false
}

// SetIBMCloudStorages gets a reference to the given []IBMCloudStorageImportSpec and assigns it to the IBMCloudStorages field.
func (o *RepositoryImportSpecCollection) SetIBMCloudStorages(v []IBMCloudStorageImportSpec) {
	o.IBMCloudStorages = v
}

// GetAmazonS3GlacierStorages returns the AmazonS3GlacierStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAmazonS3GlacierStorages() []AmazonS3GlacierStorageImportSpec {
	if o == nil || isNil(o.AmazonS3GlacierStorages) {
		var ret []AmazonS3GlacierStorageImportSpec
		return ret
	}
	return o.AmazonS3GlacierStorages
}

// GetAmazonS3GlacierStoragesOk returns a tuple with the AmazonS3GlacierStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAmazonS3GlacierStoragesOk() ([]AmazonS3GlacierStorageImportSpec, bool) {
	if o == nil || isNil(o.AmazonS3GlacierStorages) {
    return nil, false
	}
	return o.AmazonS3GlacierStorages, true
}

// HasAmazonS3GlacierStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAmazonS3GlacierStorages() bool {
	if o != nil && !isNil(o.AmazonS3GlacierStorages) {
		return true
	}

	return false
}

// SetAmazonS3GlacierStorages gets a reference to the given []AmazonS3GlacierStorageImportSpec and assigns it to the AmazonS3GlacierStorages field.
func (o *RepositoryImportSpecCollection) SetAmazonS3GlacierStorages(v []AmazonS3GlacierStorageImportSpec) {
	o.AmazonS3GlacierStorages = v
}

// GetAzureArchiveStorages returns the AzureArchiveStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetAzureArchiveStorages() []AzureArchiveStorageImportSpec {
	if o == nil || isNil(o.AzureArchiveStorages) {
		var ret []AzureArchiveStorageImportSpec
		return ret
	}
	return o.AzureArchiveStorages
}

// GetAzureArchiveStoragesOk returns a tuple with the AzureArchiveStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetAzureArchiveStoragesOk() ([]AzureArchiveStorageImportSpec, bool) {
	if o == nil || isNil(o.AzureArchiveStorages) {
    return nil, false
	}
	return o.AzureArchiveStorages, true
}

// HasAzureArchiveStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasAzureArchiveStorages() bool {
	if o != nil && !isNil(o.AzureArchiveStorages) {
		return true
	}

	return false
}

// SetAzureArchiveStorages gets a reference to the given []AzureArchiveStorageImportSpec and assigns it to the AzureArchiveStorages field.
func (o *RepositoryImportSpecCollection) SetAzureArchiveStorages(v []AzureArchiveStorageImportSpec) {
	o.AzureArchiveStorages = v
}

// GetWasabiCloudStorages returns the WasabiCloudStorages field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetWasabiCloudStorages() []WasabiCloudStorageImportSpec {
	if o == nil || isNil(o.WasabiCloudStorages) {
		var ret []WasabiCloudStorageImportSpec
		return ret
	}
	return o.WasabiCloudStorages
}

// GetWasabiCloudStoragesOk returns a tuple with the WasabiCloudStorages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetWasabiCloudStoragesOk() ([]WasabiCloudStorageImportSpec, bool) {
	if o == nil || isNil(o.WasabiCloudStorages) {
    return nil, false
	}
	return o.WasabiCloudStorages, true
}

// HasWasabiCloudStorages returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasWasabiCloudStorages() bool {
	if o != nil && !isNil(o.WasabiCloudStorages) {
		return true
	}

	return false
}

// SetWasabiCloudStorages gets a reference to the given []WasabiCloudStorageImportSpec and assigns it to the WasabiCloudStorages field.
func (o *RepositoryImportSpecCollection) SetWasabiCloudStorages(v []WasabiCloudStorageImportSpec) {
	o.WasabiCloudStorages = v
}

// GetLinuxHardenedRepositories returns the LinuxHardenedRepositories field value if set, zero value otherwise.
func (o *RepositoryImportSpecCollection) GetLinuxHardenedRepositories() []LinuxHardenedStorageImportSpec {
	if o == nil || isNil(o.LinuxHardenedRepositories) {
		var ret []LinuxHardenedStorageImportSpec
		return ret
	}
	return o.LinuxHardenedRepositories
}

// GetLinuxHardenedRepositoriesOk returns a tuple with the LinuxHardenedRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryImportSpecCollection) GetLinuxHardenedRepositoriesOk() ([]LinuxHardenedStorageImportSpec, bool) {
	if o == nil || isNil(o.LinuxHardenedRepositories) {
    return nil, false
	}
	return o.LinuxHardenedRepositories, true
}

// HasLinuxHardenedRepositories returns a boolean if a field has been set.
func (o *RepositoryImportSpecCollection) HasLinuxHardenedRepositories() bool {
	if o != nil && !isNil(o.LinuxHardenedRepositories) {
		return true
	}

	return false
}

// SetLinuxHardenedRepositories gets a reference to the given []LinuxHardenedStorageImportSpec and assigns it to the LinuxHardenedRepositories field.
func (o *RepositoryImportSpecCollection) SetLinuxHardenedRepositories(v []LinuxHardenedStorageImportSpec) {
	o.LinuxHardenedRepositories = v
}

func (o RepositoryImportSpecCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WindowsLocalRepositories) {
		toSerialize["WindowsLocalRepositories"] = o.WindowsLocalRepositories
	}
	if !isNil(o.LinuxLocalRepositories) {
		toSerialize["LinuxLocalRepositories"] = o.LinuxLocalRepositories
	}
	if !isNil(o.SmbRepositories) {
		toSerialize["SmbRepositories"] = o.SmbRepositories
	}
	if !isNil(o.NfsRepositories) {
		toSerialize["NfsRepositories"] = o.NfsRepositories
	}
	if !isNil(o.AzureBlobStorages) {
		toSerialize["AzureBlobStorages"] = o.AzureBlobStorages
	}
	if !isNil(o.AzureDataBoxStorages) {
		toSerialize["AzureDataBoxStorages"] = o.AzureDataBoxStorages
	}
	if !isNil(o.AmazonS3Storages) {
		toSerialize["AmazonS3Storages"] = o.AmazonS3Storages
	}
	if !isNil(o.AmazonSnowballEdgeStorages) {
		toSerialize["AmazonSnowballEdgeStorages"] = o.AmazonSnowballEdgeStorages
	}
	if !isNil(o.S3CompatibleStorages) {
		toSerialize["S3CompatibleStorages"] = o.S3CompatibleStorages
	}
	if !isNil(o.GoogleCloudStorages) {
		toSerialize["GoogleCloudStorages"] = o.GoogleCloudStorages
	}
	if !isNil(o.IBMCloudStorages) {
		toSerialize["IBMCloudStorages"] = o.IBMCloudStorages
	}
	if !isNil(o.AmazonS3GlacierStorages) {
		toSerialize["AmazonS3GlacierStorages"] = o.AmazonS3GlacierStorages
	}
	if !isNil(o.AzureArchiveStorages) {
		toSerialize["AzureArchiveStorages"] = o.AzureArchiveStorages
	}
	if !isNil(o.WasabiCloudStorages) {
		toSerialize["WasabiCloudStorages"] = o.WasabiCloudStorages
	}
	if !isNil(o.LinuxHardenedRepositories) {
		toSerialize["LinuxHardenedRepositories"] = o.LinuxHardenedRepositories
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoryImportSpecCollection struct {
	value *RepositoryImportSpecCollection
	isSet bool
}

func (v NullableRepositoryImportSpecCollection) Get() *RepositoryImportSpecCollection {
	return v.value
}

func (v *NullableRepositoryImportSpecCollection) Set(val *RepositoryImportSpecCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryImportSpecCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryImportSpecCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryImportSpecCollection(val *RepositoryImportSpecCollection) *NullableRepositoryImportSpecCollection {
	return &NullableRepositoryImportSpecCollection{value: val, isSet: true}
}

func (v NullableRepositoryImportSpecCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryImportSpecCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


