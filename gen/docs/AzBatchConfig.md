# AzBatchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkDir** | **string** |  | [optional] 
**PreRunScript** | **string** |  | [optional] 
**PostRunScript** | **string** |  | [optional] 
**Region** | **string** |  | [optional] 
**HeadPool** | **string** |  | [optional] 
**AutoPoolMode** | **bool** |  | [optional] 
**Forge** | [**AzBatchForgeConfig**](AzBatchForgeConfig.md) |  | [optional] 
**TokenDuration** | **string** |  | [optional] 
**DeleteJobsOnCompletion** | [**JobCleanupPolicy**](JobCleanupPolicy.md) |  | [optional] 
**DeletePoolsOnCompletion** | **bool** |  | [optional] 
**Environment** | [**[]ConfigEnvVariable**](ConfigEnvVariable.md) |  | [optional] 
**Discriminator** | **string** | property to select the compute config platform | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


