/**
 * (C) Copyright IBM Corp. 2025.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package clouddatabasesv5_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`CloudDatabasesV5`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(cloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(cloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
				URL: "https://clouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(cloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_URL": "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{
				})
				Expect(cloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := cloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(cloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{
				})
				err := cloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := cloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != cloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(cloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(cloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_URL": "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = clouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := clouddatabasesv5.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://api.us-south.databases.cloud.ibm.com/v5/ibm"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := clouddatabasesv5.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListDeployables(listDeployablesOptions *ListDeployablesOptions) - Operation response error`, func() {
		listDeployablesPath := "/deployables"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeployablesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDeployables with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := new(clouddatabasesv5.ListDeployablesOptions)
				listDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeployables(listDeployablesOptions *ListDeployablesOptions)`, func() {
		listDeployablesPath := "/deployables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeployablesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
				}))
			})
			It(`Invoke ListDeployables successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := new(clouddatabasesv5.ListDeployablesOptions)
				listDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListDeployablesWithContext(ctx, listDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListDeployablesWithContext(ctx, listDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeployablesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
				}))
			})
			It(`Invoke ListDeployables successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListDeployables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := new(clouddatabasesv5.ListDeployablesOptions)
				listDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDeployables with error: Operation request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := new(clouddatabasesv5.ListDeployablesOptions)
				listDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDeployables successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := new(clouddatabasesv5.ListDeployablesOptions)
				listDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListDeployables(listDeployablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRegions(listRegionsOptions *ListRegionsOptions) - Operation response error`, func() {
		listRegionsPath := "/regions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRegionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRegions with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(clouddatabasesv5.ListRegionsOptions)
				listRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRegions(listRegionsOptions *ListRegionsOptions)`, func() {
		listRegionsPath := "/regions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regions": ["Regions"]}`)
				}))
			})
			It(`Invoke ListRegions successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(clouddatabasesv5.ListRegionsOptions)
				listRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListRegionsWithContext(ctx, listRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListRegionsWithContext(ctx, listRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regions": ["Regions"]}`)
				}))
			})
			It(`Invoke ListRegions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(clouddatabasesv5.ListRegionsOptions)
				listRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRegions with error: Operation request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(clouddatabasesv5.ListRegionsOptions)
				listRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListRegions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := new(clouddatabasesv5.ListRegionsOptions)
				listRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListRegions(listRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions) - Operation response error`, func() {
		getDeploymentInfoPath := "/deployments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentInfo with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(clouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions)`, func() {
		getDeploymentInfoPath := "/deployments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform": "satellite, classic", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": {"mapKey": "Inner"}, "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
				}))
			})
			It(`Invoke GetDeploymentInfo successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(clouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetDeploymentInfoWithContext(ctx, getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetDeploymentInfoWithContext(ctx, getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform": "satellite, classic", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": {"mapKey": "Inner"}, "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
				}))
			})
			It(`Invoke GetDeploymentInfo successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetDeploymentInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(clouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentInfo with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(clouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentInfoOptions model with no property values
				getDeploymentInfoOptionsModelNew := new(clouddatabasesv5.GetDeploymentInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDeploymentInfo successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(clouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions) - Operation response error`, func() {
		createDatabaseUserPath := "/deployments/testString/users/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDatabaseUser with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = userModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions)`, func() {
		createDatabaseUserPath := "/deployments/testString/users/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateDatabaseUser successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = userModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.CreateDatabaseUserWithContext(ctx, createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.CreateDatabaseUserWithContext(ctx, createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateDatabaseUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.CreateDatabaseUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = userModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDatabaseUser with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = userModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDatabaseUserOptions model with no property values
				createDatabaseUserOptionsModelNew := new(clouddatabasesv5.CreateDatabaseUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateDatabaseUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = userModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUser(updateUserOptions *UpdateUserOptions) - Operation response error`, func() {
		updateUserPath := "/deployments/testString/users/database/user"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateUser with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the UpdateUserOptions model
				updateUserOptionsModel := new(clouddatabasesv5.UpdateUserOptions)
				updateUserOptionsModel.ID = core.StringPtr("testString")
				updateUserOptionsModel.UserType = core.StringPtr("database")
				updateUserOptionsModel.Username = core.StringPtr("user")
				updateUserOptionsModel.User = userUpdateModel
				updateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUser(updateUserOptions *UpdateUserOptions)`, func() {
		updateUserPath := "/deployments/testString/users/database/user"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateUser successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the UpdateUserOptions model
				updateUserOptionsModel := new(clouddatabasesv5.UpdateUserOptions)
				updateUserOptionsModel.ID = core.StringPtr("testString")
				updateUserOptionsModel.UserType = core.StringPtr("database")
				updateUserOptionsModel.Username = core.StringPtr("user")
				updateUserOptionsModel.User = userUpdateModel
				updateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.UpdateUserWithContext(ctx, updateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.UpdateUserWithContext(ctx, updateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateUserPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.UpdateUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the UpdateUserOptions model
				updateUserOptionsModel := new(clouddatabasesv5.UpdateUserOptions)
				updateUserOptionsModel.ID = core.StringPtr("testString")
				updateUserOptionsModel.UserType = core.StringPtr("database")
				updateUserOptionsModel.Username = core.StringPtr("user")
				updateUserOptionsModel.User = userUpdateModel
				updateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateUser with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the UpdateUserOptions model
				updateUserOptionsModel := new(clouddatabasesv5.UpdateUserOptions)
				updateUserOptionsModel.ID = core.StringPtr("testString")
				updateUserOptionsModel.UserType = core.StringPtr("database")
				updateUserOptionsModel.Username = core.StringPtr("user")
				updateUserOptionsModel.User = userUpdateModel
				updateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateUserOptions model with no property values
				updateUserOptionsModelNew := new(clouddatabasesv5.UpdateUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.UpdateUser(updateUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

				// Construct an instance of the UpdateUserOptions model
				updateUserOptionsModel := new(clouddatabasesv5.UpdateUserOptions)
				updateUserOptionsModel.ID = core.StringPtr("testString")
				updateUserOptionsModel.UserType = core.StringPtr("database")
				updateUserOptionsModel.Username = core.StringPtr("user")
				updateUserOptionsModel.User = userUpdateModel
				updateUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.UpdateUser(updateUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions) - Operation response error`, func() {
		deleteDatabaseUserPath := "/deployments/testString/users/database/user"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDatabaseUser with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("database")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("user")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions)`, func() {
		deleteDatabaseUserPath := "/deployments/testString/users/database/user"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteDatabaseUser successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("database")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("user")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.DeleteDatabaseUserWithContext(ctx, deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.DeleteDatabaseUserWithContext(ctx, deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteDatabaseUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.DeleteDatabaseUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("database")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("user")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDatabaseUser with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("database")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("user")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDatabaseUserOptions model with no property values
				deleteDatabaseUserOptionsModelNew := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteDatabaseUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(clouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("database")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("user")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions *UpdateDatabaseConfigurationOptions) - Operation response error`, func() {
		updateDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDatabaseConfiguration with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				updateDatabaseConfigurationOptionsModel := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				updateDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				updateDatabaseConfigurationOptionsModel.Configuration = configurationModel
				updateDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions *UpdateDatabaseConfigurationOptions)`, func() {
		updateDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateDatabaseConfiguration successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				updateDatabaseConfigurationOptionsModel := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				updateDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				updateDatabaseConfigurationOptionsModel.Configuration = configurationModel
				updateDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.UpdateDatabaseConfigurationWithContext(ctx, updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.UpdateDatabaseConfigurationWithContext(ctx, updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateDatabaseConfiguration successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.UpdateDatabaseConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				updateDatabaseConfigurationOptionsModel := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				updateDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				updateDatabaseConfigurationOptionsModel.Configuration = configurationModel
				updateDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDatabaseConfiguration with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				updateDatabaseConfigurationOptionsModel := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				updateDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				updateDatabaseConfigurationOptionsModel.Configuration = configurationModel
				updateDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDatabaseConfigurationOptions model with no property values
				updateDatabaseConfigurationOptionsModelNew := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDatabaseConfiguration successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				updateDatabaseConfigurationOptionsModel := new(clouddatabasesv5.UpdateDatabaseConfigurationOptions)
				updateDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				updateDatabaseConfigurationOptionsModel.Configuration = configurationModel
				updateDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRemotes(listRemotesOptions *ListRemotesOptions) - Operation response error`, func() {
		listRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRemotesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRemotes with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRemotesOptions model
				listRemotesOptionsModel := new(clouddatabasesv5.ListRemotesOptions)
				listRemotesOptionsModel.ID = core.StringPtr("testString")
				listRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRemotes(listRemotesOptions *ListRemotesOptions)`, func() {
		listRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRemotesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"remotes": {"leader": "01f30581-54f8-41a4-8193-4a04cc022e9b-h", "replicas": ["Replicas"]}}`)
				}))
			})
			It(`Invoke ListRemotes successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListRemotesOptions model
				listRemotesOptionsModel := new(clouddatabasesv5.ListRemotesOptions)
				listRemotesOptionsModel.ID = core.StringPtr("testString")
				listRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListRemotesWithContext(ctx, listRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListRemotesWithContext(ctx, listRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRemotesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"remotes": {"leader": "01f30581-54f8-41a4-8193-4a04cc022e9b-h", "replicas": ["Replicas"]}}`)
				}))
			})
			It(`Invoke ListRemotes successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListRemotes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRemotesOptions model
				listRemotesOptionsModel := new(clouddatabasesv5.ListRemotesOptions)
				listRemotesOptionsModel.ID = core.StringPtr("testString")
				listRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRemotes with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRemotesOptions model
				listRemotesOptionsModel := new(clouddatabasesv5.ListRemotesOptions)
				listRemotesOptionsModel.ID = core.StringPtr("testString")
				listRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRemotesOptions model with no property values
				listRemotesOptionsModelNew := new(clouddatabasesv5.ListRemotesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ListRemotes(listRemotesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListRemotes successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListRemotesOptions model
				listRemotesOptionsModel := new(clouddatabasesv5.ListRemotesOptions)
				listRemotesOptionsModel.ID = core.StringPtr("testString")
				listRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListRemotes(listRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResyncReplica(resyncReplicaOptions *ResyncReplicaOptions) - Operation response error`, func() {
		resyncReplicaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resyncReplicaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ResyncReplica with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ResyncReplicaOptions model
				resyncReplicaOptionsModel := new(clouddatabasesv5.ResyncReplicaOptions)
				resyncReplicaOptionsModel.ID = core.StringPtr("testString")
				resyncReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ResyncReplica(resyncReplicaOptions *ResyncReplicaOptions)`, func() {
		resyncReplicaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resyncReplicaPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ResyncReplica successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ResyncReplicaOptions model
				resyncReplicaOptionsModel := new(clouddatabasesv5.ResyncReplicaOptions)
				resyncReplicaOptionsModel.ID = core.StringPtr("testString")
				resyncReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ResyncReplicaWithContext(ctx, resyncReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ResyncReplicaWithContext(ctx, resyncReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(resyncReplicaPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ResyncReplica successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ResyncReplica(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResyncReplicaOptions model
				resyncReplicaOptionsModel := new(clouddatabasesv5.ResyncReplicaOptions)
				resyncReplicaOptionsModel.ID = core.StringPtr("testString")
				resyncReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ResyncReplica with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ResyncReplicaOptions model
				resyncReplicaOptionsModel := new(clouddatabasesv5.ResyncReplicaOptions)
				resyncReplicaOptionsModel.ID = core.StringPtr("testString")
				resyncReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ResyncReplicaOptions model with no property values
				resyncReplicaOptionsModelNew := new(clouddatabasesv5.ResyncReplicaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ResyncReplica successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ResyncReplicaOptions model
				resyncReplicaOptionsModel := new(clouddatabasesv5.ResyncReplicaOptions)
				resyncReplicaOptionsModel.ID = core.StringPtr("testString")
				resyncReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ResyncReplica(resyncReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions *PromoteReadOnlyReplicaOptions) - Operation response error`, func() {
		promoteReadOnlyReplicaPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteReadOnlyReplicaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PromoteReadOnlyReplica with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				promoteReadOnlyReplicaOptionsModel := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				promoteReadOnlyReplicaOptionsModel.ID = core.StringPtr("testString")
				promoteReadOnlyReplicaOptionsModel.Promotion = map[string]interface{}{"anyKey": "anyValue"}
				promoteReadOnlyReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions *PromoteReadOnlyReplicaOptions)`, func() {
		promoteReadOnlyReplicaPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteReadOnlyReplicaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke PromoteReadOnlyReplica successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				promoteReadOnlyReplicaOptionsModel := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				promoteReadOnlyReplicaOptionsModel.ID = core.StringPtr("testString")
				promoteReadOnlyReplicaOptionsModel.Promotion = map[string]interface{}{"anyKey": "anyValue"}
				promoteReadOnlyReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.PromoteReadOnlyReplicaWithContext(ctx, promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.PromoteReadOnlyReplicaWithContext(ctx, promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(promoteReadOnlyReplicaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke PromoteReadOnlyReplica successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.PromoteReadOnlyReplica(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				promoteReadOnlyReplicaOptionsModel := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				promoteReadOnlyReplicaOptionsModel.ID = core.StringPtr("testString")
				promoteReadOnlyReplicaOptionsModel.Promotion = map[string]interface{}{"anyKey": "anyValue"}
				promoteReadOnlyReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PromoteReadOnlyReplica with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				promoteReadOnlyReplicaOptionsModel := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				promoteReadOnlyReplicaOptionsModel.ID = core.StringPtr("testString")
				promoteReadOnlyReplicaOptionsModel.Promotion = map[string]interface{}{"anyKey": "anyValue"}
				promoteReadOnlyReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PromoteReadOnlyReplicaOptions model with no property values
				promoteReadOnlyReplicaOptionsModelNew := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PromoteReadOnlyReplica successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				promoteReadOnlyReplicaOptionsModel := new(clouddatabasesv5.PromoteReadOnlyReplicaOptions)
				promoteReadOnlyReplicaOptionsModel.ID = core.StringPtr("testString")
				promoteReadOnlyReplicaOptionsModel.Promotion = map[string]interface{}{"anyKey": "anyValue"}
				promoteReadOnlyReplicaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentTasks(listDeploymentTasksOptions *ListDeploymentTasksOptions) - Operation response error`, func() {
		listDeploymentTasksPath := "/deployments/testString/tasks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDeploymentTasks with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentTasksOptions model
				listDeploymentTasksOptionsModel := new(clouddatabasesv5.ListDeploymentTasksOptions)
				listDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				listDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentTasks(listDeploymentTasksOptions *ListDeploymentTasksOptions)`, func() {
		listDeploymentTasksPath := "/deployments/testString/tasks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDeploymentTasks successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListDeploymentTasksOptions model
				listDeploymentTasksOptionsModel := new(clouddatabasesv5.ListDeploymentTasksOptions)
				listDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				listDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListDeploymentTasksWithContext(ctx, listDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListDeploymentTasksWithContext(ctx, listDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDeploymentTasks successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListDeploymentTasks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDeploymentTasksOptions model
				listDeploymentTasksOptionsModel := new(clouddatabasesv5.ListDeploymentTasksOptions)
				listDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				listDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDeploymentTasks with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentTasksOptions model
				listDeploymentTasksOptionsModel := new(clouddatabasesv5.ListDeploymentTasksOptions)
				listDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				listDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDeploymentTasksOptions model with no property values
				listDeploymentTasksOptionsModelNew := new(clouddatabasesv5.ListDeploymentTasksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDeploymentTasks successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentTasksOptions model
				listDeploymentTasksOptionsModel := new(clouddatabasesv5.ListDeploymentTasksOptions)
				listDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				listDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTask(getTaskOptions *GetTaskOptions) - Operation response error`, func() {
		getTaskPath := "/tasks/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTask with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(clouddatabasesv5.GetTaskOptions)
				getTaskOptionsModel.ID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTask(getTaskOptions *GetTaskOptions)`, func() {
		getTaskPath := "/tasks/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetTask successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(clouddatabasesv5.GetTaskOptions)
				getTaskOptionsModel.ID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetTaskWithContext(ctx, getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetTaskWithContext(ctx, getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetTask successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetTask(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(clouddatabasesv5.GetTaskOptions)
				getTaskOptionsModel.ID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTask with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(clouddatabasesv5.GetTaskOptions)
				getTaskOptionsModel.ID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTaskOptions model with no property values
				getTaskOptionsModelNew := new(clouddatabasesv5.GetTaskOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetTask(getTaskOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTask successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(clouddatabasesv5.GetTaskOptions)
				getTaskOptionsModel.ID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions) - Operation response error`, func() {
		getBackupInfoPath := "/backups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBackupInfo with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(clouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions)`, func() {
		getBackupInfoPath := "/backups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "download_link": "https://securedownloadservice.com/backup-2018-02-28T19:25:12Z.tgz", "created_at": "2018-02-28T19:25:12.000Z"}}`)
				}))
			})
			It(`Invoke GetBackupInfo successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(clouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetBackupInfoWithContext(ctx, getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetBackupInfoWithContext(ctx, getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "download_link": "https://securedownloadservice.com/backup-2018-02-28T19:25:12Z.tgz", "created_at": "2018-02-28T19:25:12.000Z"}}`)
				}))
			})
			It(`Invoke GetBackupInfo successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetBackupInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(clouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBackupInfo with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(clouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBackupInfoOptions model with no property values
				getBackupInfoOptionsModelNew := new(clouddatabasesv5.GetBackupInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBackupInfo successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(clouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentBackups(listDeploymentBackupsOptions *ListDeploymentBackupsOptions) - Operation response error`, func() {
		listDeploymentBackupsPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDeploymentBackups with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentBackupsOptions model
				listDeploymentBackupsOptionsModel := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				listDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentBackups(listDeploymentBackupsOptions *ListDeploymentBackupsOptions)`, func() {
		listDeploymentBackupsPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "download_link": "https://securedownloadservice.com/backup-2018-02-28T19:25:12Z.tgz", "created_at": "2018-02-28T19:25:12.000Z"}]}`)
				}))
			})
			It(`Invoke ListDeploymentBackups successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListDeploymentBackupsOptions model
				listDeploymentBackupsOptionsModel := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				listDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListDeploymentBackupsWithContext(ctx, listDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListDeploymentBackupsWithContext(ctx, listDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "download_link": "https://securedownloadservice.com/backup-2018-02-28T19:25:12Z.tgz", "created_at": "2018-02-28T19:25:12.000Z"}]}`)
				}))
			})
			It(`Invoke ListDeploymentBackups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListDeploymentBackups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDeploymentBackupsOptions model
				listDeploymentBackupsOptionsModel := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				listDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDeploymentBackups with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentBackupsOptions model
				listDeploymentBackupsOptionsModel := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				listDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDeploymentBackupsOptions model with no property values
				listDeploymentBackupsOptionsModelNew := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDeploymentBackups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentBackupsOptions model
				listDeploymentBackupsOptionsModel := new(clouddatabasesv5.ListDeploymentBackupsOptions)
				listDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions) - Operation response error`, func() {
		startOndemandBackupPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke StartOndemandBackup with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(clouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions)`, func() {
		startOndemandBackupPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke StartOndemandBackup successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(clouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.StartOndemandBackupWithContext(ctx, startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.StartOndemandBackupWithContext(ctx, startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke StartOndemandBackup successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.StartOndemandBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(clouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke StartOndemandBackup with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(clouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the StartOndemandBackupOptions model with no property values
				startOndemandBackupOptionsModelNew := new(clouddatabasesv5.StartOndemandBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke StartOndemandBackup successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(clouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPitrData(getPitrDataOptions *GetPitrDataOptions) - Operation response error`, func() {
		getPitrDataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitrDataPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPitrData with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPitrDataOptions model
				getPitrDataOptionsModel := new(clouddatabasesv5.GetPitrDataOptions)
				getPitrDataOptionsModel.ID = core.StringPtr("testString")
				getPitrDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPitrData(getPitrDataOptions *GetPitrDataOptions)`, func() {
		getPitrDataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitrDataPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"point_in_time_recovery_data": {"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}}`)
				}))
			})
			It(`Invoke GetPitrData successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetPitrDataOptions model
				getPitrDataOptionsModel := new(clouddatabasesv5.GetPitrDataOptions)
				getPitrDataOptionsModel.ID = core.StringPtr("testString")
				getPitrDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetPitrDataWithContext(ctx, getPitrDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetPitrDataWithContext(ctx, getPitrDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitrDataPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"point_in_time_recovery_data": {"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}}`)
				}))
			})
			It(`Invoke GetPitrData successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetPitrData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPitrDataOptions model
				getPitrDataOptionsModel := new(clouddatabasesv5.GetPitrDataOptions)
				getPitrDataOptionsModel.ID = core.StringPtr("testString")
				getPitrDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPitrData with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPitrDataOptions model
				getPitrDataOptionsModel := new(clouddatabasesv5.GetPitrDataOptions)
				getPitrDataOptionsModel.ID = core.StringPtr("testString")
				getPitrDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPitrDataOptions model with no property values
				getPitrDataOptionsModelNew := new(clouddatabasesv5.GetPitrDataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetPitrData(getPitrDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPitrData successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPitrDataOptions model
				getPitrDataOptionsModel := new(clouddatabasesv5.GetPitrDataOptions)
				getPitrDataOptionsModel.ID = core.StringPtr("testString")
				getPitrDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetPitrData(getPitrDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConnection(getConnectionOptions *GetConnectionOptions) - Operation response error`, func() {
		getConnectionPath := "/deployments/testString/users/database/testString/connections/public"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConnection with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(clouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("database")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConnection(getConnectionOptions *GetConnectionOptions)`, func() {
		getConnectionPath := "/deployments/testString/users/database/testString/connections/public"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "Path", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "ssl": false, "browser_accessible": false, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"anyKey": "anyValue"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnection successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(clouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("database")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetConnectionWithContext(ctx, getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetConnectionWithContext(ctx, getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "Path", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "ssl": false, "browser_accessible": false, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"anyKey": "anyValue"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnection successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(clouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("database")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConnection with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(clouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("database")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConnectionOptions model with no property values
				getConnectionOptionsModelNew := new(clouddatabasesv5.GetConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetConnection(getConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetConnection successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(clouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("database")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions) - Operation response error`, func() {
		completeConnectionPath := "/deployments/testString/users/database/testString/connections/public"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteConnection with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(clouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("database")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("providedpassword")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions)`, func() {
		completeConnectionPath := "/deployments/testString/users/database/testString/connections/public"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "Path", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "ssl": false, "browser_accessible": false, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"anyKey": "anyValue"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnection successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(clouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("database")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("providedpassword")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.CompleteConnectionWithContext(ctx, completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.CompleteConnectionWithContext(ctx, completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "Path", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "ssl": false, "browser_accessible": false, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"anyKey": "anyValue"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnection successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.CompleteConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(clouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("database")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("providedpassword")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteConnection with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(clouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("database")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("providedpassword")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteConnectionOptions model with no property values
				completeConnectionOptionsModelNew := new(clouddatabasesv5.CompleteConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.CompleteConnection(completeConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CompleteConnection successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(clouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("database")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("providedpassword")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions) - Operation response error`, func() {
		listDeploymentScalingGroupsPath := "/deployments/testString/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDeploymentScalingGroups with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				listDeploymentScalingGroupsOptionsModel := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				listDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions)`, func() {
		listDeploymentScalingGroupsPath := "/deployments/testString/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}]}`)
				}))
			})
			It(`Invoke ListDeploymentScalingGroups successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				listDeploymentScalingGroupsOptionsModel := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				listDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ListDeploymentScalingGroupsWithContext(ctx, listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ListDeploymentScalingGroupsWithContext(ctx, listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}]}`)
				}))
			})
			It(`Invoke ListDeploymentScalingGroups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ListDeploymentScalingGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				listDeploymentScalingGroupsOptionsModel := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				listDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDeploymentScalingGroups with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				listDeploymentScalingGroupsOptionsModel := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				listDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDeploymentScalingGroupsOptions model with no property values
				listDeploymentScalingGroupsOptionsModelNew := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDeploymentScalingGroups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				listDeploymentScalingGroupsOptionsModel := new(clouddatabasesv5.ListDeploymentScalingGroupsOptions)
				listDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				listDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) - Operation response error`, func() {
		getDefaultScalingGroupsPath := "/deployables/postgresql/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"multitenant"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.HostFlavor = core.StringPtr("multitenant")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions)`, func() {
		getDefaultScalingGroupsPath := "/deployables/postgresql/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"multitenant"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}]}`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.HostFlavor = core.StringPtr("multitenant")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetDefaultScalingGroupsWithContext(ctx, getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetDefaultScalingGroupsWithContext(ctx, getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"multitenant"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}]}`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetDefaultScalingGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.HostFlavor = core.StringPtr("multitenant")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDefaultScalingGroups with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.HostFlavor = core.StringPtr("multitenant")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDefaultScalingGroupsOptions model with no property values
				getDefaultScalingGroupsOptionsModelNew := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDefaultScalingGroups successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(clouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.HostFlavor = core.StringPtr("multitenant")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) - Operation response error`, func() {
		setDeploymentScalingGroupPath := "/deployments/testString/groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.Group = groupScalingModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions)`, func() {
		setDeploymentScalingGroupPath := "/deployments/testString/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.Group = groupScalingModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetDeploymentScalingGroupWithContext(ctx, setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetDeploymentScalingGroupWithContext(ctx, setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetDeploymentScalingGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.Group = groupScalingModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetDeploymentScalingGroup with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.Group = groupScalingModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetDeploymentScalingGroupOptions model with no property values
				setDeploymentScalingGroupOptionsModelNew := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.Group = groupScalingModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) - Operation response error`, func() {
		getAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAutoscalingConditions with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions)`, func() {
		getAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"autoscaling": {"disk": {"scalers": {"capacity": {"enabled": true, "free_space_less_than_percent": 10}, "io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 20, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "memory": {"scalers": {"io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "cpu": {"scalers": {"anyKey": "anyValue"}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_count_per_member": 10, "units": "count"}}}}`)
				}))
			})
			It(`Invoke GetAutoscalingConditions successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetAutoscalingConditionsWithContext(ctx, getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetAutoscalingConditionsWithContext(ctx, getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"autoscaling": {"disk": {"scalers": {"capacity": {"enabled": true, "free_space_less_than_percent": 10}, "io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 20, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "memory": {"scalers": {"io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "cpu": {"scalers": {"anyKey": "anyValue"}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_count_per_member": 10, "units": "count"}}}}`)
				}))
			})
			It(`Invoke GetAutoscalingConditions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetAutoscalingConditions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAutoscalingConditions with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAutoscalingConditionsOptions model with no property values
				getAutoscalingConditionsOptionsModelNew := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAutoscalingConditions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(clouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) - Operation response error`, func() {
		setAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetAutoscalingConditions with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupAutoscalingModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions)`, func() {
		setAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetAutoscalingConditions successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupAutoscalingModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetAutoscalingConditionsWithContext(ctx, setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetAutoscalingConditionsWithContext(ctx, setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetAutoscalingConditions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetAutoscalingConditions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupAutoscalingModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetAutoscalingConditions with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupAutoscalingModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetAutoscalingConditionsOptions model with no property values
				setAutoscalingConditionsOptionsModelNew := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke SetAutoscalingConditions successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(clouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupAutoscalingModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`KillConnections(killConnectionsOptions *KillConnectionsOptions) - Operation response error`, func() {
		killConnectionsPath := "/deployments/testString/management/database_connections"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke KillConnections with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(clouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`KillConnections(killConnectionsOptions *KillConnectionsOptions)`, func() {
		killConnectionsPath := "/deployments/testString/management/database_connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke KillConnections successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(clouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.KillConnectionsWithContext(ctx, killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.KillConnectionsWithContext(ctx, killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke KillConnections successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.KillConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(clouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke KillConnections with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(clouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the KillConnectionsOptions model with no property values
				killConnectionsOptionsModelNew := new(clouddatabasesv5.KillConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.KillConnections(killConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke KillConnections successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(clouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions) - Operation response error`, func() {
		createLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions)`, func() {
		createLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.CreateLogicalReplicationSlotWithContext(ctx, createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.CreateLogicalReplicationSlotWithContext(ctx, createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.CreateLogicalReplicationSlot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLogicalReplicationSlot with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLogicalReplicationSlotOptions model with no property values
				createLogicalReplicationSlotOptionsModelNew := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions) - Operation response error`, func() {
		deleteLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions)`, func() {
		deleteLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlotWithContext(ctx, deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.DeleteLogicalReplicationSlotWithContext(ctx, deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteLogicalReplicationSlot with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLogicalReplicationSlotOptions model with no property values
				deleteLogicalReplicationSlotOptionsModelNew := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllowlist(getAllowlistOptions *GetAllowlistOptions) - Operation response error`, func() {
		getAllowlistPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllowlistPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAllowlist with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAllowlistOptions model
				getAllowlistOptionsModel := new(clouddatabasesv5.GetAllowlistOptions)
				getAllowlistOptionsModel.ID = core.StringPtr("testString")
				getAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllowlist(getAllowlistOptions *GetAllowlistOptions)`, func() {
		getAllowlistPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllowlistPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "Address", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetAllowlist successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetAllowlistOptions model
				getAllowlistOptionsModel := new(clouddatabasesv5.GetAllowlistOptions)
				getAllowlistOptionsModel.ID = core.StringPtr("testString")
				getAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetAllowlistWithContext(ctx, getAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetAllowlistWithContext(ctx, getAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllowlistPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "Address", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetAllowlist successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetAllowlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllowlistOptions model
				getAllowlistOptionsModel := new(clouddatabasesv5.GetAllowlistOptions)
				getAllowlistOptionsModel.ID = core.StringPtr("testString")
				getAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAllowlist with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAllowlistOptions model
				getAllowlistOptionsModel := new(clouddatabasesv5.GetAllowlistOptions)
				getAllowlistOptionsModel.ID = core.StringPtr("testString")
				getAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAllowlistOptions model with no property values
				getAllowlistOptionsModelNew := new(clouddatabasesv5.GetAllowlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetAllowlist(getAllowlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAllowlist successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAllowlistOptions model
				getAllowlistOptionsModel := new(clouddatabasesv5.GetAllowlistOptions)
				getAllowlistOptionsModel.ID = core.StringPtr("testString")
				getAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetAllowlist(getAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetAllowlist(setAllowlistOptions *SetAllowlistOptions) - Operation response error`, func() {
		setAllowlistPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAllowlistPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetAllowlist with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")

				// Construct an instance of the SetAllowlistOptions model
				setAllowlistOptionsModel := new(clouddatabasesv5.SetAllowlistOptions)
				setAllowlistOptionsModel.ID = core.StringPtr("testString")
				setAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				setAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				setAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetAllowlist(setAllowlistOptions *SetAllowlistOptions)`, func() {
		setAllowlistPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAllowlistPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetAllowlist successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")

				// Construct an instance of the SetAllowlistOptions model
				setAllowlistOptionsModel := new(clouddatabasesv5.SetAllowlistOptions)
				setAllowlistOptionsModel.ID = core.StringPtr("testString")
				setAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				setAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				setAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetAllowlistWithContext(ctx, setAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetAllowlistWithContext(ctx, setAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAllowlistPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetAllowlist successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetAllowlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")

				// Construct an instance of the SetAllowlistOptions model
				setAllowlistOptionsModel := new(clouddatabasesv5.SetAllowlistOptions)
				setAllowlistOptionsModel.ID = core.StringPtr("testString")
				setAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				setAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				setAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetAllowlist with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")

				// Construct an instance of the SetAllowlistOptions model
				setAllowlistOptionsModel := new(clouddatabasesv5.SetAllowlistOptions)
				setAllowlistOptionsModel.ID = core.StringPtr("testString")
				setAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				setAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				setAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetAllowlistOptions model with no property values
				setAllowlistOptionsModelNew := new(clouddatabasesv5.SetAllowlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetAllowlist(setAllowlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke SetAllowlist successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")

				// Construct an instance of the SetAllowlistOptions model
				setAllowlistOptionsModel := new(clouddatabasesv5.SetAllowlistOptions)
				setAllowlistOptionsModel.ID = core.StringPtr("testString")
				setAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				setAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				setAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.SetAllowlist(setAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions) - Operation response error`, func() {
		addAllowlistEntryPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAllowlistEntryPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddAllowlistEntry with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")

				// Construct an instance of the AddAllowlistEntryOptions model
				addAllowlistEntryOptionsModel := new(clouddatabasesv5.AddAllowlistEntryOptions)
				addAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				addAllowlistEntryOptionsModel.IPAddress = allowlistEntryModel
				addAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions)`, func() {
		addAllowlistEntryPath := "/deployments/testString/allowlists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAllowlistEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke AddAllowlistEntry successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")

				// Construct an instance of the AddAllowlistEntryOptions model
				addAllowlistEntryOptionsModel := new(clouddatabasesv5.AddAllowlistEntryOptions)
				addAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				addAllowlistEntryOptionsModel.IPAddress = allowlistEntryModel
				addAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.AddAllowlistEntryWithContext(ctx, addAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.AddAllowlistEntryWithContext(ctx, addAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAllowlistEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke AddAllowlistEntry successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.AddAllowlistEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")

				// Construct an instance of the AddAllowlistEntryOptions model
				addAllowlistEntryOptionsModel := new(clouddatabasesv5.AddAllowlistEntryOptions)
				addAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				addAllowlistEntryOptionsModel.IPAddress = allowlistEntryModel
				addAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddAllowlistEntry with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")

				// Construct an instance of the AddAllowlistEntryOptions model
				addAllowlistEntryOptionsModel := new(clouddatabasesv5.AddAllowlistEntryOptions)
				addAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				addAllowlistEntryOptionsModel.IPAddress = allowlistEntryModel
				addAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddAllowlistEntryOptions model with no property values
				addAllowlistEntryOptionsModelNew := new(clouddatabasesv5.AddAllowlistEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke AddAllowlistEntry successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")

				// Construct an instance of the AddAllowlistEntryOptions model
				addAllowlistEntryOptionsModel := new(clouddatabasesv5.AddAllowlistEntryOptions)
				addAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				addAllowlistEntryOptionsModel.IPAddress = allowlistEntryModel
				addAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions) - Operation response error`, func() {
		deleteAllowlistEntryPath := "/deployments/testString/allowlists/ip_addresses/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllowlistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllowlistEntry with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteAllowlistEntryOptions model
				deleteAllowlistEntryOptionsModel := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				deleteAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions)`, func() {
		deleteAllowlistEntryPath := "/deployments/testString/allowlists/ip_addresses/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllowlistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteAllowlistEntry successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAllowlistEntryOptions model
				deleteAllowlistEntryOptionsModel := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				deleteAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.DeleteAllowlistEntryWithContext(ctx, deleteAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.DeleteAllowlistEntryWithContext(ctx, deleteAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllowlistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke DeleteAllowlistEntry successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.DeleteAllowlistEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllowlistEntryOptions model
				deleteAllowlistEntryOptionsModel := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				deleteAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteAllowlistEntry with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteAllowlistEntryOptions model
				deleteAllowlistEntryOptionsModel := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				deleteAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteAllowlistEntryOptions model with no property values
				deleteAllowlistEntryOptionsModelNew := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteAllowlistEntry successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteAllowlistEntryOptions model
				deleteAllowlistEntryOptionsModel := new(clouddatabasesv5.DeleteAllowlistEntryOptions)
				deleteAllowlistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteAllowlistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCapability(createCapabilityOptions *CreateCapabilityOptions) - Operation response error`, func() {
		createCapabilityPath := "/capability/autoscaling"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCapabilityPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCapability with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")

				// Construct an instance of the CreateCapabilityOptions model
				createCapabilityOptionsModel := new(clouddatabasesv5.CreateCapabilityOptions)
				createCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				createCapabilityOptionsModel.Deployment = createCapabilityRequestDeploymentModel
				createCapabilityOptionsModel.Backup = createCapabilityRequestBackupModel
				createCapabilityOptionsModel.Options = createCapabilityRequestOptionsModel
				createCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCapability(createCapabilityOptions *CreateCapabilityOptions)`, func() {
		createCapabilityPath := "/capability/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCapabilityPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"capability": {"autoscaling": {"autoscaling_supported": true}, "encryption": {"disk_encryption_supported": true}, "endpoints": {"direct_endpoints_supported": false, "public_endpoints_supported": true, "private_endpoints_supported": true, "multiple_endpoints_supported": true}, "groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}], "flavors": [{"id": "b3c.4x16.encrypted", "name": "4x16", "cpu": {"allocation_count": 2}, "memory": {"allocation_mb": 12288}, "hosting_size": "xs"}], "locations": {"locations": ["us-south"]}, "point_in_time_recovery": {"point_in_time_recovery_supported": true}, "remotes": {"read_only_replicas_supported": true}, "restores": {"backup_restore_supported": true}, "versions": [{"type": "elasticsearch", "version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}}`)
				}))
			})
			It(`Invoke CreateCapability successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")

				// Construct an instance of the CreateCapabilityOptions model
				createCapabilityOptionsModel := new(clouddatabasesv5.CreateCapabilityOptions)
				createCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				createCapabilityOptionsModel.Deployment = createCapabilityRequestDeploymentModel
				createCapabilityOptionsModel.Backup = createCapabilityRequestBackupModel
				createCapabilityOptionsModel.Options = createCapabilityRequestOptionsModel
				createCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.CreateCapabilityWithContext(ctx, createCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.CreateCapabilityWithContext(ctx, createCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCapabilityPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"capability": {"autoscaling": {"autoscaling_supported": true}, "encryption": {"disk_encryption_supported": true}, "endpoints": {"direct_endpoints_supported": false, "public_endpoints_supported": true, "private_endpoints_supported": true, "multiple_endpoints_supported": true}, "groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}], "flavors": [{"id": "b3c.4x16.encrypted", "name": "4x16", "cpu": {"allocation_count": 2}, "memory": {"allocation_mb": 12288}, "hosting_size": "xs"}], "locations": {"locations": ["us-south"]}, "point_in_time_recovery": {"point_in_time_recovery_supported": true}, "remotes": {"read_only_replicas_supported": true}, "restores": {"backup_restore_supported": true}, "versions": [{"type": "elasticsearch", "version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}}`)
				}))
			})
			It(`Invoke CreateCapability successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.CreateCapability(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")

				// Construct an instance of the CreateCapabilityOptions model
				createCapabilityOptionsModel := new(clouddatabasesv5.CreateCapabilityOptions)
				createCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				createCapabilityOptionsModel.Deployment = createCapabilityRequestDeploymentModel
				createCapabilityOptionsModel.Backup = createCapabilityRequestBackupModel
				createCapabilityOptionsModel.Options = createCapabilityRequestOptionsModel
				createCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCapability with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")

				// Construct an instance of the CreateCapabilityOptions model
				createCapabilityOptionsModel := new(clouddatabasesv5.CreateCapabilityOptions)
				createCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				createCapabilityOptionsModel.Deployment = createCapabilityRequestDeploymentModel
				createCapabilityOptionsModel.Backup = createCapabilityRequestBackupModel
				createCapabilityOptionsModel.Options = createCapabilityRequestOptionsModel
				createCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCapabilityOptions model with no property values
				createCapabilityOptionsModelNew := new(clouddatabasesv5.CreateCapabilityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.CreateCapability(createCapabilityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateCapability successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")

				// Construct an instance of the CreateCapabilityOptions model
				createCapabilityOptionsModel := new(clouddatabasesv5.CreateCapabilityOptions)
				createCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				createCapabilityOptionsModel.Deployment = createCapabilityRequestDeploymentModel
				createCapabilityOptionsModel.Backup = createCapabilityRequestBackupModel
				createCapabilityOptionsModel.Options = createCapabilityRequestOptionsModel
				createCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.CreateCapability(createCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentCapability(getDeploymentCapabilityOptions *GetDeploymentCapabilityOptions) - Operation response error`, func() {
		getDeploymentCapabilityPath := "/deployments/testString/capability/autoscaling"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentCapabilityPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["target_platform"]).To(Equal([]string{"target_platform=classic"}))
					Expect(req.URL.Query()["target_location"]).To(Equal([]string{"target_location=us-east"}))
					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"host_flavor=multitenant"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentCapability with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentCapabilityOptions model
				getDeploymentCapabilityOptionsModel := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				getDeploymentCapabilityOptionsModel.ID = core.StringPtr("testString")
				getDeploymentCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				getDeploymentCapabilityOptionsModel.TargetPlatform = core.StringPtr("target_platform=classic")
				getDeploymentCapabilityOptionsModel.TargetLocation = core.StringPtr("target_location=us-east")
				getDeploymentCapabilityOptionsModel.HostFlavor = core.StringPtr("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentCapability(getDeploymentCapabilityOptions *GetDeploymentCapabilityOptions)`, func() {
		getDeploymentCapabilityPath := "/deployments/testString/capability/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentCapabilityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["target_platform"]).To(Equal([]string{"target_platform=classic"}))
					Expect(req.URL.Query()["target_location"]).To(Equal([]string{"target_location=us-east"}))
					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"host_flavor=multitenant"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"capability": {"autoscaling": {"autoscaling_supported": true}, "encryption": {"disk_encryption_supported": true}, "endpoints": {"direct_endpoints_supported": false, "public_endpoints_supported": true, "private_endpoints_supported": true, "multiple_endpoints_supported": true}, "groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}], "flavors": [{"id": "b3c.4x16.encrypted", "name": "4x16", "cpu": {"allocation_count": 2}, "memory": {"allocation_mb": 12288}, "hosting_size": "xs"}], "locations": {"locations": ["us-south"]}, "point_in_time_recovery": {"point_in_time_recovery_supported": true}, "remotes": {"read_only_replicas_supported": true}, "restores": {"backup_restore_supported": true}, "versions": [{"type": "elasticsearch", "version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}}`)
				}))
			})
			It(`Invoke GetDeploymentCapability successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentCapabilityOptions model
				getDeploymentCapabilityOptionsModel := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				getDeploymentCapabilityOptionsModel.ID = core.StringPtr("testString")
				getDeploymentCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				getDeploymentCapabilityOptionsModel.TargetPlatform = core.StringPtr("target_platform=classic")
				getDeploymentCapabilityOptionsModel.TargetLocation = core.StringPtr("target_location=us-east")
				getDeploymentCapabilityOptionsModel.HostFlavor = core.StringPtr("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetDeploymentCapabilityWithContext(ctx, getDeploymentCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetDeploymentCapabilityWithContext(ctx, getDeploymentCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentCapabilityPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["target_platform"]).To(Equal([]string{"target_platform=classic"}))
					Expect(req.URL.Query()["target_location"]).To(Equal([]string{"target_location=us-east"}))
					Expect(req.URL.Query()["host_flavor"]).To(Equal([]string{"host_flavor=multitenant"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"capability": {"autoscaling": {"autoscaling_supported": true}, "encryption": {"disk_encryption_supported": true}, "endpoints": {"direct_endpoints_supported": false, "public_endpoints_supported": true, "private_endpoints_supported": true, "multiple_endpoints_supported": true}, "groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true, "cpu_enforcement_ratio_ceiling_mb": 16384, "cpu_enforcement_ratio_mb": 8192}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "host_flavor": {"id": "b3c.4x16.encrypted", "name": "4x16", "hosting_size": "xs"}}], "flavors": [{"id": "b3c.4x16.encrypted", "name": "4x16", "cpu": {"allocation_count": 2}, "memory": {"allocation_mb": 12288}, "hosting_size": "xs"}], "locations": {"locations": ["us-south"]}, "point_in_time_recovery": {"point_in_time_recovery_supported": true}, "remotes": {"read_only_replicas_supported": true}, "restores": {"backup_restore_supported": true}, "versions": [{"type": "elasticsearch", "version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "skip_backup_supported": false, "from_version": "5.6", "to_version": "6.7"}]}]}}`)
				}))
			})
			It(`Invoke GetDeploymentCapability successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetDeploymentCapability(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentCapabilityOptions model
				getDeploymentCapabilityOptionsModel := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				getDeploymentCapabilityOptionsModel.ID = core.StringPtr("testString")
				getDeploymentCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				getDeploymentCapabilityOptionsModel.TargetPlatform = core.StringPtr("target_platform=classic")
				getDeploymentCapabilityOptionsModel.TargetLocation = core.StringPtr("target_location=us-east")
				getDeploymentCapabilityOptionsModel.HostFlavor = core.StringPtr("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentCapability with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentCapabilityOptions model
				getDeploymentCapabilityOptionsModel := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				getDeploymentCapabilityOptionsModel.ID = core.StringPtr("testString")
				getDeploymentCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				getDeploymentCapabilityOptionsModel.TargetPlatform = core.StringPtr("target_platform=classic")
				getDeploymentCapabilityOptionsModel.TargetLocation = core.StringPtr("target_location=us-east")
				getDeploymentCapabilityOptionsModel.HostFlavor = core.StringPtr("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentCapabilityOptions model with no property values
				getDeploymentCapabilityOptionsModelNew := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDeploymentCapability successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentCapabilityOptions model
				getDeploymentCapabilityOptionsModel := new(clouddatabasesv5.GetDeploymentCapabilityOptions)
				getDeploymentCapabilityOptionsModel.ID = core.StringPtr("testString")
				getDeploymentCapabilityOptionsModel.CapabilityID = core.StringPtr("autoscaling")
				getDeploymentCapabilityOptionsModel.TargetPlatform = core.StringPtr("target_platform=classic")
				getDeploymentCapabilityOptionsModel.TargetLocation = core.StringPtr("target_location=us-east")
				getDeploymentCapabilityOptionsModel.HostFlavor = core.StringPtr("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.GetDeploymentCapability(getDeploymentCapabilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptions *SetDatabaseInplaceVersionUpgradeOptions) - Operation response error`, func() {
		setDatabaseInplaceVersionUpgradePath := "/deployments/testString/version"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseInplaceVersionUpgradePath))
					Expect(req.Method).To(Equal("PATCH"))
					// TODO: Add check for skip_backup query parameter
					// TODO: Add check for expiration_datetime query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDatabaseInplaceVersionUpgrade with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				setDatabaseInplaceVersionUpgradeOptionsModel := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				setDatabaseInplaceVersionUpgradeOptionsModel.ID = core.StringPtr("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.Version = core.StringPtr("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup = core.BoolPtr(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime = CreateMockDateTime("2025-10-05T17:17:17Z")
				setDatabaseInplaceVersionUpgradeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptions *SetDatabaseInplaceVersionUpgradeOptions)`, func() {
		setDatabaseInplaceVersionUpgradePath := "/deployments/testString/version"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseInplaceVersionUpgradePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for skip_backup query parameter
					// TODO: Add check for expiration_datetime query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDatabaseInplaceVersionUpgrade successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				setDatabaseInplaceVersionUpgradeOptionsModel := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				setDatabaseInplaceVersionUpgradeOptionsModel.ID = core.StringPtr("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.Version = core.StringPtr("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup = core.BoolPtr(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime = CreateMockDateTime("2025-10-05T17:17:17Z")
				setDatabaseInplaceVersionUpgradeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgradeWithContext(ctx, setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetDatabaseInplaceVersionUpgradeWithContext(ctx, setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseInplaceVersionUpgradePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for skip_backup query parameter
					// TODO: Add check for expiration_datetime query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "resource_type": "backup", "description": "Description", "status": "queued", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDatabaseInplaceVersionUpgrade successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				setDatabaseInplaceVersionUpgradeOptionsModel := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				setDatabaseInplaceVersionUpgradeOptionsModel.ID = core.StringPtr("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.Version = core.StringPtr("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup = core.BoolPtr(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime = CreateMockDateTime("2025-10-05T17:17:17Z")
				setDatabaseInplaceVersionUpgradeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetDatabaseInplaceVersionUpgrade with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				setDatabaseInplaceVersionUpgradeOptionsModel := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				setDatabaseInplaceVersionUpgradeOptionsModel.ID = core.StringPtr("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.Version = core.StringPtr("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup = core.BoolPtr(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime = CreateMockDateTime("2025-10-05T17:17:17Z")
				setDatabaseInplaceVersionUpgradeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetDatabaseInplaceVersionUpgradeOptions model with no property values
				setDatabaseInplaceVersionUpgradeOptionsModelNew := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SetDatabaseInplaceVersionUpgrade successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				setDatabaseInplaceVersionUpgradeOptionsModel := new(clouddatabasesv5.SetDatabaseInplaceVersionUpgradeOptions)
				setDatabaseInplaceVersionUpgradeOptionsModel.ID = core.StringPtr("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.Version = core.StringPtr("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup = core.BoolPtr(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime = CreateMockDateTime("2025-10-05T17:17:17Z")
				setDatabaseInplaceVersionUpgradeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := cloudDatabasesService.SetDatabaseInplaceVersionUpgrade(setDatabaseInplaceVersionUpgradeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			cloudDatabasesService, _ := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
				URL:           "http://clouddatabasesv5modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddAllowlistEntryOptions successfully`, func() {
				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				Expect(allowlistEntryModel).ToNot(BeNil())
				allowlistEntryModel.Address = core.StringPtr("172.16.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 3")
				Expect(allowlistEntryModel.Address).To(Equal(core.StringPtr("172.16.0.0/16")))
				Expect(allowlistEntryModel.Description).To(Equal(core.StringPtr("Dev IP space 3")))

				// Construct an instance of the AddAllowlistEntryOptions model
				id := "testString"
				addAllowlistEntryOptionsModel := cloudDatabasesService.NewAddAllowlistEntryOptions(id)
				addAllowlistEntryOptionsModel.SetID("testString")
				addAllowlistEntryOptionsModel.SetIPAddress(allowlistEntryModel)
				addAllowlistEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addAllowlistEntryOptionsModel).ToNot(BeNil())
				Expect(addAllowlistEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(addAllowlistEntryOptionsModel.IPAddress).To(Equal(allowlistEntryModel))
				Expect(addAllowlistEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompleteConnectionOptions successfully`, func() {
				// Construct an instance of the CompleteConnectionOptions model
				id := "testString"
				userType := "database"
				userID := "testString"
				endpointType := "public"
				completeConnectionOptionsModel := cloudDatabasesService.NewCompleteConnectionOptions(id, userType, userID, endpointType)
				completeConnectionOptionsModel.SetID("testString")
				completeConnectionOptionsModel.SetUserType("database")
				completeConnectionOptionsModel.SetUserID("testString")
				completeConnectionOptionsModel.SetEndpointType("public")
				completeConnectionOptionsModel.SetPassword("providedpassword")
				completeConnectionOptionsModel.SetCertificateRoot("testString")
				completeConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeConnectionOptionsModel).ToNot(BeNil())
				Expect(completeConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(completeConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(completeConnectionOptionsModel.Password).To(Equal(core.StringPtr("providedpassword")))
				Expect(completeConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCapabilityOptions successfully`, func() {
				// Construct an instance of the CreateCapabilityRequestDeployment model
				createCapabilityRequestDeploymentModel := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
				Expect(createCapabilityRequestDeploymentModel).ToNot(BeNil())
				createCapabilityRequestDeploymentModel.Type = core.StringPtr("postgresql")
				createCapabilityRequestDeploymentModel.Version = core.StringPtr("10")
				createCapabilityRequestDeploymentModel.Platform = core.StringPtr("classic")
				createCapabilityRequestDeploymentModel.Location = core.StringPtr("us-south")
				createCapabilityRequestDeploymentModel.Plan = core.StringPtr("standard")
				Expect(createCapabilityRequestDeploymentModel.Type).To(Equal(core.StringPtr("postgresql")))
				Expect(createCapabilityRequestDeploymentModel.Version).To(Equal(core.StringPtr("10")))
				Expect(createCapabilityRequestDeploymentModel.Platform).To(Equal(core.StringPtr("classic")))
				Expect(createCapabilityRequestDeploymentModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createCapabilityRequestDeploymentModel.Plan).To(Equal(core.StringPtr("standard")))

				// Construct an instance of the CreateCapabilityRequestBackup model
				createCapabilityRequestBackupModel := new(clouddatabasesv5.CreateCapabilityRequestBackup)
				Expect(createCapabilityRequestBackupModel).ToNot(BeNil())
				createCapabilityRequestBackupModel.Type = core.StringPtr("PostgreSQL")
				createCapabilityRequestBackupModel.Version = core.StringPtr("10")
				createCapabilityRequestBackupModel.Platform = core.StringPtr("satellite")
				createCapabilityRequestBackupModel.Location = core.StringPtr("us-south")
				Expect(createCapabilityRequestBackupModel.Type).To(Equal(core.StringPtr("PostgreSQL")))
				Expect(createCapabilityRequestBackupModel.Version).To(Equal(core.StringPtr("10")))
				Expect(createCapabilityRequestBackupModel.Platform).To(Equal(core.StringPtr("satellite")))
				Expect(createCapabilityRequestBackupModel.Location).To(Equal(core.StringPtr("us-south")))

				// Construct an instance of the CreateCapabilityRequestOptions model
				createCapabilityRequestOptionsModel := new(clouddatabasesv5.CreateCapabilityRequestOptions)
				Expect(createCapabilityRequestOptionsModel).ToNot(BeNil())
				createCapabilityRequestOptionsModel.TargetPlatform = core.StringPtr("classic")
				createCapabilityRequestOptionsModel.TargetLocation = core.StringPtr("us-east")
				createCapabilityRequestOptionsModel.HostFlavor = core.StringPtr("multitenant")
				Expect(createCapabilityRequestOptionsModel.TargetPlatform).To(Equal(core.StringPtr("classic")))
				Expect(createCapabilityRequestOptionsModel.TargetLocation).To(Equal(core.StringPtr("us-east")))
				Expect(createCapabilityRequestOptionsModel.HostFlavor).To(Equal(core.StringPtr("multitenant")))

				// Construct an instance of the CreateCapabilityOptions model
				capabilityID := "autoscaling"
				createCapabilityOptionsModel := cloudDatabasesService.NewCreateCapabilityOptions(capabilityID)
				createCapabilityOptionsModel.SetCapabilityID("autoscaling")
				createCapabilityOptionsModel.SetDeployment(createCapabilityRequestDeploymentModel)
				createCapabilityOptionsModel.SetBackup(createCapabilityRequestBackupModel)
				createCapabilityOptionsModel.SetOptions(createCapabilityRequestOptionsModel)
				createCapabilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCapabilityOptionsModel).ToNot(BeNil())
				Expect(createCapabilityOptionsModel.CapabilityID).To(Equal(core.StringPtr("autoscaling")))
				Expect(createCapabilityOptionsModel.Deployment).To(Equal(createCapabilityRequestDeploymentModel))
				Expect(createCapabilityOptionsModel.Backup).To(Equal(createCapabilityRequestBackupModel))
				Expect(createCapabilityOptionsModel.Options).To(Equal(createCapabilityRequestOptionsModel))
				Expect(createCapabilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDatabaseUserOptions successfully`, func() {
				// Construct an instance of the UserDatabaseUser model
				userModel := new(clouddatabasesv5.UserDatabaseUser)
				Expect(userModel).ToNot(BeNil())
				userModel.Username = core.StringPtr("user")
				userModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
				Expect(userModel.Username).To(Equal(core.StringPtr("user")))
				Expect(userModel.Password).To(Equal(core.StringPtr("v3ry-1-secUre-pAssword-2")))

				// Construct an instance of the CreateDatabaseUserOptions model
				id := "testString"
				userType := "testString"
				createDatabaseUserOptionsModel := cloudDatabasesService.NewCreateDatabaseUserOptions(id, userType)
				createDatabaseUserOptionsModel.SetID("testString")
				createDatabaseUserOptionsModel.SetUserType("testString")
				createDatabaseUserOptionsModel.SetUser(userModel)
				createDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(createDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.User).To(Equal(userModel))
				Expect(createDatabaseUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLogicalReplicationSlotOptions successfully`, func() {
				// Construct an instance of the LogicalReplicationSlot model
				logicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlot)
				Expect(logicalReplicationSlotModel).ToNot(BeNil())
				logicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")
				Expect(logicalReplicationSlotModel.Name).To(Equal(core.StringPtr("customer_replication")))
				Expect(logicalReplicationSlotModel.DatabaseName).To(Equal(core.StringPtr("customers")))
				Expect(logicalReplicationSlotModel.PluginType).To(Equal(core.StringPtr("wal2json")))

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				id := "testString"
				createLogicalReplicationSlotOptionsModel := cloudDatabasesService.NewCreateLogicalReplicationSlotOptions(id)
				createLogicalReplicationSlotOptionsModel.SetID("testString")
				createLogicalReplicationSlotOptionsModel.SetLogicalReplicationSlot(logicalReplicationSlotModel)
				createLogicalReplicationSlotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogicalReplicationSlotOptionsModel).ToNot(BeNil())
				Expect(createLogicalReplicationSlotOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot).To(Equal(logicalReplicationSlotModel))
				Expect(createLogicalReplicationSlotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllowlistEntryOptions successfully`, func() {
				// Construct an instance of the DeleteAllowlistEntryOptions model
				id := "testString"
				ipaddress := "testString"
				deleteAllowlistEntryOptionsModel := cloudDatabasesService.NewDeleteAllowlistEntryOptions(id, ipaddress)
				deleteAllowlistEntryOptionsModel.SetID("testString")
				deleteAllowlistEntryOptionsModel.SetIpaddress("testString")
				deleteAllowlistEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllowlistEntryOptionsModel).ToNot(BeNil())
				Expect(deleteAllowlistEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllowlistEntryOptionsModel.Ipaddress).To(Equal(core.StringPtr("testString")))
				Expect(deleteAllowlistEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDatabaseUserOptions successfully`, func() {
				// Construct an instance of the DeleteDatabaseUserOptions model
				id := "testString"
				userType := "database"
				username := "user"
				deleteDatabaseUserOptionsModel := cloudDatabasesService.NewDeleteDatabaseUserOptions(id, userType, username)
				deleteDatabaseUserOptionsModel.SetID("testString")
				deleteDatabaseUserOptionsModel.SetUserType("database")
				deleteDatabaseUserOptionsModel.SetUsername("user")
				deleteDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(deleteDatabaseUserOptionsModel.Username).To(Equal(core.StringPtr("user")))
				Expect(deleteDatabaseUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLogicalReplicationSlotOptions successfully`, func() {
				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				id := "testString"
				name := "testString"
				deleteLogicalReplicationSlotOptionsModel := cloudDatabasesService.NewDeleteLogicalReplicationSlotOptions(id, name)
				deleteLogicalReplicationSlotOptionsModel.SetID("testString")
				deleteLogicalReplicationSlotOptionsModel.SetName("testString")
				deleteLogicalReplicationSlotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLogicalReplicationSlotOptionsModel).ToNot(BeNil())
				Expect(deleteLogicalReplicationSlotOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLogicalReplicationSlotOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deleteLogicalReplicationSlotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllowlistOptions successfully`, func() {
				// Construct an instance of the GetAllowlistOptions model
				id := "testString"
				getAllowlistOptionsModel := cloudDatabasesService.NewGetAllowlistOptions(id)
				getAllowlistOptionsModel.SetID("testString")
				getAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllowlistOptionsModel).ToNot(BeNil())
				Expect(getAllowlistOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAutoscalingConditionsOptions successfully`, func() {
				// Construct an instance of the GetAutoscalingConditionsOptions model
				id := "testString"
				groupID := "testString"
				getAutoscalingConditionsOptionsModel := cloudDatabasesService.NewGetAutoscalingConditionsOptions(id, groupID)
				getAutoscalingConditionsOptionsModel.SetID("testString")
				getAutoscalingConditionsOptionsModel.SetGroupID("testString")
				getAutoscalingConditionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAutoscalingConditionsOptionsModel).ToNot(BeNil())
				Expect(getAutoscalingConditionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAutoscalingConditionsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAutoscalingConditionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBackupInfoOptions successfully`, func() {
				// Construct an instance of the GetBackupInfoOptions model
				backupID := "testString"
				getBackupInfoOptionsModel := cloudDatabasesService.NewGetBackupInfoOptions(backupID)
				getBackupInfoOptionsModel.SetBackupID("testString")
				getBackupInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBackupInfoOptionsModel).ToNot(BeNil())
				Expect(getBackupInfoOptionsModel.BackupID).To(Equal(core.StringPtr("testString")))
				Expect(getBackupInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConnectionOptions successfully`, func() {
				// Construct an instance of the GetConnectionOptions model
				id := "testString"
				userType := "database"
				userID := "testString"
				endpointType := "public"
				getConnectionOptionsModel := cloudDatabasesService.NewGetConnectionOptions(id, userType, userID, endpointType)
				getConnectionOptionsModel.SetID("testString")
				getConnectionOptionsModel.SetUserType("database")
				getConnectionOptionsModel.SetUserID("testString")
				getConnectionOptionsModel.SetEndpointType("public")
				getConnectionOptionsModel.SetCertificateRoot("testString")
				getConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConnectionOptionsModel).ToNot(BeNil())
				Expect(getConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(getConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(getConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDefaultScalingGroupsOptions successfully`, func() {
				// Construct an instance of the GetDefaultScalingGroupsOptions model
				typeVar := "postgresql"
				getDefaultScalingGroupsOptionsModel := cloudDatabasesService.NewGetDefaultScalingGroupsOptions(typeVar)
				getDefaultScalingGroupsOptionsModel.SetType("postgresql")
				getDefaultScalingGroupsOptionsModel.SetHostFlavor("multitenant")
				getDefaultScalingGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDefaultScalingGroupsOptionsModel).ToNot(BeNil())
				Expect(getDefaultScalingGroupsOptionsModel.Type).To(Equal(core.StringPtr("postgresql")))
				Expect(getDefaultScalingGroupsOptionsModel.HostFlavor).To(Equal(core.StringPtr("multitenant")))
				Expect(getDefaultScalingGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentCapabilityOptions successfully`, func() {
				// Construct an instance of the GetDeploymentCapabilityOptions model
				id := "testString"
				capabilityID := "autoscaling"
				getDeploymentCapabilityOptionsModel := cloudDatabasesService.NewGetDeploymentCapabilityOptions(id, capabilityID)
				getDeploymentCapabilityOptionsModel.SetID("testString")
				getDeploymentCapabilityOptionsModel.SetCapabilityID("autoscaling")
				getDeploymentCapabilityOptionsModel.SetTargetPlatform("target_platform=classic")
				getDeploymentCapabilityOptionsModel.SetTargetLocation("target_location=us-east")
				getDeploymentCapabilityOptionsModel.SetHostFlavor("host_flavor=multitenant")
				getDeploymentCapabilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentCapabilityOptionsModel).ToNot(BeNil())
				Expect(getDeploymentCapabilityOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentCapabilityOptionsModel.CapabilityID).To(Equal(core.StringPtr("autoscaling")))
				Expect(getDeploymentCapabilityOptionsModel.TargetPlatform).To(Equal(core.StringPtr("target_platform=classic")))
				Expect(getDeploymentCapabilityOptionsModel.TargetLocation).To(Equal(core.StringPtr("target_location=us-east")))
				Expect(getDeploymentCapabilityOptionsModel.HostFlavor).To(Equal(core.StringPtr("host_flavor=multitenant")))
				Expect(getDeploymentCapabilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentInfoOptions successfully`, func() {
				// Construct an instance of the GetDeploymentInfoOptions model
				id := "testString"
				getDeploymentInfoOptionsModel := cloudDatabasesService.NewGetDeploymentInfoOptions(id)
				getDeploymentInfoOptionsModel.SetID("testString")
				getDeploymentInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentInfoOptionsModel).ToNot(BeNil())
				Expect(getDeploymentInfoOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPitrDataOptions successfully`, func() {
				// Construct an instance of the GetPitrDataOptions model
				id := "testString"
				getPitrDataOptionsModel := cloudDatabasesService.NewGetPitrDataOptions(id)
				getPitrDataOptionsModel.SetID("testString")
				getPitrDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPitrDataOptionsModel).ToNot(BeNil())
				Expect(getPitrDataOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPitrDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTaskOptions successfully`, func() {
				// Construct an instance of the GetTaskOptions model
				id := "testString"
				getTaskOptionsModel := cloudDatabasesService.NewGetTaskOptions(id)
				getTaskOptionsModel.SetID("testString")
				getTaskOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTaskOptionsModel).ToNot(BeNil())
				Expect(getTaskOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTaskOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewKillConnectionsOptions successfully`, func() {
				// Construct an instance of the KillConnectionsOptions model
				id := "testString"
				killConnectionsOptionsModel := cloudDatabasesService.NewKillConnectionsOptions(id)
				killConnectionsOptionsModel.SetID("testString")
				killConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(killConnectionsOptionsModel).ToNot(BeNil())
				Expect(killConnectionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(killConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDeployablesOptions successfully`, func() {
				// Construct an instance of the ListDeployablesOptions model
				listDeployablesOptionsModel := cloudDatabasesService.NewListDeployablesOptions()
				listDeployablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDeployablesOptionsModel).ToNot(BeNil())
				Expect(listDeployablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDeploymentBackupsOptions successfully`, func() {
				// Construct an instance of the ListDeploymentBackupsOptions model
				id := "testString"
				listDeploymentBackupsOptionsModel := cloudDatabasesService.NewListDeploymentBackupsOptions(id)
				listDeploymentBackupsOptionsModel.SetID("testString")
				listDeploymentBackupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDeploymentBackupsOptionsModel).ToNot(BeNil())
				Expect(listDeploymentBackupsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listDeploymentBackupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDeploymentScalingGroupsOptions successfully`, func() {
				// Construct an instance of the ListDeploymentScalingGroupsOptions model
				id := "testString"
				listDeploymentScalingGroupsOptionsModel := cloudDatabasesService.NewListDeploymentScalingGroupsOptions(id)
				listDeploymentScalingGroupsOptionsModel.SetID("testString")
				listDeploymentScalingGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDeploymentScalingGroupsOptionsModel).ToNot(BeNil())
				Expect(listDeploymentScalingGroupsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listDeploymentScalingGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDeploymentTasksOptions successfully`, func() {
				// Construct an instance of the ListDeploymentTasksOptions model
				id := "testString"
				listDeploymentTasksOptionsModel := cloudDatabasesService.NewListDeploymentTasksOptions(id)
				listDeploymentTasksOptionsModel.SetID("testString")
				listDeploymentTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDeploymentTasksOptionsModel).ToNot(BeNil())
				Expect(listDeploymentTasksOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listDeploymentTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRegionsOptions successfully`, func() {
				// Construct an instance of the ListRegionsOptions model
				listRegionsOptionsModel := cloudDatabasesService.NewListRegionsOptions()
				listRegionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRegionsOptionsModel).ToNot(BeNil())
				Expect(listRegionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRemotesOptions successfully`, func() {
				// Construct an instance of the ListRemotesOptions model
				id := "testString"
				listRemotesOptionsModel := cloudDatabasesService.NewListRemotesOptions(id)
				listRemotesOptionsModel.SetID("testString")
				listRemotesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRemotesOptionsModel).ToNot(BeNil())
				Expect(listRemotesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listRemotesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLogicalReplicationSlot successfully`, func() {
				name := "customer_replication"
				databaseName := "customers"
				pluginType := "wal2json"
				_model, err := cloudDatabasesService.NewLogicalReplicationSlot(name, databaseName, pluginType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPromoteReadOnlyReplicaOptions successfully`, func() {
				// Construct an instance of the PromoteReadOnlyReplicaOptions model
				id := "testString"
				promoteReadOnlyReplicaOptionsModel := cloudDatabasesService.NewPromoteReadOnlyReplicaOptions(id)
				promoteReadOnlyReplicaOptionsModel.SetID("testString")
				promoteReadOnlyReplicaOptionsModel.SetPromotion(map[string]interface{}{"anyKey": "anyValue"})
				promoteReadOnlyReplicaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(promoteReadOnlyReplicaOptionsModel).ToNot(BeNil())
				Expect(promoteReadOnlyReplicaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(promoteReadOnlyReplicaOptionsModel.Promotion).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(promoteReadOnlyReplicaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResyncReplicaOptions successfully`, func() {
				// Construct an instance of the ResyncReplicaOptions model
				id := "testString"
				resyncReplicaOptionsModel := cloudDatabasesService.NewResyncReplicaOptions(id)
				resyncReplicaOptionsModel.SetID("testString")
				resyncReplicaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(resyncReplicaOptionsModel).ToNot(BeNil())
				Expect(resyncReplicaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(resyncReplicaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetAllowlistOptions successfully`, func() {
				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				Expect(allowlistEntryModel).ToNot(BeNil())
				allowlistEntryModel.Address = core.StringPtr("195.212.0.0/16")
				allowlistEntryModel.Description = core.StringPtr("Dev IP space 1")
				Expect(allowlistEntryModel.Address).To(Equal(core.StringPtr("195.212.0.0/16")))
				Expect(allowlistEntryModel.Description).To(Equal(core.StringPtr("Dev IP space 1")))

				// Construct an instance of the SetAllowlistOptions model
				id := "testString"
				setAllowlistOptionsModel := cloudDatabasesService.NewSetAllowlistOptions(id)
				setAllowlistOptionsModel.SetID("testString")
				setAllowlistOptionsModel.SetIPAddresses([]clouddatabasesv5.AllowlistEntry{*allowlistEntryModel})
				setAllowlistOptionsModel.SetIfMatch("testString")
				setAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setAllowlistOptionsModel).ToNot(BeNil())
				Expect(setAllowlistOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setAllowlistOptionsModel.IPAddresses).To(Equal([]clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}))
				Expect(setAllowlistOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(setAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetAutoscalingConditionsOptions successfully`, func() {
				// Construct an instance of the AutoscalingMemoryGroupMemoryScalersIoUtilization model
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
				Expect(autoscalingMemoryGroupMemoryScalersIoUtilizationModel).ToNot(BeNil())
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod = core.StringPtr("5m")
				autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(90))
				Expect(autoscalingMemoryGroupMemoryScalersIoUtilizationModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(autoscalingMemoryGroupMemoryScalersIoUtilizationModel.OverPeriod).To(Equal(core.StringPtr("5m")))
				Expect(autoscalingMemoryGroupMemoryScalersIoUtilizationModel.AbovePercent).To(Equal(core.Int64Ptr(int64(90))))

				// Construct an instance of the AutoscalingMemoryGroupMemoryScalers model
				autoscalingMemoryGroupMemoryScalersModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
				Expect(autoscalingMemoryGroupMemoryScalersModel).ToNot(BeNil())
				autoscalingMemoryGroupMemoryScalersModel.IoUtilization = autoscalingMemoryGroupMemoryScalersIoUtilizationModel
				Expect(autoscalingMemoryGroupMemoryScalersModel.IoUtilization).To(Equal(autoscalingMemoryGroupMemoryScalersIoUtilizationModel))

				// Construct an instance of the AutoscalingMemoryGroupMemoryRate model
				autoscalingMemoryGroupMemoryRateModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
				Expect(autoscalingMemoryGroupMemoryRateModel).ToNot(BeNil())
				autoscalingMemoryGroupMemoryRateModel.IncreasePercent = core.Float64Ptr(float64(10))
				autoscalingMemoryGroupMemoryRateModel.PeriodSeconds = core.Int64Ptr(int64(300))
				autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember = core.Float64Ptr(float64(125952))
				autoscalingMemoryGroupMemoryRateModel.Units = core.StringPtr("mb")
				Expect(autoscalingMemoryGroupMemoryRateModel.IncreasePercent).To(Equal(core.Float64Ptr(float64(10))))
				Expect(autoscalingMemoryGroupMemoryRateModel.PeriodSeconds).To(Equal(core.Int64Ptr(int64(300))))
				Expect(autoscalingMemoryGroupMemoryRateModel.LimitMbPerMember).To(Equal(core.Float64Ptr(float64(125952))))
				Expect(autoscalingMemoryGroupMemoryRateModel.Units).To(Equal(core.StringPtr("mb")))

				// Construct an instance of the AutoscalingMemoryGroupMemory model
				autoscalingMemoryGroupMemoryModel := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
				Expect(autoscalingMemoryGroupMemoryModel).ToNot(BeNil())
				autoscalingMemoryGroupMemoryModel.Scalers = autoscalingMemoryGroupMemoryScalersModel
				autoscalingMemoryGroupMemoryModel.Rate = autoscalingMemoryGroupMemoryRateModel
				Expect(autoscalingMemoryGroupMemoryModel.Scalers).To(Equal(autoscalingMemoryGroupMemoryScalersModel))
				Expect(autoscalingMemoryGroupMemoryModel.Rate).To(Equal(autoscalingMemoryGroupMemoryRateModel))

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
				Expect(autoscalingSetGroupAutoscalingModel).ToNot(BeNil())
				autoscalingSetGroupAutoscalingModel.Memory = autoscalingMemoryGroupMemoryModel
				Expect(autoscalingSetGroupAutoscalingModel.Memory).To(Equal(autoscalingMemoryGroupMemoryModel))

				// Construct an instance of the SetAutoscalingConditionsOptions model
				id := "testString"
				groupID := "testString"
				var setAutoscalingConditionsOptionsAutoscaling clouddatabasesv5.AutoscalingSetGroupAutoscalingIntf = nil
				setAutoscalingConditionsOptionsModel := cloudDatabasesService.NewSetAutoscalingConditionsOptions(id, groupID, setAutoscalingConditionsOptionsAutoscaling)
				setAutoscalingConditionsOptionsModel.SetID("testString")
				setAutoscalingConditionsOptionsModel.SetGroupID("testString")
				setAutoscalingConditionsOptionsModel.SetAutoscaling(autoscalingSetGroupAutoscalingModel)
				setAutoscalingConditionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setAutoscalingConditionsOptionsModel).ToNot(BeNil())
				Expect(setAutoscalingConditionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setAutoscalingConditionsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(setAutoscalingConditionsOptionsModel.Autoscaling).To(Equal(autoscalingSetGroupAutoscalingModel))
				Expect(setAutoscalingConditionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetDatabaseInplaceVersionUpgradeOptions successfully`, func() {
				// Construct an instance of the SetDatabaseInplaceVersionUpgradeOptions model
				id := "testString"
				setDatabaseInplaceVersionUpgradeOptionsVersion := "7"
				setDatabaseInplaceVersionUpgradeOptionsModel := cloudDatabasesService.NewSetDatabaseInplaceVersionUpgradeOptions(id, setDatabaseInplaceVersionUpgradeOptionsVersion)
				setDatabaseInplaceVersionUpgradeOptionsModel.SetID("testString")
				setDatabaseInplaceVersionUpgradeOptionsModel.SetVersion("7")
				setDatabaseInplaceVersionUpgradeOptionsModel.SetSkipBackup(true)
				setDatabaseInplaceVersionUpgradeOptionsModel.SetExpirationDatetime(CreateMockDateTime("2025-10-05T17:17:17Z"))
				setDatabaseInplaceVersionUpgradeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel).ToNot(BeNil())
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel.Version).To(Equal(core.StringPtr("7")))
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel.SkipBackup).To(Equal(core.BoolPtr(true)))
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel.ExpirationDatetime).To(Equal(CreateMockDateTime("2025-10-05T17:17:17Z")))
				Expect(setDatabaseInplaceVersionUpgradeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetDeploymentScalingGroupOptions successfully`, func() {
				// Construct an instance of the GroupScalingMembers model
				groupScalingMembersModel := new(clouddatabasesv5.GroupScalingMembers)
				Expect(groupScalingMembersModel).ToNot(BeNil())
				groupScalingMembersModel.AllocationCount = core.Int64Ptr(int64(4))
				Expect(groupScalingMembersModel.AllocationCount).To(Equal(core.Int64Ptr(int64(4))))

				// Construct an instance of the GroupScalingMemory model
				groupScalingMemoryModel := new(clouddatabasesv5.GroupScalingMemory)
				Expect(groupScalingMemoryModel).ToNot(BeNil())
				groupScalingMemoryModel.AllocationMb = core.Int64Ptr(int64(12288))
				Expect(groupScalingMemoryModel.AllocationMb).To(Equal(core.Int64Ptr(int64(12288))))

				// Construct an instance of the GroupScalingCPU model
				groupScalingCPUModel := new(clouddatabasesv5.GroupScalingCPU)
				Expect(groupScalingCPUModel).ToNot(BeNil())
				groupScalingCPUModel.AllocationCount = core.Int64Ptr(int64(2))
				Expect(groupScalingCPUModel.AllocationCount).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the GroupScalingDisk model
				groupScalingDiskModel := new(clouddatabasesv5.GroupScalingDisk)
				Expect(groupScalingDiskModel).ToNot(BeNil())
				groupScalingDiskModel.AllocationMb = core.Int64Ptr(int64(20480))
				Expect(groupScalingDiskModel.AllocationMb).To(Equal(core.Int64Ptr(int64(20480))))

				// Construct an instance of the GroupScalingHostFlavor model
				groupScalingHostFlavorModel := new(clouddatabasesv5.GroupScalingHostFlavor)
				Expect(groupScalingHostFlavorModel).ToNot(BeNil())
				groupScalingHostFlavorModel.ID = core.StringPtr("b3c.16x64.encrypted")
				Expect(groupScalingHostFlavorModel.ID).To(Equal(core.StringPtr("b3c.16x64.encrypted")))

				// Construct an instance of the GroupScaling model
				groupScalingModel := new(clouddatabasesv5.GroupScaling)
				Expect(groupScalingModel).ToNot(BeNil())
				groupScalingModel.Members = groupScalingMembersModel
				groupScalingModel.Memory = groupScalingMemoryModel
				groupScalingModel.CPU = groupScalingCPUModel
				groupScalingModel.Disk = groupScalingDiskModel
				groupScalingModel.HostFlavor = groupScalingHostFlavorModel
				Expect(groupScalingModel.Members).To(Equal(groupScalingMembersModel))
				Expect(groupScalingModel.Memory).To(Equal(groupScalingMemoryModel))
				Expect(groupScalingModel.CPU).To(Equal(groupScalingCPUModel))
				Expect(groupScalingModel.Disk).To(Equal(groupScalingDiskModel))
				Expect(groupScalingModel.HostFlavor).To(Equal(groupScalingHostFlavorModel))

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				id := "testString"
				groupID := "testString"
				setDeploymentScalingGroupOptionsModel := cloudDatabasesService.NewSetDeploymentScalingGroupOptions(id, groupID)
				setDeploymentScalingGroupOptionsModel.SetID("testString")
				setDeploymentScalingGroupOptionsModel.SetGroupID("testString")
				setDeploymentScalingGroupOptionsModel.SetGroup(groupScalingModel)
				setDeploymentScalingGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDeploymentScalingGroupOptionsModel).ToNot(BeNil())
				Expect(setDeploymentScalingGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.Group).To(Equal(groupScalingModel))
				Expect(setDeploymentScalingGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewStartOndemandBackupOptions successfully`, func() {
				// Construct an instance of the StartOndemandBackupOptions model
				id := "testString"
				startOndemandBackupOptionsModel := cloudDatabasesService.NewStartOndemandBackupOptions(id)
				startOndemandBackupOptionsModel.SetID("testString")
				startOndemandBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(startOndemandBackupOptionsModel).ToNot(BeNil())
				Expect(startOndemandBackupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(startOndemandBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDatabaseConfigurationOptions successfully`, func() {
				// Construct an instance of the ConfigurationPgConfiguration model
				configurationModel := new(clouddatabasesv5.ConfigurationPgConfiguration)
				Expect(configurationModel).ToNot(BeNil())
				configurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				configurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				configurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				configurationModel.LogConnections = core.StringPtr("off")
				configurationModel.LogDisconnections = core.StringPtr("off")
				configurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				configurationModel.MaxConnections = core.Int64Ptr(int64(200))
				configurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
				configurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				configurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				configurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				configurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				configurationModel.SynchronousCommit = core.StringPtr("local")
				configurationModel.TCPKeepalivesCount = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
				configurationModel.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
				configurationModel.WalLevel = core.StringPtr("replica")
				Expect(configurationModel.ArchiveTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(configurationModel.DeadlockTimeout).To(Equal(core.Int64Ptr(int64(100))))
				Expect(configurationModel.EffectiveIoConcurrency).To(Equal(core.Int64Ptr(int64(1))))
				Expect(configurationModel.LogConnections).To(Equal(core.StringPtr("off")))
				Expect(configurationModel.LogDisconnections).To(Equal(core.StringPtr("off")))
				Expect(configurationModel.LogMinDurationStatement).To(Equal(core.Int64Ptr(int64(100))))
				Expect(configurationModel.MaxConnections).To(Equal(core.Int64Ptr(int64(200))))
				Expect(configurationModel.MaxLocksPerTransaction).To(Equal(core.Int64Ptr(int64(10))))
				Expect(configurationModel.MaxPreparedTransactions).To(Equal(core.Int64Ptr(int64(0))))
				Expect(configurationModel.MaxReplicationSlots).To(Equal(core.Int64Ptr(int64(10))))
				Expect(configurationModel.MaxWalSenders).To(Equal(core.Int64Ptr(int64(12))))
				Expect(configurationModel.SharedBuffers).To(Equal(core.Int64Ptr(int64(16))))
				Expect(configurationModel.SynchronousCommit).To(Equal(core.StringPtr("local")))
				Expect(configurationModel.TCPKeepalivesCount).To(Equal(core.Int64Ptr(int64(0))))
				Expect(configurationModel.TCPKeepalivesIdle).To(Equal(core.Int64Ptr(int64(0))))
				Expect(configurationModel.TCPKeepalivesInterval).To(Equal(core.Int64Ptr(int64(0))))
				Expect(configurationModel.WalLevel).To(Equal(core.StringPtr("replica")))

				// Construct an instance of the UpdateDatabaseConfigurationOptions model
				id := "testString"
				updateDatabaseConfigurationOptionsModel := cloudDatabasesService.NewUpdateDatabaseConfigurationOptions(id)
				updateDatabaseConfigurationOptionsModel.SetID("testString")
				updateDatabaseConfigurationOptionsModel.SetConfiguration(configurationModel)
				updateDatabaseConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDatabaseConfigurationOptionsModel).ToNot(BeNil())
				Expect(updateDatabaseConfigurationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDatabaseConfigurationOptionsModel.Configuration).To(Equal(configurationModel))
				Expect(updateDatabaseConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserOptions successfully`, func() {
				// Construct an instance of the UserUpdatePasswordSetting model
				userUpdateModel := new(clouddatabasesv5.UserUpdatePasswordSetting)
				Expect(userUpdateModel).ToNot(BeNil())
				userUpdateModel.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
				Expect(userUpdateModel.Password).To(Equal(core.StringPtr("v3ry-1-secUre-pAssword-2")))

				// Construct an instance of the UpdateUserOptions model
				id := "testString"
				userType := "database"
				username := "user"
				updateUserOptionsModel := cloudDatabasesService.NewUpdateUserOptions(id, userType, username)
				updateUserOptionsModel.SetID("testString")
				updateUserOptionsModel.SetUserType("database")
				updateUserOptionsModel.SetUsername("user")
				updateUserOptionsModel.SetUser(userUpdateModel)
				updateUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserOptionsModel).ToNot(BeNil())
				Expect(updateUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserOptionsModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(updateUserOptionsModel.Username).To(Equal(core.StringPtr("user")))
				Expect(updateUserOptionsModel.User).To(Equal(userUpdateModel))
				Expect(updateUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUserUpdatePasswordSetting successfully`, func() {
				password := "v3ry-1-secUre-pAssword-2"
				_model, err := cloudDatabasesService.NewUserUpdatePasswordSetting(password)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUserUpdateRedisRoleSetting successfully`, func() {
				role := "-@all -@read"
				_model, err := cloudDatabasesService.NewUserUpdateRedisRoleSetting(role)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUserDatabaseUser successfully`, func() {
				username := "user"
				password := "v3ry-1-secUre-pAssword-2"
				_model, err := cloudDatabasesService.NewUserDatabaseUser(username, password)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUserOpsManagerUser successfully`, func() {
				username := "user"
				password := "v3ry-1-secUre-pAssword-2"
				_model, err := cloudDatabasesService.NewUserOpsManagerUser(username, password)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUserRedisDatabaseUser successfully`, func() {
				username := "user"
				password := "v3ry-1-secUre-pAssword-2"
				_model, err := cloudDatabasesService.NewUserRedisDatabaseUser(username, password)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAllowlistEntry successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AllowlistEntry)
			model.Address = core.StringPtr("testString")
			model.Description = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AllowlistEntry
			err = clouddatabasesv5.UnmarshalAllowlistEntry(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingCPUGroupCPU successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingCPUGroupCPU)
			model.Scalers = map[string]interface{}{"anyKey": "anyValue"}
			model.Rate = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingCPUGroupCPU
			err = clouddatabasesv5.UnmarshalAutoscalingCPUGroupCPU(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingCPUGroupCPURate successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingCPUGroupCPURate)
			model.IncreasePercent = core.Float64Ptr(float64(10))
			model.PeriodSeconds = core.Int64Ptr(int64(900))
			model.LimitCountPerMember = core.Int64Ptr(int64(10))
			model.Units = core.StringPtr("count")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingCPUGroupCPURate
			err = clouddatabasesv5.UnmarshalAutoscalingCPUGroupCPURate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingDiskGroupDisk successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
			model.Scalers = nil
			model.Rate = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingDiskGroupDisk
			err = clouddatabasesv5.UnmarshalAutoscalingDiskGroupDisk(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingDiskGroupDiskRate successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
			model.IncreasePercent = core.Float64Ptr(float64(20))
			model.PeriodSeconds = core.Int64Ptr(int64(900))
			model.LimitMbPerMember = core.Float64Ptr(float64(3670016))
			model.Units = core.StringPtr("mb")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingDiskGroupDiskRate
			err = clouddatabasesv5.UnmarshalAutoscalingDiskGroupDiskRate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingDiskGroupDiskScalers successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
			model.Capacity = nil
			model.IoUtilization = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingDiskGroupDiskScalers
			err = clouddatabasesv5.UnmarshalAutoscalingDiskGroupDiskScalers(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingDiskGroupDiskScalersCapacity successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
			model.Enabled = core.BoolPtr(true)
			model.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity
			err = clouddatabasesv5.UnmarshalAutoscalingDiskGroupDiskScalersCapacity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
			model.Enabled = core.BoolPtr(true)
			model.OverPeriod = core.StringPtr("30m")
			model.AbovePercent = core.Int64Ptr(int64(45))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization
			err = clouddatabasesv5.UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingMemoryGroupMemory successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingMemoryGroupMemory)
			model.Scalers = nil
			model.Rate = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingMemoryGroupMemory
			err = clouddatabasesv5.UnmarshalAutoscalingMemoryGroupMemory(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingMemoryGroupMemoryRate successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryRate)
			model.IncreasePercent = core.Float64Ptr(float64(10))
			model.PeriodSeconds = core.Int64Ptr(int64(900))
			model.LimitMbPerMember = core.Float64Ptr(float64(3670016))
			model.Units = core.StringPtr("mb")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingMemoryGroupMemoryRate
			err = clouddatabasesv5.UnmarshalAutoscalingMemoryGroupMemoryRate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingMemoryGroupMemoryScalers successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers)
			model.IoUtilization = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers
			err = clouddatabasesv5.UnmarshalAutoscalingMemoryGroupMemoryScalers(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization)
			model.Enabled = core.BoolPtr(true)
			model.OverPeriod = core.StringPtr("30m")
			model.AbovePercent = core.Int64Ptr(int64(45))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization
			err = clouddatabasesv5.UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingSetGroupAutoscaling successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingSetGroupAutoscaling)
			model.Disk = nil
			model.Memory = nil
			model.CPU = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingSetGroupAutoscaling
			err = clouddatabasesv5.UnmarshalAutoscalingSetGroupAutoscaling(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConfiguration successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.Configuration)
			model.ArchiveTimeout = core.Int64Ptr(int64(300))
			model.DeadlockTimeout = core.Int64Ptr(int64(100))
			model.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
			model.LogConnections = core.StringPtr("off")
			model.LogDisconnections = core.StringPtr("off")
			model.LogMinDurationStatement = core.Int64Ptr(int64(100))
			model.MaxConnections = core.Int64Ptr(int64(115))
			model.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
			model.MaxPreparedTransactions = core.Int64Ptr(int64(0))
			model.MaxReplicationSlots = core.Int64Ptr(int64(10))
			model.MaxWalSenders = core.Int64Ptr(int64(12))
			model.SharedBuffers = core.Int64Ptr(int64(16))
			model.SynchronousCommit = core.StringPtr("local")
			model.TCPKeepalivesCount = core.Int64Ptr(int64(0))
			model.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
			model.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
			model.WalLevel = core.StringPtr("replica")
			model.Maxmemory = core.Int64Ptr(int64(0))
			model.MaxmemoryPolicy = core.StringPtr("volatile-lru")
			model.Appendonly = core.StringPtr("yes")
			model.MaxmemorySamples = core.Int64Ptr(int64(0))
			model.StopWritesOnBgsaveError = core.StringPtr("yes")
			model.DeleteUndefinedQueues = core.BoolPtr(true)
			model.DefaultAuthenticationPlugin = core.StringPtr("sha256_password")
			model.InnodbBufferPoolSizePercentage = core.Int64Ptr(int64(10))
			model.InnodbFlushLogAtTrxCommit = core.Int64Ptr(int64(0))
			model.InnodbLogBufferSize = core.Int64Ptr(int64(1048576))
			model.InnodbLogFileSize = core.Int64Ptr(int64(4194304))
			model.InnodbLruScanDepth = core.Int64Ptr(int64(128))
			model.InnodbReadIoThreads = core.Int64Ptr(int64(1))
			model.InnodbWriteIoThreads = core.Int64Ptr(int64(1))
			model.MaxAllowedPacket = core.Int64Ptr(int64(1024))
			model.MysqlMaxBinlogAgeSec = core.Int64Ptr(int64(1800))
			model.NetReadTimeout = core.Int64Ptr(int64(1))
			model.NetWriteTimeout = core.Int64Ptr(int64(1))
			model.SQLMode = core.StringPtr("testString")
			model.WaitTimeout = core.Int64Ptr(int64(1))
			model.MaxPreparedStmtCount = core.Int64Ptr(int64(0))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.Configuration
			err = clouddatabasesv5.UnmarshalConfiguration(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateCapabilityRequestBackup successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.CreateCapabilityRequestBackup)
			model.Type = core.StringPtr("PostgreSQL")
			model.Version = core.StringPtr("10")
			model.Platform = core.StringPtr("satellite")
			model.Location = core.StringPtr("us-south")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.CreateCapabilityRequestBackup
			err = clouddatabasesv5.UnmarshalCreateCapabilityRequestBackup(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateCapabilityRequestDeployment successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.CreateCapabilityRequestDeployment)
			model.Type = core.StringPtr("PostgreSQL")
			model.Version = core.StringPtr("10")
			model.Platform = core.StringPtr("satellite")
			model.Location = core.StringPtr("us-south")
			model.Plan = core.StringPtr("standard")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.CreateCapabilityRequestDeployment
			err = clouddatabasesv5.UnmarshalCreateCapabilityRequestDeployment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateCapabilityRequestOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.CreateCapabilityRequestOptions)
			model.TargetPlatform = core.StringPtr("classic")
			model.TargetLocation = core.StringPtr("us-east")
			model.HostFlavor = core.StringPtr("multitenant")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.CreateCapabilityRequestOptions
			err = clouddatabasesv5.UnmarshalCreateCapabilityRequestOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScaling successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScaling)
			model.Members = nil
			model.Memory = nil
			model.CPU = nil
			model.Disk = nil
			model.HostFlavor = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScaling
			err = clouddatabasesv5.UnmarshalGroupScaling(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScalingCPU successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScalingCPU)
			model.AllocationCount = core.Int64Ptr(int64(2))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScalingCPU
			err = clouddatabasesv5.UnmarshalGroupScalingCPU(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScalingDisk successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScalingDisk)
			model.AllocationMb = core.Int64Ptr(int64(20480))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScalingDisk
			err = clouddatabasesv5.UnmarshalGroupScalingDisk(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScalingHostFlavor successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScalingHostFlavor)
			model.ID = core.StringPtr("b3c.16x64.encrypted")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScalingHostFlavor
			err = clouddatabasesv5.UnmarshalGroupScalingHostFlavor(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScalingMembers successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScalingMembers)
			model.AllocationCount = core.Int64Ptr(int64(4))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScalingMembers
			err = clouddatabasesv5.UnmarshalGroupScalingMembers(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalGroupScalingMemory successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.GroupScalingMemory)
			model.AllocationMb = core.Int64Ptr(int64(12288))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.GroupScalingMemory
			err = clouddatabasesv5.UnmarshalGroupScalingMemory(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalLogicalReplicationSlot successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.LogicalReplicationSlot)
			model.Name = core.StringPtr("customer_replication")
			model.DatabaseName = core.StringPtr("customers")
			model.PluginType = core.StringPtr("wal2json")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.LogicalReplicationSlot
			err = clouddatabasesv5.UnmarshalLogicalReplicationSlot(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUser successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.User)
			model.Username = core.StringPtr("user")
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
			model.Role = core.StringPtr("-@all -@read")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.User
			err = clouddatabasesv5.UnmarshalUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserUpdate successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserUpdate)
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
			model.Role = core.StringPtr("-@all -@read")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserUpdate
			err = clouddatabasesv5.UnmarshalUserUpdate(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingCPUGroup)
			model.CPU = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingCPUGroup
			err = clouddatabasesv5.UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
			model.Disk = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup
			err = clouddatabasesv5.UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
			model.Memory = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup
			err = clouddatabasesv5.UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConfigurationMySQLConfiguration successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.ConfigurationMySQLConfiguration)
			model.DefaultAuthenticationPlugin = core.StringPtr("sha256_password")
			model.InnodbBufferPoolSizePercentage = core.Int64Ptr(int64(10))
			model.InnodbFlushLogAtTrxCommit = core.Int64Ptr(int64(0))
			model.InnodbLogBufferSize = core.Int64Ptr(int64(1048576))
			model.InnodbLogFileSize = core.Int64Ptr(int64(4194304))
			model.InnodbLruScanDepth = core.Int64Ptr(int64(128))
			model.InnodbReadIoThreads = core.Int64Ptr(int64(1))
			model.InnodbWriteIoThreads = core.Int64Ptr(int64(1))
			model.MaxAllowedPacket = core.Int64Ptr(int64(1024))
			model.MaxConnections = core.Int64Ptr(int64(100))
			model.MysqlMaxBinlogAgeSec = core.Int64Ptr(int64(1800))
			model.NetReadTimeout = core.Int64Ptr(int64(1))
			model.NetWriteTimeout = core.Int64Ptr(int64(1))
			model.SQLMode = core.StringPtr("testString")
			model.WaitTimeout = core.Int64Ptr(int64(1))
			model.MaxPreparedStmtCount = core.Int64Ptr(int64(0))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.ConfigurationMySQLConfiguration
			err = clouddatabasesv5.UnmarshalConfigurationMySQLConfiguration(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConfigurationPgConfiguration successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.ConfigurationPgConfiguration)
			model.ArchiveTimeout = core.Int64Ptr(int64(300))
			model.DeadlockTimeout = core.Int64Ptr(int64(100))
			model.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
			model.LogConnections = core.StringPtr("off")
			model.LogDisconnections = core.StringPtr("off")
			model.LogMinDurationStatement = core.Int64Ptr(int64(100))
			model.MaxConnections = core.Int64Ptr(int64(115))
			model.MaxLocksPerTransaction = core.Int64Ptr(int64(10))
			model.MaxPreparedTransactions = core.Int64Ptr(int64(0))
			model.MaxReplicationSlots = core.Int64Ptr(int64(10))
			model.MaxWalSenders = core.Int64Ptr(int64(12))
			model.SharedBuffers = core.Int64Ptr(int64(16))
			model.SynchronousCommit = core.StringPtr("local")
			model.TCPKeepalivesCount = core.Int64Ptr(int64(0))
			model.TCPKeepalivesIdle = core.Int64Ptr(int64(0))
			model.TCPKeepalivesInterval = core.Int64Ptr(int64(0))
			model.WalLevel = core.StringPtr("replica")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.ConfigurationPgConfiguration
			err = clouddatabasesv5.UnmarshalConfigurationPgConfiguration(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConfigurationRabbitMqConfiguration successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.ConfigurationRabbitMqConfiguration)
			model.DeleteUndefinedQueues = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.ConfigurationRabbitMqConfiguration
			err = clouddatabasesv5.UnmarshalConfigurationRabbitMqConfiguration(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalConfigurationRedisConfiguration successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.ConfigurationRedisConfiguration)
			model.Maxmemory = core.Int64Ptr(int64(0))
			model.MaxmemoryPolicy = core.StringPtr("volatile-lru")
			model.Appendonly = core.StringPtr("yes")
			model.MaxmemorySamples = core.Int64Ptr(int64(0))
			model.StopWritesOnBgsaveError = core.StringPtr("yes")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.ConfigurationRedisConfiguration
			err = clouddatabasesv5.UnmarshalConfigurationRedisConfiguration(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserUpdatePasswordSetting successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserUpdatePasswordSetting)
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserUpdatePasswordSetting
			err = clouddatabasesv5.UnmarshalUserUpdatePasswordSetting(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserUpdateRedisRoleSetting successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserUpdateRedisRoleSetting)
			model.Role = core.StringPtr("-@all -@read")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserUpdateRedisRoleSetting
			err = clouddatabasesv5.UnmarshalUserUpdateRedisRoleSetting(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserDatabaseUser successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserDatabaseUser)
			model.Username = core.StringPtr("user")
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserDatabaseUser
			err = clouddatabasesv5.UnmarshalUserDatabaseUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserOpsManagerUser successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserOpsManagerUser)
			model.Username = core.StringPtr("user")
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
			model.Role = core.StringPtr("group_data_access_admin")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserOpsManagerUser
			err = clouddatabasesv5.UnmarshalUserOpsManagerUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUserRedisDatabaseUser successfully`, func() {
			// Construct an instance of the model.
			model := new(clouddatabasesv5.UserRedisDatabaseUser)
			model.Username = core.StringPtr("user")
			model.Password = core.StringPtr("v3ry-1-secUre-pAssword-2")
			model.Role = core.StringPtr("-@all -@read")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *clouddatabasesv5.UserRedisDatabaseUser
			err = clouddatabasesv5.UnmarshalUserRedisDatabaseUser(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
