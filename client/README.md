# Go API client for client

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>
Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br>
Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0-rev2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AutomationApi* | [**ExportCredentials**](docs/AutomationApi.md#exportcredentials) | **Post** /api/v1/automation/credentials/export | Export Credentials
*AutomationApi* | [**ExportEncryptionPasswords**](docs/AutomationApi.md#exportencryptionpasswords) | **Post** /api/v1/automation/encryptionPasswords/export | Export Encryption Passwords
*AutomationApi* | [**ExportJobs**](docs/AutomationApi.md#exportjobs) | **Post** /api/v1/automation/jobs/export | Export Jobs
*AutomationApi* | [**ExportManagedServers**](docs/AutomationApi.md#exportmanagedservers) | **Post** /api/v1/automation/managedServers/export | Export Servers
*AutomationApi* | [**ExportProxies**](docs/AutomationApi.md#exportproxies) | **Post** /api/v1/automation/proxies/export | Export Proxies
*AutomationApi* | [**ExportRepositories**](docs/AutomationApi.md#exportrepositories) | **Post** /api/v1/automation/repositories/export | Export Repositories
*AutomationApi* | [**GetAllAutomationSessions**](docs/AutomationApi.md#getallautomationsessions) | **Get** /api/v1/automation/sessions | Get All Automation Sessions
*AutomationApi* | [**GetAutomationSession**](docs/AutomationApi.md#getautomationsession) | **Get** /api/v1/automation/sessions/{id} | Get Automation Session
*AutomationApi* | [**GetAutomationSessionLogs**](docs/AutomationApi.md#getautomationsessionlogs) | **Get** /api/v1/automation/sessions/{id}/logs | Get Automation Session Logs
*AutomationApi* | [**ImportCredentials**](docs/AutomationApi.md#importcredentials) | **Post** /api/v1/automation/credentials/import | Import Credentials
*AutomationApi* | [**ImportEncryptionPasswords**](docs/AutomationApi.md#importencryptionpasswords) | **Post** /api/v1/automation/encryptionPasswords/import | Import Encryption Passwords
*AutomationApi* | [**ImportJobs**](docs/AutomationApi.md#importjobs) | **Post** /api/v1/automation/jobs/import | Import Jobs
*AutomationApi* | [**ImportManagedServers**](docs/AutomationApi.md#importmanagedservers) | **Post** /api/v1/automation/managedServers/import | Import Servers
*AutomationApi* | [**ImportProxies**](docs/AutomationApi.md#importproxies) | **Post** /api/v1/automation/proxies/import | Import Proxies
*AutomationApi* | [**ImportRepositories**](docs/AutomationApi.md#importrepositories) | **Post** /api/v1/automation/repositories/import | Import Repositories
*AutomationApi* | [**StopAutomationSession**](docs/AutomationApi.md#stopautomationsession) | **Post** /api/v1/automation/sessions/{id}/stop | Stop Automation Session
*BackupObjectsApi* | [**GetAllBackupObjects**](docs/BackupObjectsApi.md#getallbackupobjects) | **Get** /api/v1/backupObjects | Get All Backup Objects
*BackupObjectsApi* | [**GetBackupObject**](docs/BackupObjectsApi.md#getbackupobject) | **Get** /api/v1/backupObjects/{id} | Get Backup Object
*BackupObjectsApi* | [**GetBackupObjectRestorePoints**](docs/BackupObjectsApi.md#getbackupobjectrestorepoints) | **Get** /api/v1/backupObjects/{id}/restorePoints | Get Restore Points
*BackupsApi* | [**GetAllBackups**](docs/BackupsApi.md#getallbackups) | **Get** /api/v1/backups | Get All Backups
*BackupsApi* | [**GetBackup**](docs/BackupsApi.md#getbackup) | **Get** /api/v1/backups/{id} | Get Backup
*BackupsApi* | [**GetBackupObjects**](docs/BackupsApi.md#getbackupobjects) | **Get** /api/v1/backups/{id}/objects | Get Backup Objects
*ConfigurationBackupApi* | [**GetConfigBackupOptions**](docs/ConfigurationBackupApi.md#getconfigbackupoptions) | **Get** /api/v1/configBackup | Get Configuration Backup
*ConfigurationBackupApi* | [**StartConfigBackup**](docs/ConfigurationBackupApi.md#startconfigbackup) | **Post** /api/v1/configBackup/backup | Start Configuration Backup
*ConfigurationBackupApi* | [**UpdateConfigBackupOptions**](docs/ConfigurationBackupApi.md#updateconfigbackupoptions) | **Put** /api/v1/configBackup | Edit Configuration Backup
*ConnectionApi* | [**GetConnectionCertificate**](docs/ConnectionApi.md#getconnectioncertificate) | **Post** /api/v1/connectionCertificate | Request TLS Certificate or SSH Fingerprint
*CredentialsApi* | [**ChangePasswordForCreds**](docs/CredentialsApi.md#changepasswordforcreds) | **Post** /api/v1/credentials/{id}/changepassword | Change Password
*CredentialsApi* | [**ChangePrivateKeyForCreds**](docs/CredentialsApi.md#changeprivatekeyforcreds) | **Post** /api/v1/credentials/{id}/changeprivatekey | Change Linux Private Key
*CredentialsApi* | [**ChangeRootPasswordForCreds**](docs/CredentialsApi.md#changerootpasswordforcreds) | **Post** /api/v1/credentials/{id}/changerootpassword | Change Linux Root Password
*CredentialsApi* | [**CreateCreds**](docs/CredentialsApi.md#createcreds) | **Post** /api/v1/credentials | Add Credentials Record
*CredentialsApi* | [**DeleteCreds**](docs/CredentialsApi.md#deletecreds) | **Delete** /api/v1/credentials/{id} | Remove Credentials Record
*CredentialsApi* | [**GetAllCreds**](docs/CredentialsApi.md#getallcreds) | **Get** /api/v1/credentials | Get All Credentials
*CredentialsApi* | [**GetCreds**](docs/CredentialsApi.md#getcreds) | **Get** /api/v1/credentials/{id} | Get Credentials Record
*CredentialsApi* | [**UpdateCreds**](docs/CredentialsApi.md#updatecreds) | **Put** /api/v1/credentials/{id} | Edit Credentials Record
*EncryptionApi* | [**CreateEncryptionPassword**](docs/EncryptionApi.md#createencryptionpassword) | **Post** /api/v1/encryptionPasswords | Add Encryption Password
*EncryptionApi* | [**DeleteEncryptionPassword**](docs/EncryptionApi.md#deleteencryptionpassword) | **Delete** /api/v1/encryptionPasswords/{id} | Remove Encryption Password
*EncryptionApi* | [**GetAllEncryptionPasswords**](docs/EncryptionApi.md#getallencryptionpasswords) | **Get** /api/v1/encryptionPasswords | Get All Encryption Passwords
*EncryptionApi* | [**GetEncryptionPassword**](docs/EncryptionApi.md#getencryptionpassword) | **Get** /api/v1/encryptionPasswords/{id} | Get Encryption Password
*EncryptionApi* | [**UpdateEncryptionPassword**](docs/EncryptionApi.md#updateencryptionpassword) | **Put** /api/v1/encryptionPasswords/{id} | Edit Encryption Password
*GeneralOptionsApi* | [**GetGeneralOptions**](docs/GeneralOptionsApi.md#getgeneraloptions) | **Get** /api/v1/generalOptions | Get General Options
*GeneralOptionsApi* | [**UpdateGeneralOptions**](docs/GeneralOptionsApi.md#updategeneraloptions) | **Put** /api/v1/generalOptions | Edit General Options
*InventoryBrowserApi* | [**GetAllInventoryVmwareHosts**](docs/InventoryBrowserApi.md#getallinventoryvmwarehosts) | **Get** /api/v1/inventory/vmware/hosts | Get All VMware vSphere Servers
*InventoryBrowserApi* | [**GetVmwareHostObject**](docs/InventoryBrowserApi.md#getvmwarehostobject) | **Get** /api/v1/inventory/vmware/hosts/{name} | Get VMware vSphere Server Objects
*JobsApi* | [**CreateJob**](docs/JobsApi.md#createjob) | **Post** /api/v1/jobs | Create Job
*JobsApi* | [**DeleteJob**](docs/JobsApi.md#deletejob) | **Delete** /api/v1/jobs/{id} | Delete Job
*JobsApi* | [**DisableJob**](docs/JobsApi.md#disablejob) | **Post** /api/v1/jobs/{id}/disable | Disable Job
*JobsApi* | [**EnableJob**](docs/JobsApi.md#enablejob) | **Post** /api/v1/jobs/{id}/enable | Enable Job
*JobsApi* | [**GetAllJobs**](docs/JobsApi.md#getalljobs) | **Get** /api/v1/jobs | Get All Jobs
*JobsApi* | [**GetAllJobsStates**](docs/JobsApi.md#getalljobsstates) | **Get** /api/v1/jobs/states | Get All Job States
*JobsApi* | [**GetJob**](docs/JobsApi.md#getjob) | **Get** /api/v1/jobs/{id} | Get Job
*JobsApi* | [**StartJob**](docs/JobsApi.md#startjob) | **Post** /api/v1/jobs/{id}/start | Start Job
*JobsApi* | [**StopJob**](docs/JobsApi.md#stopjob) | **Post** /api/v1/jobs/{id}/stop | Stop Job
*JobsApi* | [**UpdateJob**](docs/JobsApi.md#updatejob) | **Put** /api/v1/jobs/{id} | Edit Job
*LoginApi* | [**CreateAuthorizationCode**](docs/LoginApi.md#createauthorizationcode) | **Post** /api/oauth2/authorization_code | Get Authorization Code
*LoginApi* | [**CreateToken**](docs/LoginApi.md#createtoken) | **Post** /api/oauth2/token | Get Access Token
*LoginApi* | [**Logout**](docs/LoginApi.md#logout) | **Post** /api/oauth2/logout | Log Out
*ManagedServersApi* | [**CreateManagedServer**](docs/ManagedServersApi.md#createmanagedserver) | **Post** /api/v1/backupInfrastructure/managedServers | Add Server
*ManagedServersApi* | [**DeleteManagedServer**](docs/ManagedServersApi.md#deletemanagedserver) | **Delete** /api/v1/backupInfrastructure/managedServers/{id} | Remove Server
*ManagedServersApi* | [**GetAllManagedServers**](docs/ManagedServersApi.md#getallmanagedservers) | **Get** /api/v1/backupInfrastructure/managedServers | Get All Servers
*ManagedServersApi* | [**GetManagedServer**](docs/ManagedServersApi.md#getmanagedserver) | **Get** /api/v1/backupInfrastructure/managedServers/{id} | Get Server
*ManagedServersApi* | [**UpdateManagedServer**](docs/ManagedServersApi.md#updatemanagedserver) | **Put** /api/v1/backupInfrastructure/managedServers/{id} | Edit Server
*ObjectRestorePointsApi* | [**GetAllObjectRestorePoints**](docs/ObjectRestorePointsApi.md#getallobjectrestorepoints) | **Get** /api/v1/objectRestorePoints | Get All Restore Points
*ObjectRestorePointsApi* | [**GetObjectRestorePoint**](docs/ObjectRestorePointsApi.md#getobjectrestorepoint) | **Get** /api/v1/objectRestorePoints/{id} | Get Restore Point
*ObjectRestorePointsApi* | [**GetObjectRestorePointDisks**](docs/ObjectRestorePointsApi.md#getobjectrestorepointdisks) | **Get** /api/v1/objectRestorePoints/{id}/disks | Get Restore Point Disks
*ProxiesApi* | [**CreateProxy**](docs/ProxiesApi.md#createproxy) | **Post** /api/v1/backupInfrastructure/proxies | Add Proxy
*ProxiesApi* | [**DeleteProxy**](docs/ProxiesApi.md#deleteproxy) | **Delete** /api/v1/backupInfrastructure/proxies/{id} | Remove Proxy
*ProxiesApi* | [**GetAllProxies**](docs/ProxiesApi.md#getallproxies) | **Get** /api/v1/backupInfrastructure/proxies | Get All Proxies
*ProxiesApi* | [**GetProxy**](docs/ProxiesApi.md#getproxy) | **Get** /api/v1/backupInfrastructure/proxies/{id} | Get Proxy
*ProxiesApi* | [**UpdateProxy**](docs/ProxiesApi.md#updateproxy) | **Put** /api/v1/backupInfrastructure/proxies/{id} | Edit Proxy
*RepositoriesApi* | [**CreateRepository**](docs/RepositoriesApi.md#createrepository) | **Post** /api/v1/backupInfrastructure/repositories | Add Repository
*RepositoriesApi* | [**DeleteRepository**](docs/RepositoriesApi.md#deleterepository) | **Delete** /api/v1/backupInfrastructure/repositories/{id} | Remove Repository
*RepositoriesApi* | [**DisableScaleOutExtentMaintenanceMode**](docs/RepositoriesApi.md#disablescaleoutextentmaintenancemode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/disableMaintenanceMode | Disable Maintenance Mode
*RepositoriesApi* | [**EnableScaleOutExtentMaintenanceMode**](docs/RepositoriesApi.md#enablescaleoutextentmaintenancemode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode | Enable Maintenance Mode
*RepositoriesApi* | [**GetAllRepositories**](docs/RepositoriesApi.md#getallrepositories) | **Get** /api/v1/backupInfrastructure/repositories | Get All Repositories
*RepositoriesApi* | [**GetAllRepositoriesStates**](docs/RepositoriesApi.md#getallrepositoriesstates) | **Get** /api/v1/backupInfrastructure/repositories/states | Get All Repository States
*RepositoriesApi* | [**GetAllScaleOutRepositories**](docs/RepositoriesApi.md#getallscaleoutrepositories) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories | Get All Scale-Out Backup Repositories
*RepositoriesApi* | [**GetRepository**](docs/RepositoriesApi.md#getrepository) | **Get** /api/v1/backupInfrastructure/repositories/{id} | Get Repository
*RepositoriesApi* | [**GetScaleOutRepository**](docs/RepositoriesApi.md#getscaleoutrepository) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories/{id} | Get Scale-Out Backup Repository
*RepositoriesApi* | [**UpdateRepository**](docs/RepositoriesApi.md#updaterepository) | **Put** /api/v1/backupInfrastructure/repositories/{id} | Edit Repository
*RestoreApi* | [**GetAllVmwareFcdInstantRecoveryMountModels**](docs/RestoreApi.md#getallvmwarefcdinstantrecoverymountmodels) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/ | Get All FCD Mounts Information
*RestoreApi* | [**GetVmwareFcdInstantRecoveryMountModel**](docs/RestoreApi.md#getvmwarefcdinstantrecoverymountmodel) | **Get** /api/v1/restore/instantRecovery/vmware/fcd/{mountId} | Get Mount Information
*RestoreApi* | [**InstantRecoveryVmwareFcdDismountWithSession**](docs/RestoreApi.md#instantrecoveryvmwarefcddismountwithsession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/dismount | Stop FCD Publishing
*RestoreApi* | [**InstantRecoveryVmwareFcdMigrateWithSession**](docs/RestoreApi.md#instantrecoveryvmwarefcdmigratewithsession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/{mountId}/migrate | Start FCD Migration
*RestoreApi* | [**InstantRecoveryVmwareFcdMountWithSession**](docs/RestoreApi.md#instantrecoveryvmwarefcdmountwithsession) | **Post** /api/v1/restore/instantRecovery/vmware/fcd/ | Start Instant FCD Recovery
*ServiceApi* | [**GetServerCertificate**](docs/ServiceApi.md#getservercertificate) | **Get** /api/v1/serverCertificate | Get Server Certificate
*ServiceApi* | [**GetServerTime**](docs/ServiceApi.md#getservertime) | **Get** /api/v1/serverTime | Get Server Time
*ServicesApi* | [**GetAllServices**](docs/ServicesApi.md#getallservices) | **Get** /api/v1/services | Get Associated Services
*SessionsApi* | [**GetAllSessions**](docs/SessionsApi.md#getallsessions) | **Get** /api/v1/sessions | Get All Sessions
*SessionsApi* | [**GetSession**](docs/SessionsApi.md#getsession) | **Get** /api/v1/sessions/{id} | Get Session
*SessionsApi* | [**GetSessionLogs**](docs/SessionsApi.md#getsessionlogs) | **Get** /api/v1/sessions/{id}/logs | Get Session Logs
*SessionsApi* | [**StopSession**](docs/SessionsApi.md#stopsession) | **Post** /api/v1/sessions/{id}/stop | Stop Session
*TrafficRulesApi* | [**GetAllTrafficRules**](docs/TrafficRulesApi.md#getalltrafficrules) | **Get** /api/v1/trafficRules | Get Traffic Rules
*TrafficRulesApi* | [**UpdateTrafficRules**](docs/TrafficRulesApi.md#updatetrafficrules) | **Put** /api/v1/trafficRules | Edit Traffic Rules


## Documentation For Models

 - [ActiveFullSettingsModel](docs/ActiveFullSettingsModel.md)
 - [AdvancedSmtpOptionsModel](docs/AdvancedSmtpOptionsModel.md)
 - [AdvancedStorageScheduleMonthlyModel](docs/AdvancedStorageScheduleMonthlyModel.md)
 - [AdvancedStorageScheduleWeeklyModel](docs/AdvancedStorageScheduleWeeklyModel.md)
 - [ArchiveTierAdvancedSettingsModel](docs/ArchiveTierAdvancedSettingsModel.md)
 - [ArchiveTierModel](docs/ArchiveTierModel.md)
 - [AuthorizationCodeModel](docs/AuthorizationCodeModel.md)
 - [BackupApplicationAwareProcessingImportModel](docs/BackupApplicationAwareProcessingImportModel.md)
 - [BackupApplicationAwareProcessingModel](docs/BackupApplicationAwareProcessingModel.md)
 - [BackupApplicationSettingsImportModel](docs/BackupApplicationSettingsImportModel.md)
 - [BackupApplicationSettingsModel](docs/BackupApplicationSettingsModel.md)
 - [BackupFSExclusionsModel](docs/BackupFSExclusionsModel.md)
 - [BackupHealthCheckSettingsModels](docs/BackupHealthCheckSettingsModels.md)
 - [BackupIndexingSettingsModel](docs/BackupIndexingSettingsModel.md)
 - [BackupJobAdvancedSettingsModel](docs/BackupJobAdvancedSettingsModel.md)
 - [BackupJobAdvancedSettingsVSphereModel](docs/BackupJobAdvancedSettingsVSphereModel.md)
 - [BackupJobExclusions](docs/BackupJobExclusions.md)
 - [BackupJobExclusionsSpec](docs/BackupJobExclusionsSpec.md)
 - [BackupJobExclusionsTemplates](docs/BackupJobExclusionsTemplates.md)
 - [BackupJobGuestProcessingImportModel](docs/BackupJobGuestProcessingImportModel.md)
 - [BackupJobGuestProcessingModel](docs/BackupJobGuestProcessingModel.md)
 - [BackupJobImportProxiesModel](docs/BackupJobImportProxiesModel.md)
 - [BackupJobModel](docs/BackupJobModel.md)
 - [BackupJobModelAllOf](docs/BackupJobModelAllOf.md)
 - [BackupJobRetentionPolicySettingsModel](docs/BackupJobRetentionPolicySettingsModel.md)
 - [BackupJobSpec](docs/BackupJobSpec.md)
 - [BackupJobSpecAllOf](docs/BackupJobSpecAllOf.md)
 - [BackupJobStorageImportModel](docs/BackupJobStorageImportModel.md)
 - [BackupJobStorageModel](docs/BackupJobStorageModel.md)
 - [BackupJobVirtualMachinesModel](docs/BackupJobVirtualMachinesModel.md)
 - [BackupJobVirtualMachinesSpec](docs/BackupJobVirtualMachinesSpec.md)
 - [BackupLinuxScriptModel](docs/BackupLinuxScriptModel.md)
 - [BackupLogShippingServersImportModel](docs/BackupLogShippingServersImportModel.md)
 - [BackupLogShippingServersModel](docs/BackupLogShippingServersModel.md)
 - [BackupModel](docs/BackupModel.md)
 - [BackupObjectIndexingModel](docs/BackupObjectIndexingModel.md)
 - [BackupObjectModel](docs/BackupObjectModel.md)
 - [BackupObjectsFilters](docs/BackupObjectsFilters.md)
 - [BackupObjectsResult](docs/BackupObjectsResult.md)
 - [BackupOracleSettingsImportModel](docs/BackupOracleSettingsImportModel.md)
 - [BackupOracleSettingsModel](docs/BackupOracleSettingsModel.md)
 - [BackupPlacementSettingsModel](docs/BackupPlacementSettingsModel.md)
 - [BackupProxiesSettingsModel](docs/BackupProxiesSettingsModel.md)
 - [BackupProxyImportModel](docs/BackupProxyImportModel.md)
 - [BackupRepositoryImportModel](docs/BackupRepositoryImportModel.md)
 - [BackupSQLSettingsImportModel](docs/BackupSQLSettingsImportModel.md)
 - [BackupSQLSettingsModel](docs/BackupSQLSettingsModel.md)
 - [BackupScheduleModel](docs/BackupScheduleModel.md)
 - [BackupScriptSettingsModel](docs/BackupScriptSettingsModel.md)
 - [BackupStorageSettingModel](docs/BackupStorageSettingModel.md)
 - [BackupStorageSettingsEncryptionModel](docs/BackupStorageSettingsEncryptionModel.md)
 - [BackupWindowDayHoursModel](docs/BackupWindowDayHoursModel.md)
 - [BackupWindowSettingModel](docs/BackupWindowSettingModel.md)
 - [BackupWindowsScriptModel](docs/BackupWindowsScriptModel.md)
 - [BackupsFilters](docs/BackupsFilters.md)
 - [BackupsResult](docs/BackupsResult.md)
 - [CapacityTierModel](docs/CapacityTierModel.md)
 - [CapacityTierOverridePolicyModel](docs/CapacityTierOverridePolicyModel.md)
 - [CertificateModel](docs/CertificateModel.md)
 - [ConfigBackupEncryptionModel](docs/ConfigBackupEncryptionModel.md)
 - [ConfigBackupLastSuccessfulModel](docs/ConfigBackupLastSuccessfulModel.md)
 - [ConfigBackupModel](docs/ConfigBackupModel.md)
 - [ConfigBackupNotificationsModel](docs/ConfigBackupNotificationsModel.md)
 - [ConfigBackupSMTPSettigsModel](docs/ConfigBackupSMTPSettigsModel.md)
 - [ConfigBackupScheduleModel](docs/ConfigBackupScheduleModel.md)
 - [ConnectionCertificateModel](docs/ConnectionCertificateModel.md)
 - [CredentialsExportSpec](docs/CredentialsExportSpec.md)
 - [CredentialsFilters](docs/CredentialsFilters.md)
 - [CredentialsImportModel](docs/CredentialsImportModel.md)
 - [CredentialsImportSpec](docs/CredentialsImportSpec.md)
 - [CredentialsImportSpecCollection](docs/CredentialsImportSpecCollection.md)
 - [CredentialsLinuxSettingsImportModel](docs/CredentialsLinuxSettingsImportModel.md)
 - [CredentialsModel](docs/CredentialsModel.md)
 - [CredentialsPasswordChangeSpec](docs/CredentialsPasswordChangeSpec.md)
 - [CredentialsResult](docs/CredentialsResult.md)
 - [CredentialsSpec](docs/CredentialsSpec.md)
 - [DeleteRepositoryFilters](docs/DeleteRepositoryFilters.md)
 - [EAllowedBackupsType](docs/EAllowedBackupsType.md)
 - [EApplicationSettingsVSS](docs/EApplicationSettingsVSS.md)
 - [EBackupExclusionPolicy](docs/EBackupExclusionPolicy.md)
 - [EBackupModeType](docs/EBackupModeType.md)
 - [EBackupObjectsFiltersOrderColumn](docs/EBackupObjectsFiltersOrderColumn.md)
 - [EBackupOracleLogsSettings](docs/EBackupOracleLogsSettings.md)
 - [EBackupProxyImportType](docs/EBackupProxyImportType.md)
 - [EBackupProxyTransportMode](docs/EBackupProxyTransportMode.md)
 - [EBackupScriptProcessingMode](docs/EBackupScriptProcessingMode.md)
 - [EBackupsFiltersOrderColumn](docs/EBackupsFiltersOrderColumn.md)
 - [ECompressionLevel](docs/ECompressionLevel.md)
 - [EConfigBackupSMTPSettingsType](docs/EConfigBackupSMTPSettingsType.md)
 - [ECredentialsFiltersOrderColumn](docs/ECredentialsFiltersOrderColumn.md)
 - [ECredentialsType](docs/ECredentialsType.md)
 - [EDailyKinds](docs/EDailyKinds.md)
 - [EDayNumberInMonth](docs/EDayNumberInMonth.md)
 - [EDayOfWeek](docs/EDayOfWeek.md)
 - [EDiskInfoProcessState](docs/EDiskInfoProcessState.md)
 - [EDiskInfoType](docs/EDiskInfoType.md)
 - [EEmailNotificationType](docs/EEmailNotificationType.md)
 - [EEncryptionPasswordsFiltersOrderColumn](docs/EEncryptionPasswordsFiltersOrderColumn.md)
 - [EGuestFSIndexingMode](docs/EGuestFSIndexingMode.md)
 - [EGuestOSType](docs/EGuestOSType.md)
 - [EHierarchyType](docs/EHierarchyType.md)
 - [EInstantRecoveryMountState](docs/EInstantRecoveryMountState.md)
 - [EJobFiltersOrderColumn](docs/EJobFiltersOrderColumn.md)
 - [EJobStatesFiltersOrderColumn](docs/EJobStatesFiltersOrderColumn.md)
 - [EJobStatus](docs/EJobStatus.md)
 - [EJobType](docs/EJobType.md)
 - [EJobWorkload](docs/EJobWorkload.md)
 - [ELoginGrantType](docs/ELoginGrantType.md)
 - [EManagedServerType](docs/EManagedServerType.md)
 - [EManagedServersFiltersOrderColumn](docs/EManagedServersFiltersOrderColumn.md)
 - [EMonth](docs/EMonth.md)
 - [EObjectRestorePointOperation](docs/EObjectRestorePointOperation.md)
 - [EObjectRestorePointsFiltersOrderColumn](docs/EObjectRestorePointsFiltersOrderColumn.md)
 - [EPeriodicallyKinds](docs/EPeriodicallyKinds.md)
 - [EPlacementPolicyType](docs/EPlacementPolicyType.md)
 - [EPlatform](docs/EPlatform.md)
 - [EPlatformType](docs/EPlatformType.md)
 - [EProxiesFiltersOrderColumn](docs/EProxiesFiltersOrderColumn.md)
 - [EProxyType](docs/EProxyType.md)
 - [ERepositoryExtentStatusType](docs/ERepositoryExtentStatusType.md)
 - [ERepositoryFiltersOrderColumn](docs/ERepositoryFiltersOrderColumn.md)
 - [ERepositoryStatesFiltersOrderColumn](docs/ERepositoryStatesFiltersOrderColumn.md)
 - [ERepositoryType](docs/ERepositoryType.md)
 - [ERetainLogBackupsType](docs/ERetainLogBackupsType.md)
 - [ERetentionPolicyType](docs/ERetentionPolicyType.md)
 - [ESQLLogsProcessing](docs/ESQLLogsProcessing.md)
 - [EScaleOutRepositoryFiltersOrderColumn](docs/EScaleOutRepositoryFiltersOrderColumn.md)
 - [EScriptPeriodicityType](docs/EScriptPeriodicityType.md)
 - [ESennightOfMonth](docs/ESennightOfMonth.md)
 - [EServicesFiltersOrderColumn](docs/EServicesFiltersOrderColumn.md)
 - [ESessionResult](docs/ESessionResult.md)
 - [ESessionState](docs/ESessionState.md)
 - [ESessionType](docs/ESessionType.md)
 - [ESessionsFiltersOrderColumn](docs/ESessionsFiltersOrderColumn.md)
 - [ESpeedUnit](docs/ESpeedUnit.md)
 - [EStorageOptimization](docs/EStorageOptimization.md)
 - [ETaskLogRecordStatus](docs/ETaskLogRecordStatus.md)
 - [ETransactionLogsSettings](docs/ETransactionLogsSettings.md)
 - [EViHostType](docs/EViHostType.md)
 - [EViRootFiltersOrderColumn](docs/EViRootFiltersOrderColumn.md)
 - [EVmwareDisksTypeToProcess](docs/EVmwareDisksTypeToProcess.md)
 - [EVmwareFcdInstantRecoveryMountsFiltersOrderColumn](docs/EVmwareFcdInstantRecoveryMountsFiltersOrderColumn.md)
 - [EVmwareInventoryType](docs/EVmwareInventoryType.md)
 - [EWindowsHostComponentType](docs/EWindowsHostComponentType.md)
 - [EmailCustomNotificationType](docs/EmailCustomNotificationType.md)
 - [EmailNotificationSettingsModel](docs/EmailNotificationSettingsModel.md)
 - [EncryptionPasswordExportSpec](docs/EncryptionPasswordExportSpec.md)
 - [EncryptionPasswordImportSpec](docs/EncryptionPasswordImportSpec.md)
 - [EncryptionPasswordImportSpecCollection](docs/EncryptionPasswordImportSpecCollection.md)
 - [EncryptionPasswordModel](docs/EncryptionPasswordModel.md)
 - [EncryptionPasswordSpec](docs/EncryptionPasswordSpec.md)
 - [EncryptionPasswordsFilters](docs/EncryptionPasswordsFilters.md)
 - [EncryptionPasswordsResult](docs/EncryptionPasswordsResult.md)
 - [Error](docs/Error.md)
 - [EvCentersInventoryFiltersOrderColumn](docs/EvCentersInventoryFiltersOrderColumn.md)
 - [FullBackupMaintenanceDefragmentAndCompactModel](docs/FullBackupMaintenanceDefragmentAndCompactModel.md)
 - [FullBackupMaintenanceModel](docs/FullBackupMaintenanceModel.md)
 - [FullBackupMaintenanceRemoveDataModel](docs/FullBackupMaintenanceRemoveDataModel.md)
 - [GFSPolicySettingsModel](docs/GFSPolicySettingsModel.md)
 - [GFSPolicySettingsMonthlyModel](docs/GFSPolicySettingsMonthlyModel.md)
 - [GFSPolicySettingsWeeklyModel](docs/GFSPolicySettingsWeeklyModel.md)
 - [GFSPolicySettingsYearlyModel](docs/GFSPolicySettingsYearlyModel.md)
 - [GeneralOptionsEmailNotificationsModel](docs/GeneralOptionsEmailNotificationsModel.md)
 - [GeneralOptionsModel](docs/GeneralOptionsModel.md)
 - [GeneralOptionsNotificationsModel](docs/GeneralOptionsNotificationsModel.md)
 - [GlobalNetworkTrafficRulesModel](docs/GlobalNetworkTrafficRulesModel.md)
 - [GuestFileSystemIndexingModel](docs/GuestFileSystemIndexingModel.md)
 - [GuestInteractionProxiesSettingsImportModel](docs/GuestInteractionProxiesSettingsImportModel.md)
 - [GuestInteractionProxiesSettingsModel](docs/GuestInteractionProxiesSettingsModel.md)
 - [GuestOsCredentialsImportModel](docs/GuestOsCredentialsImportModel.md)
 - [GuestOsCredentialsModel](docs/GuestOsCredentialsModel.md)
 - [GuestOsCredentialsPerMachineImportModel](docs/GuestOsCredentialsPerMachineImportModel.md)
 - [GuestOsCredentialsPerMachineModel](docs/GuestOsCredentialsPerMachineModel.md)
 - [HostConnectionSpec](docs/HostConnectionSpec.md)
 - [JobExportSpec](docs/JobExportSpec.md)
 - [JobImportSpec](docs/JobImportSpec.md)
 - [JobImportSpecCollection](docs/JobImportSpecCollection.md)
 - [JobModel](docs/JobModel.md)
 - [JobScriptsSettingsModel](docs/JobScriptsSettingsModel.md)
 - [JobSpec](docs/JobSpec.md)
 - [JobStartSpec](docs/JobStartSpec.md)
 - [JobStateModel](docs/JobStateModel.md)
 - [JobStatesFilters](docs/JobStatesFilters.md)
 - [JobStatesResult](docs/JobStatesResult.md)
 - [JobStopSpec](docs/JobStopSpec.md)
 - [JobsFilters](docs/JobsFilters.md)
 - [JobsResult](docs/JobsResult.md)
 - [LinuxCredentialsModel](docs/LinuxCredentialsModel.md)
 - [LinuxCredentialsModelAllOf](docs/LinuxCredentialsModelAllOf.md)
 - [LinuxCredentialsSpec](docs/LinuxCredentialsSpec.md)
 - [LinuxCredentialsSpecAllOf](docs/LinuxCredentialsSpecAllOf.md)
 - [LinuxHostImportSpec](docs/LinuxHostImportSpec.md)
 - [LinuxHostModel](docs/LinuxHostModel.md)
 - [LinuxHostModelAllOf](docs/LinuxHostModelAllOf.md)
 - [LinuxHostSSHSettingsModel](docs/LinuxHostSSHSettingsModel.md)
 - [LinuxHostSpec](docs/LinuxHostSpec.md)
 - [LinuxLocalRepositorySettingsModel](docs/LinuxLocalRepositorySettingsModel.md)
 - [LinuxLocalStorageImportSpec](docs/LinuxLocalStorageImportSpec.md)
 - [LinuxLocalStorageModel](docs/LinuxLocalStorageModel.md)
 - [LinuxLocalStorageSpec](docs/LinuxLocalStorageSpec.md)
 - [LinuxLocalStorageSpecAllOf](docs/LinuxLocalStorageSpecAllOf.md)
 - [ManageServerExportSpec](docs/ManageServerExportSpec.md)
 - [ManageServerImportSpecCollection](docs/ManageServerImportSpecCollection.md)
 - [ManagedServerModel](docs/ManagedServerModel.md)
 - [ManagedServerSpec](docs/ManagedServerSpec.md)
 - [ManagedServersFilters](docs/ManagedServersFilters.md)
 - [ManagedServersResult](docs/ManagedServersResult.md)
 - [MountServerSettingsImportSpec](docs/MountServerSettingsImportSpec.md)
 - [MountServerSettingsModel](docs/MountServerSettingsModel.md)
 - [NetworkRepositorySettingsModel](docs/NetworkRepositorySettingsModel.md)
 - [NfsRepositoryShareSettingsModel](docs/NfsRepositoryShareSettingsModel.md)
 - [NfsRepositoryShareSettingsSpec](docs/NfsRepositoryShareSettingsSpec.md)
 - [NfsStorageImportSpec](docs/NfsStorageImportSpec.md)
 - [NfsStorageModel](docs/NfsStorageModel.md)
 - [NfsStorageSpec](docs/NfsStorageSpec.md)
 - [NfsStorageSpecAllOf](docs/NfsStorageSpecAllOf.md)
 - [NotificationSettingsModel](docs/NotificationSettingsModel.md)
 - [NotificationVmAttributeSettingsModel](docs/NotificationVmAttributeSettingsModel.md)
 - [ObjectRestorePointDiskModel](docs/ObjectRestorePointDiskModel.md)
 - [ObjectRestorePointDisksResult](docs/ObjectRestorePointDisksResult.md)
 - [ObjectRestorePointModel](docs/ObjectRestorePointModel.md)
 - [ObjectRestorePointsFilters](docs/ObjectRestorePointsFilters.md)
 - [ObjectRestorePointsResult](docs/ObjectRestorePointsResult.md)
 - [PaginationResult](docs/PaginationResult.md)
 - [PerformanceExtentModel](docs/PerformanceExtentModel.md)
 - [PerformanceTierAdvancedSettingsModel](docs/PerformanceTierAdvancedSettingsModel.md)
 - [PerformanceTierModel](docs/PerformanceTierModel.md)
 - [PlacementPolicyModel](docs/PlacementPolicyModel.md)
 - [PreferredNetworkModel](docs/PreferredNetworkModel.md)
 - [PreferredNetworksModel](docs/PreferredNetworksModel.md)
 - [PrimaryStorageIntegrationSettingsModel](docs/PrimaryStorageIntegrationSettingsModel.md)
 - [PrivateKeyChangeSpec](docs/PrivateKeyChangeSpec.md)
 - [ProxiesFilters](docs/ProxiesFilters.md)
 - [ProxiesResult](docs/ProxiesResult.md)
 - [ProxyDatastoreModel](docs/ProxyDatastoreModel.md)
 - [ProxyDatastoreSettingsModel](docs/ProxyDatastoreSettingsModel.md)
 - [ProxyExportSpec](docs/ProxyExportSpec.md)
 - [ProxyImportSpec](docs/ProxyImportSpec.md)
 - [ProxyImportSpecCollection](docs/ProxyImportSpecCollection.md)
 - [ProxyModel](docs/ProxyModel.md)
 - [ProxyServerSettingsImportSpec](docs/ProxyServerSettingsImportSpec.md)
 - [ProxyServerSettingsModel](docs/ProxyServerSettingsModel.md)
 - [ProxySpec](docs/ProxySpec.md)
 - [RepositoriesFilters](docs/RepositoriesFilters.md)
 - [RepositoriesResult](docs/RepositoriesResult.md)
 - [RepositoryAdvancedSettingsModel](docs/RepositoryAdvancedSettingsModel.md)
 - [RepositoryExportSpec](docs/RepositoryExportSpec.md)
 - [RepositoryImportSpecCollection](docs/RepositoryImportSpecCollection.md)
 - [RepositoryModel](docs/RepositoryModel.md)
 - [RepositoryShareGatewayImportSpec](docs/RepositoryShareGatewayImportSpec.md)
 - [RepositoryShareGatewayModel](docs/RepositoryShareGatewayModel.md)
 - [RepositorySpec](docs/RepositorySpec.md)
 - [RepositoryStateModel](docs/RepositoryStateModel.md)
 - [RepositoryStatesFilters](docs/RepositoryStatesFilters.md)
 - [RepositoryStatesResult](docs/RepositoryStatesResult.md)
 - [ScaleOutExtentMaintenanceSpec](docs/ScaleOutExtentMaintenanceSpec.md)
 - [ScaleOutRepositoriesFilters](docs/ScaleOutRepositoriesFilters.md)
 - [ScaleOutRepositoriesResult](docs/ScaleOutRepositoriesResult.md)
 - [ScaleOutRepositoryModel](docs/ScaleOutRepositoryModel.md)
 - [ScheduleAfterThisJobModel](docs/ScheduleAfterThisJobModel.md)
 - [ScheduleBackupWindowModel](docs/ScheduleBackupWindowModel.md)
 - [ScheduleDailyModel](docs/ScheduleDailyModel.md)
 - [ScheduleMonthlyModel](docs/ScheduleMonthlyModel.md)
 - [SchedulePeriodicallyModel](docs/SchedulePeriodicallyModel.md)
 - [ScheduleRetryModel](docs/ScheduleRetryModel.md)
 - [ScriptCommand](docs/ScriptCommand.md)
 - [ServerTimeModel](docs/ServerTimeModel.md)
 - [ServicesFilters](docs/ServicesFilters.md)
 - [ServicesModel](docs/ServicesModel.md)
 - [ServicesResult](docs/ServicesResult.md)
 - [SessionLogRecordModel](docs/SessionLogRecordModel.md)
 - [SessionLogResult](docs/SessionLogResult.md)
 - [SessionModel](docs/SessionModel.md)
 - [SessionResultModel](docs/SessionResultModel.md)
 - [SessionsFilters](docs/SessionsFilters.md)
 - [SessionsResult](docs/SessionsResult.md)
 - [SmbRepositoryShareSettingsModel](docs/SmbRepositoryShareSettingsModel.md)
 - [SmbRepositoryShareSettingsSpec](docs/SmbRepositoryShareSettingsSpec.md)
 - [SmbStorageImportSpec](docs/SmbStorageImportSpec.md)
 - [SmbStorageModel](docs/SmbStorageModel.md)
 - [SmbStorageSpec](docs/SmbStorageSpec.md)
 - [SmbStorageSpecAllOf](docs/SmbStorageSpecAllOf.md)
 - [StandardCredentialsModel](docs/StandardCredentialsModel.md)
 - [StandardCredentialsSpec](docs/StandardCredentialsSpec.md)
 - [StandardCredentialsSpecAllOf](docs/StandardCredentialsSpecAllOf.md)
 - [SyntheticFullSettingsModel](docs/SyntheticFullSettingsModel.md)
 - [TokenLoginSpec](docs/TokenLoginSpec.md)
 - [TokenModel](docs/TokenModel.md)
 - [TrafficRuleModel](docs/TrafficRuleModel.md)
 - [VCenterInventoryFilters](docs/VCenterInventoryFilters.md)
 - [VCenterInventoryResult](docs/VCenterInventoryResult.md)
 - [VPowerNFSPortSettingsModel](docs/VPowerNFSPortSettingsModel.md)
 - [VSphereChangedBlockTrackingSettingsModel](docs/VSphereChangedBlockTrackingSettingsModel.md)
 - [ViBackupObjectModel](docs/ViBackupObjectModel.md)
 - [ViBackupObjectModelAllOf](docs/ViBackupObjectModelAllOf.md)
 - [ViHostImportSpec](docs/ViHostImportSpec.md)
 - [ViHostModel](docs/ViHostModel.md)
 - [ViHostModelAllOf](docs/ViHostModelAllOf.md)
 - [ViHostSpec](docs/ViHostSpec.md)
 - [ViHostSpecAllOf](docs/ViHostSpecAllOf.md)
 - [ViProxyModel](docs/ViProxyModel.md)
 - [ViProxySpec](docs/ViProxySpec.md)
 - [ViProxySpecAllOf](docs/ViProxySpecAllOf.md)
 - [ViRootFilters](docs/ViRootFilters.md)
 - [ViRootsResult](docs/ViRootsResult.md)
 - [VmbApiFilterModel](docs/VmbApiFilterModel.md)
 - [VmwareFcdInstantRecoveryDiskInfo](docs/VmwareFcdInstantRecoveryDiskInfo.md)
 - [VmwareFcdInstantRecoveryDiskSpec](docs/VmwareFcdInstantRecoveryDiskSpec.md)
 - [VmwareFcdInstantRecoveryMount](docs/VmwareFcdInstantRecoveryMount.md)
 - [VmwareFcdInstantRecoveryMountsFilters](docs/VmwareFcdInstantRecoveryMountsFilters.md)
 - [VmwareFcdInstantRecoveryMountsResult](docs/VmwareFcdInstantRecoveryMountsResult.md)
 - [VmwareFcdInstantRecoverySpec](docs/VmwareFcdInstantRecoverySpec.md)
 - [VmwareFcdQuickMigrationSpec](docs/VmwareFcdQuickMigrationSpec.md)
 - [VmwareFcdWriteCacheSpec](docs/VmwareFcdWriteCacheSpec.md)
 - [VmwareObjectDiskModel](docs/VmwareObjectDiskModel.md)
 - [VmwareObjectModel](docs/VmwareObjectModel.md)
 - [VmwareObjectSizeModel](docs/VmwareObjectSizeModel.md)
 - [WindowsHostComponentPortModel](docs/WindowsHostComponentPortModel.md)
 - [WindowsHostImportSpec](docs/WindowsHostImportSpec.md)
 - [WindowsHostModel](docs/WindowsHostModel.md)
 - [WindowsHostPortsModel](docs/WindowsHostPortsModel.md)
 - [WindowsHostSpec](docs/WindowsHostSpec.md)
 - [WindowsHostSpecAllOf](docs/WindowsHostSpecAllOf.md)
 - [WindowsLocalRepositorySettingsModel](docs/WindowsLocalRepositorySettingsModel.md)
 - [WindowsLocalStorageImportSpec](docs/WindowsLocalStorageImportSpec.md)
 - [WindowsLocalStorageModel](docs/WindowsLocalStorageModel.md)
 - [WindowsLocalStorageSpec](docs/WindowsLocalStorageSpec.md)
 - [WindowsLocalStorageSpecAllOf](docs/WindowsLocalStorageSpecAllOf.md)


## Documentation For Authorization



### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@veeam.com
