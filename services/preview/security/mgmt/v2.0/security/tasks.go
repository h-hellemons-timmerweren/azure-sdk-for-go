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

// TasksClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type TasksClient struct {
	BaseClient
}

// NewTasksClient creates an instance of the TasksClient client.
func NewTasksClient(subscriptionID string, ascLocation string) TasksClient {
	return NewTasksClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewTasksClientWithBaseURI creates an instance of the TasksClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) TasksClient {
	return TasksClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// GetResourceGroupLevelTask recommended tasks that will help improve the security of the subscription proactively
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// taskName - name of the task object, will be a GUID
func (client TasksClient) GetResourceGroupLevelTask(ctx context.Context, resourceGroupName string, taskName string) (result Task, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.GetResourceGroupLevelTask")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "GetResourceGroupLevelTask", err.Error())
	}

	req, err := client.GetResourceGroupLevelTaskPreparer(ctx, resourceGroupName, taskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetResourceGroupLevelTask", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceGroupLevelTaskSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetResourceGroupLevelTask", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceGroupLevelTaskResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetResourceGroupLevelTask", resp, "Failure responding to request")
		return
	}

	return
}

// GetResourceGroupLevelTaskPreparer prepares the GetResourceGroupLevelTask request.
func (client TasksClient) GetResourceGroupLevelTaskPreparer(ctx context.Context, resourceGroupName string, taskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":       autorest.Encode("path", client.AscLocation),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"taskName":          autorest.Encode("path", taskName),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceGroupLevelTaskSender sends the GetResourceGroupLevelTask request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) GetResourceGroupLevelTaskSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResourceGroupLevelTaskResponder handles the response to the GetResourceGroupLevelTask request. The method always
// closes the http.Response Body.
func (client TasksClient) GetResourceGroupLevelTaskResponder(resp *http.Response) (result Task, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSubscriptionLevelTask recommended tasks that will help improve the security of the subscription proactively
// Parameters:
// taskName - name of the task object, will be a GUID
func (client TasksClient) GetSubscriptionLevelTask(ctx context.Context, taskName string) (result Task, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.GetSubscriptionLevelTask")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "GetSubscriptionLevelTask", err.Error())
	}

	req, err := client.GetSubscriptionLevelTaskPreparer(ctx, taskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetSubscriptionLevelTask", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSubscriptionLevelTaskSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetSubscriptionLevelTask", resp, "Failure sending request")
		return
	}

	result, err = client.GetSubscriptionLevelTaskResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "GetSubscriptionLevelTask", resp, "Failure responding to request")
		return
	}

	return
}

