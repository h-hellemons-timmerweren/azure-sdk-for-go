package logic

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

// IntegrationAccountsClient is the REST API for Azure Logic Apps.
type IntegrationAccountsClient struct {
	BaseClient
}

// NewIntegrationAccountsClient creates an instance of the IntegrationAccountsClient client.
func NewIntegrationAccountsClient(subscriptionID string) IntegrationAccountsClient {
	return NewIntegrationAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationAccountsClientWithBaseURI creates an instance of the IntegrationAccountsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewIntegrationAccountsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountsClient {
	return IntegrationAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// integrationAccount - the integration account.
func (client IntegrationAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount) (result IntegrationAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, integrationAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", pathParameters),
		autorest.WithJSON(integrationAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
func (client IntegrationAccountsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationAccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
func (client IntegrationAccountsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string) (result IntegrationAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationAccountsClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) GetResponder(resp *http.Response) (result IntegrationAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of integration accounts by resource group.
// Parameters:
// resourceGroupName - the resource group name.
// top - the number of items to be included in the result.
func (client IntegrationAccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result IntegrationAccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ialr.Response.Response != nil {
				sc = result.ialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ialr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ialr.hasNextLink() && result.ialr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client IntegrationAccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) ListByResourceGroupResponder(resp *http.Response) (result IntegrationAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client IntegrationAccountsClient) listByResourceGroupNextResults(ctx context.Context, lastResults IntegrationAccountListResult) (result IntegrationAccountListResult, err error) {
	req, err := lastResults.integrationAccountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationAccountsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result IntegrationAccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top)
	return
}

// ListBySubscription gets a list of integration accounts by subscription.
// Parameters:
// top - the number of items to be included in the result.
func (client IntegrationAccountsClient) ListBySubscription(ctx context.Context, top *int32) (result IntegrationAccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.ialr.Response.Response != nil {
				sc = result.ialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.ialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.ialr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.ialr.hasNextLink() && result.ialr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client IntegrationAccountsClient) ListBySubscriptionPreparer(ctx context.Context, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) ListBySubscriptionResponder(resp *http.Response) (result IntegrationAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client IntegrationAccountsClient) listBySubscriptionNextResults(ctx context.Context, lastResults IntegrationAccountListResult) (result IntegrationAccountListResult, err error) {
	req, err := lastResults.integrationAccountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationAccountsClient) ListBySubscriptionComplete(ctx context.Context, top *int32) (result IntegrationAccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, top)
	return
}

// ListCallbackURL gets the integration account callback URL.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// parameters - the callback URL parameters.
func (client IntegrationAccountsClient) ListCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters) (result CallbackURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListCallbackURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListCallbackURLPreparer(ctx, resourceGroupName, integrationAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListCallbackURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCallbackURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListCallbackURL", resp, "Failure sending request")
		return
	}

	result, err = client.ListCallbackURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListCallbackURL", resp, "Failure responding to request")
		return
	}

	return
}

// ListCallbackURLPreparer prepares the ListCallbackURL request.
func (client IntegrationAccountsClient) ListCallbackURLPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listCallbackUrl", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCallbackURLSender sends the ListCallbackURL request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) ListCallbackURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCallbackURLResponder handles the response to the ListCallbackURL request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) ListCallbackURLResponder(resp *http.Response) (result CallbackURL, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListKeyVaultKeys gets the integration account's Key Vault keys.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// listKeyVaultKeys - the key vault parameters.
func (client IntegrationAccountsClient) ListKeyVaultKeys(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition) (result KeyVaultKeyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.ListKeyVaultKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: listKeyVaultKeys,
			Constraints: []validation.Constraint{{Target: "listKeyVaultKeys.KeyVault", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("logic.IntegrationAccountsClient", "ListKeyVaultKeys", err.Error())
	}

	req, err := client.ListKeyVaultKeysPreparer(ctx, resourceGroupName, integrationAccountName, listKeyVaultKeys)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListKeyVaultKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeyVaultKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListKeyVaultKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeyVaultKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "ListKeyVaultKeys", resp, "Failure responding to request")
		return
	}

	return
}

// ListKeyVaultKeysPreparer prepares the ListKeyVaultKeys request.
func (client IntegrationAccountsClient) ListKeyVaultKeysPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listKeyVaultKeys", pathParameters),
		autorest.WithJSON(listKeyVaultKeys),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeyVaultKeysSender sends the ListKeyVaultKeys request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) ListKeyVaultKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListKeyVaultKeysResponder handles the response to the ListKeyVaultKeys request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) ListKeyVaultKeysResponder(resp *http.Response) (result KeyVaultKeyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// LogTrackingEvents logs the integration account's tracking events.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// logTrackingEvents - the callback URL parameters.
func (client IntegrationAccountsClient) LogTrackingEvents(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.LogTrackingEvents")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: logTrackingEvents,
			Constraints: []validation.Constraint{{Target: "logTrackingEvents.SourceType", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "logTrackingEvents.Events", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("logic.IntegrationAccountsClient", "LogTrackingEvents", err.Error())
	}

	req, err := client.LogTrackingEventsPreparer(ctx, resourceGroupName, integrationAccountName, logTrackingEvents)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "LogTrackingEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.LogTrackingEventsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "LogTrackingEvents", resp, "Failure sending request")
		return
	}

	result, err = client.LogTrackingEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "LogTrackingEvents", resp, "Failure responding to request")
		return
	}

	return
}

// LogTrackingEventsPreparer prepares the LogTrackingEvents request.
func (client IntegrationAccountsClient) LogTrackingEventsPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/logTrackingEvents", pathParameters),
		autorest.WithJSON(logTrackingEvents),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// LogTrackingEventsSender sends the LogTrackingEvents request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) LogTrackingEventsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// LogTrackingEventsResponder handles the response to the LogTrackingEvents request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) LogTrackingEventsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateAccessKey regenerates the integration account access key.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// regenerateAccessKey - the access key type.
func (client IntegrationAccountsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter) (result IntegrationAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.RegenerateAccessKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateAccessKeyPreparer(ctx, resourceGroupName, integrationAccountName, regenerateAccessKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "RegenerateAccessKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateAccessKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "RegenerateAccessKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateAccessKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "RegenerateAccessKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateAccessKeyPreparer prepares the RegenerateAccessKey request.
func (client IntegrationAccountsClient) RegenerateAccessKeyPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/regenerateAccessKey", pathParameters),
		autorest.WithJSON(regenerateAccessKey),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateAccessKeySender sends the RegenerateAccessKey request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) RegenerateAccessKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateAccessKeyResponder handles the response to the RegenerateAccessKey request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) RegenerateAccessKeyResponder(resp *http.Response) (result IntegrationAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// integrationAccount - the integration account.
func (client IntegrationAccountsClient) Update(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount) (result IntegrationAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, integrationAccountName, integrationAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client IntegrationAccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", pathParameters),
		autorest.WithJSON(integrationAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client IntegrationAccountsClient) UpdateResponder(resp *http.Response) (result IntegrationAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
