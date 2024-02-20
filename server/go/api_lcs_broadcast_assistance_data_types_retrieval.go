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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// LCSBroadcastAssistanceDataTypesRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type LCSBroadcastAssistanceDataTypesRetrievalAPIController struct {
	service      LCSBroadcastAssistanceDataTypesRetrievalAPIServicer
	errorHandler ErrorHandler
}

// LCSBroadcastAssistanceDataTypesRetrievalAPIOption for how the controller is set up.
type LCSBroadcastAssistanceDataTypesRetrievalAPIOption func(*LCSBroadcastAssistanceDataTypesRetrievalAPIController)

// WithLCSBroadcastAssistanceDataTypesRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithLCSBroadcastAssistanceDataTypesRetrievalAPIErrorHandler(h ErrorHandler) LCSBroadcastAssistanceDataTypesRetrievalAPIOption {
	return func(c *LCSBroadcastAssistanceDataTypesRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewLCSBroadcastAssistanceDataTypesRetrievalAPIController creates a default api controller
func NewLCSBroadcastAssistanceDataTypesRetrievalAPIController(s LCSBroadcastAssistanceDataTypesRetrievalAPIServicer, opts ...LCSBroadcastAssistanceDataTypesRetrievalAPIOption) Router {
	controller := &LCSBroadcastAssistanceDataTypesRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LCSBroadcastAssistanceDataTypesRetrievalAPIController
func (c *LCSBroadcastAssistanceDataTypesRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetLcsBcaData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/lcs-bca-data",
			c.GetLcsBcaData,
		},
	}
}

// GetLcsBcaData - retrieve a UE's LCS Broadcast Assistance Data Types Subscription Data
func (c *LCSBroadcastAssistanceDataTypesRetrievalAPIController) GetLcsBcaData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	supiParam := params["supi"]
	if supiParam == "" {
		c.errorHandler(w, r, &RequiredError{"supi"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var plmnIdParam string
	if query.Has("plmn-id") {
		plmnIdParam = query.Get("plmn-id")
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetLcsBcaData(r.Context(), supiParam, supportedFeaturesParam, plmnIdParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