// GetSubscriptionLevelTaskPreparer prepares the GetSubscriptionLevelTask request.
func (client TasksClient) GetSubscriptionLevelTaskPreparer(ctx context.Context, taskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"taskName":       autorest.Encode("path", taskName),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSubscriptionLevelTaskSender sends the GetSubscriptionLevelTask request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) GetSubscriptionLevelTaskSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSubscriptionLevelTaskResponder handles the response to the GetSubscriptionLevelTask request. The method always
// closes the http.Response Body.
func (client TasksClient) GetSubscriptionLevelTaskResponder(resp *http.Response) (result Task, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List recommended tasks that will help improve the security of the subscription proactively
// Parameters:
// filter - oData filter. Optional.
func (client TasksClient) List(ctx context.Context, filter string) (result TaskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.List")
		defer func() {
			sc := -1
			if result.tl.Response.Response != nil {
				sc = result.tl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TasksClient", "List", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tl.hasNextLink() && result.tl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client TasksClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TasksClient) ListResponder(resp *http.Response) (result TaskList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TasksClient) listNextResults(ctx context.Context, lastResults TaskList) (result TaskList, err error) {
	req, err := lastResults.taskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TasksClient) ListComplete(ctx context.Context, filter string) (result TaskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.List")
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

// ListByHomeRegion recommended tasks that will help improve the security of the subscription proactively
// Parameters:
// filter - oData filter. Optional.
func (client TasksClient) ListByHomeRegion(ctx context.Context, filter string) (result TaskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.tl.Response.Response != nil {
				sc = result.tl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "ListByHomeRegion", err.Error())
	}

	result.fn = client.listByHomeRegionNextResults
	req, err := client.ListByHomeRegionPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByHomeRegion", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByHomeRegion", resp, "Failure responding to request")
		return
	}
	if result.tl.hasNextLink() && result.tl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByHomeRegionPreparer prepares the ListByHomeRegion request.
func (client TasksClient) ListByHomeRegionPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHomeRegionSender sends the ListByHomeRegion request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) ListByHomeRegionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHomeRegionResponder handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (client TasksClient) ListByHomeRegionResponder(resp *http.Response) (result TaskList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHomeRegionNextResults retrieves the next set of results, if any.
func (client TasksClient) listByHomeRegionNextResults(ctx context.Context, lastResults TaskList) (result TaskList, err error) {
	req, err := lastResults.taskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listByHomeRegionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listByHomeRegionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "listByHomeRegionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByHomeRegionComplete enumerates all values, automatically crossing page boundaries as required.
func (client TasksClient) ListByHomeRegionComplete(ctx context.Context, filter string) (result TaskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHomeRegion(ctx, filter)
	return
}

// ListByResourceGroup recommended tasks that will help improve the security of the subscription proactively
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// filter - oData filter. Optional.
func (client TasksClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result TaskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.tl.Response.Response != nil {
				sc = result.tl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.tl.hasNextLink() && result.tl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client TasksClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":       autorest.Encode("path", client.AscLocation),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client TasksClient) ListByResourceGroupResponder(resp *http.Response) (result TaskList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client TasksClient) listByResourceGroupNextResults(ctx context.Context, lastResults TaskList) (result TaskList, err error) {
	req, err := lastResults.taskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.TasksClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client TasksClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result TaskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.ListByResourceGroup")
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

// UpdateResourceGroupLevelTaskState recommended tasks that will help improve the security of the subscription
// proactively
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// taskName - name of the task object, will be a GUID
// taskUpdateActionType - type of the action to do on the task
func (client TasksClient) UpdateResourceGroupLevelTaskState(ctx context.Context, resourceGroupName string, taskName string, taskUpdateActionType string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.UpdateResourceGroupLevelTaskState")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "UpdateResourceGroupLevelTaskState", err.Error())
	}

	req, err := client.UpdateResourceGroupLevelTaskStatePreparer(ctx, resourceGroupName, taskName, taskUpdateActionType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateResourceGroupLevelTaskState", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateResourceGroupLevelTaskStateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateResourceGroupLevelTaskState", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResourceGroupLevelTaskStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateResourceGroupLevelTaskState", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateResourceGroupLevelTaskStatePreparer prepares the UpdateResourceGroupLevelTaskState request.
func (client TasksClient) UpdateResourceGroupLevelTaskStatePreparer(ctx context.Context, resourceGroupName string, taskName string, taskUpdateActionType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":          autorest.Encode("path", client.AscLocation),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"taskName":             autorest.Encode("path", taskName),
		"taskUpdateActionType": autorest.Encode("path", taskUpdateActionType),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}/{taskUpdateActionType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateResourceGroupLevelTaskStateSender sends the UpdateResourceGroupLevelTaskState request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) UpdateResourceGroupLevelTaskStateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResourceGroupLevelTaskStateResponder handles the response to the UpdateResourceGroupLevelTaskState request. The method always
// closes the http.Response Body.
func (client TasksClient) UpdateResourceGroupLevelTaskStateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdateSubscriptionLevelTaskState recommended tasks that will help improve the security of the subscription
// proactively
// Parameters:
// taskName - name of the task object, will be a GUID
// taskUpdateActionType - type of the action to do on the task
func (client TasksClient) UpdateSubscriptionLevelTaskState(ctx context.Context, taskName string, taskUpdateActionType string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TasksClient.UpdateSubscriptionLevelTaskState")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.TasksClient", "UpdateSubscriptionLevelTaskState", err.Error())
	}

	req, err := client.UpdateSubscriptionLevelTaskStatePreparer(ctx, taskName, taskUpdateActionType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateSubscriptionLevelTaskState", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSubscriptionLevelTaskStateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateSubscriptionLevelTaskState", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateSubscriptionLevelTaskStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.TasksClient", "UpdateSubscriptionLevelTaskState", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateSubscriptionLevelTaskStatePreparer prepares the UpdateSubscriptionLevelTaskState request.
func (client TasksClient) UpdateSubscriptionLevelTaskStatePreparer(ctx context.Context, taskName string, taskUpdateActionType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":          autorest.Encode("path", client.AscLocation),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"taskName":             autorest.Encode("path", taskName),
		"taskUpdateActionType": autorest.Encode("path", taskUpdateActionType),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}/{taskUpdateActionType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSubscriptionLevelTaskStateSender sends the UpdateSubscriptionLevelTaskState request. The method will close the
// http.Response Body if it receives an error.
func (client TasksClient) UpdateSubscriptionLevelTaskStateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateSubscriptionLevelTaskStateResponder handles the response to the UpdateSubscriptionLevelTaskState request. The method always
// closes the http.Response Body.
func (client TasksClient) UpdateSubscriptionLevelTaskStateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
