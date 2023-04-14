# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLabelsToActions**](DefaultApi.md#AddLabelsToActions) | **Post** /actions/labels/add | Add some labels to some actions
[**AddLabelsToPipelines**](DefaultApi.md#AddLabelsToPipelines) | **Post** /pipelines/labels/add | Add some labels to some pipelines
[**AddLabelsToWorkflows**](DefaultApi.md#AddLabelsToWorkflows) | **Post** /workflow/labels/add | Add some labels to some workflows
[**ApplyLabelsToActions**](DefaultApi.md#ApplyLabelsToActions) | **Post** /actions/labels/apply | Apply some labels to some actions
[**ApplyLabelsToPipelines**](DefaultApi.md#ApplyLabelsToPipelines) | **Post** /pipelines/labels/apply | Apply some labels to some pipelines
[**ApplyLabelsToWorkflows**](DefaultApi.md#ApplyLabelsToWorkflows) | **Post** /workflow/labels/apply | Apply some labels to some workflows
[**CancelWorkflow**](DefaultApi.md#CancelWorkflow) | **Post** /workflow/{workflowId}/cancel | Cancel a Workflow execution
[**CreateAction**](DefaultApi.md#CreateAction) | **Post** /actions | Create a new Pipeline Action
[**CreateAvatar**](DefaultApi.md#CreateAvatar) | **Post** /avatars | Create the avatar image
[**CreateComputeEnv**](DefaultApi.md#CreateComputeEnv) | **Post** /compute-envs | Create a new Tower compute environment
[**CreateCredentials**](DefaultApi.md#CreateCredentials) | **Post** /credentials | Create a new credentials record
[**CreateDataset**](DefaultApi.md#CreateDataset) | **Post** /workspaces/{workspaceId}/datasets | Create a dataset
[**CreateLabel**](DefaultApi.md#CreateLabel) | **Post** /labels | Create a new label
[**CreateOrganization**](DefaultApi.md#CreateOrganization) | **Post** /orgs | Create a new organization
[**CreateOrganizationMember**](DefaultApi.md#CreateOrganizationMember) | **Put** /orgs/{orgId}/members/add | Add a new organization member
[**CreateOrganizationTeam**](DefaultApi.md#CreateOrganizationTeam) | **Post** /orgs/{orgId}/teams | Create a new organization team
[**CreateOrganizationTeamMember**](DefaultApi.md#CreateOrganizationTeamMember) | **Post** /orgs/{orgId}/teams/{teamId}/members | Add a new team member
[**CreatePipeline**](DefaultApi.md#CreatePipeline) | **Post** /pipelines | Create a new Pipeline in a workspace
[**CreatePipelineSecret**](DefaultApi.md#CreatePipelineSecret) | **Post** /pipeline-secrets | Create a Pipeline Secrets inside the given workspace
[**CreateToken**](DefaultApi.md#CreateToken) | **Post** /tokens | Create an API token
[**CreateTrace**](DefaultApi.md#CreateTrace) | **Post** /trace/create | Create a new Workflow execution trace
[**CreateWorkflowLaunch**](DefaultApi.md#CreateWorkflowLaunch) | **Post** /workflow/launch | Submit a Workflow execution
[**CreateWorkflowStar**](DefaultApi.md#CreateWorkflowStar) | **Post** /workflow/{workflowId}/star | Star a workflow
[**CreateWorkspace**](DefaultApi.md#CreateWorkspace) | **Post** /orgs/{orgId}/workspaces | Create a new workspace
[**CreateWorkspaceParticipant**](DefaultApi.md#CreateWorkspaceParticipant) | **Put** /orgs/{orgId}/workspaces/{workspaceId}/participants/add | Create a new workspace participant
[**DeleteAction**](DefaultApi.md#DeleteAction) | **Delete** /actions/{actionId} | Delete a Pipeline Action
[**DeleteAllTokens**](DefaultApi.md#DeleteAllTokens) | **Delete** /tokens/delete-all | Delete all user API tokens
[**DeleteComputeEnv**](DefaultApi.md#DeleteComputeEnv) | **Delete** /compute-envs/{computeEnvId} | Delete an existing Tower compute environment
[**DeleteCredentials**](DefaultApi.md#DeleteCredentials) | **Delete** /credentials/{credentialsId} | Delete the credentials record for the given id
[**DeleteDataset**](DefaultApi.md#DeleteDataset) | **Delete** /workspaces/{workspaceId}/datasets/{datasetId} | Delete a dataset
[**DeleteLabel**](DefaultApi.md#DeleteLabel) | **Delete** /labels/{labelId} | Delete a label
[**DeleteOrganization**](DefaultApi.md#DeleteOrganization) | **Delete** /orgs/{orgId} | Delete an organization
[**DeleteOrganizationMember**](DefaultApi.md#DeleteOrganizationMember) | **Delete** /orgs/{orgId}/members/{memberId} | Delete an organization member
[**DeleteOrganizationTeam**](DefaultApi.md#DeleteOrganizationTeam) | **Delete** /orgs/{orgId}/teams/{teamId} | Delete an organization team
[**DeleteOrganizationTeamMember**](DefaultApi.md#DeleteOrganizationTeamMember) | **Delete** /orgs/{orgId}/teams/{teamId}/members/{memberId}/delete | Delete a team member
[**DeletePipeline**](DefaultApi.md#DeletePipeline) | **Delete** /pipelines/{pipelineId} | Delete a Pipeline
[**DeletePipelineSecret**](DefaultApi.md#DeletePipelineSecret) | **Delete** /pipeline-secrets/{secretId} | Delete the pipeline secret record for the given id
[**DeleteToken**](DefaultApi.md#DeleteToken) | **Delete** /tokens/{tokenId} | Delete an API token
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /users/{userId} | Delete a user entity
[**DeleteWorkflow**](DefaultApi.md#DeleteWorkflow) | **Delete** /workflow/{workflowId} | Delete the Workflow entity with the given ID
[**DeleteWorkflowMany**](DefaultApi.md#DeleteWorkflowMany) | **Post** /workflow/delete | Delete several workflow entities given their ids
[**DeleteWorkflowStar**](DefaultApi.md#DeleteWorkflowStar) | **Delete** /workflow/{workflowId}/star | Unstar a workflow
[**DeleteWorkspace**](DefaultApi.md#DeleteWorkspace) | **Delete** /orgs/{orgId}/workspaces/{workspaceId} | Delete a workspace
[**DeleteWorkspaceParticipant**](DefaultApi.md#DeleteWorkspaceParticipant) | **Delete** /orgs/{orgId}/workspaces/{workspaceId}/participants/{participantId} | Delete a workspace participant
[**DescribeAction**](DefaultApi.md#DescribeAction) | **Get** /actions/{actionId} | Describe an existing Pipeline Action
[**DescribeComputeEnv**](DefaultApi.md#DescribeComputeEnv) | **Get** /compute-envs/{computeEnvId} | Describe a Tower compute environment
[**DescribeCredentials**](DefaultApi.md#DescribeCredentials) | **Get** /credentials/{credentialsId} | Describe the credentials for the given id
[**DescribeDataset**](DefaultApi.md#DescribeDataset) | **Get** /workspaces/{workspaceId}/datasets/{datasetId}/metadata | Describe a dataset
[**DescribeLaunch**](DefaultApi.md#DescribeLaunch) | **Get** /launch/{launchId} | Describe the Launch record for the given id
[**DescribeOrganization**](DefaultApi.md#DescribeOrganization) | **Get** /orgs/{orgId} | Describe organization details
[**DescribePipeline**](DefaultApi.md#DescribePipeline) | **Get** /pipelines/{pipelineId} | Describe a Pipeline
[**DescribePipelineLaunch**](DefaultApi.md#DescribePipelineLaunch) | **Get** /pipelines/{pipelineId}/launch | Describe a Pipeline launch
[**DescribePipelineRepository**](DefaultApi.md#DescribePipelineRepository) | **Get** /pipelines/info | Describe the Pipeline entity for the given id
[**DescribePipelineSchema**](DefaultApi.md#DescribePipelineSchema) | **Get** /pipelines/{pipelineId}/schema | Retrieve the Pipeline input schema
[**DescribePipelineSecret**](DefaultApi.md#DescribePipelineSecret) | **Get** /pipeline-secrets/{secretId} | Fetch a single pipeline secret
[**DescribePlatform**](DefaultApi.md#DescribePlatform) | **Get** /platforms/{platformId} | Describe the platform entity for the given id
[**DescribeUser**](DefaultApi.md#DescribeUser) | **Get** /users/{userId} | Describe a user entity
[**DescribeWorkflow**](DefaultApi.md#DescribeWorkflow) | **Get** /workflow/{workflowId} | Describe the Workflow entity for the given ID
[**DescribeWorkflowLaunch**](DefaultApi.md#DescribeWorkflowLaunch) | **Get** /workflow/{workflowId}/launch | Describe a Workflow launch
[**DescribeWorkflowMetrics**](DefaultApi.md#DescribeWorkflowMetrics) | **Get** /workflow/{workflowId}/metrics | Get the execution metrics for the given Workflow ID
[**DescribeWorkflowProgress**](DefaultApi.md#DescribeWorkflowProgress) | **Get** /workflow/{workflowId}/progress | Retrieve the execution progress for the given Workflow ID
[**DescribeWorkflowStar**](DefaultApi.md#DescribeWorkflowStar) | **Get** /workflow/{workflowId}/star | Check starred status of a workflow
[**DescribeWorkflowTask**](DefaultApi.md#DescribeWorkflowTask) | **Get** /workflow/{workflowId}/task/{taskId} | Describe a task entity with the given ID
[**DescribeWorkspace**](DefaultApi.md#DescribeWorkspace) | **Get** /orgs/{orgId}/workspaces/{workspaceId} | Describe a workspace
[**DownloadAvatar**](DefaultApi.md#DownloadAvatar) | **Get** /avatars/{avatarId} | Download the avatar image
[**DownloadDataset**](DefaultApi.md#DownloadDataset) | **Get** /workspaces/{workspaceId}/datasets/{datasetId}/v/{version}/n/{fileName} | Download the content of a dataset version
[**DownloadWorkflowLog**](DefaultApi.md#DownloadWorkflowLog) | **Get** /workflow/{workflowId}/download | Download Workflow files of the Nextflow main job
[**DownloadWorkflowTaskLog**](DefaultApi.md#DownloadWorkflowTaskLog) | **Get** /workflow/{workflowId}/download/{taskId} | Download Workflow files of a given task
[**GaRunCancel**](DefaultApi.md#GaRunCancel) | **Post** /ga4gh/wes/v1/runs/{run_id}/cancel | GA4GH cancel a run
[**GaRunCreate**](DefaultApi.md#GaRunCreate) | **Post** /ga4gh/wes/v1/runs | GA4GH create a new run
[**GaRunDescribe**](DefaultApi.md#GaRunDescribe) | **Get** /ga4gh/wes/v1/runs/{run_id} | GA4GH describe run
[**GaRunList**](DefaultApi.md#GaRunList) | **Get** /ga4gh/wes/v1/runs | GA4GH list runs
[**GaRunStatus**](DefaultApi.md#GaRunStatus) | **Get** /ga4gh/wes/v1/runs/{run_id}/status | GA4GH retrieve run status
[**GaServiceInfo**](DefaultApi.md#GaServiceInfo) | **Get** /ga4gh/wes/v1/service-info | GA4GH service info
[**GenerateRandomWorkflowName**](DefaultApi.md#GenerateRandomWorkflowName) | **Get** /workflow/random-name | Generates a random name
[**GetWorkflowTaskLog**](DefaultApi.md#GetWorkflowTaskLog) | **Get** /workflow/{workflowId}/log/{taskId} | Retrieve Workflow output logs of a given task
[**Info**](DefaultApi.md#Info) | **Get** /service-info | General Tower service features and version
[**LaunchAction**](DefaultApi.md#LaunchAction) | **Post** /actions/{actionId}/launch | Trigger the execution of a Tower Launch Action
[**LeaveOrganization**](DefaultApi.md#LeaveOrganization) | **Delete** /orgs/{orgId}/members/leave | Leave an organization
[**LeaveWorkspaceParticipant**](DefaultApi.md#LeaveWorkspaceParticipant) | **Delete** /orgs/{orgId}/workspaces/{workspaceId}/participants | Leave a workspace
[**ListActionTypes**](DefaultApi.md#ListActionTypes) | **Get** /actions/types | List the supported event types that can trigger a Pipeline Action
[**ListActions**](DefaultApi.md#ListActions) | **Get** /actions | List the available Pipeline Actions for the authenticated user or given workspace
[**ListComputeEnvs**](DefaultApi.md#ListComputeEnvs) | **Get** /compute-envs | List all Tower compute environments for the authenticated user or given workspace
[**ListCredentials**](DefaultApi.md#ListCredentials) | **Get** /credentials | List available credentials for the authenticated user or given workspace
[**ListDatasetVersions**](DefaultApi.md#ListDatasetVersions) | **Get** /workspaces/{workspaceId}/datasets/{datasetId}/versions | List all versions of a dataset
[**ListDatasets**](DefaultApi.md#ListDatasets) | **Get** /workspaces/{workspaceId}/datasets | List all datasets in the workspace
[**ListLabels**](DefaultApi.md#ListLabels) | **Get** /labels | List the labels of the authenticated user or a workspace
[**ListLaunchDatasetVersions**](DefaultApi.md#ListLaunchDatasetVersions) | **Get** /launch/{launchId}/datasets | Describe the datasets used in a launch
[**ListOrganizationCollaborators**](DefaultApi.md#ListOrganizationCollaborators) | **Get** /orgs/{orgId}/collaborators | List all collaborators of an organization
[**ListOrganizationMembers**](DefaultApi.md#ListOrganizationMembers) | **Get** /orgs/{orgId}/members | List all members of an organization
[**ListOrganizationTeamMembers**](DefaultApi.md#ListOrganizationTeamMembers) | **Get** /orgs/{orgId}/teams/{teamId}/members | List all the members of a team
[**ListOrganizationTeams**](DefaultApi.md#ListOrganizationTeams) | **Get** /orgs/{orgId}/teams | List all the teams of a given organization
[**ListOrganizations**](DefaultApi.md#ListOrganizations) | **Get** /orgs | List available organizations
[**ListPipelineRepositories**](DefaultApi.md#ListPipelineRepositories) | **Get** /pipelines/repositories | List Pipelines accessible to the authenticated user
[**ListPipelineSecrets**](DefaultApi.md#ListPipelineSecrets) | **Get** /pipeline-secrets | List all pipeline secrets in the workspace
[**ListPipelines**](DefaultApi.md#ListPipelines) | **Get** /pipelines | List all the Pipelines of a workspace
[**ListPlatformRegions**](DefaultApi.md#ListPlatformRegions) | **Get** /platforms/{platformId}/regions | List available regions for the platform specified
[**ListPlatforms**](DefaultApi.md#ListPlatforms) | **Get** /platforms | List available computing platforms
[**ListWorkflowTasks**](DefaultApi.md#ListWorkflowTasks) | **Get** /workflow/{workflowId}/tasks | List the tasks for the given Workflow ID and filter parameters
[**ListWorkflows**](DefaultApi.md#ListWorkflows) | **Get** /workflow | List Workflow records matching the filter parameters
[**ListWorkspaceDatasetVersions**](DefaultApi.md#ListWorkspaceDatasetVersions) | **Get** /workspaces/{workspaceId}/datasets/versions | List the latest version of each dataset in the workspace
[**ListWorkspaceParticipants**](DefaultApi.md#ListWorkspaceParticipants) | **Get** /orgs/{orgId}/workspaces/{workspaceId}/participants | List workspace participants
[**ListWorkspaces**](DefaultApi.md#ListWorkspaces) | **Get** /orgs/{orgId}/workspaces | List the workspaces of a given organization accessible by the authenticated user
[**ListWorkspacesUser**](DefaultApi.md#ListWorkspacesUser) | **Get** /user/{userId}/workspaces | List the workspaces and organizations of a given user
[**PauseAction**](DefaultApi.md#PauseAction) | **Post** /actions/{actionId}/pause | Toggle the pause status of an existing Pipeline Action
[**Profile**](DefaultApi.md#Profile) | **Get** /user | 
[**RemoveLabelsFromActions**](DefaultApi.md#RemoveLabelsFromActions) | **Post** /actions/labels/remove | Remove some labels from some actions
[**RemoveLabelsFromPipelines**](DefaultApi.md#RemoveLabelsFromPipelines) | **Post** /pipelines/labels/remove | Remove some labels from some pipelines
[**RemoveLabelsFromWorkflows**](DefaultApi.md#RemoveLabelsFromWorkflows) | **Post** /workflow/labels/remove | Remove some labels from some workflows
[**TokenList**](DefaultApi.md#TokenList) | **Get** /tokens | List all available API tokens
[**UpdateAction**](DefaultApi.md#UpdateAction) | **Put** /actions/{actionId} | Update a Pipeline Action
[**UpdateComputeEnvPrimary**](DefaultApi.md#UpdateComputeEnvPrimary) | **Post** /compute-envs/{computeEnvId}/primary | Defines the primary Tower compute environment
[**UpdateCredentials**](DefaultApi.md#UpdateCredentials) | **Put** /credentials/{credentialsId} | Update the credentials data for the given id
[**UpdateDataset**](DefaultApi.md#UpdateDataset) | **Put** /workspaces/{workspaceId}/datasets/{datasetId} | Update a dataset
[**UpdateLabel**](DefaultApi.md#UpdateLabel) | **Put** /labels/{labelId} | Update an existing label
[**UpdateOrganization**](DefaultApi.md#UpdateOrganization) | **Put** /orgs/{orgId} | Update organization details
[**UpdateOrganizationMemberRole**](DefaultApi.md#UpdateOrganizationMemberRole) | **Put** /orgs/{orgId}/members/{memberId}/role | Update the role of an organization member
[**UpdatePipeline**](DefaultApi.md#UpdatePipeline) | **Put** /pipelines/{pipelineId} | Update a Pipeline
[**UpdatePipelineSecret**](DefaultApi.md#UpdatePipelineSecret) | **Put** /pipeline-secrets/{secretId} | Update the secrets data for the given id
[**UpdateTraceBegin**](DefaultApi.md#UpdateTraceBegin) | **Put** /trace/{workflowId}/begin | Signal the Workflow execution has started
[**UpdateTraceComplete**](DefaultApi.md#UpdateTraceComplete) | **Put** /trace/{workflowId}/complete | Signal the Workflow execution has completed
[**UpdateTraceHeartbeat**](DefaultApi.md#UpdateTraceHeartbeat) | **Put** /trace/{workflowId}/heartbeat | Period request to signal the execution is still on-going
[**UpdateTraceProgress**](DefaultApi.md#UpdateTraceProgress) | **Put** /trace/{workflowId}/progress | Store one or more task executions metadata for the specified Workflow
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /users/{userId} | Update an user entity
[**UpdateWorkspace**](DefaultApi.md#UpdateWorkspace) | **Put** /orgs/{orgId}/workspaces/{workspaceId} | Update workspace details
[**UpdateWorkspaceParticipantRole**](DefaultApi.md#UpdateWorkspaceParticipantRole) | **Put** /orgs/{orgId}/workspaces/{workspaceId}/participants/{participantId}/role | Update a participant role
[**UploadDataset**](DefaultApi.md#UploadDataset) | **Post** /workspaces/{workspaceId}/datasets/{datasetId}/upload | Upload the content of a new dataset version
[**UserInfo**](DefaultApi.md#UserInfo) | **Get** /user-info | Describe current user
[**ValidateActionName**](DefaultApi.md#ValidateActionName) | **Get** /actions/validate | Check that an action name is valid
[**ValidateComputeEnvName**](DefaultApi.md#ValidateComputeEnvName) | **Get** /compute-envs/validate | Check that is a valid compute env name
[**ValidateCredentialsName**](DefaultApi.md#ValidateCredentialsName) | **Get** /credentials/validate | Check that is a valid credentials name
[**ValidateOrganizationName**](DefaultApi.md#ValidateOrganizationName) | **Get** /orgs/validate | Check that is a valid organization name
[**ValidatePipelineName**](DefaultApi.md#ValidatePipelineName) | **Get** /pipelines/validate | Check that a Pipeline name is valid
[**ValidatePipelineSecretName**](DefaultApi.md#ValidatePipelineSecretName) | **Get** /pipeline-secrets/validate | Check that is a valid pipeline secret name
[**ValidateTeamName**](DefaultApi.md#ValidateTeamName) | **Get** /orgs/{orgId}/teams/validate | Check that is a valid team name
[**ValidateUserName**](DefaultApi.md#ValidateUserName) | **Get** /users/validate | Check that the user name is valid
[**ValidateWorkflowConstraints**](DefaultApi.md#ValidateWorkflowConstraints) | **Get** /workflow/validate | Check that the given run name of a workflow has a valid format. When the session ID is given: check that no other workflow in the system exists with the combination of both elements
[**WorkflowLogs**](DefaultApi.md#WorkflowLogs) | **Get** /workflow/{workflowId}/log | Retrieve Workflow output logs of the Nextflow main job
[**WorkspaceValidate**](DefaultApi.md#WorkspaceValidate) | **Get** /orgs/{orgId}/workspaces/validate | Check that is a valid workspace name



## AddLabelsToActions

> AddLabelsToActions(ctx, associateActionLabelsRequest, optional)

Add some labels to some actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateActionLabelsRequest** | [**AssociateActionLabelsRequest**](AssociateActionLabelsRequest.md)| Labels add request | 
 **optional** | ***AddLabelsToActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddLabelsToActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLabelsToPipelines

> AddLabelsToPipelines(ctx, associatePipelineLabelsRequest, optional)

Add some labels to some pipelines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associatePipelineLabelsRequest** | [**AssociatePipelineLabelsRequest**](AssociatePipelineLabelsRequest.md)| Labels add request | 
 **optional** | ***AddLabelsToPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddLabelsToPipelinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLabelsToWorkflows

> AddLabelsToWorkflows(ctx, associateWorkflowLabelsRequest, optional)

Add some labels to some workflows

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateWorkflowLabelsRequest** | [**AssociateWorkflowLabelsRequest**](AssociateWorkflowLabelsRequest.md)| Labels add request | 
 **optional** | ***AddLabelsToWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddLabelsToWorkflowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyLabelsToActions

> ApplyLabelsToActions(ctx, associateActionLabelsRequest, optional)

Apply some labels to some actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateActionLabelsRequest** | [**AssociateActionLabelsRequest**](AssociateActionLabelsRequest.md)| Labels apply request | 
 **optional** | ***ApplyLabelsToActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplyLabelsToActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyLabelsToPipelines

> ApplyLabelsToPipelines(ctx, associatePipelineLabelsRequest, optional)

Apply some labels to some pipelines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associatePipelineLabelsRequest** | [**AssociatePipelineLabelsRequest**](AssociatePipelineLabelsRequest.md)| Labels apply request | 
 **optional** | ***ApplyLabelsToPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplyLabelsToPipelinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyLabelsToWorkflows

> ApplyLabelsToWorkflows(ctx, associateWorkflowLabelsRequest, optional)

Apply some labels to some workflows

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateWorkflowLabelsRequest** | [**AssociateWorkflowLabelsRequest**](AssociateWorkflowLabelsRequest.md)| Labels apply request | 
 **optional** | ***ApplyLabelsToWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplyLabelsToWorkflowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelWorkflow

> CancelWorkflow(ctx, workflowId, optional)

Cancel a Workflow execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***CancelWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **body** | **optional.Map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAction

> CreateActionResponse CreateAction(ctx, createActionRequest, optional)

Create a new Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createActionRequest** | [**CreateActionRequest**](CreateActionRequest.md)| Action create request | 
 **optional** | ***CreateActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateActionResponse**](CreateActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAvatar

> CreateAvatarResponse CreateAvatar(ctx, optional)

Create the avatar image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAvatarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**CreateAvatarResponse**](CreateAvatarResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComputeEnv

> CreateComputeEnvResponse CreateComputeEnv(ctx, createComputeEnvRequest, optional)

Create a new Tower compute environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createComputeEnvRequest** | [**CreateComputeEnvRequest**](CreateComputeEnvRequest.md)| Compute environment create request | 
 **optional** | ***CreateComputeEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateComputeEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateComputeEnvResponse**](CreateComputeEnvResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentials

> CreateCredentialsResponse CreateCredentials(ctx, createCredentialsRequest, optional)

Create a new credentials record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createCredentialsRequest** | [**CreateCredentialsRequest**](CreateCredentialsRequest.md)| Credentials create request | 
 **optional** | ***CreateCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateCredentialsResponse**](CreateCredentialsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataset

> CreateDatasetResponse CreateDataset(ctx, workspaceId, createDatasetRequest)

Create a dataset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**createDatasetRequest** | [**CreateDatasetRequest**](CreateDatasetRequest.md)| Dataset create request | 

### Return type

[**CreateDatasetResponse**](CreateDatasetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> CreateLabelResponse CreateLabel(ctx, createLabelRequest, optional)

Create a new label

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createLabelRequest** | [**CreateLabelRequest**](CreateLabelRequest.md)| Label create request | 
 **optional** | ***CreateLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateLabelResponse**](CreateLabelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> CreateOrganizationResponse CreateOrganization(ctx, createOrganizationRequest)

Create a new organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md)| Organization create request | 

### Return type

[**CreateOrganizationResponse**](CreateOrganizationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationMember

> AddMemberResponse CreateOrganizationMember(ctx, orgId, addMemberRequest)

Add a new organization member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**addMemberRequest** | [**AddMemberRequest**](AddMemberRequest.md)| Member addition request | 

### Return type

[**AddMemberResponse**](AddMemberResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationTeam

> CreateTeamResponse CreateOrganizationTeam(ctx, orgId, createTeamRequest)

Create a new organization team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md)| Team create request | 

### Return type

[**CreateTeamResponse**](CreateTeamResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationTeamMember

> AddTeamMemberResponse CreateOrganizationTeamMember(ctx, orgId, teamId, createTeamMemberRequest)

Add a new team member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**teamId** | **int64**| Team numeric identifier | 
**createTeamMemberRequest** | [**CreateTeamMemberRequest**](CreateTeamMemberRequest.md)| Team create request | 

### Return type

[**AddTeamMemberResponse**](AddTeamMemberResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipeline

> CreatePipelineResponse CreatePipeline(ctx, createPipelineRequest, optional)

Create a new Pipeline in a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPipelineRequest** | [**CreatePipelineRequest**](CreatePipelineRequest.md)| Pipeline creation request | 
 **optional** | ***CreatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreatePipelineResponse**](CreatePipelineResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePipelineSecret

> CreatePipelineSecretResponse CreatePipelineSecret(ctx, createPipelineSecretRequest, optional)

Create a Pipeline Secrets inside the given workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPipelineSecretRequest** | [**CreatePipelineSecretRequest**](CreatePipelineSecretRequest.md)| Pipeline secret create request | 
 **optional** | ***CreatePipelineSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePipelineSecretOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreatePipelineSecretResponse**](CreatePipelineSecretResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> CreateAccessTokenResponse CreateToken(ctx, createAccessTokenRequest)

Create an API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAccessTokenRequest** | [**CreateAccessTokenRequest**](CreateAccessTokenRequest.md)| Access token create request | 

### Return type

[**CreateAccessTokenResponse**](CreateAccessTokenResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrace

> TraceCreateResponse CreateTrace(ctx, traceCreateRequest, optional)

Create a new Workflow execution trace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceCreateRequest** | [**TraceCreateRequest**](TraceCreateRequest.md)| Trace create request | 
 **optional** | ***CreateTraceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTraceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**TraceCreateResponse**](TraceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowLaunch

> SubmitWorkflowLaunchResponse CreateWorkflowLaunch(ctx, submitWorkflowLaunchRequest, optional)

Submit a Workflow execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submitWorkflowLaunchRequest** | [**SubmitWorkflowLaunchRequest**](SubmitWorkflowLaunchRequest.md)| Workflow launch request | 
 **optional** | ***CreateWorkflowLaunchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWorkflowLaunchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **optimized** | **optional.Bool**| Launch with optimized configuration from groundswell service | 
 **sourceWorkspaceId** | **optional.Int64**|  | 

### Return type

[**SubmitWorkflowLaunchResponse**](SubmitWorkflowLaunchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowStar

> CreateWorkflowStarResponse CreateWorkflowStar(ctx, workflowId, optional)

Star a workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***CreateWorkflowStarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWorkflowStarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateWorkflowStarResponse**](CreateWorkflowStarResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> CreateWorkspaceResponse CreateWorkspace(ctx, orgId, createWorkspaceRequest)

Create a new workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**createWorkspaceRequest** | [**CreateWorkspaceRequest**](CreateWorkspaceRequest.md)| Workspace create request | 

### Return type

[**CreateWorkspaceResponse**](CreateWorkspaceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceParticipant

> AddParticipantResponse CreateWorkspaceParticipant(ctx, orgId, workspaceId, addParticipantRequest)

Create a new workspace participant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 
**addParticipantRequest** | [**AddParticipantRequest**](AddParticipantRequest.md)| Participant addition request | 

### Return type

[**AddParticipantResponse**](AddParticipantResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAction

> DeleteAction(ctx, actionId, optional)

Delete a Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string**| Action string identifier | 
 **optional** | ***DeleteActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllTokens

> DeleteAllTokens(ctx, )

Delete all user API tokens

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComputeEnv

> DeleteComputeEnv(ctx, computeEnvId, optional)

Delete an existing Tower compute environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**computeEnvId** | **string**| Compute environment string identifier | 
 **optional** | ***DeleteComputeEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteComputeEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentials

> DeleteCredentials(ctx, credentialsId, optional)

Delete the credentials record for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string**| Credentials string identifier | 
 **optional** | ***DeleteCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataset

> DeleteDataset(ctx, workspaceId, datasetId)

Delete a dataset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, labelId, optional)

Delete a label

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **int64**| Label numeric identifier | 
 **optional** | ***DeleteLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, orgId)

Delete an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationMember

> DeleteOrganizationMember(ctx, orgId, memberId)

Delete an organization member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**memberId** | **int64**| Organization member numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationTeam

> DeleteOrganizationTeam(ctx, orgId, teamId)

Delete an organization team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**teamId** | **int64**| Team numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationTeamMember

> DeleteOrganizationTeamMember(ctx, orgId, teamId, memberId)

Delete a team member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**teamId** | **int64**| Team numeric identifier | 
**memberId** | **int64**| Member numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipeline

> DeletePipeline(ctx, pipelineId, optional)

Delete a Pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **int64**| Pipeline numeric identifier | 
 **optional** | ***DeletePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeletePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineSecret

> DeletePipelineSecret(ctx, secretId, optional)

Delete the pipeline secret record for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **int64**| Secret numeric identifier | 
 **optional** | ***DeletePipelineSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeletePipelineSecretOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, tokenId)

Delete an API token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int64**| Token numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId)

Delete a user entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| User numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx, workflowId, optional)

Delete the Workflow entity with the given ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DeleteWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **force** | **optional.Bool**| Force the deletion even if the workflow is active | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowMany

> DeleteWorkflowsResponse DeleteWorkflowMany(ctx, deleteWorkflowsRequest, optional)

Delete several workflow entities given their ids

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deleteWorkflowsRequest** | [**DeleteWorkflowsRequest**](DeleteWorkflowsRequest.md)| Delete workflows request | 
 **optional** | ***DeleteWorkflowManyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWorkflowManyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **force** | **optional.Bool**| Force the deletion even if any workflows are active | 

### Return type

[**DeleteWorkflowsResponse**](DeleteWorkflowsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowStar

> CreateWorkflowStarResponse DeleteWorkflowStar(ctx, workflowId, optional)

Unstar a workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DeleteWorkflowStarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWorkflowStarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateWorkflowStarResponse**](CreateWorkflowStarResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspace(ctx, orgId, workspaceId)

Delete a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspaceParticipant

> DeleteWorkspaceParticipant(ctx, orgId, workspaceId, participantId)

Delete a workspace participant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 
**participantId** | **int64**| Participant numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAction

> DescribeActionResponse DescribeAction(ctx, actionId, optional)

Describe an existing Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string**| Action string identifier | 
 **optional** | ***DescribeActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **attributes** | [**optional.Interface of []ActionQueryAttribute**](ActionQueryAttribute.md)| Comma-separated list of attributes to retrieve: &#x60;labels&#x60;. Empty to retrieve nothing  | 

### Return type

[**DescribeActionResponse**](DescribeActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeComputeEnv

> DescribeComputeEnvResponse DescribeComputeEnv(ctx, computeEnvId, optional)

Describe a Tower compute environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**computeEnvId** | **string**| Compute environment string identifier | 
 **optional** | ***DescribeComputeEnvOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeComputeEnvOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **attributes** | [**optional.Interface of []ComputeEnvQueryAttribute**](ComputeEnvQueryAttribute.md)| Comma-separated list of attributes to retrieve: &#x60;labels&#x60;. Empty to retrieve nothing  | 

### Return type

[**DescribeComputeEnvResponse**](DescribeComputeEnvResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCredentials

> DescribeCredentialsResponse DescribeCredentials(ctx, credentialsId, optional)

Describe the credentials for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string**| Credentials string identifier | 
 **optional** | ***DescribeCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribeCredentialsResponse**](DescribeCredentialsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDataset

> DescribeDatasetResponse DescribeDataset(ctx, workspaceId, datasetId)

Describe a dataset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 

### Return type

[**DescribeDatasetResponse**](DescribeDatasetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeLaunch

> DescribeLaunchResponse DescribeLaunch(ctx, launchId, optional)

Describe the Launch record for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launchId** | **string**| Launch string identifier | 
 **optional** | ***DescribeLaunchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeLaunchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribeLaunchResponse**](DescribeLaunchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeOrganization

> DescribeOrganizationResponse DescribeOrganization(ctx, orgId)

Describe organization details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 

### Return type

[**DescribeOrganizationResponse**](DescribeOrganizationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePipeline

> DescribePipelineResponse DescribePipeline(ctx, pipelineId, optional)

Describe a Pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **int64**| Pipeline numeric identifier | 
 **optional** | ***DescribePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | [**optional.Interface of []PipelineQueryAttribute**](PipelineQueryAttribute.md)| Comma-separated list of attributes to retrieve. Empty to retrieve nothing  | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **sourceWorkspaceId** | **optional.Int64**| Source workspace numeric identifier | 

### Return type

[**DescribePipelineResponse**](DescribePipelineResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePipelineLaunch

> DescribeLaunchResponse DescribePipelineLaunch(ctx, pipelineId, optional)

Describe a Pipeline launch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **int64**| Pipeline numeric identifier | 
 **optional** | ***DescribePipelineLaunchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePipelineLaunchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **sourceWorkspaceId** | **optional.Int64**| Source workspace numeric identifier | 

### Return type

[**DescribeLaunchResponse**](DescribeLaunchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePipelineRepository

> DescribePipelineInfoResponse DescribePipelineRepository(ctx, optional)

Describe the Pipeline entity for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DescribePipelineRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePipelineRepositoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Pipeline name | 
 **revision** | **optional.String**| Pipeline revision | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribePipelineInfoResponse**](DescribePipelineInfoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePipelineSchema

> PipelineSchemaResponse DescribePipelineSchema(ctx, pipelineId, optional)

Retrieve the Pipeline input schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **int64**| Pipeline numeric identifier | 
 **optional** | ***DescribePipelineSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePipelineSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **sourceWorkspaceId** | **optional.Int64**| Source workspace numeric identifier | 
 **attributes** | [**optional.Interface of []PipelineSchemaAttributes**](PipelineSchemaAttributes.md)| Comma-separated list of attributes to retrieve: &#x60;schema&#x60; or &#x60;params&#x60;. Empty to retrieve all of them | 

### Return type

[**PipelineSchemaResponse**](PipelineSchemaResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePipelineSecret

> DescribePipelineSecretResponse DescribePipelineSecret(ctx, secretId, optional)

Fetch a single pipeline secret

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **int64**| Secret numeric identifier | 
 **optional** | ***DescribePipelineSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePipelineSecretOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribePipelineSecretResponse**](DescribePipelineSecretResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePlatform

> DescribePlatformResponse DescribePlatform(ctx, platformId, regionId, credentialsId, optional)

Describe the platform entity for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string**| Platform string identifier | 
**regionId** | **string**|  | 
**credentialsId** | **string**|  | 
 **optional** | ***DescribePlatformOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribePlatformOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribePlatformResponse**](DescribePlatformResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUser

> DescribeUserResponse DescribeUser(ctx, userId)

Describe a user entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| User numeric identifier | 

### Return type

[**DescribeUserResponse**](DescribeUserResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflow

> DescribeWorkflowResponse DescribeWorkflow(ctx, workflowId, optional)

Describe the Workflow entity for the given ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DescribeWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **attributes** | [**optional.Interface of []WorkflowQueryAttribute**](WorkflowQueryAttribute.md)| Comma-separated list of attributes to retrieve. Empty to retrieve nothing  | 

### Return type

[**DescribeWorkflowResponse**](DescribeWorkflowResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflowLaunch

> DescribeWorkflowLaunchResponse DescribeWorkflowLaunch(ctx, workflowId, optional)

Describe a Workflow launch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DescribeWorkflowLaunchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowLaunchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribeWorkflowLaunchResponse**](DescribeWorkflowLaunchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflowMetrics

> GetWorkflowMetricsResponse DescribeWorkflowMetrics(ctx, workflowId, optional)

Get the execution metrics for the given Workflow ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DescribeWorkflowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**GetWorkflowMetricsResponse**](GetWorkflowMetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflowProgress

> GetProgressResponse DescribeWorkflowProgress(ctx, workflowId, optional)

Retrieve the execution progress for the given Workflow ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DescribeWorkflowProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**GetProgressResponse**](GetProgressResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflowStar

> CreateWorkflowStarResponse DescribeWorkflowStar(ctx, workflowId, optional)

Check starred status of a workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DescribeWorkflowStarOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowStarOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**CreateWorkflowStarResponse**](CreateWorkflowStarResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkflowTask

> DescribeTaskResponse DescribeWorkflowTask(ctx, workflowId, taskId, optional)

Describe a task entity with the given ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**taskId** | **int64**| Task numeric identifier | 
 **optional** | ***DescribeWorkflowTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DescribeWorkflowTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribeTaskResponse**](DescribeTaskResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWorkspace

> DescribeWorkspaceResponse DescribeWorkspace(ctx, orgId, workspaceId)

Describe a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 

### Return type

[**DescribeWorkspaceResponse**](DescribeWorkspaceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAvatar

> *os.File DownloadAvatar(ctx, avatarId)

Download the avatar image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string**| Avatar string identifier | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDataset

> *os.File DownloadDataset(ctx, workspaceId, datasetId, version, fileName)

Download the content of a dataset version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 
**version** | **string**| Version number to download | 
**fileName** | **string**| File name | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadWorkflowLog

> *os.File DownloadWorkflowLog(ctx, workflowId, optional)

Download Workflow files of the Nextflow main job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***DownloadWorkflowLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadWorkflowLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileName** | **optional.String**| Filename to download | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadWorkflowTaskLog

> *os.File DownloadWorkflowTaskLog(ctx, workflowId, taskId, optional)

Download Workflow files of a given task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**taskId** | **int64**| Task numeric identifier | 
 **optional** | ***DownloadWorkflowTaskLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadWorkflowTaskLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileName** | **optional.String**| Filename to download | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaRunCancel

> RunId GaRunCancel(ctx, runId, optional)

GA4GH cancel a run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string**| Run string identifier | 
 **optional** | ***GaRunCancelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaRunCancelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**RunId**](RunId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaRunCreate

> RunId GaRunCreate(ctx, runRequest)

GA4GH create a new run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runRequest** | [**RunRequest**](RunRequest.md)| Run request | 

### Return type

[**RunId**](RunId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaRunDescribe

> RunLog GaRunDescribe(ctx, runId)

GA4GH describe run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string**| Run string identifier | 

### Return type

[**RunLog**](RunLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaRunList

> RunListResponse GaRunList(ctx, optional)

GA4GH list runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GaRunListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaRunListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| Page size | 
 **pageToken** | **optional.String**| Page token | 

### Return type

[**RunListResponse**](RunListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaRunStatus

> RunStatus GaRunStatus(ctx, runId)

GA4GH retrieve run status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string**| Run string identifier | 

### Return type

[**RunStatus**](RunStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaServiceInfo

> ServiceInfo GaServiceInfo(ctx, )

GA4GH service info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceInfo**](ServiceInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateRandomWorkflowName

> RandomWorkflowNameResponse GenerateRandomWorkflowName(ctx, )

Generates a random name

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RandomWorkflowNameResponse**](RandomWorkflowNameResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTaskLog

> WorkflowLogResponse GetWorkflowTaskLog(ctx, workflowId, taskId, optional)

Retrieve Workflow output logs of a given task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**taskId** | **int64**| Task numeric identifier | 
 **optional** | ***GetWorkflowTaskLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWorkflowTaskLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **next** | **optional.String**| Workflow log cursor | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**WorkflowLogResponse**](WorkflowLogResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Info

> ServiceInfoResponse Info(ctx, )

General Tower service features and version

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceInfoResponse**](ServiceInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchAction

> LaunchActionResponse LaunchAction(ctx, actionId, launchActionRequest, optional)

Trigger the execution of a Tower Launch Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string**| Action string identifier | 
**launchActionRequest** | [**LaunchActionRequest**](LaunchActionRequest.md)| Action launch request | 
 **optional** | ***LaunchActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LaunchActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**LaunchActionResponse**](LaunchActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveOrganization

> LeaveOrganization(ctx, orgId)

Leave an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveWorkspaceParticipant

> LeaveWorkspaceParticipant(ctx, orgId, workspaceId)

Leave a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionTypes

> ListEventTypesResponse ListActionTypes(ctx, optional)

List the supported event types that can trigger a Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListActionTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListActionTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListEventTypesResponse**](ListEventTypesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActions

> ListActionsResponse ListActions(ctx, optional)

List the available Pipeline Actions for the authenticated user or given workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **attributes** | [**optional.Interface of []ActionQueryAttribute**](ActionQueryAttribute.md)| Comma-separated list of attributes to retrieve: &#x60;labels&#x60;. Empty to retrieve nothing  | 

### Return type

[**ListActionsResponse**](ListActionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComputeEnvs

> ListComputeEnvsResponse ListComputeEnvs(ctx, optional)

List all Tower compute environments for the authenticated user or given workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListComputeEnvsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListComputeEnvsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Compute environment status | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListComputeEnvsResponse**](ListComputeEnvsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentials

> ListCredentialsResponse ListCredentials(ctx, optional)

List available credentials for the authenticated user or given workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **platformId** | **optional.String**| Platform string identifier | 

### Return type

[**ListCredentialsResponse**](ListCredentialsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatasetVersions

> ListDatasetVersionsResponse ListDatasetVersions(ctx, workspaceId, datasetId, optional)

List all versions of a dataset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 
 **optional** | ***ListDatasetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDatasetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mimeType** | **optional.String**| Optional MIME type filter | 

### Return type

[**ListDatasetVersionsResponse**](ListDatasetVersionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatasets

> ListDatasetsResponse ListDatasets(ctx, workspaceId)

List all datasets in the workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 

### Return type

[**ListDatasetsResponse**](ListDatasetsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

> ListLabelsResponse ListLabels(ctx, optional)

List the labels of the authenticated user or a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLabelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 
 **type_** | [**optional.Interface of LabelType**](.md)| Labels type | 

### Return type

[**ListLabelsResponse**](ListLabelsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLaunchDatasetVersions

> DescribeLaunchResponse ListLaunchDatasetVersions(ctx, launchId, optional)

Describe the datasets used in a launch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launchId** | **string**| Launch string identifier | 
 **optional** | ***ListLaunchDatasetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLaunchDatasetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**DescribeLaunchResponse**](DescribeLaunchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCollaborators

> ListMembersResponse ListOrganizationCollaborators(ctx, orgId, optional)

List all collaborators of an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
 **optional** | ***ListOrganizationCollaboratorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrganizationCollaboratorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 

### Return type

[**ListMembersResponse**](ListMembersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationMembers

> ListMembersResponse ListOrganizationMembers(ctx, orgId, optional)

List all members of an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
 **optional** | ***ListOrganizationMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrganizationMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 

### Return type

[**ListMembersResponse**](ListMembersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationTeamMembers

> ListMembersResponse ListOrganizationTeamMembers(ctx, orgId, teamId, optional)

List all the members of a team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**teamId** | **int64**| Team numeric identifier | 
 **optional** | ***ListOrganizationTeamMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrganizationTeamMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **max** | **optional.Int32**|  | 
 **offset** | **optional.Int32**|  | 
 **search** | **optional.String**|  | 

### Return type

[**ListMembersResponse**](ListMembersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationTeams

> ListTeamResponse ListOrganizationTeams(ctx, orgId, optional)

List all the teams of a given organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
 **optional** | ***ListOrganizationTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrganizationTeamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 

### Return type

[**ListTeamResponse**](ListTeamResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> ListOrganizationsResponse ListOrganizations(ctx, optional)

List available organizations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOrganizationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **optional.String**| Organization user role identifier | 

### Return type

[**ListOrganizationsResponse**](ListOrganizationsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineRepositories

> ListPipelineInfoResponse ListPipelineRepositories(ctx, optional)

List Pipelines accessible to the authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPipelineRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelineRepositoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListPipelineInfoResponse**](ListPipelineInfoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineSecrets

> ListPipelineSecretsResponse ListPipelineSecrets(ctx, optional)

List all pipeline secrets in the workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPipelineSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelineSecretsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListPipelineSecretsResponse**](ListPipelineSecretsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelines

> ListPipelinesResponse ListPipelines(ctx, optional)

List all the Pipelines of a workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | [**optional.Interface of []PipelineQueryAttribute**](PipelineQueryAttribute.md)| Comma-separated list of attributes to retrieve. Empty to retrieve nothing  | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 
 **visibility** | **optional.String**| Filter visibility param | 

### Return type

[**ListPipelinesResponse**](ListPipelinesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlatformRegions

> ListRegionsResponse ListPlatformRegions(ctx, platformId, optional)

List available regions for the platform specified

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string**| Platform string identifier | 
 **optional** | ***ListPlatformRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPlatformRegionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListRegionsResponse**](ListRegionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlatforms

> ListPlatformsResponse ListPlatforms(ctx, optional)

List available computing platforms

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPlatformsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPlatformsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**ListPlatformsResponse**](ListPlatformsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowTasks

> ListTasksResponse ListWorkflowTasks(ctx, workflowId, optional)

List the tasks for the given Workflow ID and filter parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***ListWorkflowTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkflowTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **sortBy** | **optional.String**| Field to sort by | 
 **sortDir** | **optional.String**| Sorting direction (asc|desc) | 
 **search** | **optional.String**| Search tasks by name | 

### Return type

[**ListTasksResponse**](ListTasksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> ListWorkflowsResponse ListWorkflows(ctx, optional)

List Workflow records matching the filter parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkflowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | [**optional.Interface of []WorkflowQueryAttribute**](WorkflowQueryAttribute.md)| Comma-separated list of attributes to retrieve. Empty to retrieve nothing  | 
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 

### Return type

[**ListWorkflowsResponse**](ListWorkflowsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceDatasetVersions

> ListDatasetVersionsResponse ListWorkspaceDatasetVersions(ctx, workspaceId, optional)

List the latest version of each dataset in the workspace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
 **optional** | ***ListWorkspaceDatasetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkspaceDatasetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mimeType** | **optional.String**|  | 

### Return type

[**ListDatasetVersionsResponse**](ListDatasetVersionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceParticipants

> ListParticipantsResponse ListWorkspaceParticipants(ctx, orgId, workspaceId, optional)

List workspace participants

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 
 **optional** | ***ListWorkspaceParticipantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkspaceParticipantsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **max** | **optional.Int32**| Pagination result max result | 
 **offset** | **optional.Int32**| Pagination offset | 
 **search** | **optional.String**| Filter search param | 

### Return type

[**ListParticipantsResponse**](ListParticipantsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> ListWorkspacesResponse ListWorkspaces(ctx, orgId)

List the workspaces of a given organization accessible by the authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 

### Return type

[**ListWorkspacesResponse**](ListWorkspacesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspacesUser

> ListWorkspacesAndOrgResponse ListWorkspacesUser(ctx, userId)

List the workspaces and organizations of a given user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| User numeric identifier | 

### Return type

[**ListWorkspacesAndOrgResponse**](ListWorkspacesAndOrgResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseAction

> PauseAction(ctx, actionId, optional)

Toggle the pause status of an existing Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string**| Action string identifier | 
 **optional** | ***PauseActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PauseActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **body** | **optional.Map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Profile

> DescribeUserResponse Profile(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DescribeUserResponse**](DescribeUserResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLabelsFromActions

> RemoveLabelsFromActions(ctx, associateActionLabelsRequest, optional)

Remove some labels from some actions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateActionLabelsRequest** | [**AssociateActionLabelsRequest**](AssociateActionLabelsRequest.md)| Labels remove request | 
 **optional** | ***RemoveLabelsFromActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveLabelsFromActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLabelsFromPipelines

> RemoveLabelsFromPipelines(ctx, associatePipelineLabelsRequest, optional)

Remove some labels from some pipelines

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associatePipelineLabelsRequest** | [**AssociatePipelineLabelsRequest**](AssociatePipelineLabelsRequest.md)| Labels remove request | 
 **optional** | ***RemoveLabelsFromPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveLabelsFromPipelinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLabelsFromWorkflows

> RemoveLabelsFromWorkflows(ctx, associateWorkflowLabelsRequest, optional)

Remove some labels from some workflows

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**associateWorkflowLabelsRequest** | [**AssociateWorkflowLabelsRequest**](AssociateWorkflowLabelsRequest.md)| Labels remove request | 
 **optional** | ***RemoveLabelsFromWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveLabelsFromWorkflowsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenList

> ListAccessTokensResponse TokenList(ctx, )

List all available API tokens

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ListAccessTokensResponse**](ListAccessTokensResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAction

> UpdateAction(ctx, actionId, updateActionRequest, optional)

Update a Pipeline Action

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionId** | **string**| Action string identifier | 
**updateActionRequest** | [**UpdateActionRequest**](UpdateActionRequest.md)| Action update request | 
 **optional** | ***UpdateActionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateActionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComputeEnvPrimary

> UpdateComputeEnvPrimary(ctx, computeEnvId, optional)

Defines the primary Tower compute environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**computeEnvId** | **string**| Compute environment string identifier | 
 **optional** | ***UpdateComputeEnvPrimaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateComputeEnvPrimaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **body** | **optional.Map[string]interface{}**|  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentials

> UpdateCredentials(ctx, credentialsId, updateCredentialsRequest, optional)

Update the credentials data for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string**| Credentials string identifier | 
**updateCredentialsRequest** | [**UpdateCredentialsRequest**](UpdateCredentialsRequest.md)| Credentials create request | 
 **optional** | ***UpdateCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataset

> UpdateDataset(ctx, workspaceId, datasetId, updateDatasetRequest)

Update a dataset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 
**updateDatasetRequest** | [**UpdateDatasetRequest**](UpdateDatasetRequest.md)| Dataset update request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> UpdateLabelResponse UpdateLabel(ctx, labelId, updateLabelRequest, optional)

Update an existing label

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **int64**| Label numeric identifier | 
**updateLabelRequest** | [**UpdateLabelRequest**](UpdateLabelRequest.md)| Label update request | 
 **optional** | ***UpdateLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLabelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**UpdateLabelResponse**](UpdateLabelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> UpdateOrganization(ctx, orgId, updateOrganizationRequest)

Update organization details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**updateOrganizationRequest** | [**UpdateOrganizationRequest**](UpdateOrganizationRequest.md)| Organization update request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationMemberRole

> UpdateOrganizationMemberRole(ctx, orgId, memberId, updateMemberRoleRequest)

Update the role of an organization member

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**memberId** | **int64**| Organization member numeric identifier | 
**updateMemberRoleRequest** | [**UpdateMemberRoleRequest**](UpdateMemberRoleRequest.md)| Member role update request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipeline

> UpdatePipelineResponse UpdatePipeline(ctx, pipelineId, updatePipelineRequest, optional)

Update a Pipeline

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **int64**| Pipeline numeric identifier | 
**updatePipelineRequest** | [**UpdatePipelineRequest**](UpdatePipelineRequest.md)| Pipeline update request | 
 **optional** | ***UpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePipelineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**UpdatePipelineResponse**](UpdatePipelineResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineSecret

> UpdatePipelineSecret(ctx, secretId, updatePipelineSecretRequest, optional)

Update the secrets data for the given id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **int64**| Secret numeric identifier | 
**updatePipelineSecretRequest** | [**UpdatePipelineSecretRequest**](UpdatePipelineSecretRequest.md)| Secret update request | 
 **optional** | ***UpdatePipelineSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePipelineSecretOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceBegin

> TraceBeginResponse UpdateTraceBegin(ctx, workflowId, traceBeginRequest, optional)

Signal the Workflow execution has started

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**traceBeginRequest** | [**TraceBeginRequest**](TraceBeginRequest.md)| Trace begin request | 
 **optional** | ***UpdateTraceBeginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTraceBeginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**TraceBeginResponse**](TraceBeginResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceComplete

> TraceCompleteResponse UpdateTraceComplete(ctx, workflowId, traceCompleteRequest, optional)

Signal the Workflow execution has completed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**traceCompleteRequest** | [**TraceCompleteRequest**](TraceCompleteRequest.md)| Trace complete request | 
 **optional** | ***UpdateTraceCompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTraceCompleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**TraceCompleteResponse**](TraceCompleteResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceHeartbeat

> TraceHeartbeatResponse UpdateTraceHeartbeat(ctx, workflowId, traceHeartbeatRequest, optional)

Period request to signal the execution is still on-going

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**traceHeartbeatRequest** | [**TraceHeartbeatRequest**](TraceHeartbeatRequest.md)| Trace heartbeat request | 
 **optional** | ***UpdateTraceHeartbeatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTraceHeartbeatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**TraceHeartbeatResponse**](TraceHeartbeatResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceProgress

> TraceProgressResponse UpdateTraceProgress(ctx, workflowId, traceProgressRequest, optional)

Store one or more task executions metadata for the specified Workflow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
**traceProgressRequest** | [**TraceProgressRequest**](TraceProgressRequest.md)| Trace progress request | 
 **optional** | ***UpdateTraceProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTraceProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 

### Return type

[**TraceProgressResponse**](TraceProgressResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, userId, user)

Update an user entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| User numeric identifier | 
**user** | [**User**](User.md)| User update request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspace

> DescribeWorkspaceResponse UpdateWorkspace(ctx, orgId, workspaceId, updateWorkspaceRequest)

Update workspace details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 
**updateWorkspaceRequest** | [**UpdateWorkspaceRequest**](UpdateWorkspaceRequest.md)| Workspace update request | 

### Return type

[**DescribeWorkspaceResponse**](DescribeWorkspaceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceParticipantRole

> UpdateWorkspaceParticipantRole(ctx, orgId, workspaceId, participantId, updateParticipantRoleRequest)

Update a participant role

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
**workspaceId** | **int64**| Workspace numeric identifier | 
**participantId** | **int64**| Participant numeric identifier | 
**updateParticipantRoleRequest** | [**UpdateParticipantRoleRequest**](UpdateParticipantRoleRequest.md)| Participant role update request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDataset

> UploadDatasetVersionResponse UploadDataset(ctx, workspaceId, datasetId, optional)

Upload the content of a new dataset version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **int64**| Workspace numeric identifier | 
**datasetId** | **string**| Dataset string identifier | 
 **optional** | ***UploadDatasetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadDatasetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **header** | **optional.Bool**| Uploaded file has header | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**UploadDatasetVersionResponse**](UploadDatasetVersionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInfo

> DescribeUserResponse UserInfo(ctx, )

Describe current user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DescribeUserResponse**](DescribeUserResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateActionName

> ValidateActionName(ctx, optional)

Check that an action name is valid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateActionNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateActionNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **name** | **optional.String**| Action name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateComputeEnvName

> ValidateComputeEnvName(ctx, optional)

Check that is a valid compute env name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateComputeEnvNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateComputeEnvNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **name** | **optional.String**| Compute env name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCredentialsName

> ValidateCredentialsName(ctx, optional)

Check that is a valid credentials name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateCredentialsNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateCredentialsNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **name** | **optional.String**| Credentials name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateOrganizationName

> ValidateOrganizationName(ctx, optional)

Check that is a valid organization name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateOrganizationNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateOrganizationNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Organization name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePipelineName

> ValidatePipelineName(ctx, optional)

Check that a Pipeline name is valid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatePipelineNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePipelineNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **orgId** | **optional.Int64**| Organization numeric identifier | 
 **name** | **optional.String**| Pipeline name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePipelineSecretName

> ValidatePipelineSecretName(ctx, optional)

Check that is a valid pipeline secret name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatePipelineSecretNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatePipelineSecretNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **name** | **optional.String**| Secret name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTeamName

> ValidateTeamName(ctx, orgId, optional)

Check that is a valid team name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
 **optional** | ***ValidateTeamNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateTeamNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Organization name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateUserName

> ValidateUserName(ctx, optional)

Check that the user name is valid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateUserNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateUserNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| User name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateWorkflowConstraints

> ValidateWorkflowConstraints(ctx, optional)

Check that the given run name of a workflow has a valid format. When the session ID is given: check that no other workflow in the system exists with the combination of both elements

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateWorkflowConstraintsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateWorkflowConstraintsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runName** | **optional.String**| Workflow run name to validate | 
 **sessionId** | **optional.String**| Workflow session ID to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowLogs

> WorkflowLogResponse WorkflowLogs(ctx, workflowId, optional)

Retrieve Workflow output logs of the Nextflow main job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string**| Workflow string identifier | 
 **optional** | ***WorkflowLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkflowLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **optional.Int64**| Workspace numeric identifier | 
 **next** | **optional.String**| Workflow log cursor | 

### Return type

[**WorkflowLogResponse**](WorkflowLogResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceValidate

> WorkspaceValidate(ctx, orgId, optional)

Check that is a valid workspace name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int64**| Organization numeric identifier | 
 **optional** | ***WorkspaceValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WorkspaceValidateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name to validate | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

