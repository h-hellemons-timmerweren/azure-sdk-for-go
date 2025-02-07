package adhybridhealthservice

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

// AdDomainServiceMembersClient is the REST APIs for Azure Active Directory Connect Health
type AdDomainServiceMembersClient struct {
	BaseClient
}

// NewAdDomainServiceMembersClient creates an instance of the AdDomainServiceMembersClient client.
func NewAdDomainServiceMembersClient() AdDomainServiceMembersClient {
	return NewAdDomainServiceMembersClientWithBaseURI(DefaultBaseURI)
}

// NewAdDomainServiceMembersClientWithBaseURI creates an instance of the AdDomainServiceMembersClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAdDomainServiceMembersClientWithBaseURI(baseURI string) AdDomainServiceMembersClient {
	return AdDomainServiceMembersClient{NewWithBaseURI(baseURI)}
}

// List gets the details of the servers, for a given Active Directory Domain Service, that are onboarded to Azure
// Active Directory Connect Health.
// Parameters:
// serviceName - the name of the service.
// isGroupbySite - indicates if the result should be grouped by site or not.
// filter - the server property filter to apply.
// query - the custom query.
// takeCount - the take count , which specifies the number of elements that can be returned from a sequence.
func (client AdDomainServiceMembersClient) List(ctx context.Context, serviceName string, isGroupbySite bool, filter string, query string, takeCount *int32) (result AddsServiceMembersPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdDomainServiceMembersClient.List")
		defer func() {
			sc := -1
			if result.asm.Response.Response != nil {
				sc = result.asm.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, serviceName, isGroupbySite, filter, query, takeCount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.asm.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "List", resp, "Failure sending request")
		return
	}

	result.asm, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.asm.hasNextLink() && result.asm.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client AdDomainServiceMembersClient) ListPreparer(ctx context.Context, serviceName string, isGroupbySite bool, filter string, query string, takeCount *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version":      APIVersion,
		"isGroupbySite":    autorest.Encode("query", isGroupbySite),
		"nextPartitionKey": autorest.Encode("query", ""),
		"nextRowKey":       autorest.Encode("query", ""),
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(query) > 0 {
		queryParameters["query"] = autorest.Encode("query", query)
	}
	if takeCount != nil {
		queryParameters["takeCount"] = autorest.Encode("query", *takeCount)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addomainservicemembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AdDomainServiceMembersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AdDomainServiceMembersClient) ListResponder(resp *http.Response) (result AddsServiceMembers, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AdDomainServiceMembersClient) listNextResults(ctx context.Context, lastResults AddsServiceMembers) (result AddsServiceMembers, err error) {
	req, err := lastResults.addsServiceMembersPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.AdDomainServiceMembersClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AdDomainServiceMembersClient) ListComplete(ctx context.Context, serviceName string, isGroupbySite bool, filter string, query string, takeCount *int32) (result AddsServiceMembersIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdDomainServiceMembersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, serviceName, isGroupbySite, filter, query, takeCount)
	return
}
