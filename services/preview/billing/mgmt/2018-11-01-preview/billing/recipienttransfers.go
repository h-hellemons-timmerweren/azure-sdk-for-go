package billing

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

// RecipientTransfersClient is the billing client provides access to billing resources for Azure subscriptions.
type RecipientTransfersClient struct {
	BaseClient
}

// NewRecipientTransfersClient creates an instance of the RecipientTransfersClient client.
func NewRecipientTransfersClient(subscriptionID string) RecipientTransfersClient {
	return NewRecipientTransfersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecipientTransfersClientWithBaseURI creates an instance of the RecipientTransfersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRecipientTransfersClientWithBaseURI(baseURI string, subscriptionID string) RecipientTransfersClient {
	return RecipientTransfersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Accept sends the accept request.
// Parameters:
// transferName - transfer Name.
// body - accept transfer parameters.
func (client RecipientTransfersClient) Accept(ctx context.Context, transferName string, body AcceptTransferRequest) (result RecipientTransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecipientTransfersClient.Accept")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AcceptPreparer(ctx, transferName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Accept", nil, "Failure preparing request")
		return
	}

	resp, err := client.AcceptSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Accept", resp, "Failure sending request")
		return
	}

	result, err = client.AcceptResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Accept", resp, "Failure responding to request")
		return
	}

	return
}

// AcceptPreparer prepares the Accept request.
func (client RecipientTransfersClient) AcceptPreparer(ctx context.Context, transferName string, body AcceptTransferRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"transferName": autorest.Encode("path", transferName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/transfers/{transferName}/acceptTransfer", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AcceptSender sends the Accept request. The method will close the
// http.Response Body if it receives an error.
func (client RecipientTransfersClient) AcceptSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AcceptResponder handles the response to the Accept request. The method always
// closes the http.Response Body.
func (client RecipientTransfersClient) AcceptResponder(resp *http.Response) (result RecipientTransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Decline sends the decline request.
// Parameters:
// transferName - transfer Name.
func (client RecipientTransfersClient) Decline(ctx context.Context, transferName string) (result RecipientTransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecipientTransfersClient.Decline")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeclinePreparer(ctx, transferName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Decline", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeclineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Decline", resp, "Failure sending request")
		return
	}

	result, err = client.DeclineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Decline", resp, "Failure responding to request")
		return
	}

	return
}

// DeclinePreparer prepares the Decline request.
func (client RecipientTransfersClient) DeclinePreparer(ctx context.Context, transferName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"transferName": autorest.Encode("path", transferName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/transfers/{transferName}/declineTransfer", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeclineSender sends the Decline request. The method will close the
// http.Response Body if it receives an error.
func (client RecipientTransfersClient) DeclineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeclineResponder handles the response to the Decline request. The method always
// closes the http.Response Body.
func (client RecipientTransfersClient) DeclineResponder(resp *http.Response) (result RecipientTransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// transferName - transfer Name.
func (client RecipientTransfersClient) Get(ctx context.Context, transferName string) (result RecipientTransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecipientTransfersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, transferName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecipientTransfersClient) GetPreparer(ctx context.Context, transferName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"transferName": autorest.Encode("path", transferName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/transfers/{transferName}/", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecipientTransfersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecipientTransfersClient) GetResponder(resp *http.Response) (result RecipientTransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client RecipientTransfersClient) List(ctx context.Context) (result RecipientTransferDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecipientTransfersClient.List")
		defer func() {
			sc := -1
			if result.rtdlr.Response.Response != nil {
				sc = result.rtdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rtdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "List", resp, "Failure sending request")
		return
	}

	result.rtdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rtdlr.hasNextLink() && result.rtdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client RecipientTransfersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Billing/transfers"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecipientTransfersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecipientTransfersClient) ListResponder(resp *http.Response) (result RecipientTransferDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecipientTransfersClient) listNextResults(ctx context.Context, lastResults RecipientTransferDetailsListResult) (result RecipientTransferDetailsListResult, err error) {
	req, err := lastResults.recipientTransferDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RecipientTransfersClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecipientTransfersClient) ListComplete(ctx context.Context) (result RecipientTransferDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecipientTransfersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
