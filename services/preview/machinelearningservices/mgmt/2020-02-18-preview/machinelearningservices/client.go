// Package machinelearningservices implements the Azure ARM Machinelearningservices service API version
// 2020-02-18-preview.
//
// These APIs allow end users to operate on Azure Machine Learning Workspace resources.
package machinelearningservices

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

const (
	// DefaultBaseURI is the default URI used for the service Machinelearningservices
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Machinelearningservices.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// ListSkus lists all skus with associated features
func (client BaseClient) ListSkus(ctx context.Context) (result SkuListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListSkus")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listSkusNextResults
	req, err := client.ListSkusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "ListSkus", resp, "Failure responding to request")
		return
	}
	if result.slr.hasNextLink() && result.slr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client BaseClient) ListSkusPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client BaseClient) ListSkusResponder(resp *http.Response) (result SkuListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSkusNextResults retrieves the next set of results, if any.
func (client BaseClient) listSkusNextResults(ctx context.Context, lastResults SkuListResult) (result SkuListResult, err error) {
	req, err := lastResults.skuListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "listSkusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "listSkusNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.BaseClient", "listSkusNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListSkusComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) ListSkusComplete(ctx context.Context) (result SkuListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ListSkus")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSkus(ctx)
	return
}
