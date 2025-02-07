package resourcehealth

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

// ChildResourcesClient is the the Resource Health Client.
type ChildResourcesClient struct {
	BaseClient
}

// NewChildResourcesClient creates an instance of the ChildResourcesClient client.
func NewChildResourcesClient(subscriptionID string) ChildResourcesClient {
	return NewChildResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewChildResourcesClientWithBaseURI creates an instance of the ChildResourcesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewChildResourcesClientWithBaseURI(baseURI string, subscriptionID string) ChildResourcesClient {
	return ChildResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the all the children and its current health status for a parent resource. Use the nextLink property in
// the response to get the next page of children current health
// Parameters:
// resourceURI - the fully qualified ID of the resource, including the resource name and resource type.
// Currently the API only support not nested parent resource type:
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// filter - the filter to apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand - setting $expand=recommendedactions in url query expands the recommendedactions in the response.
func (client ChildResourcesClient) List(ctx context.Context, resourceURI string, filter string, expand string) (result AvailabilityStatusListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChildResourcesClient.List")
		defer func() {
			sc := -1
			if result.aslr.Response.Response != nil {
				sc = result.aslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceURI, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.aslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.aslr.hasNextLink() && result.aslr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ChildResourcesClient) ListPreparer(ctx context.Context, resourceURI string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2015-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.ResourceHealth/childResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ChildResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ChildResourcesClient) ListResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ChildResourcesClient) listNextResults(ctx context.Context, lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.availabilityStatusListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.ChildResourcesClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ChildResourcesClient) ListComplete(ctx context.Context, resourceURI string, filter string, expand string) (result AvailabilityStatusListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChildResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceURI, filter, expand)
	return
}
