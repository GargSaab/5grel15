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

// Class5MBSSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type Class5MBSSubscriptionDataRetrievalAPIController struct {
	service      Class5MBSSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// Class5MBSSubscriptionDataRetrievalAPIOption for how the controller is set up.
type Class5MBSSubscriptionDataRetrievalAPIOption func(*Class5MBSSubscriptionDataRetrievalAPIController)

// WithClass5MBSSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithClass5MBSSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) Class5MBSSubscriptionDataRetrievalAPIOption {
	return func(c *Class5MBSSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewClass5MBSSubscriptionDataRetrievalAPIController creates a default api controller
func NewClass5MBSSubscriptionDataRetrievalAPIController(s Class5MBSSubscriptionDataRetrievalAPIServicer, opts ...Class5MBSSubscriptionDataRetrievalAPIOption) Router {
	controller := &Class5MBSSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the Class5MBSSubscriptionDataRetrievalAPIController
func (c *Class5MBSSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetMbsData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/5mbs-data",
			c.GetMbsData,
		},
	}
}

// GetMbsData - retrieve a UE's 5MBS Subscription Data
func (c *Class5MBSSubscriptionDataRetrievalAPIController) GetMbsData(w http.ResponseWriter, r *http.Request) {
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
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetMbsData(r.Context(), supiParam, supportedFeaturesParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
