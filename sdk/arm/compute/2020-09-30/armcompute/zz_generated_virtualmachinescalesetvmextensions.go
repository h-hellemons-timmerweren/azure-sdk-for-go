// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetVMExtensionsClient contains the methods for the VirtualMachineScaleSetVMExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMExtensionsClient() instead.
type VirtualMachineScaleSetVMExtensionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetVMExtensionsClient creates a new instance of VirtualMachineScaleSetVMExtensionsClient with the specified values.
func NewVirtualMachineScaleSetVMExtensionsClient(con *armcore.Connection, subscriptionID string) VirtualMachineScaleSetVMExtensionsClient {
	return VirtualMachineScaleSetVMExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualMachineScaleSetVMExtensionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (VirtualMachineScaleSetVMExtensionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionPollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionPollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualMachineScaleSetVMExtensionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineScaleSetVMExtensionPoller.ResumeToken().
func (client VirtualMachineScaleSetVMExtensionsClient) ResumeCreateOrUpdate(token string) (VirtualMachineScaleSetVMExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineScaleSetVMExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client VirtualMachineScaleSetVMExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtension, options *VirtualMachineScaleSetVMExtensionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VirtualMachineScaleSetVMExtensionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionResponse, error) {
	result := VirtualMachineScaleSetVMExtensionResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtension)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VirtualMachineScaleSetVMExtensionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - The operation to delete the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client VirtualMachineScaleSetVMExtensionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - The operation to delete the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) delete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client VirtualMachineScaleSetVMExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client VirtualMachineScaleSetVMExtensionsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - The operation to get the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (VirtualMachineScaleSetVMExtensionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineScaleSetVMExtensionResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualMachineScaleSetVMExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, options *VirtualMachineScaleSetVMExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VirtualMachineScaleSetVMExtensionsClient) getHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionResponse, error) {
	result := VirtualMachineScaleSetVMExtensionResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtension)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualMachineScaleSetVMExtensionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - The operation to get all extensions of an instance in Virtual Machine Scaleset.
func (client VirtualMachineScaleSetVMExtensionsClient) List(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, options *VirtualMachineScaleSetVMExtensionsListOptions) (VirtualMachineScaleSetVMExtensionsListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsListResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineScaleSetVMExtensionsListResultResponse{}, client.listHandleError(resp)
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionsListResultResponse{}, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client VirtualMachineScaleSetVMExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, options *VirtualMachineScaleSetVMExtensionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client VirtualMachineScaleSetVMExtensionsClient) listHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionsListResultResponse, error) {
	result := VirtualMachineScaleSetVMExtensionsListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtensionsListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client VirtualMachineScaleSetVMExtensionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - The operation to update the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (VirtualMachineScaleSetVMExtensionPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionPollerResponse{}, err
	}
	result := VirtualMachineScaleSetVMExtensionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMExtensionsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return VirtualMachineScaleSetVMExtensionPollerResponse{}, err
	}
	poller := &virtualMachineScaleSetVMExtensionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineScaleSetVMExtensionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new VirtualMachineScaleSetVMExtensionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineScaleSetVMExtensionPoller.ResumeToken().
func (client VirtualMachineScaleSetVMExtensionsClient) ResumeUpdate(token string) (VirtualMachineScaleSetVMExtensionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMExtensionsClient.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineScaleSetVMExtensionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - The operation to update the VMSS VM extension.
func (client VirtualMachineScaleSetVMExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client VirtualMachineScaleSetVMExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string, extensionParameters VirtualMachineScaleSetVMExtensionUpdate, options *VirtualMachineScaleSetVMExtensionsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(extensionParameters)
}

// updateHandleResponse handles the Update response.
func (client VirtualMachineScaleSetVMExtensionsClient) updateHandleResponse(resp *azcore.Response) (VirtualMachineScaleSetVMExtensionResponse, error) {
	result := VirtualMachineScaleSetVMExtensionResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineScaleSetVMExtension)
	return result, err
}

// updateHandleError handles the Update error response.
func (client VirtualMachineScaleSetVMExtensionsClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
