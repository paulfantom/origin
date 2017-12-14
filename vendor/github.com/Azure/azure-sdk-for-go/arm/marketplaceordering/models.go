package marketplaceordering

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AgreementProperties is agreement Terms definition
type AgreementProperties struct {
	Publisher         *string    `json:"publisher,omitempty"`
	Product           *string    `json:"product,omitempty"`
	Plan              *string    `json:"plan,omitempty"`
	LicenseTextLink   *string    `json:"licenseTextLink,omitempty"`
	PrivacyPolicyLink *string    `json:"privacyPolicyLink,omitempty"`
	RetrieveDatetime  *date.Time `json:"retrieveDatetime,omitempty"`
	Signature         *string    `json:"signature,omitempty"`
	Accepted          *bool      `json:"accepted,omitempty"`
}

// AgreementTerms is terms properties for provided Publisher/Offer/Plan tuple
type AgreementTerms struct {
	autorest.Response    `json:"-"`
	ID                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Type                 *string `json:"type,omitempty"`
	*AgreementProperties `json:"properties,omitempty"`
}

// ErrorResponse is error reponse indicates Microsoft.MarketplaceOrdering service is not able to process the incoming
// request. The reason is provided in the error message.
type ErrorResponse struct {
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError is the details of the error.
type ErrorResponseError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Operation is microsoft.MarketplaceOrdering REST API operation
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of the request to list MarketplaceOrdering operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is ARM resource.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}
