Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewListResultPage` parameter(s) have been changed from `(func(context.Context, ListResult) (ListResult, error))` to `(ListResult, func(context.Context, ListResult) (ListResult, error))`
- Function `NewOpenShiftManagedClusterListResultPage` parameter(s) have been changed from `(func(context.Context, OpenShiftManagedClusterListResult) (OpenShiftManagedClusterListResult, error))` to `(OpenShiftManagedClusterListResult, func(context.Context, OpenShiftManagedClusterListResult) (OpenShiftManagedClusterListResult, error))`
- Function `NewManagedClusterListResultPage` parameter(s) have been changed from `(func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error))` to `(ManagedClusterListResult, func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error))`
- Function `NewAgentPoolListResultPage` parameter(s) have been changed from `(func(context.Context, AgentPoolListResult) (AgentPoolListResult, error))` to `(AgentPoolListResult, func(context.Context, AgentPoolListResult) (AgentPoolListResult, error))`
- Function `ManagedClustersClient.UpgradeNodeImageVersionResponder` has been removed
- Function `ManagedClustersClient.UpgradeNodeImageVersionPreparer` has been removed
- Function `ManagedClustersClient.UpgradeNodeImageVersionSender` has been removed
- Function `ManagedClustersClient.UpgradeNodeImageVersion` has been removed
- Function `*ManagedClustersUpgradeNodeImageVersionFuture.Result` has been removed
- Struct `ManagedClustersUpgradeNodeImageVersionFuture` has been removed

## New Content

- New function `AgentPoolsClient.UpgradeNodeImageVersion(context.Context, string, string, string) (AgentPoolsUpgradeNodeImageVersionFuture, error)`
- New function `AgentPoolsClient.UpgradeNodeImageVersionPreparer(context.Context, string, string, string) (*http.Request, error)`
- New function `AgentPoolsClient.UpgradeNodeImageVersionResponder(*http.Response) (AgentPool, error)`
- New function `*AgentPoolsUpgradeNodeImageVersionFuture.Result(AgentPoolsClient) (AgentPool, error)`
- New function `AgentPoolsClient.UpgradeNodeImageVersionSender(*http.Request) (AgentPoolsUpgradeNodeImageVersionFuture, error)`
- New struct `AgentPoolsUpgradeNodeImageVersionFuture`
