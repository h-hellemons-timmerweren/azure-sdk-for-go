package databox

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

// ServiceClient is the client for the Service methods of the Databox service.
type ServiceClient struct {
	BaseClient
}

// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient(subscriptionID string) ServiceClient {
	return NewServiceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceClientWithBaseURI creates an instance of the ServiceClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return ServiceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListAvailableSkusByResourceGroup this method provides the list of available skus for the given subscription,
// resource group and location.
// Parameters:
// resourceGroupName - the Resource Group Name
// location - the location of the resource
// availableSkuRequest - filters for showing the available skus.
func (client ServiceClient) ListAvailableSkusByResourceGroup(ctx context.Context, resourceGroupName string, location string, availableSkuRequest AvailableSkuRequest) (result AvailableSkusResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ListAvailableSkusByResourceGroup")
		defer func() {
			sc := -1
			if result.asr.Response.Response != nil {
				sc = result.asr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: availableSkuRequest,
			Constraints: []validation.Constraint{{Target: "availableSkuRequest.Country", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "availableSkuRequest.Location", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "ListAvailableSkusByResourceGroup", err.Error())
	}

	result.fn = client.listAvailableSkusByResourceGroupNextResults
	req, err := client.ListAvailableSkusByResourceGroupPreparer(ctx, resourceGroupName, location, availableSkuRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ListAvailableSkusByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableSkusByResourceGroupSender(req)
	if err != nil {
		result.asr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ListAvailableSkusByResourceGroup", resp, "Failure sending request")
		return
	}

	result.asr, err = client.ListAvailableSkusByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ListAvailableSkusByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.asr.hasNextLink() && result.asr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListAvailableSkusByResourceGroupPreparer prepares the ListAvailableSkusByResourceGroup request.
func (client ServiceClient) ListAvailableSkusByResourceGroupPreparer(ctx context.Context, resourceGroupName string, location string, availableSkuRequest AvailableSkuRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/locations/{location}/availableSkus", pathParameters),
		autorest.WithJSON(availableSkuRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAvailableSkusByResourceGroupSender sends the ListAvailableSkusByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ListAvailableSkusByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAvailableSkusByResourceGroupResponder handles the response to the ListAvailableSkusByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceClient) ListAvailableSkusByResourceGroupResponder(resp *http.Response) (result AvailableSkusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAvailableSkusByResourceGroupNextResults retrieves the next set of results, if any.
func (client ServiceClient) listAvailableSkusByResourceGroupNextResults(ctx context.Context, lastResults AvailableSkusResult) (result AvailableSkusResult, err error) {
	req, err := lastResults.availableSkusResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databox.ServiceClient", "listAvailableSkusByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAvailableSkusByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databox.ServiceClient", "listAvailableSkusByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAvailableSkusByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "listAvailableSkusByResourceGroupNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListAvailableSkusByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceClient) ListAvailableSkusByResourceGroupComplete(ctx context.Context, resourceGroupName string, location string, availableSkuRequest AvailableSkuRequest) (result AvailableSkusResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ListAvailableSkusByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAvailableSkusByResourceGroup(ctx, resourceGroupName, location, availableSkuRequest)
	return
}

// RegionConfiguration this API provides configuration details specific to given region/location at Subscription level.
// Parameters:
// location - the location of the resource
// regionConfigurationRequest - request body to get the configuration for the region.
func (client ServiceClient) RegionConfiguration(ctx context.Context, location string, regionConfigurationRequest RegionConfigurationRequest) (result RegionConfigurationResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.RegionConfiguration")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: regionConfigurationRequest,
			Constraints: []validation.Constraint{{Target: "regionConfigurationRequest.ScheduleAvailabilityRequest", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "regionConfigurationRequest.ScheduleAvailabilityRequest.StorageLocation", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "RegionConfiguration", err.Error())
	}

	req, err := client.RegionConfigurationPreparer(ctx, location, regionConfigurationRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfiguration", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegionConfigurationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfiguration", resp, "Failure sending request")
		return
	}

	result, err = client.RegionConfigurationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfiguration", resp, "Failure responding to request")
		return
	}

	return
}

// RegionConfigurationPreparer prepares the RegionConfiguration request.
func (client ServiceClient) RegionConfigurationPreparer(ctx context.Context, location string, regionConfigurationRequest RegionConfigurationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/locations/{location}/regionConfiguration", pathParameters),
		autorest.WithJSON(regionConfigurationRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegionConfigurationSender sends the RegionConfiguration request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) RegionConfigurationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegionConfigurationResponder handles the response to the RegionConfiguration request. The method always
// closes the http.Response Body.
func (client ServiceClient) RegionConfigurationResponder(resp *http.Response) (result RegionConfigurationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegionConfigurationByResourceGroup this API provides configuration details specific to given region/location at
// Resource group level.
// Parameters:
// resourceGroupName - the Resource Group Name
// location - the location of the resource
// regionConfigurationRequest - request body to get the configuration for the region at resource group level.
func (client ServiceClient) RegionConfigurationByResourceGroup(ctx context.Context, resourceGroupName string, location string, regionConfigurationRequest RegionConfigurationRequest) (result RegionConfigurationResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.RegionConfigurationByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: regionConfigurationRequest,
			Constraints: []validation.Constraint{{Target: "regionConfigurationRequest.ScheduleAvailabilityRequest", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "regionConfigurationRequest.ScheduleAvailabilityRequest.StorageLocation", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "RegionConfigurationByResourceGroup", err.Error())
	}

	req, err := client.RegionConfigurationByResourceGroupPreparer(ctx, resourceGroupName, location, regionConfigurationRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfigurationByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegionConfigurationByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfigurationByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.RegionConfigurationByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "RegionConfigurationByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// RegionConfigurationByResourceGroupPreparer prepares the RegionConfigurationByResourceGroup request.
func (client ServiceClient) RegionConfigurationByResourceGroupPreparer(ctx context.Context, resourceGroupName string, location string, regionConfigurationRequest RegionConfigurationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/locations/{location}/regionConfiguration", pathParameters),
		autorest.WithJSON(regionConfigurationRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegionConfigurationByResourceGroupSender sends the RegionConfigurationByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) RegionConfigurationByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegionConfigurationByResourceGroupResponder handles the response to the RegionConfigurationByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceClient) RegionConfigurationByResourceGroupResponder(resp *http.Response) (result RegionConfigurationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateAddressMethod [DEPRECATED NOTICE: This operation will soon be removed]. This method validates the customer
// shipping address and provide alternate addresses if any.
// Parameters:
// location - the location of the resource
// validateAddress - shipping address of the customer.
func (client ServiceClient) ValidateAddressMethod(ctx context.Context, location string, validateAddress ValidateAddress) (result AddressValidationOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ValidateAddressMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: validateAddress,
			Constraints: []validation.Constraint{{Target: "validateAddress.ShippingAddress", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "validateAddress.ShippingAddress.StreetAddress1", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "validateAddress.ShippingAddress.Country", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "ValidateAddressMethod", err.Error())
	}

	req, err := client.ValidateAddressMethodPreparer(ctx, location, validateAddress)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateAddressMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateAddressMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateAddressMethod", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateAddressMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateAddressMethod", resp, "Failure responding to request")
		return
	}

	return
}

// ValidateAddressMethodPreparer prepares the ValidateAddressMethod request.
func (client ServiceClient) ValidateAddressMethodPreparer(ctx context.Context, location string, validateAddress ValidateAddress) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/locations/{location}/validateAddress", pathParameters),
		autorest.WithJSON(validateAddress),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateAddressMethodSender sends the ValidateAddressMethod request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ValidateAddressMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateAddressMethodResponder handles the response to the ValidateAddressMethod request. The method always
// closes the http.Response Body.
func (client ServiceClient) ValidateAddressMethodResponder(resp *http.Response) (result AddressValidationOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateInputs this method does all necessary pre-job creation validation under subscription.
// Parameters:
// location - the location of the resource
// validationRequest - inputs of the customer.
func (client ServiceClient) ValidateInputs(ctx context.Context, location string, validationRequest BasicValidationRequest) (result ValidationResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ValidateInputs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: validationRequest,
			Constraints: []validation.Constraint{{Target: "validationRequest.IndividualRequestDetails", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "ValidateInputs", err.Error())
	}

	req, err := client.ValidateInputsPreparer(ctx, location, validationRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateInputsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputs", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateInputsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputs", resp, "Failure responding to request")
		return
	}

	return
}

// ValidateInputsPreparer prepares the ValidateInputs request.
func (client ServiceClient) ValidateInputsPreparer(ctx context.Context, location string, validationRequest BasicValidationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/locations/{location}/validateInputs", pathParameters),
		autorest.WithJSON(validationRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateInputsSender sends the ValidateInputs request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ValidateInputsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateInputsResponder handles the response to the ValidateInputs request. The method always
// closes the http.Response Body.
func (client ServiceClient) ValidateInputsResponder(resp *http.Response) (result ValidationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateInputsByResourceGroup this method does all necessary pre-job creation validation under resource group.
// Parameters:
// resourceGroupName - the Resource Group Name
// location - the location of the resource
// validationRequest - inputs of the customer.
func (client ServiceClient) ValidateInputsByResourceGroup(ctx context.Context, resourceGroupName string, location string, validationRequest BasicValidationRequest) (result ValidationResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceClient.ValidateInputsByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: validationRequest,
			Constraints: []validation.Constraint{{Target: "validationRequest.IndividualRequestDetails", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("databox.ServiceClient", "ValidateInputsByResourceGroup", err.Error())
	}

	req, err := client.ValidateInputsByResourceGroupPreparer(ctx, resourceGroupName, location, validationRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputsByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateInputsByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputsByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateInputsByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databox.ServiceClient", "ValidateInputsByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ValidateInputsByResourceGroupPreparer prepares the ValidateInputsByResourceGroup request.
func (client ServiceClient) ValidateInputsByResourceGroupPreparer(ctx context.Context, resourceGroupName string, location string, validationRequest BasicValidationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/locations/{location}/validateInputs", pathParameters),
		autorest.WithJSON(validationRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateInputsByResourceGroupSender sends the ValidateInputsByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceClient) ValidateInputsByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateInputsByResourceGroupResponder handles the response to the ValidateInputsByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceClient) ValidateInputsByResourceGroupResponder(resp *http.Response) (result ValidationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
