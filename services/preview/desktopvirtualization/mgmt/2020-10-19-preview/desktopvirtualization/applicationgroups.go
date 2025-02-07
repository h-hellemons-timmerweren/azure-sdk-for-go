package desktopvirtualization

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

// ApplicationGroupsClient is the client for the ApplicationGroups methods of the Desktopvirtualization service.
type ApplicationGroupsClient struct {
	BaseClient
}

// NewApplicationGroupsClient creates an instance of the ApplicationGroupsClient client.
func NewApplicationGroupsClient(subscriptionID string) ApplicationGroupsClient {
	return NewApplicationGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGroupsClientWithBaseURI creates an instance of the ApplicationGroupsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewApplicationGroupsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGroupsClient {
	return ApplicationGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update an applicationGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
// applicationGroup - object containing ApplicationGroup definitions.
func (client ApplicationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup) (result ApplicationGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: applicationGroup,
			Constraints: []validation.Constraint{{Target: "applicationGroup.ApplicationGroupProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "applicationGroup.ApplicationGroupProperties.HostPoolArmPath", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, applicationGroupName, applicationGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithJSON(applicationGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete remove an applicationGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
func (client ApplicationGroupsClient) Delete(ctx context.Context, resourceGroupName string, applicationGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an application group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
func (client ApplicationGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string) (result ApplicationGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, applicationGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) GetResponder(resp *http.Response) (result ApplicationGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list applicationGroups.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// filter - oData filter expression. Valid properties for filtering are applicationGroupType.
func (client ApplicationGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result ApplicationGroupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.agl.Response.Response != nil {
				sc = result.agl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.agl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.agl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.agl.hasNextLink() && result.agl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplicationGroupsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) ListByResourceGroupResponder(resp *http.Response) (result ApplicationGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ApplicationGroupsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ApplicationGroupList) (result ApplicationGroupList, err error) {
	req, err := lastResults.applicationGroupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGroupsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result ApplicationGroupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter)
	return
}

// ListBySubscription list applicationGroups in subscription.
// Parameters:
// filter - oData filter expression. Valid properties for filtering are applicationGroupType.
func (client ApplicationGroupsClient) ListBySubscription(ctx context.Context, filter string) (result ApplicationGroupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.agl.Response.Response != nil {
				sc = result.agl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.agl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.agl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.agl.hasNextLink() && result.agl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ApplicationGroupsClient) ListBySubscriptionPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/applicationGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) ListBySubscriptionResponder(resp *http.Response) (result ApplicationGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client ApplicationGroupsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ApplicationGroupList) (result ApplicationGroupList, err error) {
	req, err := lastResults.applicationGroupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGroupsClient) ListBySubscriptionComplete(ctx context.Context, filter string) (result ApplicationGroupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter)
	return
}

// Update update an applicationGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
// applicationGroup - object containing ApplicationGroup definitions.
func (client ApplicationGroupsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup *ApplicationGroupPatch) (result ApplicationGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGroupsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ApplicationGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, applicationGroupName, applicationGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ApplicationGroupsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ApplicationGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup *ApplicationGroupPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if applicationGroup != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(applicationGroup))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ApplicationGroupsClient) UpdateResponder(resp *http.Response) (result ApplicationGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
