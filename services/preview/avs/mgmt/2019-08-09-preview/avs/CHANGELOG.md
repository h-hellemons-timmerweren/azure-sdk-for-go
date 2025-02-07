Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewOperationListPage` parameter(s) have been changed from `(func(context.Context, OperationList) (OperationList, error))` to `(OperationList, func(context.Context, OperationList) (OperationList, error))`
- Function `NewClusterListPage` parameter(s) have been changed from `(func(context.Context, ClusterList) (ClusterList, error))` to `(ClusterList, func(context.Context, ClusterList) (ClusterList, error))`
- Function `NewPrivateCloudListPage` parameter(s) have been changed from `(func(context.Context, PrivateCloudList) (PrivateCloudList, error))` to `(PrivateCloudList, func(context.Context, PrivateCloudList) (PrivateCloudList, error))`

## New Content

- New function `Operation.MarshalJSON() ([]byte, error)`
- New struct `LogSpecification`
- New struct `MetricDimension`
- New struct `MetricSpecification`
- New struct `OperationProperties`
- New struct `ServiceSpecification`
- New field `Origin` in struct `Operation`
- New field `Properties` in struct `Operation`
- New field `IsDataAction` in struct `Operation`
