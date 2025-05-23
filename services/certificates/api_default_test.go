/*
Load Balancer Certificates API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package certificates

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

func Test_certificates_DefaultApiService(t *testing.T) {

	t.Run("Test DefaultApiService CreateCertificate", func(t *testing.T) {
		_apiUrlPath := "/v2beta/projects/{projectId}/regions/{region}/certificates"
		projectIdValue := "projectId-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := CreateCertificateResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for certificates_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := projectIdValue
		region := regionValue
		createCertificatePayload := CreateCertificatePayload{}

		resp, reqErr := apiClient.CreateCertificate(context.Background(), projectId, region).CreateCertificatePayload(createCertificatePayload).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService DeleteCertificate", func(t *testing.T) {
		_apiUrlPath := "/v2beta/projects/{projectId}/regions/{region}/certificates/{id}"
		projectIdValue := "projectId-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		idValue := "id-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"id"+"}", url.PathEscape(ParameterValueToString(idValue, "id")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := map[string]interface{}{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for certificates_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := projectIdValue
		region := regionValue
		id := idValue

		resp, reqErr := apiClient.DeleteCertificate(context.Background(), projectId, region, id).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService GetCertificate", func(t *testing.T) {
		_apiUrlPath := "/v2beta/projects/{projectId}/regions/{region}/certificates/{id}"
		projectIdValue := "projectId-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)
		idValue := "id-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"id"+"}", url.PathEscape(ParameterValueToString(idValue, "id")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := GetCertificateResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for certificates_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := projectIdValue
		region := regionValue
		id := idValue

		resp, reqErr := apiClient.GetCertificate(context.Background(), projectId, region, id).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

	t.Run("Test DefaultApiService ListCertificates", func(t *testing.T) {
		_apiUrlPath := "/v2beta/projects/{projectId}/regions/{region}/certificates"
		projectIdValue := "projectId-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(projectIdValue, "projectId")), -1)
		regionValue := "region-value"
		_apiUrlPath = strings.Replace(_apiUrlPath, "{"+"region"+"}", url.PathEscape(ParameterValueToString(regionValue, "region")), -1)

		testDefaultApiServeMux := http.NewServeMux()
		testDefaultApiServeMux.HandleFunc(_apiUrlPath, func(w http.ResponseWriter, req *http.Request) {
			data := ListCertificatesResponse{}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
		testServer := httptest.NewServer(testDefaultApiServeMux)
		defer testServer.Close()

		configuration := &config.Configuration{
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Region:        "test_region",
			Servers: config.ServerConfigurations{
				{
					URL:         testServer.URL,
					Description: "Localhost for certificates_DefaultApi",
					Variables: map[string]config.ServerVariable{
						"region": {
							DefaultValue: "test_region.",
							EnumValues: []string{
								"test_region.",
							},
						},
					},
				},
			},
			OperationServers: map[string]config.ServerConfigurations{},
		}
		apiClient, err := NewAPIClient(config.WithCustomConfiguration(configuration), config.WithoutAuthentication())
		if err != nil {
			t.Fatalf("creating API client: %v", err)
		}

		projectId := projectIdValue
		region := regionValue

		resp, reqErr := apiClient.ListCertificates(context.Background(), projectId, region).Execute()

		if reqErr != nil {
			t.Fatalf("error in call: %v", reqErr)
		}
		if IsNil(resp) {
			t.Fatalf("response not present")
		}
	})

}
