# GoogleLifeSciencesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** |  | [optional] 
**Zones** | **[]string** |  | [optional] 
**Location** | **string** |  | [optional] 
**WorkDir** | **string** |  | [optional] 
**Preemptible** | **bool** |  | [optional] 
**BootDiskSizeGb** | **int32** |  | [optional] 
**ProjectId** | **string** |  | [optional] 
**SshDaemon** | **bool** |  | [optional] 
**SshImage** | **string** |  | [optional] 
**DebugMode** | **int32** |  | [optional] 
**CopyImage** | **string** |  | [optional] 
**UsePrivateAddress** | **bool** |  | [optional] 
**Labels** | **map[string]string** |  | [optional] 
**PreRunScript** | **string** |  | [optional] 
**PostRunScript** | **string** |  | [optional] 
**HeadJobCpus** | **int32** |  | [optional] 
**HeadJobMemoryMb** | **int32** |  | [optional] 
**NfsTarget** | **string** |  | [optional] 
**NfsMount** | **string** |  | [optional] 
**Environment** | [**[]ConfigEnvVariable**](ConfigEnvVariable.md) |  | [optional] 
**Discriminator** | **string** | property to select the compute config platform | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


