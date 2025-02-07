package keyvault

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

// HSMSecurityDomainClient is the the key vault client performs cryptographic key operations and vault operations
// against the Key Vault service.
type HSMSecurityDomainClient struct {
	BaseClient
}

// NewHSMSecurityDomainClient creates an instance of the HSMSecurityDomainClient client.
func NewHSMSecurityDomainClient() HSMSecurityDomainClient {
	return HSMSecurityDomainClient{New()}
}

// Download retrieves Security domain from HSM enclave
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// certificateInfoObject - security domain download operation requires customer to provide N certificates
// (minimum 3 and maximum 10) containing public key in JWK format.
func (client HSMSecurityDomainClient) Download(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject) (result SecurityDomainObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HSMSecurityDomainClient.Download")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateInfoObject,
			Constraints: []validation.Constraint{{Target: "certificateInfoObject.Certificates", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "certificateInfoObject.Certificates", Name: validation.MaxItems, Rule: 10, Chain: nil},
					{Target: "certificateInfoObject.Certificates", Name: validation.MinItems, Rule: 3, Chain: nil},
					{Target: "certificateInfoObject.Certificates", Name: validation.UniqueItems, Rule: true, Chain: nil},
				}},
				{Target: "certificateInfoObject.Required", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "certificateInfoObject.Required", Name: validation.InclusiveMaximum, Rule: int64(10), Chain: nil},
						{Target: "certificateInfoObject.Required", Name: validation.InclusiveMinimum, Rule: int64(2), Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("keyvault.HSMSecurityDomainClient", "Download", err.Error())
	}

	req, err := client.DownloadPreparer(ctx, vaultBaseURL, certificateInfoObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "Download", nil, "Failure preparing request")
		return
	}

	resp, err := client.DownloadSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "Download", resp, "Failure sending request")
		return
	}

	result, err = client.DownloadResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "Download", resp, "Failure responding to request")
		return
	}

	return
}

// DownloadPreparer prepares the Download request.
func (client HSMSecurityDomainClient) DownloadPreparer(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	const APIVersion = "7.2-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/securitydomain/download"),
		autorest.WithJSON(certificateInfoObject),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DownloadSender sends the Download request. The method will close the
// http.Response Body if it receives an error.
func (client HSMSecurityDomainClient) DownloadSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DownloadResponder handles the response to the Download request. The method always
// closes the http.Response Body.
func (client HSMSecurityDomainClient) DownloadResponder(resp *http.Response) (result SecurityDomainObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TransferKeyMethod retrieve security domain transfer key
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
func (client HSMSecurityDomainClient) TransferKeyMethod(ctx context.Context, vaultBaseURL string) (result TransferKey, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HSMSecurityDomainClient.TransferKeyMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TransferKeyMethodPreparer(ctx, vaultBaseURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "TransferKeyMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.TransferKeyMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "TransferKeyMethod", resp, "Failure sending request")
		return
	}

	result, err = client.TransferKeyMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "TransferKeyMethod", resp, "Failure responding to request")
		return
	}

	return
}

// TransferKeyMethodPreparer prepares the TransferKeyMethod request.
func (client HSMSecurityDomainClient) TransferKeyMethodPreparer(ctx context.Context, vaultBaseURL string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	const APIVersion = "7.2-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/securitydomain/transferkey"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TransferKeyMethodSender sends the TransferKeyMethod request. The method will close the
// http.Response Body if it receives an error.
func (client HSMSecurityDomainClient) TransferKeyMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TransferKeyMethodResponder handles the response to the TransferKeyMethod request. The method always
// closes the http.Response Body.
func (client HSMSecurityDomainClient) TransferKeyMethodResponder(resp *http.Response) (result TransferKey, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Upload request Security domain upload operation
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// securityDomain - security domain
func (client HSMSecurityDomainClient) Upload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainUploadObject) (result HSMSecurityDomainUploadFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HSMSecurityDomainClient.Upload")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: securityDomain,
			Constraints: []validation.Constraint{{Target: "securityDomain.Value", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "securityDomain.Value.EncData", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "securityDomain.Value.EncData.Data", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "securityDomain.Value.EncData.Data", Name: validation.UniqueItems, Rule: true, Chain: nil}}},
						{Target: "securityDomain.Value.EncData.Kdf", Name: validation.Null, Rule: true, Chain: nil},
					}},
					{Target: "securityDomain.Value.WrappedKey", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "securityDomain.Value.WrappedKey.EncKey", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "securityDomain.Value.WrappedKey.X5t256", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("keyvault.HSMSecurityDomainClient", "Upload", err.Error())
	}

	req, err := client.UploadPreparer(ctx, vaultBaseURL, securityDomain)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "Upload", nil, "Failure preparing request")
		return
	}

	result, err = client.UploadSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "Upload", result.Response(), "Failure sending request")
		return
	}

	return
}

// UploadPreparer prepares the Upload request.
func (client HSMSecurityDomainClient) UploadPreparer(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainUploadObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/securitydomain/upload"),
		autorest.WithJSON(securityDomain))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadSender sends the Upload request. The method will close the
// http.Response Body if it receives an error.
func (client HSMSecurityDomainClient) UploadSender(req *http.Request) (future HSMSecurityDomainUploadFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UploadResponder handles the response to the Upload request. The method always
// closes the http.Response Body.
func (client HSMSecurityDomainClient) UploadResponder(resp *http.Response) (result SecurityDomainOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UploadPending get Security domain upload operation status
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
func (client HSMSecurityDomainClient) UploadPending(ctx context.Context, vaultBaseURL string) (result SecurityDomainOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HSMSecurityDomainClient.UploadPending")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UploadPendingPreparer(ctx, vaultBaseURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "UploadPending", nil, "Failure preparing request")
		return
	}

	resp, err := client.UploadPendingSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "UploadPending", resp, "Failure sending request")
		return
	}

	result, err = client.UploadPendingResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.HSMSecurityDomainClient", "UploadPending", resp, "Failure responding to request")
		return
	}

	return
}

// UploadPendingPreparer prepares the UploadPending request.
func (client HSMSecurityDomainClient) UploadPendingPreparer(ctx context.Context, vaultBaseURL string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPath("/securitydomain/upload/pending"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UploadPendingSender sends the UploadPending request. The method will close the
// http.Response Body if it receives an error.
func (client HSMSecurityDomainClient) UploadPendingSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UploadPendingResponder handles the response to the UploadPending request. The method always
// closes the http.Response Body.
func (client HSMSecurityDomainClient) UploadPendingResponder(resp *http.Response) (result SecurityDomainOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
