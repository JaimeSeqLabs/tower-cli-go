# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**WorkflowStatus**](WorkflowStatus.md) |  | [optional] 
**OwnerId** | **int64** |  | [optional] [readonly] 
**Repository** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Submit** | [**time.Time**](time.Time.md) |  | 
**Start** | [**time.Time**](time.Time.md) |  | [optional] 
**Complete** | [**time.Time**](time.Time.md) |  | [optional] 
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**RunName** | **string** |  | 
**SessionId** | **string** |  | 
**Profile** | **string** |  | [optional] 
**WorkDir** | **string** |  | 
**CommitId** | **string** |  | [optional] 
**UserName** | **string** |  | 
**ScriptId** | **string** |  | [optional] 
**Revision** | **string** |  | [optional] 
**CommandLine** | **string** |  | 
**ProjectName** | **string** |  | 
**ScriptName** | **string** |  | [optional] 
**LaunchId** | **string** |  | [optional] 
**ConfigFiles** | **[]string** |  | [optional] 
**Params** | **map[string]map[string]interface{}** |  | [optional] 
**ConfigText** | **string** |  | [optional] 
**Manifest** | [**WfManifest**](WfManifest.md) |  | [optional] 
**Nextflow** | [**WfNextflow**](WfNextflow.md) |  | [optional] 
**Stats** | [**WfStats**](WfStats.md) |  | [optional] 
**ErrorMessage** | **string** |  | [optional] 
**ErrorReport** | **string** |  | [optional] 
**Deleted** | **bool** |  | [optional] [readonly] 
**ProjectDir** | **string** |  | [optional] 
**HomeDir** | **string** |  | [optional] 
**Container** | **string** |  | [optional] 
**ContainerEngine** | **string** |  | [optional] 
**ScriptFile** | **string** |  | [optional] 
**LaunchDir** | **string** |  | [optional] 
**Duration** | **int64** |  | [optional] 
**ExitStatus** | **int32** |  | [optional] 
**Resume** | **bool** |  | [optional] 
**Success** | **bool** |  | [optional] 
**LogFile** | **string** |  | [optional] 
**OutFile** | **string** |  | [optional] 
**OperationId** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


