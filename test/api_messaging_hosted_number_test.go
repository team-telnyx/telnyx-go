/*
Telnyx API

Testing MessagingHostedNumberAPIService

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

func Test_telnyx_MessagingHostedNumberAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MessagingHostedNumberAPIService CheckEligibilityNumbers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.CheckEligibilityNumbers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService CreateMessagingHostedNumberOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.CreateMessagingHostedNumberOrder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService CreateVerificationCodesForMessagingHostedNumberOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.CreateVerificationCodesForMessagingHostedNumberOrder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService DeleteMessagingHostedNumber", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.DeleteMessagingHostedNumber(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService GetMessagingHostedNumberOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.GetMessagingHostedNumberOrder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService ListMessagingHostedNumberOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.ListMessagingHostedNumberOrders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService UploadMessagingHostedNumberOrderFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.UploadMessagingHostedNumberOrderFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagingHostedNumberAPIService ValidateVerificationCodesForMessagingHostedNumberOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MessagingHostedNumberAPI.ValidateVerificationCodesForMessagingHostedNumberOrder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
