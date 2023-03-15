# AwsBatchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** |  | [optional] 
**ComputeQueue** | **string** |  | [optional] 
**DragenQueue** | **string** |  | [optional] 
**ComputeJobRole** | **string** |  | [optional] 
**ExecutionRole** | **string** |  | [optional] 
**HeadQueue** | **string** |  | [optional] 
**HeadJobRole** | **string** |  | [optional] 
**CliPath** | **string** |  | [optional] 
**Volumes** | **[]string** |  | [optional] 
**WorkDir** | **string** |  | [optional] 
**PreRunScript** | **string** |  | [optional] 
**PostRunScript** | **string** |  | [optional] 
**HeadJobCpus** | **int32** |  | [optional] 
**HeadJobMemoryMb** | **int32** |  | [optional] 
**Environment** | [**[]ConfigEnvVariable**](ConfigEnvVariable.md) |  | [optional] 
**WaveEnabled** | **bool** |  | [optional] 
**Fusion2Enabled** | **bool** |  | [optional] 
**NvnmeStorageEnabled** | **bool** |  | [optional] 
**LogsGroup** | **string** |  | [optional] 
**Forge** | [**ForgeConfig**](ForgeConfig.md) |  | [optional] 
**ForgedResources** | [**[]map[string]map[string]interface{}**](map.md) |  | [optional] 
**Discriminator** | **string** | property to select the compute config platform | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


