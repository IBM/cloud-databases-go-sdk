/**
 * (C) Copyright IBM Corp. 2021.
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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/icd-go-sdk/clouddatabasesv5"
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`ListDeployables(listDeployablesOptions *ListDeployablesOptions) - Operation response error`, func() {
		listDeployablesPath := "/deployables"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeployablesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
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
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
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
	})
	Describe(`ListRegions(listRegionsOptions *ListRegionsOptions) - Operation response error`, func() {
		listRegionsPath := "/regions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRegionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
	})
	Describe(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions) - Operation response error`, func() {
		getDeploymentInfoPath := "/deployments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": {"mapKey": "Inner"}, "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
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
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": {"mapKey": "Inner"}, "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions) - Operation response error`, func() {
		createDatabaseUserPath := "/deployments/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDatabaseUser with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(clouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(clouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(clouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
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

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(clouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(clouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
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
	})
	Describe(`ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions) - Operation response error`, func() {
		changeUserPasswordPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeUserPassword with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(clouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(clouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions)`, func() {
		changeUserPasswordPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ChangeUserPassword successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(clouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(clouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ChangeUserPasswordWithContext(ctx, changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ChangeUserPasswordWithContext(ctx, changeUserPasswordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ChangeUserPassword successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ChangeUserPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(clouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(clouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ChangeUserPassword with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(clouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(clouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ChangeUserPasswordOptions model with no property values
				changeUserPasswordOptionsModelNew := new(clouddatabasesv5.ChangeUserPasswordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions) - Operation response error`, func() {
		deleteDatabaseUserPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
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
		deleteDatabaseUserPath := "/deployments/testString/users/testString/testString"
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
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
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
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
	})
	Describe(`GetUser(getUserOptions *GetUserOptions) - Operation response error`, func() {
		getUserPath := "/deployments/testString/users/testString/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUser with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(clouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUser(getUserOptions *GetUserOptions)`, func() {
		getUserPath := "/deployments/testString/users/testString/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetUser successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(clouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetUserWithContext(ctx, getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetUserWithContext(ctx, getUserOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetUser successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(clouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUser with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(clouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserOptions model with no property values
				getUserOptionsModelNew := new(clouddatabasesv5.GetUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetUser(getUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions) - Operation response error`, func() {
		setDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPgConfiguration model
				setConfigurationConfigurationModel := new(clouddatabasesv5.SetConfigurationConfigurationPgConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(clouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions)`, func() {
		setDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetConfigurationConfigurationPgConfiguration model
				setConfigurationConfigurationModel := new(clouddatabasesv5.SetConfigurationConfigurationPgConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(clouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetDatabaseConfigurationWithContext(ctx, setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetDatabaseConfigurationWithContext(ctx, setDatabaseConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetDatabaseConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPgConfiguration model
				setConfigurationConfigurationModel := new(clouddatabasesv5.SetConfigurationConfigurationPgConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(clouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetDatabaseConfiguration with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPgConfiguration model
				setConfigurationConfigurationModel := new(clouddatabasesv5.SetConfigurationConfigurationPgConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(clouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetDatabaseConfigurationOptions model with no property values
				setDatabaseConfigurationOptionsModelNew := new(clouddatabasesv5.SetDatabaseConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions) - Operation response error`, func() {
		getDatabaseConfigurationSchemaPath := "/deployments/testString/configuration/schema"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(clouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions)`, func() {
		getDatabaseConfigurationSchemaPath := "/deployments/testString/configuration/schema"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"max_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_prepared_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "backup_retention_period": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "deadlock_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "effective_io_concurrency": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_replication_slots": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_wal_senders": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "shared_buffers": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "synchronous_commit": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "wal_level": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "archive_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "log_min_duration_statement": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}}}`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(clouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetDatabaseConfigurationSchemaWithContext(ctx, getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetDatabaseConfigurationSchemaWithContext(ctx, getDatabaseConfigurationSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"max_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_prepared_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "backup_retention_period": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "deadlock_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "effective_io_concurrency": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_replication_slots": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_wal_senders": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "shared_buffers": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "synchronous_commit": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "wal_level": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "archive_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "log_min_duration_statement": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}}}`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetDatabaseConfigurationSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(clouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDatabaseConfigurationSchema with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(clouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDatabaseConfigurationSchemaOptions model with no property values
				getDatabaseConfigurationSchemaOptionsModelNew := new(clouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`ListRemotes(listRemotesOptions *ListRemotesOptions) - Operation response error`, func() {
		listRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRemotesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
	})
	Describe(`GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions) - Operation response error`, func() {
		getRemotesSchemaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRemotesSchema with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(clouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions)`, func() {
		getRemotesSchemaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetRemotesSchema successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(clouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetRemotesSchemaWithContext(ctx, getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetRemotesSchemaWithContext(ctx, getRemotesSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetRemotesSchema successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetRemotesSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(clouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRemotesSchema with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(clouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRemotesSchemaOptions model with no property values
				getRemotesSchemaOptionsModelNew := new(clouddatabasesv5.GetRemotesSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetPromotion(setPromotionOptions *SetPromotionOptions) - Operation response error`, func() {
		setPromotionPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetPromotion with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(clouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(clouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetPromotion(setPromotionOptions *SetPromotionOptions)`, func() {
		setPromotionPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetPromotion successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(clouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(clouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.SetPromotionWithContext(ctx, setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.SetPromotionWithContext(ctx, setPromotionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke SetPromotion successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.SetPromotion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(clouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(clouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetPromotion with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(clouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(clouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetPromotionOptions model with no property values
				setPromotionOptionsModelNew := new(clouddatabasesv5.SetPromotionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.SetPromotion(setPromotionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`ListDeploymentTasks(listDeploymentTasksOptions *ListDeploymentTasksOptions) - Operation response error`, func() {
		listDeploymentTasksPath := "/deployments/testString/tasks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}]}`)
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
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}]}`)
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
	})
	Describe(`GetTask(getTaskOptions *GetTaskOptions) - Operation response error`, func() {
		getTaskPath := "/tasks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions) - Operation response error`, func() {
		getBackupInfoPath := "/backups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2018-02-28T19:25:12.000Z"}}`)
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
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2018-02-28T19:25:12.000Z"}}`)
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
	})
	Describe(`ListDeploymentBackups(listDeploymentBackupsOptions *ListDeploymentBackupsOptions) - Operation response error`, func() {
		listDeploymentBackupsPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2018-02-28T19:25:12.000Z"}]}`)
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
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2018-02-28T19:25:12.000Z"}]}`)
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
	})
	Describe(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions) - Operation response error`, func() {
		startOndemandBackupPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
	})
	Describe(`GetPitRdata(getPitRdataOptions *GetPitRdataOptions) - Operation response error`, func() {
		getPitRdataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPitRdata with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPitRdataOptions model
				getPitRdataOptionsModel := new(clouddatabasesv5.GetPitRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.GetPitRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.GetPitRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPitRdata(getPitRdataOptions *GetPitRdataOptions)`, func() {
		getPitRdataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}`)
				}))
			})
			It(`Invoke GetPitRdata successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetPitRdataOptions model
				getPitRdataOptionsModel := new(clouddatabasesv5.GetPitRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.GetPitRdataWithContext(ctx, getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.GetPitRdata(getPitRdataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.GetPitRdataWithContext(ctx, getPitRdataOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}`)
				}))
			})
			It(`Invoke GetPitRdata successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.GetPitRdata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPitRdataOptions model
				getPitRdataOptionsModel := new(clouddatabasesv5.GetPitRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.GetPitRdata(getPitRdataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPitRdata with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPitRdataOptions model
				getPitRdataOptionsModel := new(clouddatabasesv5.GetPitRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.GetPitRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPitRdataOptions model with no property values
				getPitRdataOptionsModelNew := new(clouddatabasesv5.GetPitRdataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.GetPitRdata(getPitRdataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`GetConnection(getConnectionOptions *GetConnectionOptions) - Operation response error`, func() {
		getConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
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
		getConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
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
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
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
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
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
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
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
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
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
	})
	Describe(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions) - Operation response error`, func() {
		completeConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
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
		completeConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
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
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
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
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
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
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
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
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions) - Operation response error`, func() {
		listDeploymentScalingGroupsPath := "/deployments/testString/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
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
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
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
	})
	Describe(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) - Operation response error`, func() {
		getDefaultScalingGroupsPath := "/deployables/postgresql/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
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
	})
	Describe(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) - Operation response error`, func() {
		setDeploymentScalingGroupPath := "/deployments/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(clouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(clouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(clouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
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

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(clouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(clouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) - Operation response error`, func() {
		getAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
	})
	Describe(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) - Operation response error`, func() {
		setAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetAutoscalingConditions with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingDiskGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
				autoscalingSetGroupAutoscalingModel.Disk = autoscalingDiskGroupDiskModel

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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingDiskGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
				autoscalingSetGroupAutoscalingModel.Disk = autoscalingDiskGroupDiskModel

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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingDiskGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
				autoscalingSetGroupAutoscalingModel.Disk = autoscalingDiskGroupDiskModel

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

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingDiskGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
				autoscalingSetGroupAutoscalingModel.Disk = autoscalingDiskGroupDiskModel

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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`KillConnections(killConnectionsOptions *KillConnectionsOptions) - Operation response error`, func() {
		killConnectionsPath := "/deployments/testString/management/database_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
	})
	Describe(`FileSync(fileSyncOptions *FileSyncOptions) - Operation response error`, func() {
		fileSyncPath := "/deployments/testString/elasticsearch/file_syncs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke FileSync with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(clouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`FileSync(fileSyncOptions *FileSyncOptions)`, func() {
		fileSyncPath := "/deployments/testString/elasticsearch/file_syncs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke FileSync successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(clouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.FileSyncWithContext(ctx, fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.FileSyncWithContext(ctx, fileSyncOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke FileSync successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.FileSync(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(clouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke FileSync with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(clouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the FileSyncOptions model with no property values
				fileSyncOptionsModelNew := new(clouddatabasesv5.FileSyncOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.FileSync(fileSyncOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions) - Operation response error`, func() {
		createLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
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

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(clouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
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
	})
	Describe(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions) - Operation response error`, func() {
		deleteLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
	})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})
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
				"CLOUD_DATABASES_URL":       "https://clouddatabasesv5/api",
				"CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(&clouddatabasesv5.CloudDatabasesV5Options{})

			It(`Instantiate service client with error`, func() {
				Expect(cloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"CLOUD_DATABASES_AUTH_TYPE": "NOAuth",
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
	Describe(`GetAllowlist(getAllowlistOptions *GetAllowlistOptions) - Operation response error`, func() {
		getAllowlistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllowlistPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
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
		getAllowlistPath := "/deployments/testString/whitelists/ip_addresses"
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
	})
	Describe(`ReplaceAllowlist(replaceAllowlistOptions *ReplaceAllowlistOptions) - Operation response error`, func() {
		replaceAllowlistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAllowlistPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceAllowlist with error: Operation response processing error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceAllowlistOptions model
				replaceAllowlistOptionsModel := new(clouddatabasesv5.ReplaceAllowlistOptions)
				replaceAllowlistOptionsModel.ID = core.StringPtr("testString")
				replaceAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				replaceAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				cloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceAllowlist(replaceAllowlistOptions *ReplaceAllowlistOptions)`, func() {
		replaceAllowlistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAllowlistPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ReplaceAllowlist successfully with retries`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())
				cloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceAllowlistOptions model
				replaceAllowlistOptionsModel := new(clouddatabasesv5.ReplaceAllowlistOptions)
				replaceAllowlistOptionsModel.ID = core.StringPtr("testString")
				replaceAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				replaceAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := cloudDatabasesService.ReplaceAllowlistWithContext(ctx, replaceAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				cloudDatabasesService.DisableRetries()
				result, response, operationErr := cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = cloudDatabasesService.ReplaceAllowlistWithContext(ctx, replaceAllowlistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceAllowlistPath))
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke ReplaceAllowlist successfully`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := cloudDatabasesService.ReplaceAllowlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceAllowlistOptions model
				replaceAllowlistOptionsModel := new(clouddatabasesv5.ReplaceAllowlistOptions)
				replaceAllowlistOptionsModel.ID = core.StringPtr("testString")
				replaceAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				replaceAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceAllowlist with error: Operation validation and request error`, func() {
				cloudDatabasesService, serviceErr := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(cloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceAllowlistOptions model
				replaceAllowlistOptionsModel := new(clouddatabasesv5.ReplaceAllowlistOptions)
				replaceAllowlistOptionsModel.ID = core.StringPtr("testString")
				replaceAllowlistOptionsModel.IPAddresses = []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}
				replaceAllowlistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := cloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceAllowlistOptions model with no property values
				replaceAllowlistOptionsModelNew := new(clouddatabasesv5.ReplaceAllowlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions) - Operation response error`, func() {
		addAllowlistEntryPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAllowlistEntryPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
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
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

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
		addAllowlistEntryPath := "/deployments/testString/whitelists/ip_addresses"
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

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
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")

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
	})
	Describe(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions) - Operation response error`, func() {
		deleteAllowlistEntryPath := "/deployments/testString/whitelists/ip_addresses/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllowlistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
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
		deleteAllowlistEntryPath := "/deployments/testString/whitelists/ip_addresses/testString"
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00.000Z"}}`)
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
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")
				Expect(allowlistEntryModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(allowlistEntryModel.Description).To(Equal(core.StringPtr("testString")))

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
			It(`Invoke NewChangeUserPasswordOptions successfully`, func() {
				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(clouddatabasesv5.APasswordSettingUser)
				Expect(aPasswordSettingUserModel).ToNot(BeNil())
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")
				Expect(aPasswordSettingUserModel.Password).To(Equal(core.StringPtr("xyzzy")))

				// Construct an instance of the ChangeUserPasswordOptions model
				id := "testString"
				userType := "testString"
				username := "testString"
				changeUserPasswordOptionsModel := cloudDatabasesService.NewChangeUserPasswordOptions(id, userType, username)
				changeUserPasswordOptionsModel.SetID("testString")
				changeUserPasswordOptionsModel.SetUserType("testString")
				changeUserPasswordOptionsModel.SetUsername("testString")
				changeUserPasswordOptionsModel.SetUser(aPasswordSettingUserModel)
				changeUserPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeUserPasswordOptionsModel).ToNot(BeNil())
				Expect(changeUserPasswordOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.User).To(Equal(aPasswordSettingUserModel))
				Expect(changeUserPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompleteConnectionOptions successfully`, func() {
				// Construct an instance of the CompleteConnectionOptions model
				id := "testString"
				userType := "testString"
				userID := "testString"
				endpointType := "public"
				completeConnectionOptionsModel := cloudDatabasesService.NewCompleteConnectionOptions(id, userType, userID, endpointType)
				completeConnectionOptionsModel.SetID("testString")
				completeConnectionOptionsModel.SetUserType("testString")
				completeConnectionOptionsModel.SetUserID("testString")
				completeConnectionOptionsModel.SetEndpointType("public")
				completeConnectionOptionsModel.SetPassword("testString")
				completeConnectionOptionsModel.SetCertificateRoot("testString")
				completeConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeConnectionOptionsModel).ToNot(BeNil())
				Expect(completeConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(completeConnectionOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDatabaseUserOptions successfully`, func() {
				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(clouddatabasesv5.CreateDatabaseUserRequestUser)
				Expect(createDatabaseUserRequestUserModel).ToNot(BeNil())
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")
				Expect(createDatabaseUserRequestUserModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(createDatabaseUserRequestUserModel.Username).To(Equal(core.StringPtr("james")))
				Expect(createDatabaseUserRequestUserModel.Password).To(Equal(core.StringPtr("kickoutthe")))

				// Construct an instance of the CreateDatabaseUserOptions model
				id := "testString"
				userType := "testString"
				createDatabaseUserOptionsModel := cloudDatabasesService.NewCreateDatabaseUserOptions(id, userType)
				createDatabaseUserOptionsModel.SetID("testString")
				createDatabaseUserOptionsModel.SetUserType("testString")
				createDatabaseUserOptionsModel.SetUser(createDatabaseUserRequestUserModel)
				createDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(createDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.User).To(Equal(createDatabaseUserRequestUserModel))
				Expect(createDatabaseUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLogicalReplicationSlotOptions successfully`, func() {
				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				Expect(logicalReplicationSlotLogicalReplicationSlotModel).ToNot(BeNil())
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.Name).To(Equal(core.StringPtr("customer_replication")))
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName).To(Equal(core.StringPtr("customers")))
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.PluginType).To(Equal(core.StringPtr("wal2json")))

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				id := "testString"
				createLogicalReplicationSlotOptionsModel := cloudDatabasesService.NewCreateLogicalReplicationSlotOptions(id)
				createLogicalReplicationSlotOptionsModel.SetID("testString")
				createLogicalReplicationSlotOptionsModel.SetLogicalReplicationSlot(logicalReplicationSlotLogicalReplicationSlotModel)
				createLogicalReplicationSlotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogicalReplicationSlotOptionsModel).ToNot(BeNil())
				Expect(createLogicalReplicationSlotOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot).To(Equal(logicalReplicationSlotLogicalReplicationSlotModel))
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
				userType := "testString"
				username := "testString"
				deleteDatabaseUserOptionsModel := cloudDatabasesService.NewDeleteDatabaseUserOptions(id, userType, username)
				deleteDatabaseUserOptionsModel.SetID("testString")
				deleteDatabaseUserOptionsModel.SetUserType("testString")
				deleteDatabaseUserOptionsModel.SetUsername("testString")
				deleteDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.Username).To(Equal(core.StringPtr("testString")))
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
			It(`Invoke NewFileSyncOptions successfully`, func() {
				// Construct an instance of the FileSyncOptions model
				id := "testString"
				fileSyncOptionsModel := cloudDatabasesService.NewFileSyncOptions(id)
				fileSyncOptionsModel.SetID("testString")
				fileSyncOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(fileSyncOptionsModel).ToNot(BeNil())
				Expect(fileSyncOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(fileSyncOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				userType := "testString"
				userID := "testString"
				endpointType := "public"
				getConnectionOptionsModel := cloudDatabasesService.NewGetConnectionOptions(id, userType, userID, endpointType)
				getConnectionOptionsModel.SetID("testString")
				getConnectionOptionsModel.SetUserType("testString")
				getConnectionOptionsModel.SetUserID("testString")
				getConnectionOptionsModel.SetEndpointType("public")
				getConnectionOptionsModel.SetCertificateRoot("testString")
				getConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConnectionOptionsModel).ToNot(BeNil())
				Expect(getConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(getConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDatabaseConfigurationSchemaOptions successfully`, func() {
				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				id := "testString"
				getDatabaseConfigurationSchemaOptionsModel := cloudDatabasesService.NewGetDatabaseConfigurationSchemaOptions(id)
				getDatabaseConfigurationSchemaOptionsModel.SetID("testString")
				getDatabaseConfigurationSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDatabaseConfigurationSchemaOptionsModel).ToNot(BeNil())
				Expect(getDatabaseConfigurationSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDatabaseConfigurationSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDefaultScalingGroupsOptions successfully`, func() {
				// Construct an instance of the GetDefaultScalingGroupsOptions model
				typeVar := "postgresql"
				getDefaultScalingGroupsOptionsModel := cloudDatabasesService.NewGetDefaultScalingGroupsOptions(typeVar)
				getDefaultScalingGroupsOptionsModel.SetType("postgresql")
				getDefaultScalingGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDefaultScalingGroupsOptionsModel).ToNot(BeNil())
				Expect(getDefaultScalingGroupsOptionsModel.Type).To(Equal(core.StringPtr("postgresql")))
				Expect(getDefaultScalingGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetPitRdataOptions successfully`, func() {
				// Construct an instance of the GetPitRdataOptions model
				id := "testString"
				getPitRdataOptionsModel := cloudDatabasesService.NewGetPitRdataOptions(id)
				getPitRdataOptionsModel.SetID("testString")
				getPitRdataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPitRdataOptionsModel).ToNot(BeNil())
				Expect(getPitRdataOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPitRdataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRemotesSchemaOptions successfully`, func() {
				// Construct an instance of the GetRemotesSchemaOptions model
				id := "testString"
				getRemotesSchemaOptionsModel := cloudDatabasesService.NewGetRemotesSchemaOptions(id)
				getRemotesSchemaOptionsModel.SetID("testString")
				getRemotesSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRemotesSchemaOptionsModel).ToNot(BeNil())
				Expect(getRemotesSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRemotesSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetUserOptions successfully`, func() {
				// Construct an instance of the GetUserOptions model
				id := "testString"
				userID := "testString"
				getUserOptionsModel := cloudDatabasesService.NewGetUserOptions(id, userID)
				getUserOptionsModel.SetID("testString")
				getUserOptionsModel.SetUserID("testString")
				getUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserOptionsModel).ToNot(BeNil())
				Expect(getUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewReplaceAllowlistOptions successfully`, func() {
				// Construct an instance of the AllowlistEntry model
				allowlistEntryModel := new(clouddatabasesv5.AllowlistEntry)
				Expect(allowlistEntryModel).ToNot(BeNil())
				allowlistEntryModel.Address = core.StringPtr("testString")
				allowlistEntryModel.Description = core.StringPtr("testString")
				Expect(allowlistEntryModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(allowlistEntryModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceAllowlistOptions model
				id := "testString"
				replaceAllowlistOptionsModel := cloudDatabasesService.NewReplaceAllowlistOptions(id)
				replaceAllowlistOptionsModel.SetID("testString")
				replaceAllowlistOptionsModel.SetIPAddresses([]clouddatabasesv5.AllowlistEntry{*allowlistEntryModel})
				replaceAllowlistOptionsModel.SetIfMatch("testString")
				replaceAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAllowlistOptionsModel).ToNot(BeNil())
				Expect(replaceAllowlistOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAllowlistOptionsModel.IPAddresses).To(Equal([]clouddatabasesv5.AllowlistEntry{*allowlistEntryModel}))
				Expect(replaceAllowlistOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetAutoscalingConditionsOptions successfully`, func() {
				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				Expect(autoscalingDiskGroupDiskScalersCapacityModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))
				Expect(autoscalingDiskGroupDiskScalersCapacityModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod).To(Equal(core.StringPtr("30m")))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent).To(Equal(core.Int64Ptr(int64(45))))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				Expect(autoscalingDiskGroupDiskScalersModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel
				Expect(autoscalingDiskGroupDiskScalersModel.Capacity).To(Equal(autoscalingDiskGroupDiskScalersCapacityModel))
				Expect(autoscalingDiskGroupDiskScalersModel.IoUtilization).To(Equal(autoscalingDiskGroupDiskScalersIoUtilizationModel))

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(clouddatabasesv5.AutoscalingDiskGroupDiskRate)
				Expect(autoscalingDiskGroupDiskRateModel).ToNot(BeNil())
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")
				Expect(autoscalingDiskGroupDiskRateModel.IncreasePercent).To(Equal(core.Float64Ptr(float64(20))))
				Expect(autoscalingDiskGroupDiskRateModel.PeriodSeconds).To(Equal(core.Int64Ptr(int64(900))))
				Expect(autoscalingDiskGroupDiskRateModel.LimitMbPerMember).To(Equal(core.Float64Ptr(float64(3670016))))
				Expect(autoscalingDiskGroupDiskRateModel.Units).To(Equal(core.StringPtr("mb")))

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(clouddatabasesv5.AutoscalingDiskGroupDisk)
				Expect(autoscalingDiskGroupDiskModel).ToNot(BeNil())
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel
				Expect(autoscalingDiskGroupDiskModel.Scalers).To(Equal(autoscalingDiskGroupDiskScalersModel))
				Expect(autoscalingDiskGroupDiskModel.Rate).To(Equal(autoscalingDiskGroupDiskRateModel))

				// Construct an instance of the AutoscalingSetGroupAutoscalingAutoscalingDiskGroup model
				autoscalingSetGroupAutoscalingModel := new(clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
				Expect(autoscalingSetGroupAutoscalingModel).ToNot(BeNil())
				autoscalingSetGroupAutoscalingModel.Disk = autoscalingDiskGroupDiskModel
				Expect(autoscalingSetGroupAutoscalingModel.Disk).To(Equal(autoscalingDiskGroupDiskModel))

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
			It(`Invoke NewSetDatabaseConfigurationOptions successfully`, func() {
				// Construct an instance of the SetConfigurationConfigurationPgConfiguration model
				setConfigurationConfigurationModel := new(clouddatabasesv5.SetConfigurationConfigurationPgConfiguration)
				Expect(setConfigurationConfigurationModel).ToNot(BeNil())
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				Expect(setConfigurationConfigurationModel.MaxConnections).To(Equal(core.Int64Ptr(int64(115))))
				Expect(setConfigurationConfigurationModel.MaxPreparedTransactions).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setConfigurationConfigurationModel.DeadlockTimeout).To(Equal(core.Int64Ptr(int64(100))))
				Expect(setConfigurationConfigurationModel.EffectiveIoConcurrency).To(Equal(core.Int64Ptr(int64(1))))
				Expect(setConfigurationConfigurationModel.MaxReplicationSlots).To(Equal(core.Int64Ptr(int64(10))))
				Expect(setConfigurationConfigurationModel.MaxWalSenders).To(Equal(core.Int64Ptr(int64(12))))
				Expect(setConfigurationConfigurationModel.SharedBuffers).To(Equal(core.Int64Ptr(int64(16))))
				Expect(setConfigurationConfigurationModel.SynchronousCommit).To(Equal(core.StringPtr("local")))
				Expect(setConfigurationConfigurationModel.WalLevel).To(Equal(core.StringPtr("hot_standby")))
				Expect(setConfigurationConfigurationModel.ArchiveTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(setConfigurationConfigurationModel.LogMinDurationStatement).To(Equal(core.Int64Ptr(int64(100))))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				id := "testString"
				var setDatabaseConfigurationOptionsConfiguration clouddatabasesv5.SetConfigurationConfigurationIntf = nil
				setDatabaseConfigurationOptionsModel := cloudDatabasesService.NewSetDatabaseConfigurationOptions(id, setDatabaseConfigurationOptionsConfiguration)
				setDatabaseConfigurationOptionsModel.SetID("testString")
				setDatabaseConfigurationOptionsModel.SetConfiguration(setConfigurationConfigurationModel)
				setDatabaseConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDatabaseConfigurationOptionsModel).ToNot(BeNil())
				Expect(setDatabaseConfigurationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDatabaseConfigurationOptionsModel.Configuration).To(Equal(setConfigurationConfigurationModel))
				Expect(setDatabaseConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetDeploymentScalingGroupOptions successfully`, func() {
				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(clouddatabasesv5.SetMembersGroupMembers)
				Expect(setMembersGroupMembersModel).ToNot(BeNil())
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))
				Expect(setMembersGroupMembersModel.AllocationCount).To(Equal(core.Int64Ptr(int64(4))))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				Expect(setDeploymentScalingGroupRequestModel).ToNot(BeNil())
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel
				Expect(setDeploymentScalingGroupRequestModel.Members).To(Equal(setMembersGroupMembersModel))

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				id := "testString"
				groupID := "testString"
				var setDeploymentScalingGroupRequest clouddatabasesv5.SetDeploymentScalingGroupRequestIntf = nil
				setDeploymentScalingGroupOptionsModel := cloudDatabasesService.NewSetDeploymentScalingGroupOptions(id, groupID, setDeploymentScalingGroupRequest)
				setDeploymentScalingGroupOptionsModel.SetID("testString")
				setDeploymentScalingGroupOptionsModel.SetGroupID("testString")
				setDeploymentScalingGroupOptionsModel.SetSetDeploymentScalingGroupRequest(setDeploymentScalingGroupRequestModel)
				setDeploymentScalingGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDeploymentScalingGroupOptionsModel).ToNot(BeNil())
				Expect(setDeploymentScalingGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest).To(Equal(setDeploymentScalingGroupRequestModel))
				Expect(setDeploymentScalingGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetPromotionOptions successfully`, func() {
				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(clouddatabasesv5.SetPromotionPromotionPromote)
				Expect(setPromotionPromotionModel).ToNot(BeNil())
				setPromotionPromotionModel.Promotion = make(map[string]interface{})
				Expect(setPromotionPromotionModel.Promotion).To(Equal(make(map[string]interface{})))

				// Construct an instance of the SetPromotionOptions model
				id := "testString"
				var setPromotionOptionsPromotion clouddatabasesv5.SetPromotionPromotionIntf = nil
				setPromotionOptionsModel := cloudDatabasesService.NewSetPromotionOptions(id, setPromotionOptionsPromotion)
				setPromotionOptionsModel.SetID("testString")
				setPromotionOptionsModel.SetPromotion(setPromotionPromotionModel)
				setPromotionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setPromotionOptionsModel).ToNot(BeNil())
				Expect(setPromotionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setPromotionOptionsModel.Promotion).To(Equal(setPromotionPromotionModel))
				Expect(setPromotionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
