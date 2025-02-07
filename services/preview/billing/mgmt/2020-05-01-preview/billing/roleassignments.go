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

// RoleAssignmentsClient is the billing client provides access to billing resources for Azure subscriptions.
type RoleAssignmentsClient struct {
	BaseClient
}

// NewRoleAssignmentsClient creates an instance of the RoleAssignmentsClient client.
func NewRoleAssignmentsClient(subscriptionID string) RoleAssignmentsClient {
	return NewRoleAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleAssignmentsClientWithBaseURI creates an instance of the RoleAssignmentsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRoleAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentsClient {
	return RoleAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteByBillingAccount deletes a role assignment for the caller on a billing account. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) DeleteByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.DeleteByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByBillingAccountPreparer(ctx, billingAccountName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingAccount", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteByBillingAccountPreparer prepares the DeleteByBillingAccount request.
func (client RoleAssignmentsClient) DeleteByBillingAccountPreparer(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByBillingAccountSender sends the DeleteByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByBillingAccountResponder handles the response to the DeleteByBillingAccount request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteByBillingAccountResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByBillingProfile deletes a role assignment for the caller on a billing profile. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) DeleteByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.DeleteByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByBillingProfile", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteByBillingProfilePreparer prepares the DeleteByBillingProfile request.
func (client RoleAssignmentsClient) DeleteByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByBillingProfileSender sends the DeleteByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByBillingProfileResponder handles the response to the DeleteByBillingProfile request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteByBillingProfileResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByInvoiceSection deletes a role assignment for the caller on an invoice section. The operation is supported
// for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) DeleteByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.DeleteByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByInvoiceSection", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "DeleteByInvoiceSection", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteByInvoiceSectionPreparer prepares the DeleteByInvoiceSection request.
func (client RoleAssignmentsClient) DeleteByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
		"invoiceSectionName":        autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByInvoiceSectionSender sends the DeleteByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) DeleteByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByInvoiceSectionResponder handles the response to the DeleteByInvoiceSection request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) DeleteByInvoiceSectionResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByBillingAccount gets a role assignment for the caller on a billing account. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.GetByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingAccountPreparer(ctx, billingAccountName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingAccount", resp, "Failure responding to request")
		return
	}

	return
}

// GetByBillingAccountPreparer prepares the GetByBillingAccount request.
func (client RoleAssignmentsClient) GetByBillingAccountPreparer(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingAccountSender sends the GetByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingAccountResponder handles the response to the GetByBillingAccount request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetByBillingAccountResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByBillingProfile gets a role assignment for the caller on a billing profile. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.GetByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByBillingProfile", resp, "Failure responding to request")
		return
	}

	return
}

// GetByBillingProfilePreparer prepares the GetByBillingProfile request.
func (client RoleAssignmentsClient) GetByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingProfileSender sends the GetByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingProfileResponder handles the response to the GetByBillingProfile request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetByBillingProfileResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByInvoiceSection gets a role assignment for the caller on an invoice section. The operation is supported for
// billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// billingRoleAssignmentName - the ID that uniquely identifies a role assignment.
func (client RoleAssignmentsClient) GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (result RoleAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.GetByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByInvoiceSection", resp, "Failure sending request")
		return
	}

	result, err = client.GetByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "GetByInvoiceSection", resp, "Failure responding to request")
		return
	}

	return
}

// GetByInvoiceSectionPreparer prepares the GetByInvoiceSection request.
func (client RoleAssignmentsClient) GetByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleAssignmentName": autorest.Encode("path", billingRoleAssignmentName),
		"invoiceSectionName":        autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByInvoiceSectionSender sends the GetByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) GetByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByInvoiceSectionResponder handles the response to the GetByInvoiceSection request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) GetByInvoiceSectionResponder(resp *http.Response) (result RoleAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the role assignments for the caller on a billing account. The operation is supported for
// billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client RoleAssignmentsClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result RoleAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.ralr.Response.Response != nil {
				sc = result.ralr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.ralr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.ralr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingAccount", resp, "Failure responding to request")
		return
	}
	if result.ralr.hasNextLink() && result.ralr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client RoleAssignmentsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListByBillingAccountResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) listByBillingAccountNextResults(ctx context.Context, lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.roleAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleAssignmentsClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result RoleAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName)
	return
}

// ListByBillingProfile lists the role assignments for the caller on a billing profile. The operation is supported for
// billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client RoleAssignmentsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result RoleAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.ralr.Response.Response != nil {
				sc = result.ralr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.ralr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.ralr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByBillingProfile", resp, "Failure responding to request")
		return
	}
	if result.ralr.hasNextLink() && result.ralr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client RoleAssignmentsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListByBillingProfileResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) listByBillingProfileNextResults(ctx context.Context, lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.roleAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleAssignmentsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result RoleAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName)
	return
}

// ListByInvoiceSection lists the role assignments for the caller on an invoice section. The operation is supported for
// billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client RoleAssignmentsClient) ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result RoleAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.ralr.Response.Response != nil {
				sc = result.ralr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionNextResults
	req, err := client.ListByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.ralr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByInvoiceSection", resp, "Failure sending request")
		return
	}

	result.ralr, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "ListByInvoiceSection", resp, "Failure responding to request")
		return
	}
	if result.ralr.hasNextLink() && result.ralr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByInvoiceSectionPreparer prepares the ListByInvoiceSection request.
func (client RoleAssignmentsClient) ListByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionSender sends the ListByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentsClient) ListByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionResponder handles the response to the ListByInvoiceSection request. The method always
// closes the http.Response Body.
func (client RoleAssignmentsClient) ListByInvoiceSectionResponder(resp *http.Response) (result RoleAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionNextResults retrieves the next set of results, if any.
func (client RoleAssignmentsClient) listByInvoiceSectionNextResults(ctx context.Context, lastResults RoleAssignmentListResult) (result RoleAssignmentListResult, err error) {
	req, err := lastResults.roleAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByInvoiceSectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByInvoiceSectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleAssignmentsClient", "listByInvoiceSectionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByInvoiceSectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleAssignmentsClient) ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result RoleAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSection(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	return
}
