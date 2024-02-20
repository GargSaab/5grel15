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

// TimeSyncSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type TimeSyncSubscriptionDataRetrievalAPIController struct {
	service      TimeSyncSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// TimeSyncSubscriptionDataRetrievalAPIOption for how the controller is set up.
type TimeSyncSubscriptionDataRetrievalAPIOption func(*TimeSyncSubscriptionDataRetrievalAPIController)

// WithTimeSyncSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithTimeSyncSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) TimeSyncSubscriptionDataRetrievalAPIOption {
	return func(c *TimeSyncSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewTimeSyncSubscriptionDataRetrievalAPIController creates a default api controller
func NewTimeSyncSubscriptionDataRetrievalAPIController(s TimeSyncSubscriptionDataRetrievalAPIServicer, opts ...TimeSyncSubscriptionDataRetrievalAPIOption) Router {
	controller := &TimeSyncSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TimeSyncSubscriptionDataRetrievalAPIController
func (c *TimeSyncSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetTimeSyncSubscriptionData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/time-sync-data",
			c.GetTimeSyncSubscriptionData,
		},
	}
}

// GetTimeSyncSubscriptionData - retrieve a UE's Time Synchronization Subscription Data
func (c *TimeSyncSubscriptionDataRetrievalAPIController) GetTimeSyncSubscriptionData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supiParam := params["supi"]
	if supiParam == "" {
		c.errorHandler(w, r, &RequiredError{"supi"}, nil)
		return
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetTimeSyncSubscriptionData(r.Context(), supiParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
