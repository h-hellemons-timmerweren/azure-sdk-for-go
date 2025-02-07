package security

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

// IoTSecuritySolutionsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type IoTSecuritySolutionsClient struct {
	BaseClient
}

// NewIoTSecuritySolutionsClient creates an instance of the IoTSecuritySolutionsClient client.
func NewIoTSecuritySolutionsClient(subscriptionID string, ascLocation string) IoTSecuritySolutionsClient {
	return NewIoTSecuritySolutionsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIoTSecuritySolutionsClientWithBaseURI creates an instance of the IoTSecuritySolutionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewIoTSecuritySolutionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IoTSecuritySolutionsClient {
	return IoTSecuritySolutionsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// List list of security solutions
// Parameters:
// filter - filter the Security Solution with OData syntax. supporting filter by iotHubs
func (client IoTSecuritySolutionsClient) List(ctx context.Context, filter string) (result IoTSecuritySolutionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsClient.List")
		defer func() {
			sc := -1
			if result.itssl.Response.Response != nil {
				sc = result.itssl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itssl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "List", resp, "Failure sending request")
		return
	}

	result.itssl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.itssl.hasNextLink() && result.itssl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client IoTSecuritySolutionsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotSecuritySolutions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsClient) ListResponder(resp *http.Response) (result IoTSecuritySolutionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IoTSecuritySolutionsClient) listNextResults(ctx context.Context, lastResults IoTSecuritySolutionsList) (result IoTSecuritySolutionsList, err error) {
	req, err := lastResults.ioTSecuritySolutionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IoTSecuritySolutionsClient) ListComplete(ctx context.Context, filter string) (result IoTSecuritySolutionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}
