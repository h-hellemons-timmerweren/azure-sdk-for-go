package datamigration

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

// ProjectsClient is the data Migration Client
type ProjectsClient struct {
	BaseClient
}

// NewProjectsClient creates an instance of the ProjectsClient client.
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return NewProjectsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProjectsClientWithBaseURI creates an instance of the ProjectsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return ProjectsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the project resource is a nested resource representing a stored migration project. The PUT method
// creates a new project or updates an existing one.
// Parameters:
// parameters - information about the project
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
func (client ProjectsClient) CreateOrUpdate(ctx context.Context, parameters Project, groupName string, serviceName string, projectName string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, groupName, serviceName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProjectsClient) CreateOrUpdatePreparer(ctx context.Context, parameters Project, groupName string, serviceName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProjectsClient) CreateOrUpdateResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the project resource is a nested resource representing a stored migration project. The DELETE method deletes
// a project.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// deleteRunningTasks - delete the resource even if it contains running tasks
func (client ProjectsClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, deleteRunningTasks *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, groupName, serviceName, projectName, deleteRunningTasks)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProjectsClient) DeletePreparer(ctx context.Context, groupName string, serviceName string, projectName string, deleteRunningTasks *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if deleteRunningTasks != nil {
		queryParameters["deleteRunningTasks"] = autorest.Encode("query", *deleteRunningTasks)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProjectsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the project resource is a nested resource representing a stored migration project. The GET method retrieves
// information about a project.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
func (client ProjectsClient) Get(ctx context.Context, groupName string, serviceName string, projectName string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, groupName, serviceName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProjectsClient) GetPreparer(ctx context.Context, groupName string, serviceName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProjectsClient) GetResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the project resource is a nested resource representing a stored migration project. This method returns a list
// of projects owned by a service resource.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
func (client ProjectsClient) List(ctx context.Context, groupName string, serviceName string) (result ProjectListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.List")
		defer func() {
			sc := -1
			if result.pl.Response.Response != nil {
				sc = result.pl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, groupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "List", resp, "Failure sending request")
		return
	}

	result.pl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pl.hasNextLink() && result.pl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ProjectsClient) ListPreparer(ctx context.Context, groupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProjectsClient) ListResponder(resp *http.Response) (result ProjectList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProjectsClient) listNextResults(ctx context.Context, lastResults ProjectList) (result ProjectList, err error) {
	req, err := lastResults.projectListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProjectsClient) ListComplete(ctx context.Context, groupName string, serviceName string) (result ProjectListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, groupName, serviceName)
	return
}

// Update the project resource is a nested resource representing a stored migration project. The PATCH method updates
// an existing project.
// Parameters:
// parameters - information about the project
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
func (client ProjectsClient) Update(ctx context.Context, parameters Project, groupName string, serviceName string, projectName string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, parameters, groupName, serviceName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.ProjectsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProjectsClient) UpdatePreparer(ctx context.Context, parameters Project, groupName string, serviceName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProjectsClient) UpdateResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
