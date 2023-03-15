# EksComputeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkDir** | **string** |  | [optional] 
**PreRunScript** | **string** |  | [optional] 
**PostRunScript** | **string** |  | [optional] 
**Server** | **string** |  | [optional] 
**SslCert** | **string** |  | [optional] 
**Namespace** | **string** |  | [optional] 
**ComputeServiceAccount** | **string** |  | [optional] 
**HeadServiceAccount** | **string** |  | [optional] 
**StorageClaimName** | **string** |  | [optional] 
**StorageMountPath** | **string** |  | [optional] 
**PodCleanup** | [**PodCleanupPolicy**](PodCleanupPolicy.md) |  | [optional] 
**HeadPodSpec** | **string** |  | [optional] 
**ServicePodSpec** | **string** |  | [optional] 
**Environment** | [**[]ConfigEnvVariable**](ConfigEnvVariable.md) |  | [optional] 
**HeadJobCpus** | **int32** |  | [optional] 
**HeadJobMemoryMb** | **int32** |  | [optional] 
**Discriminator** | **string** | property to select the compute config platform | [optional] [readonly] 
**Region** | **string** | AWS region | [optional] 
**ClusterName** | **string** | The AWS EKS cluster name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


