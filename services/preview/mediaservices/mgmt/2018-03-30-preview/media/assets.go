package media

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

// AssetsClient is the client for the Assets methods of the Media service.
type AssetsClient struct {
	BaseClient
}

// NewAssetsClient creates an instance of the AssetsClient client.
func NewAssetsClient(subscriptionID string) AssetsClient {
	return NewAssetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssetsClientWithBaseURI creates an instance of the AssetsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAssetsClientWithBaseURI(baseURI string, subscriptionID string) AssetsClient {
	return AssetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an Asset in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
// parameters - the request parameters
func (client AssetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset) (result Asset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, assetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AssetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AssetsClient) CreateOrUpdateResponder(resp *http.Response) (result Asset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an Asset in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
func (client AssetsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, assetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of an Asset in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
func (client AssetsClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result Asset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, assetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssetsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssetsClient) GetResponder(resp *http.Response) (result Asset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEncryptionKey gets the Asset storage encryption keys used to decrypt content created by version 2 of the Media
// Services API
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
func (client AssetsClient) GetEncryptionKey(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result AssetStorageEncryptionKey, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.GetEncryptionKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEncryptionKeyPreparer(ctx, resourceGroupName, accountName, assetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "GetEncryptionKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEncryptionKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "GetEncryptionKey", resp, "Failure sending request")
		return
	}

	result, err = client.GetEncryptionKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "GetEncryptionKey", resp, "Failure responding to request")
		return
	}

	return
}

// GetEncryptionKeyPreparer prepares the GetEncryptionKey request.
func (client AssetsClient) GetEncryptionKeyPreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/getEncryptionKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEncryptionKeySender sends the GetEncryptionKey request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) GetEncryptionKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEncryptionKeyResponder handles the response to the GetEncryptionKey request. The method always
// closes the http.Response Body.
func (client AssetsClient) GetEncryptionKeyResponder(resp *http.Response) (result AssetStorageEncryptionKey, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list Assets in the Media Services account with optional filtering and ordering
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// filter - restricts the set of items returned.
// top - specifies a non-negative integer n that limits the number of items returned from a collection. The
// service returns the number of available items up to but not greater than the specified value n.
// orderby - specifies the key by which the result collection should be ordered.
func (client AssetsClient) List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result AssetCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.List")
		defer func() {
			sc := -1
			if result.ac.Response.Response != nil {
				sc = result.ac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "List", resp, "Failure sending request")
		return
	}

	result.ac, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ac.hasNextLink() && result.ac.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client AssetsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssetsClient) ListResponder(resp *http.Response) (result AssetCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AssetsClient) listNextResults(ctx context.Context, lastResults AssetCollection) (result AssetCollection, err error) {
	req, err := lastResults.assetCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.AssetsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.AssetsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssetsClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result AssetCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, filter, top, orderby)
	return
}

// ListContainerSas lists storage container URLs with shared access signatures (SAS) for uploading and downloading
// Asset content. The signatures are derived from the storage account keys.
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
// parameters - the request parameters
func (client AssetsClient) ListContainerSas(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput) (result AssetContainerSas, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.ListContainerSas")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListContainerSasPreparer(ctx, resourceGroupName, accountName, assetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "ListContainerSas", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListContainerSasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "ListContainerSas", resp, "Failure sending request")
		return
	}

	result, err = client.ListContainerSasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "ListContainerSas", resp, "Failure responding to request")
		return
	}

	return
}

// ListContainerSasPreparer prepares the ListContainerSas request.
func (client AssetsClient) ListContainerSasPreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/listContainerSas", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListContainerSasSender sends the ListContainerSas request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) ListContainerSasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListContainerSasResponder handles the response to the ListContainerSas request. The method always
// closes the http.Response Body.
func (client AssetsClient) ListContainerSasResponder(resp *http.Response) (result AssetContainerSas, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing Asset in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// assetName - the Asset name.
// parameters - the request parameters
func (client AssetsClient) Update(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset) (result Asset, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssetsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, assetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.AssetsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AssetsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"assetName":         autorest.Encode("path", assetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AssetsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AssetsClient) UpdateResponder(resp *http.Response) (result Asset, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
