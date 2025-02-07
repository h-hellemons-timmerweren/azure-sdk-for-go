package hybriddata

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataManagersClient is the client for the DataManagers methods of the Hybriddata service.
type DataManagersClient struct {
	BaseClient
}

// NewDataManagersClient creates an instance of the DataManagersClient client.
func NewDataManagersClient(subscriptionID string) DataManagersClient {
	return NewDataManagersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataManagersClientWithBaseURI creates an instance of the DataManagersClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataManagersClientWithBaseURI(baseURI string, subscriptionID string) DataManagersClient {
	return DataManagersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new data manager resource with the specified parameters. Existing resources cannot be updated with
// this API
// and should instead be updated with the Update data manager resource API.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// dataManager - data manager resource details from request body.
func (client DataManagersClient) Create(ctx context.Context, resourceGroupName string, dataManagerName string, dataManager DataManager) (result DataManagersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataManagersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, dataManagerName, dataManager)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DataManagersClient) CreatePreparer(ctx context.Context, resourceGroupName string, dataManagerName string, dataManager DataManager) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", pathParameters),
		autorest.WithJSON(dataManager),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) CreateSender(req *http.Request) (future DataManagersCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DataManagersClient) CreateResponder(resp *http.Response) (result DataManager, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a data manager resource in Microsoft Azure.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client DataManagersClient) Delete(ctx context.Context, resourceGroupName string, dataManagerName string) (result DataManagersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataManagersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DataManagersClient) DeletePreparer(ctx context.Context, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) DeleteSender(req *http.Request) (future DataManagersDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DataManagersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified data manager resource.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client DataManagersClient) Get(ctx context.Context, resourceGroupName string, dataManagerName string) (result DataManager, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataManagersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataManagersClient) GetPreparer(ctx context.Context, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataManagersClient) GetResponder(resp *http.Response) (result DataManager, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the data manager resources available under the subscription.
func (client DataManagersClient) List(ctx context.Context) (result DataManagerList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DataManagersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridData/dataManagers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DataManagersClient) ListResponder(resp *http.Response) (result DataManagerList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists all the data manager resources available under the given resource group.
// Parameters:
// resourceGroupName - the Resource Group Name
func (client DataManagersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DataManagerList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DataManagersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DataManagersClient) ListByResourceGroupResponder(resp *http.Response) (result DataManagerList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the properties of an existing data manager resource.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
// dataManagerUpdateParameter - data manager resource details from request body.
// ifMatch - defines the If-Match condition. The patch will be performed only if the ETag of the data manager
// resource on the server matches this value.
func (client DataManagersClient) Update(ctx context.Context, resourceGroupName string, dataManagerName string, dataManagerUpdateParameter DataManagerUpdateParameter, ifMatch string) (result DataManagersUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataManagersClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataManagersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, dataManagerName, dataManagerUpdateParameter, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataManagersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DataManagersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, dataManagerName string, dataManagerUpdateParameter DataManagerUpdateParameter, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", pathParameters),
		autorest.WithJSON(dataManagerUpdateParameter),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DataManagersClient) UpdateSender(req *http.Request) (future DataManagersUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DataManagersClient) UpdateResponder(resp *http.Response) (result DataManager, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
