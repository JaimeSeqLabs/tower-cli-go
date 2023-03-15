# SlurmComputeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkDir** | **string** |  | [optional] 
**LaunchDir** | **string** |  | [optional] 
**UserName** | **string** |  | [optional] 
**HostName** | **string** |  | [optional] 
**Port** | **int32** |  | [optional] 
**HeadQueue** | **string** |  | [optional] 
**ComputeQueue** | **string** |  | [optional] 
**MaxQueueSize** | **int32** |  | [optional] 
**HeadJobOptions** | **string** |  | [optional] 
**PropagateHeadJobOptions** | **bool** |  | [optional] 
**PreRunScript** | **string** |  | [optional] 
**PostRunScript** | **string** |  | [optional] 
**Environment** | [**[]ConfigEnvVariable**](ConfigEnvVariable.md) |  | [optional] 
**Discriminator** | **string** | property to select the compute config platform | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


