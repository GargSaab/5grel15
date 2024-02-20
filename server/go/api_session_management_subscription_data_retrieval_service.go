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

// SessionManagementSubscriptionDataRetrievalAPIService is a service that implements the logic for the SessionManagementSubscriptionDataRetrievalAPIServicer
// This service should implement the business logic for every endpoint for the SessionManagementSubscriptionDataRetrievalAPI API.
// Include any external packages or services that will be required by this service.
type SessionManagementSubscriptionDataRetrievalAPIService struct {
}

// NewSessionManagementSubscriptionDataRetrievalAPIService creates a default api service
func NewSessionManagementSubscriptionDataRetrievalAPIService() SessionManagementSubscriptionDataRetrievalAPIServicer {
	return &SessionManagementSubscriptionDataRetrievalAPIService{}
}

// GetSmData - retrieve a UE&#39;s Session Management Subscription Data
func (s *SessionManagementSubscriptionDataRetrievalAPIService) GetSmData(ctx context.Context, supi string, supportedFeatures string, singleNssai Snssai, dnn string, plmnId PlmnId, disasterRoamingInd bool, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {
	// TODO - update GetSmData with the required logic for this service method.
	// Add api_session_management_subscription_data_retrieval_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmSubsData{}) or use other options such as http.Ok ...
	// return Response(200, SmSubsData{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("GetSmData method not implemented")
}
