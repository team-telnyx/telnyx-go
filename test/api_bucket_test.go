/*
Telnyx API

Testing BucketAPIService

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

func Test_telnyx_BucketAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BucketAPIService CreateBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketName string

		httpRes, err := apiClient.BucketAPI.CreateBucket(context.Background(), bucketName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService DeleteBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketName string

		httpRes, err := apiClient.BucketAPI.DeleteBucket(context.Background(), bucketName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService HeadBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucketName string

		httpRes, err := apiClient.BucketAPI.HeadBucket(context.Background(), bucketName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketAPIService ListBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BucketAPI.ListBuckets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
