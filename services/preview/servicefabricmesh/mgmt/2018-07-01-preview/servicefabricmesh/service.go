package servicefabricmesh

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

// ServiceClient is the service Fabric Mesh Management Client
type ServiceClient struct {
	BaseClient
}

// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient(subscriptionID string) ServiceClient {
	return NewServiceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceClientWithBaseURI creates an instance of the ServiceClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return ServiceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get the operation returns the properties of the service.
// Parameters:
// resourceGroupName - azure resource group name
// applicationName - the identity of the application.
// serviceName - the identity of the service.
func (client ServiceClient) Get(ctx context.Context, resourceGroupName string, applicationName string, serviceName string) (result ServiceResourceDescription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   applicationName,
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       serviceName,
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceClient) GetResponder(resp *http.Response) (result ServiceResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByApplicationName gets the information about all services of a given service of an application. The information
// includes the runtime properties of the service instance.
// Parameters:
// resourceGroupName - azure resource group name
// applicationName - the identity of the application.
func (client ServiceClient) ListByApplicationName(ctx context.Context, resourceGroupName string, applicationName string) (result ServiceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ListByApplicationName")
		defer func() {
			sc := -1
			if result.sl.Response.Response != nil {
				sc = result.sl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByApplicationNameNextResults
	req, err := client.ListByApplicationNamePreparer(ctx, resourceGroupName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "ListByApplicationName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByApplicationNameSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "ListByApplicationName", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListByApplicationNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "ListByApplicationName", resp, "Failure responding to request")
		return
	}
	if result.sl.hasNextLink() && result.sl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByApplicationNamePreparer prepares the ListByApplicationName request.
func (client ServiceClient) ListByApplicationNamePreparer(ctx context.Context, resourceGroupName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   applicationName,
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationName}/services", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByApplicationNameSender sends the ListByApplicationName request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ListByApplicationNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByApplicationNameResponder handles the response to the ListByApplicationName request. The method always
// closes the http.Response Body.
func (client ServiceClient) ListByApplicationNameResponder(resp *http.Response) (result ServiceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByApplicationNameNextResults retrieves the next set of results, if any.
func (client ServiceClient) listByApplicationNameNextResults(ctx context.Context, lastResults ServiceList) (result ServiceList, err error) {
	req, err := lastResults.serviceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "listByApplicationNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByApplicationNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "listByApplicationNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByApplicationNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.ServiceClient", "listByApplicationNameNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByApplicationNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceClient) ListByApplicationNameComplete(ctx context.Context, resourceGroupName string, applicationName string) (result ServiceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ListByApplicationName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByApplicationName(ctx, resourceGroupName, applicationName)
	return
}
