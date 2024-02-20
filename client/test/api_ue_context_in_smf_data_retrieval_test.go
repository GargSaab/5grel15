/*
Nudm_SDM

Testing UEContextInSMFDataRetrievalAPIService

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

func Test_openapi_UEContextInSMFDataRetrievalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UEContextInSMFDataRetrievalAPIService GetUeCtxInSmfData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supi string

		resp, httpRes, err := apiClient.UEContextInSMFDataRetrievalAPI.GetUeCtxInSmfData(context.Background(), supi).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}