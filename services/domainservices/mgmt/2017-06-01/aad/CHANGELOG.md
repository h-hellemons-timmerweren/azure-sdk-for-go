Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewDomainServiceListResultPage` parameter(s) have been changed from `(func(context.Context, DomainServiceListResult) (DomainServiceListResult, error))` to `(DomainServiceListResult, func(context.Context, DomainServiceListResult) (DomainServiceListResult, error))`
- Function `NewOperationEntityListResultPage` parameter(s) have been changed from `(func(context.Context, OperationEntityListResult) (OperationEntityListResult, error))` to `(OperationEntityListResult, func(context.Context, OperationEntityListResult) (OperationEntityListResult, error))`
- Type of `DomainServiceProperties.HealthLastEvaluated` has been changed from `*date.Time` to `*date.TimeRFC1123`

## New Content

- New const `SyncKerberosPasswordsDisabled`
- New const `SyncOnPremPasswordsEnabled`
- New const `SyncKerberosPasswordsEnabled`
- New const `SyncOnPremPasswordsDisabled`
- New function `OuContainerClient.CreateResponder(*http.Response) (OuContainer, error)`
- New function `OuContainerClient.Update(context.Context, string, string, string, ContainerAccount) (OuContainerUpdateFuture, error)`
- New function `OuContainerListResultIterator.Value() OuContainer`
- New function `OuContainerListResultPage.Values() []OuContainer`
- New function `OuContainerOperationsClient.List(context.Context) (OperationEntityListResultPage, error)`
- New function `OuContainerClient.ListSender(*http.Request) (*http.Response, error)`
- New function `OuContainerClient.GetSender(*http.Request) (*http.Response, error)`
- New function `OuContainerClient.UpdatePreparer(context.Context, string, string, string, ContainerAccount) (*http.Request, error)`
- New function `OuContainerProperties.MarshalJSON() ([]byte, error)`
- New function `NewOuContainerClient(string) OuContainerClient`
- New function `OuContainerClient.GetResponder(*http.Response) (OuContainer, error)`
- New function `*OuContainerListResultPage.Next() error`
- New function `OuContainerOperationsClient.ListPreparer(context.Context) (*http.Request, error)`
- New function `OuContainerClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)`
- New function `OuContainerListResultPage.Response() OuContainerListResult`
- New function `NewOuContainerClientWithBaseURI(string, string) OuContainerClient`
- New function `PossibleSyncKerberosPasswordsValues() []SyncKerberosPasswords`
- New function `OuContainerClient.Delete(context.Context, string, string, string) (OuContainerDeleteFuture, error)`
- New function `NewOuContainerOperationsClientWithBaseURI(string, string) OuContainerOperationsClient`
- New function `OuContainerListResult.IsEmpty() bool`
- New function `OuContainerOperationsClient.ListResponder(*http.Response) (OperationEntityListResult, error)`
- New function `OuContainerClient.CreatePreparer(context.Context, string, string, string, ContainerAccount) (*http.Request, error)`
- New function `*OuContainerListResultIterator.NextWithContext(context.Context) error`
- New function `PossibleSyncOnPremPasswordsValues() []SyncOnPremPasswords`
- New function `NewOuContainerListResultIterator(OuContainerListResultPage) OuContainerListResultIterator`
- New function `*OuContainerUpdateFuture.Result(OuContainerClient) (OuContainer, error)`
- New function `OuContainerOperationsClient.ListSender(*http.Request) (*http.Response, error)`
- New function `OuContainerListResult.MarshalJSON() ([]byte, error)`
- New function `OuContainerListResultPage.NotDone() bool`
- New function `OuContainer.MarshalJSON() ([]byte, error)`
- New function `OuContainerClient.UpdateResponder(*http.Response) (OuContainer, error)`
- New function `OuContainerClient.ListResponder(*http.Response) (OuContainerListResult, error)`
- New function `OuContainerClient.DeleteSender(*http.Request) (OuContainerDeleteFuture, error)`
- New function `OuContainerClient.Get(context.Context, string, string, string) (OuContainer, error)`
- New function `OuContainerClient.DeletePreparer(context.Context, string, string, string) (*http.Request, error)`
- New function `*OuContainer.UnmarshalJSON([]byte) error`
- New function `OuContainerClient.DeleteResponder(*http.Response) (autorest.Response, error)`
- New function `OuContainerClient.UpdateSender(*http.Request) (OuContainerUpdateFuture, error)`
- New function `NewOuContainerOperationsClient(string) OuContainerOperationsClient`
- New function `OuContainerClient.CreateSender(*http.Request) (OuContainerCreateFuture, error)`
- New function `OuContainerListResultIterator.Response() OuContainerListResult`
- New function `*OuContainerListResultPage.NextWithContext(context.Context) error`
- New function `OuContainerClient.ListPreparer(context.Context, string, string) (*http.Request, error)`
- New function `*OuContainerListResultIterator.Next() error`
- New function `NewOuContainerListResultPage(OuContainerListResult, func(context.Context, OuContainerListResult) (OuContainerListResult, error)) OuContainerListResultPage`
- New function `OuContainerClient.ListComplete(context.Context, string, string) (OuContainerListResultIterator, error)`
- New function `OuContainerClient.List(context.Context, string, string) (OuContainerListResultPage, error)`
- New function `*OuContainerCreateFuture.Result(OuContainerClient) (OuContainer, error)`
- New function `*OuContainerDeleteFuture.Result(OuContainerClient) (autorest.Response, error)`
- New function `OuContainerClient.Create(context.Context, string, string, string, ContainerAccount) (OuContainerCreateFuture, error)`
- New function `DefaultErrorResponseError.MarshalJSON() ([]byte, error)`
- New function `OuContainerListResultIterator.NotDone() bool`
- New function `OuContainerOperationsClient.ListComplete(context.Context) (OperationEntityListResultIterator, error)`
- New struct `CloudError`
- New struct `CloudErrorBody`
- New struct `ContainerAccount`
- New struct `DefaultErrorResponse`
- New struct `DefaultErrorResponseError`
- New struct `DefaultErrorResponseErrorDetailsItem`
- New struct `ForestTrust`
- New struct `OuContainer`
- New struct `OuContainerClient`
- New struct `OuContainerCreateFuture`
- New struct `OuContainerDeleteFuture`
- New struct `OuContainerListResult`
- New struct `OuContainerListResultIterator`
- New struct `OuContainerListResultPage`
- New struct `OuContainerOperationsClient`
- New struct `OuContainerProperties`
- New struct `OuContainerUpdateFuture`
- New struct `ResourceForestSettings`
- New field `Sku` in struct `DomainServiceProperties`
- New field `Version` in struct `DomainServiceProperties`
- New field `ResourceForestSettings` in struct `DomainServiceProperties`
- New field `DeploymentID` in struct `DomainServiceProperties`
- New field `DomainConfigurationType` in struct `DomainServiceProperties`
- New field `SyncKerberosPasswords` in struct `DomainSecuritySettings`
- New field `SyncOnPremPasswords` in struct `DomainSecuritySettings`
