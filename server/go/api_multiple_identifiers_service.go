/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// MultipleIdentifiersAPIService is a service that implements the logic for the MultipleIdentifiersAPIServicer
// This service should implement the business logic for every endpoint for the MultipleIdentifiersAPI API.
// Include any external packages or services that will be required by this service.
type MultipleIdentifiersAPIService struct {
}

// NewMultipleIdentifiersAPIService creates a default api service
func NewMultipleIdentifiersAPIService() MultipleIdentifiersAPIServicer {
	return &MultipleIdentifiersAPIService{}
}

// GetMultipleIdentifiers - Mapping of UE Identifiers
func (s *MultipleIdentifiersAPIService) GetMultipleIdentifiers(ctx context.Context, gpsiList []string, supportedFeatures string) (ImplResponse, error) {
	// TODO - update GetMultipleIdentifiers with the required logic for this service method.
	// Add api_multiple_identifiers_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, map[string]SupiInfo{}) or use other options such as http.Ok ...
	// return Response(200, map[string]SupiInfo{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(401, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(401, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(406, {}) or use other options such as http.Ok ...
	// return Response(406, nil),nil

	// TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(429, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(502, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(502, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMultipleIdentifiers method not implemented")
}
