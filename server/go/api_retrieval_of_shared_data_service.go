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

// RetrievalOfSharedDataAPIService is a service that implements the logic for the RetrievalOfSharedDataAPIServicer
// This service should implement the business logic for every endpoint for the RetrievalOfSharedDataAPI API.
// Include any external packages or services that will be required by this service.
type RetrievalOfSharedDataAPIService struct {
}

// NewRetrievalOfSharedDataAPIService creates a default api service
func NewRetrievalOfSharedDataAPIService() RetrievalOfSharedDataAPIServicer {
	return &RetrievalOfSharedDataAPIService{}
}

// GetSharedData - retrieve shared data
func (s *RetrievalOfSharedDataAPIService) GetSharedData(ctx context.Context, sharedDataIds []string, supportedFeatures string, supportedFeatures2 string, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {
	// TODO - update GetSharedData with the required logic for this service method.
	// Add api_retrieval_of_shared_data_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []SharedData{}) or use other options such as http.Ok ...
	// return Response(200, []SharedData{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("GetSharedData method not implemented")
}
