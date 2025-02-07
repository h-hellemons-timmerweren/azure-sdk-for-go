package databoxedge

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

// ContainersClient is the client for the Containers methods of the Databoxedge service.
type ContainersClient struct {
	BaseClient
}

// NewContainersClient creates an instance of the ContainersClient client.
func NewContainersClient(subscriptionID string) ContainersClient {
	return NewContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainersClientWithBaseURI creates an instance of the ContainersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewContainersClientWithBaseURI(baseURI string, subscriptionID string) ContainersClient {
	return ContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the Storage Account Name
// containerName - the container name.
// containerParameter - the container properties.
// resourceGroupName - the resource group name.
func (client ContainersClient) CreateOrUpdate(ctx context.Context, deviceName string, storageAccountName string, containerName string, containerParameter Container, resourceGroupName string) (result ContainersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: containerParameter,
			Constraints: []validation.Constraint{{Target: "containerParameter.ContainerProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databoxedge.ContainersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, storageAccountName, containerName, containerParameter, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainersClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, storageAccountName string, containerName string, containerParameter Container, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":      autorest.Encode("path", containerName),
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", pathParameters),
		autorest.WithJSON(containerParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) CreateOrUpdateSender(req *http.Request) (future ContainersCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ContainersClient) CreateOrUpdateResponder(resp *http.Response) (result Container, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the container on the Data Box Edge/Data Box Gateway device.
// Parameters:
// deviceName - the device name.
// storageAccountName - the Storage Account Name
// containerName - the container name.
// resourceGroupName - the resource group name.
func (client ContainersClient) Delete(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (result ContainersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, storageAccountName, containerName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContainersClient) DeletePreparer(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":      autorest.Encode("path", containerName),
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) DeleteSender(req *http.Request) (future ContainersDeleteFuture, err error) {
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
func (client ContainersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the Storage Account Name
// containerName - the container Name
// resourceGroupName - the resource group name.
func (client ContainersClient) Get(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (result Container, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, storageAccountName, containerName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ContainersClient) GetPreparer(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":      autorest.Encode("path", containerName),
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ContainersClient) GetResponder(resp *http.Response) (result Container, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStorageAccount sends the list by storage account request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the storage Account name.
// resourceGroupName - the resource group name.
func (client ContainersClient) ListByStorageAccount(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (result ContainerListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.ListByStorageAccount")
		defer func() {
			sc := -1
			if result.cl.Response.Response != nil {
				sc = result.cl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByStorageAccountNextResults
	req, err := client.ListByStorageAccountPreparer(ctx, deviceName, storageAccountName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "ListByStorageAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStorageAccountSender(req)
	if err != nil {
		result.cl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "ListByStorageAccount", resp, "Failure sending request")
		return
	}

	result.cl, err = client.ListByStorageAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "ListByStorageAccount", resp, "Failure responding to request")
		return
	}
	if result.cl.hasNextLink() && result.cl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByStorageAccountPreparer prepares the ListByStorageAccount request.
func (client ContainersClient) ListByStorageAccountPreparer(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStorageAccountSender sends the ListByStorageAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) ListByStorageAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByStorageAccountResponder handles the response to the ListByStorageAccount request. The method always
// closes the http.Response Body.
func (client ContainersClient) ListByStorageAccountResponder(resp *http.Response) (result ContainerList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByStorageAccountNextResults retrieves the next set of results, if any.
func (client ContainersClient) listByStorageAccountNextResults(ctx context.Context, lastResults ContainerList) (result ContainerList, err error) {
	req, err := lastResults.containerListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "listByStorageAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByStorageAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "listByStorageAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByStorageAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "listByStorageAccountNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByStorageAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainersClient) ListByStorageAccountComplete(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string) (result ContainerListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.ListByStorageAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByStorageAccount(ctx, deviceName, storageAccountName, resourceGroupName)
	return
}

// Refresh sends the refresh request.
// Parameters:
// deviceName - the device name.
// storageAccountName - the Storage Account Name
// containerName - the container name.
// resourceGroupName - the resource group name.
func (client ContainersClient) Refresh(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (result ContainersRefreshFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainersClient.Refresh")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshPreparer(ctx, deviceName, storageAccountName, containerName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Refresh", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.ContainersClient", "Refresh", result.Response(), "Failure sending request")
		return
	}

	return
}

// RefreshPreparer prepares the Refresh request.
func (client ContainersClient) RefreshPreparer(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":      autorest.Encode("path", containerName),
		"deviceName":         autorest.Encode("path", deviceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"storageAccountName": autorest.Encode("path", storageAccountName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}/refresh", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshSender sends the Refresh request. The method will close the
// http.Response Body if it receives an error.
func (client ContainersClient) RefreshSender(req *http.Request) (future ContainersRefreshFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RefreshResponder handles the response to the Refresh request. The method always
// closes the http.Response Body.
func (client ContainersClient) RefreshResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
