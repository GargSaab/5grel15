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

// SubscriptionDeletionForSharedDataAPIController binds http requests to an api service and writes the service results to the http response
type SubscriptionDeletionForSharedDataAPIController struct {
	service      SubscriptionDeletionForSharedDataAPIServicer
	errorHandler ErrorHandler
}

// SubscriptionDeletionForSharedDataAPIOption for how the controller is set up.
type SubscriptionDeletionForSharedDataAPIOption func(*SubscriptionDeletionForSharedDataAPIController)

// WithSubscriptionDeletionForSharedDataAPIErrorHandler inject ErrorHandler into controller
func WithSubscriptionDeletionForSharedDataAPIErrorHandler(h ErrorHandler) SubscriptionDeletionForSharedDataAPIOption {
	return func(c *SubscriptionDeletionForSharedDataAPIController) {
		c.errorHandler = h
	}
}

// NewSubscriptionDeletionForSharedDataAPIController creates a default api controller
func NewSubscriptionDeletionForSharedDataAPIController(s SubscriptionDeletionForSharedDataAPIServicer, opts ...SubscriptionDeletionForSharedDataAPIOption) Router {
	controller := &SubscriptionDeletionForSharedDataAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SubscriptionDeletionForSharedDataAPIController
func (c *SubscriptionDeletionForSharedDataAPIController) Routes() Routes {
	return Routes{
		"UnsubscribeForSharedData": Route{
			strings.ToUpper("Delete"),
			"/nudm-sdm/v2/shared-data-subscriptions/{subscriptionId}",
			c.UnsubscribeForSharedData,
		},
	}
}

// UnsubscribeForSharedData - unsubscribe from notifications for shared data
func (c *SubscriptionDeletionForSharedDataAPIController) UnsubscribeForSharedData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.UnsubscribeForSharedData(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
