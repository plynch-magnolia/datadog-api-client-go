/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func generateFastlyService(ctx context.Context, t *testing.T) datadog.FastlyService {
	serviceId := *tests.UniqueEntityName(ctx, t)
	return datadog.FastlyService{
		ServiceId:      serviceId,
		ServiceTags:    "fastly-test:test",
		ServiceAccount: "TESTING-ACCOUNT",
	}
}

func TestFastlyLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// Add single service object to the Fastly Integration
	serviceBody := generateFastlyService(ctx, t)
	serviceResponse, httpresp, err := Client(ctx).FastlyIntegrationApi.CreateFastlyIntegrationService(ctx).Body(serviceBody).Execute()
	defer deleteFastlyService(ctx, t, serviceBody.ServiceId)
	assert.NoError(err)
	assert.Equal(201, httpresp.StatusCode)
	assert.Equal(serviceBody.GetServiceId(), serviceResponse.GetServiceId())

	// Get created Service object
	serviceName, httpresp, err := Client(ctx).FastlyIntegrationApi.GetFastlyIntegrationService(ctx, serviceBody.GetServiceId()).Execute()
	assert.NoError(err)
	assert.Equal(serviceBody.GetServiceId(), serviceName.GetServiceId())
	assert.Equal(200, httpresp.StatusCode)

	// Update service object
	serviceId := datadog.FastlyServiceId{
		ServiceId: "newkey",
	}
	serviceTags := datadog.FastlyServiceTags{
		ServiceTags: "newkey:newvalue",
	}
	serviceAccount := datadog.FastlyServiceAccount{
		ServiceAccount: "NEW-TEST-ACCOUNT",
	}
	httpresp, err = Client(ctx).FastlyIntegrationApi.UpdateFastlyIntegrationService(ctx, serviceId.GetServiceId(), serviceTags.GetServiceTags(), serviceAccount.GetServiceAccount()).Execute()
	assert.NoError(err)
	assert.Equal(200, httpresp.StatusCode)

	// Delete service object
	httpresp, err = Client(ctx).FastlyIntegrationApi.DeleteFastlyIntegrationService(ctx, serviceBody.GetServiceId()).Execute()
	assert.NoError(err)
	assert.Equal(200, httpresp.StatusCode)

	// Check service object
	_, httpresp, err = Client(ctx).FastlyIntegrationApi.GetFastlyIntegrationService(ctx, serviceBody.GetServiceId()).Execute()
	assert.Error(err)
	assert.Equal(404, httpresp.StatusCode)
}

func TestFastlyServicesCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(ctx, t)
	defer finish()

	fastlyService := generateFastlyService(ctx, t)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.FastlyService
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.FastlyService{}, 400},
		"403 Forbidden":   {WithFakeAuth, fastlyService, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			service, httpresp, err := Client(ctx).FastlyIntegrationApi.CreateFastlyIntegrationService(ctx).Body(tc.Body).Execute()
			if err == nil {
				defer deleteFastlyService(ctx, t, service.GetServiceId())
			}
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestFastlyServicesGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			fastlyService := generateFastlyService(ctx, t)

			_, httpresp, err := Client(ctx).FastlyIntegrationApi.GetFastlyIntegrationService(ctx, fastlyService.GetServiceId()).Execute()
			assert.NotNil(httpresp)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			errors, _ := apiError.GetErrorsOk()
			assert.NotNil(errors)
		})
	}
}

func TestFastlyServicesUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	service := generateFastlyService(ctx, t)
	_, _, err := Client(ctx).FastlyIntegrationApi.CreateFastlyIntegrationService(ctx).Body(service).Execute()
	defer deleteFastlyService(ctx, t, service.ServiceId)
	if err != nil {
		t.Errorf("could not create service %v: %v", service, err)
	}

	serviceId := *datadog.NewFastlyServiceIdWithDefaults()
	serviceId.SetServiceId("lalaa")
	serviceTags := *datadog.NewFastlyServiceTagsWithDefaults()
	serviceTags.SetServiceTags("newtag:newvalue")
	serviceAccount := *datadog.NewFastlyServiceAccountWithDefaults()
	serviceAccount.SetServiceAccount("New-Test-Account")

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ServiceId          string
		Body               datadog.FastlyServiceId
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, service.ServiceId, datadog.FastlyServiceId{}, 400},
		"403 Forbidden":   {WithFakeAuth, service.ServiceId, serviceId, 403},
		"404 Not Found":   {WithTestAuth, service.ServiceId + "-invalid", serviceId, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			httpresp, err := Client(ctx).FastlyIntegrationApi.UpdateFastlyIntegrationService(ctx, tc.ServiceId, serviceTags.GetServiceTags(), serviceAccount.GetServiceAccount()).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestFastlyServicesDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			fastlyService := generateFastlyService(ctx, t)

			httpresp, err := Client(ctx).FastlyIntegrationApi.DeleteFastlyIntegrationService(ctx, fastlyService.GetServiceId()).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			errors, _ := apiError.GetErrorsOk()
			assert.NotNil(errors)
		})
	}
}

func deleteFastlyService(ctx context.Context, t *testing.T, serviceId string) {
	httpresp, err := Client(ctx).FastlyIntegrationApi.DeleteFastlyIntegrationService(ctx, serviceId).Execute()
	if httpresp.StatusCode != 204 || err != nil {
		t.Logf("Error deleting Fastly Service %s: Another test may have already removed this account: %v", serviceId, err)
	}
}
