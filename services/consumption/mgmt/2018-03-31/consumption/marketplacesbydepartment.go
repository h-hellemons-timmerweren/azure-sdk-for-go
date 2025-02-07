package consumption

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

// MarketplacesByDepartmentClient is the consumption management client provides access to consumption resources for
// Azure Enterprise Subscriptions.
type MarketplacesByDepartmentClient struct {
	BaseClient
}

// NewMarketplacesByDepartmentClient creates an instance of the MarketplacesByDepartmentClient client.
func NewMarketplacesByDepartmentClient(subscriptionID string) MarketplacesByDepartmentClient {
	return NewMarketplacesByDepartmentClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMarketplacesByDepartmentClientWithBaseURI creates an instance of the MarketplacesByDepartmentClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewMarketplacesByDepartmentClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesByDepartmentClient {
	return MarketplacesByDepartmentClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the marketplaces for a scope by departmentId and current billing period. Marketplaces are available via
// this API only for May 1, 2014 or later.
// Parameters:
// departmentID - department ID
// filter - may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// top - may be used to limit the number of results to the most recent N marketplaces.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
func (client MarketplacesByDepartmentClient) List(ctx context.Context, departmentID string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByDepartmentClient.List")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesByDepartmentClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, departmentID, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "List", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client MarketplacesByDepartmentClient) ListPreparer(ctx context.Context, departmentID string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"departmentId": autorest.Encode("path", departmentID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesByDepartmentClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MarketplacesByDepartmentClient) ListResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MarketplacesByDepartmentClient) listNextResults(ctx context.Context, lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesByDepartmentClient) ListComplete(ctx context.Context, departmentID string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByDepartmentClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, departmentID, filter, top, skiptoken)
	return
}

// ListByBillingPeriod lists the marketplaces for a scope by billing period and departmentId. Marketplaces are
// available via this API only for May 1, 2014 or later.
// Parameters:
// departmentID - department ID
// billingPeriodName - billing Period Name.
// filter - may be used to filter marketplaces by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// top - may be used to limit the number of results to the most recent N marketplaces.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
func (client MarketplacesByDepartmentClient) ListByBillingPeriod(ctx context.Context, departmentID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByDepartmentClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.mlr.Response.Response != nil {
				sc = result.mlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.MarketplacesByDepartmentClient", "ListByBillingPeriod", err.Error())
	}

	result.fn = client.listByBillingPeriodNextResults
	req, err := client.ListByBillingPeriodPreparer(ctx, departmentID, billingPeriodName, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "ListByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.mlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "ListByBillingPeriod", resp, "Failure sending request")
		return
	}

	result.mlr, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "ListByBillingPeriod", resp, "Failure responding to request")
		return
	}
	if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingPeriodPreparer prepares the ListByBillingPeriod request.
func (client MarketplacesByDepartmentClient) ListByBillingPeriodPreparer(ctx context.Context, departmentID string, billingPeriodName string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"departmentId":      autorest.Encode("path", departmentID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/marketplaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingPeriodSender sends the ListByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client MarketplacesByDepartmentClient) ListByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingPeriodResponder handles the response to the ListByBillingPeriod request. The method always
// closes the http.Response Body.
func (client MarketplacesByDepartmentClient) ListByBillingPeriodResponder(resp *http.Response) (result MarketplacesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingPeriodNextResults retrieves the next set of results, if any.
func (client MarketplacesByDepartmentClient) listByBillingPeriodNextResults(ctx context.Context, lastResults MarketplacesListResult) (result MarketplacesListResult, err error) {
	req, err := lastResults.marketplacesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listByBillingPeriodNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listByBillingPeriodNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.MarketplacesByDepartmentClient", "listByBillingPeriodNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByBillingPeriodComplete enumerates all values, automatically crossing page boundaries as required.
func (client MarketplacesByDepartmentClient) ListByBillingPeriodComplete(ctx context.Context, departmentID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result MarketplacesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MarketplacesByDepartmentClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingPeriod(ctx, departmentID, billingPeriodName, filter, top, skiptoken)
	return
}
