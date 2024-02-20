/*
Nudm_SDM

Testing ProvidingAcknowledgementOfCAGUpdateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ProvidingAcknowledgementOfCAGUpdateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProvidingAcknowledgementOfCAGUpdateAPIService CAGAck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supi string

		httpRes, err := apiClient.ProvidingAcknowledgementOfCAGUpdateAPI.CAGAck(context.Background(), supi).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
