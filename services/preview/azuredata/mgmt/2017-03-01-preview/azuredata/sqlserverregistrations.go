package azuredata

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

// SQLServerRegistrationsClient is the the AzureData management API provides a RESTful set of web APIs to manage Azure
// Data Resources.
type SQLServerRegistrationsClient struct {
	BaseClient
}

// NewSQLServerRegistrationsClient creates an instance of the SQLServerRegistrationsClient client.
func NewSQLServerRegistrationsClient(subscriptionID string, subscriptionID1 string) SQLServerRegistrationsClient {
	return NewSQLServerRegistrationsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewSQLServerRegistrationsClientWithBaseURI creates an instance of the SQLServerRegistrationsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSQLServerRegistrationsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) SQLServerRegistrationsClient {
	return SQLServerRegistrationsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// CreateOrUpdate creates or updates a SQL Server registration.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// parameters - the SQL Server registration to be created or updated.
func (client SQLServerRegistrationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters SQLServerRegistration) (result SQLServerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, SQLServerRegistrationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLServerRegistrationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters SQLServerRegistration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) CreateOrUpdateResponder(resp *http.Response) (result SQLServerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL Server registration.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
func (client SQLServerRegistrationsClient) Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SQLServerRegistrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLServerRegistrationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a SQL Server registration.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
func (client SQLServerRegistrationsClient) Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (result SQLServerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SQLServerRegistrationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLServerRegistrationsClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) GetResponder(resp *http.Response) (result SQLServerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all SQL Server registrations in a subscription.
func (client SQLServerRegistrationsClient) List(ctx context.Context) (result SQLServerRegistrationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.List")
		defer func() {
			sc := -1
			if result.ssrlr.Response.Response != nil {
				sc = result.ssrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ssrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "List", resp, "Failure sending request")
		return
	}

	result.ssrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ssrlr.hasNextLink() && result.ssrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLServerRegistrationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) ListResponder(resp *http.Response) (result SQLServerRegistrationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SQLServerRegistrationsClient) listNextResults(ctx context.Context, lastResults SQLServerRegistrationListResult) (result SQLServerRegistrationListResult, err error) {
	req, err := lastResults.sQLServerRegistrationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLServerRegistrationsClient) ListComplete(ctx context.Context) (result SQLServerRegistrationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets all SQL Server registrations in a resource group.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
func (client SQLServerRegistrationsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SQLServerRegistrationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.ssrlr.Response.Response != nil {
				sc = result.ssrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.ssrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.ssrlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.ssrlr.hasNextLink() && result.ssrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SQLServerRegistrationsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) ListByResourceGroupResponder(resp *http.Response) (result SQLServerRegistrationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SQLServerRegistrationsClient) listByResourceGroupNextResults(ctx context.Context, lastResults SQLServerRegistrationListResult) (result SQLServerRegistrationListResult, err error) {
	req, err := lastResults.sQLServerRegistrationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLServerRegistrationsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SQLServerRegistrationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates SQL Server Registration tags.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// parameters - the SQL Server Registration.
func (client SQLServerRegistrationsClient) Update(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters SQLServerRegistrationUpdate) (result SQLServerRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServerRegistrationsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, SQLServerRegistrationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServerRegistrationsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SQLServerRegistrationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, parameters SQLServerRegistrationUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServerRegistrationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SQLServerRegistrationsClient) UpdateResponder(resp *http.Response) (result SQLServerRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
