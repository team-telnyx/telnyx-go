/*
Telnyx API

Testing CustomerServiceRecordAPIService

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

func Test_telnyx_CustomerServiceRecordAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomerServiceRecordAPIService CreateCustomerServiceRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerServiceRecordAPI.CreateCustomerServiceRecord(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerServiceRecordAPIService GetCustomerServiceRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerServiceRecordId string

		resp, httpRes, err := apiClient.CustomerServiceRecordAPI.GetCustomerServiceRecord(context.Background(), customerServiceRecordId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerServiceRecordAPIService ListCustomerServiceRecords", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerServiceRecordAPI.ListCustomerServiceRecords(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomerServiceRecordAPIService VerifyPhoneNumberCoverage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomerServiceRecordAPI.VerifyPhoneNumberCoverage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
