Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewOperationsListPage` parameter(s) have been changed from `(func(context.Context, OperationsList) (OperationsList, error))` to `(OperationsList, func(context.Context, OperationsList) (OperationsList, error))`
- Function `NewAlertsListPage` parameter(s) have been changed from `(func(context.Context, AlertsList) (AlertsList, error))` to `(AlertsList, func(context.Context, AlertsList) (AlertsList, error))`
