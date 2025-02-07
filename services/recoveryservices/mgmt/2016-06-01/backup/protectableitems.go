package backup

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProtectableItemsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectableItemsClient struct {
	BaseClient
}

// NewProtectableItemsClient creates an instance of the ProtectableItemsClient client.
func NewProtectableItemsClient(subscriptionID string) ProtectableItemsClient {
	return NewProtectableItemsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectableItemsClientWithBaseURI creates an instance of the ProtectableItemsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewProtectableItemsClientWithBaseURI(baseURI string, subscriptionID string) ProtectableItemsClient {
	return ProtectableItemsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List based on the query filter and the pagination parameters, this operation provides a pageable list of objects
// within the subscription that can be protected.
// Parameters:
// vaultName - the name of the Recovery Services vault.
// resourceGroupName - the name of the resource group associated with the Recovery Services vault.
// filter - using the following query filters, you can sort a specific backup item based on: type of backup
// item, status, name of the item, and more.  providerType eq { AzureIaasVM, MAB, DPM, AzureBackupServer,
// AzureSql } and status eq { NotProtected , Protecting , Protected } and friendlyName {name} and skipToken eq
// {string which provides the next set of list} and topToken eq {int} and backupManagementType eq {
// AzureIaasVM, MAB, DPM, AzureBackupServer, AzureSql }.
// skipToken - the Skip Token filter.
func (client ProtectableItemsClient) List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result WorkloadProtectableItemResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectableItemsClient.List")
		defer func() {
			sc := -1
			if result.wpirl.Response.Response != nil {
				sc = result.wpirl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wpirl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "List", resp, "Failure sending request")
		return
	}

	result.wpirl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.wpirl.hasNextLink() && result.wpirl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ProtectableItemsClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectableItems", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectableItemsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectableItemsClient) ListResponder(resp *http.Response) (result WorkloadProtectableItemResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProtectableItemsClient) listNextResults(ctx context.Context, lastResults WorkloadProtectableItemResourceList) (result WorkloadProtectableItemResourceList, err error) {
	req, err := lastResults.workloadProtectableItemResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableItemsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProtectableItemsClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result WorkloadProtectableItemResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectableItemsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, filter, skipToken)
	return
}
