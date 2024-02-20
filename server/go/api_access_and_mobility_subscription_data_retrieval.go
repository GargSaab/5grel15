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

// AccessAndMobilitySubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type AccessAndMobilitySubscriptionDataRetrievalAPIController struct {
	service      AccessAndMobilitySubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// AccessAndMobilitySubscriptionDataRetrievalAPIOption for how the controller is set up.
type AccessAndMobilitySubscriptionDataRetrievalAPIOption func(*AccessAndMobilitySubscriptionDataRetrievalAPIController)

// WithAccessAndMobilitySubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithAccessAndMobilitySubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) AccessAndMobilitySubscriptionDataRetrievalAPIOption {
	return func(c *AccessAndMobilitySubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewAccessAndMobilitySubscriptionDataRetrievalAPIController creates a default api controller
func NewAccessAndMobilitySubscriptionDataRetrievalAPIController(s AccessAndMobilitySubscriptionDataRetrievalAPIServicer, opts ...AccessAndMobilitySubscriptionDataRetrievalAPIOption) Router {
	controller := &AccessAndMobilitySubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AccessAndMobilitySubscriptionDataRetrievalAPIController
func (c *AccessAndMobilitySubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetAmData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/am-data",
			c.GetAmData,
		},
	}
}

// GetAmData - retrieve a UE's Access and Mobility Subscription Data
func (c *AccessAndMobilitySubscriptionDataRetrievalAPIController) GetAmData(w http.ResponseWriter, r *http.Request) {
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
	var plmnIdParam PlmnIdNid
	if query.Has("plmn-id") {
		plmnIdParam = query.Get("plmn-id")
	}
	var adjacentPlmnsParam []string
	if query.Has("adjacent-plmns") {
		adjacentPlmnsParam = strings.Split(query.Get("adjacent-plmns"), ",")
	}
	disasterRoamingIndParam, err := parseBoolParameter(
		query.Get("disaster-roaming-ind"),
		WithDefaultOrParse[bool](false, parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var sharedDataIdsParam []string
	if query.Has("shared-data-ids") {
		sharedDataIdsParam = strings.Split(query.Get("shared-data-ids"), ",")
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetAmData(r.Context(), supiParam, supportedFeaturesParam, plmnIdParam, adjacentPlmnsParam, disasterRoamingIndParam, sharedDataIdsParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
