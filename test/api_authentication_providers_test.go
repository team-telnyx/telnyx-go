/*
Telnyx API

Testing AuthenticationProvidersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package telnyx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/telnyx/telnyx-go"
)

func Test_telnyx_AuthenticationProvidersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthenticationProvidersAPIService CreateAuthenticationProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthenticationProvidersAPI.CreateAuthenticationProvider(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationProvidersAPIService DeleteAuthenticationProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AuthenticationProvidersAPI.DeleteAuthenticationProvider(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationProvidersAPIService FindAuthenticationProviders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthenticationProvidersAPI.FindAuthenticationProviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationProvidersAPIService GetAuthenticationProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AuthenticationProvidersAPI.GetAuthenticationProvider(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationProvidersAPIService UpdateAuthenticationProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AuthenticationProvidersAPI.UpdateAuthenticationProvider(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
