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

// SessionManagementSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type SessionManagementSubscriptionDataRetrievalAPIController struct {
	service      SessionManagementSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// SessionManagementSubscriptionDataRetrievalAPIOption for how the controller is set up.
type SessionManagementSubscriptionDataRetrievalAPIOption func(*SessionManagementSubscriptionDataRetrievalAPIController)

// WithSessionManagementSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithSessionManagementSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) SessionManagementSubscriptionDataRetrievalAPIOption {
	return func(c *SessionManagementSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewSessionManagementSubscriptionDataRetrievalAPIController creates a default api controller
func NewSessionManagementSubscriptionDataRetrievalAPIController(s SessionManagementSubscriptionDataRetrievalAPIServicer, opts ...SessionManagementSubscriptionDataRetrievalAPIOption) Router {
	controller := &SessionManagementSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SessionManagementSubscriptionDataRetrievalAPIController
func (c *SessionManagementSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetSmData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/sm-data",
			c.GetSmData,
		},
	}
}

// GetSmData - retrieve a UE's Session Management Subscription Data
func (c *SessionManagementSubscriptionDataRetrievalAPIController) GetSmData(w http.ResponseWriter, r *http.Request) {
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
	var singleNssaiParam Snssai
	if query.Has("single-nssai") {
		singleNssaiParam = query.Get("single-nssai")
	}
	var dnnParam string
	if query.Has("dnn") {
		dnnParam = query.Get("dnn")
	}
	var plmnIdParam PlmnId
	if query.Has("plmn-id") {
		plmnIdParam = query.Get("plmn-id")
	}
	disasterRoamingIndParam, err := parseBoolParameter(
		query.Get("disaster-roaming-ind"),
		WithDefaultOrParse[bool](false, parseBool),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetSmData(r.Context(), supiParam, supportedFeaturesParam, singleNssaiParam, dnnParam, plmnIdParam, disasterRoamingIndParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}