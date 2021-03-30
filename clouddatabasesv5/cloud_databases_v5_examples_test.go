// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/icd-go-sdk/clouddatabasesv5"
)

//
// This file provides an example of how to use the Cloud Databases service.
//
// The following configuration properties are assumed to be defined:
// CLOUD_DATABASES_URL=<service base url>
// CLOUD_DATABASES_AUTH_TYPE=iam
// CLOUD_DATABASES_APIKEY=<IAM apikey>
// CLOUD_DATABASES_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../cloud_databases.env"

var (
	cloudDatabasesService *clouddatabasesv5.CloudDatabasesV5
	config                map[string]string
	configLoaded          bool = false
)

// Globlal variables to hold link values
var (
	taskIDLink string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`CloudDatabasesV5 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(clouddatabasesv5.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			cloudDatabasesServiceOptions := &clouddatabasesv5.CloudDatabasesV5Options{}

			cloudDatabasesService, err = clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(cloudDatabasesServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(cloudDatabasesService).ToNot(BeNil())
		})
	})

	Describe(`CloudDatabasesV5 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddAllowlistEntry request example`, func() {
			// begin-addAllowlistEntry

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     core.StringPtr("172.16.0.0/16"),
				Description: core.StringPtr("Dev IP space 3"),
			}

			addAllowlistEntryOptions := cloudDatabasesService.NewAddAllowlistEntryOptions(
				"testString",
			)
			addAllowlistEntryOptions.SetIPAddress(allowlistEntryModel)

			addAllowlistEntryResponse, response, err := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addAllowlistEntryResponse, "", "  ")
			fmt.Println(string(b))

			// end-addAllowlistEntry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(addAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *addAllowlistEntryResponse.Task.ID

		})
		It(`ChangeUserPassword request example`, func() {
			// begin-changeUserPassword

			aPasswordSettingUserModel := &clouddatabasesv5.APasswordSettingUser{
				Password: core.StringPtr("xyzzyyzzyx"),
			}

			changeUserPasswordOptions := cloudDatabasesService.NewChangeUserPasswordOptions(
				"testString",
				"database",
				"james",
			)
			changeUserPasswordOptions.SetUser(aPasswordSettingUserModel)

			changeUserPasswordResponse, response, err := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(changeUserPasswordResponse, "", "  ")
			fmt.Println(string(b))

			// end-changeUserPassword

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(changeUserPasswordResponse).ToNot(BeNil())

			taskIDLink = *changeUserPasswordResponse.Task.ID

		})
		It(`CreateDatabaseUser request example`, func() {
			// begin-createDatabaseUser

			createDatabaseUserRequestUserModel := &clouddatabasesv5.CreateDatabaseUserRequestUser{
				Username: core.StringPtr("james"),
				Password: core.StringPtr("kickoutthe"),
			}

			createDatabaseUserOptions := cloudDatabasesService.NewCreateDatabaseUserOptions(
				"testString",
				"database",
			)
			createDatabaseUserOptions.SetUser(createDatabaseUserRequestUserModel)

			createDatabaseUserResponse, response, err := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createDatabaseUserResponse, "", "  ")
			fmt.Println(string(b))

			// end-createDatabaseUser

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *createDatabaseUserResponse.Task.ID

		})
		It(`DeleteAllowlistEntry request example`, func() {
			// begin-deleteAllowlistEntry

			deleteAllowlistEntryOptions := cloudDatabasesService.NewDeleteAllowlistEntryOptions(
				"testString",
				"testString",
			)

			deleteAllowlistEntryResponse, response, err := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteAllowlistEntryResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteAllowlistEntry

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *deleteAllowlistEntryResponse.Task.ID

		})
		It(`DeleteDatabaseUser request example`, func() {
			// begin-deleteDatabaseUser

			deleteDatabaseUserOptions := cloudDatabasesService.NewDeleteDatabaseUserOptions(
				"testString",
				"database",
				"james",
			)

			deleteDatabaseUserResponse, response, err := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteDatabaseUserResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteDatabaseUser

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *deleteDatabaseUserResponse.Task.ID

		})
		It(`KillConnections request example`, func() {
			// begin-killConnections

			killConnectionsOptions := cloudDatabasesService.NewKillConnectionsOptions(
				"testString",
			)

			killConnectionsResponse, response, err := cloudDatabasesService.KillConnections(killConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(killConnectionsResponse, "", "  ")
			fmt.Println(string(b))

			// end-killConnections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(killConnectionsResponse).ToNot(BeNil())

			taskIDLink = *killConnectionsResponse.Task.ID

		})
		It(`SetAllowlist request example`, func() {
			// begin-setAllowlist

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     core.StringPtr("195.212.0.0/16"),
				Description: core.StringPtr("Dev IP space 1"),
			}

			setAllowlistOptions := cloudDatabasesService.NewSetAllowlistOptions(
				"testString",
			)
			setAllowlistOptions.SetIPAddresses([]clouddatabasesv5.AllowlistEntry{*allowlistEntryModel})

			setAllowlistResponse, response, err := cloudDatabasesService.SetAllowlist(setAllowlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setAllowlistResponse, "", "  ")
			fmt.Println(string(b))

			// end-setAllowlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setAllowlistResponse).ToNot(BeNil())

			taskIDLink = *setAllowlistResponse.Task.ID

		})
		It(`SetAutoscalingConditions request example`, func() {
			// begin-setAutoscalingConditions

			autoscalingMemoryGroupMemoryScalersIoUtilizationModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization{
				Enabled:      core.BoolPtr(true),
				OverPeriod:   core.StringPtr("5m"),
				AbovePercent: core.Int64Ptr(int64(90)),
			}

			autoscalingMemoryGroupMemoryScalersModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers{
				IoUtilization: autoscalingMemoryGroupMemoryScalersIoUtilizationModel,
			}

			autoscalingMemoryGroupMemoryRateModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryRate{
				IncreasePercent:  core.Float64Ptr(float64(10.0)),
				PeriodSeconds:    core.Int64Ptr(int64(300)),
				LimitMbPerMember: core.Float64Ptr(float64(125952)),
				Units:            core.StringPtr("mb"),
			}

			autoscalingMemoryGroupMemoryModel := &clouddatabasesv5.AutoscalingMemoryGroupMemory{
				Scalers: autoscalingMemoryGroupMemoryScalersModel,
				Rate:    autoscalingMemoryGroupMemoryRateModel,
			}

			autoscalingSetGroupAutoscalingModel := &clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup{
				Memory: autoscalingMemoryGroupMemoryModel,
			}

			setAutoscalingConditionsOptions := cloudDatabasesService.NewSetAutoscalingConditionsOptions(
				"testString",
				"testString",
				autoscalingSetGroupAutoscalingModel,
			)

			setAutoscalingConditionsResponse, response, err := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setAutoscalingConditionsResponse, "", "  ")
			fmt.Println(string(b))

			// end-setAutoscalingConditions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setAutoscalingConditionsResponse).ToNot(BeNil())

			taskIDLink = *setAutoscalingConditionsResponse.Task.ID

		})
		It(`SetDeploymentScalingGroup request example`, func() {
			// begin-setDeploymentScalingGroup

			setMemoryGroupMemoryModel := &clouddatabasesv5.SetMemoryGroupMemory{
				AllocationMb: core.Int64Ptr(int64(4096)),
			}

			setDeploymentScalingGroupRequestModel := &clouddatabasesv5.SetDeploymentScalingGroupRequestSetMemoryGroup{
				Memory: setMemoryGroupMemoryModel,
			}

			setDeploymentScalingGroupOptions := cloudDatabasesService.NewSetDeploymentScalingGroupOptions(
				"testString",
				"testString",
				setDeploymentScalingGroupRequestModel,
			)

			setDeploymentScalingGroupResponse, response, err := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setDeploymentScalingGroupResponse, "", "  ")
			fmt.Println(string(b))

			// end-setDeploymentScalingGroup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setDeploymentScalingGroupResponse).ToNot(BeNil())

			taskIDLink = *setDeploymentScalingGroupResponse.Task.ID

		})
		It(`UpdateDatabaseConfiguration request example`, func() {
			// begin-updateDatabaseConfiguration

			setConfigurationConfigurationModel := &clouddatabasesv5.SetConfigurationConfigurationPgConfiguration{
				MaxConnections: core.Int64Ptr(int64(200)),
			}

			updateDatabaseConfigurationOptions := cloudDatabasesService.NewUpdateDatabaseConfigurationOptions(
				"testString",
				setConfigurationConfigurationModel,
			)

			updateDatabaseConfigurationResponse, response, err := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updateDatabaseConfigurationResponse, "", "  ")
			fmt.Println(string(b))

			// end-updateDatabaseConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateDatabaseConfigurationResponse).ToNot(BeNil())

			taskIDLink = *updateDatabaseConfigurationResponse.Task.ID

		})
		It(`ListDeployables request example`, func() {
			// begin-listDeployables

			listDeployablesOptions := cloudDatabasesService.NewListDeployablesOptions()

			listDeployablesResponse, response, err := cloudDatabasesService.ListDeployables(listDeployablesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listDeployablesResponse, "", "  ")
			fmt.Println(string(b))

			// end-listDeployables

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDeployablesResponse).ToNot(BeNil())

		})
		It(`ListRegions request example`, func() {
			// begin-listRegions

			listRegionsOptions := cloudDatabasesService.NewListRegionsOptions()

			listRegionsResponse, response, err := cloudDatabasesService.ListRegions(listRegionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listRegionsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listRegions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRegionsResponse).ToNot(BeNil())

		})
		It(`GetDeploymentInfo request example`, func() {
			// begin-getDeploymentInfo

			getDeploymentInfoOptions := cloudDatabasesService.NewGetDeploymentInfoOptions(
				"testString",
			)

			getDeploymentInfoResponse, response, err := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getDeploymentInfoResponse, "", "  ")
			fmt.Println(string(b))

			// end-getDeploymentInfo

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDeploymentInfoResponse).ToNot(BeNil())

		})
		It(`ListRemotes request example`, func() {
			// begin-listRemotes

			listRemotesOptions := cloudDatabasesService.NewListRemotesOptions(
				"testString",
			)

			listRemotesResponse, response, err := cloudDatabasesService.ListRemotes(listRemotesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listRemotesResponse, "", "  ")
			fmt.Println(string(b))

			// end-listRemotes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRemotesResponse).ToNot(BeNil())

		})
		It(`ResyncReplica request example`, func() {
			// begin-resyncReplica

			resyncReplicaOptions := cloudDatabasesService.NewResyncReplicaOptions(
				"testString",
			)

			resyncReplicaResponse, response, err := cloudDatabasesService.ResyncReplica(resyncReplicaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resyncReplicaResponse, "", "  ")
			fmt.Println(string(b))

			// end-resyncReplica

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resyncReplicaResponse).ToNot(BeNil())

		})
		It(`SetPromotion request example`, func() {
			// begin-setPromotion

			setPromotionOptions := cloudDatabasesService.NewSetPromotionOptions(
				"testString",
			)

			setPromotionResponse, response, err := cloudDatabasesService.SetPromotion(setPromotionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setPromotionResponse, "", "  ")
			fmt.Println(string(b))

			// end-setPromotion

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setPromotionResponse).ToNot(BeNil())

		})
		It(`ListDeploymentTasks request example`, func() {
			// begin-listDeploymentTasks

			listDeploymentTasksOptions := cloudDatabasesService.NewListDeploymentTasksOptions(
				"testString",
			)

			tasks, response, err := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tasks, "", "  ")
			fmt.Println(string(b))

			// end-listDeploymentTasks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tasks).ToNot(BeNil())

		})
		It(`GetTask request example`, func() {
			// begin-getTask

			getTaskOptions := cloudDatabasesService.NewGetTaskOptions(
				taskIDLink,
			)

			getTaskResponse, response, err := cloudDatabasesService.GetTask(getTaskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getTaskResponse, "", "  ")
			fmt.Println(string(b))

			// end-getTask

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTaskResponse).ToNot(BeNil())

		})
		It(`GetBackupInfo request example`, func() {
			// begin-getBackupInfo

			getBackupInfoOptions := cloudDatabasesService.NewGetBackupInfoOptions(
				"testString",
			)

			getBackupInfoResponse, response, err := cloudDatabasesService.GetBackupInfo(getBackupInfoOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getBackupInfoResponse, "", "  ")
			fmt.Println(string(b))

			// end-getBackupInfo

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBackupInfoResponse).ToNot(BeNil())

		})
		It(`ListDeploymentBackups request example`, func() {
			// begin-listDeploymentBackups

			listDeploymentBackupsOptions := cloudDatabasesService.NewListDeploymentBackupsOptions(
				"testString",
			)

			backups, response, err := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(backups, "", "  ")
			fmt.Println(string(b))

			// end-listDeploymentBackups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backups).ToNot(BeNil())

		})
		It(`StartOndemandBackup request example`, func() {
			// begin-startOndemandBackup

			startOndemandBackupOptions := cloudDatabasesService.NewStartOndemandBackupOptions(
				"testString",
			)

			startOndemandBackupResponse, response, err := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(startOndemandBackupResponse, "", "  ")
			fmt.Println(string(b))

			// end-startOndemandBackup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(startOndemandBackupResponse).ToNot(BeNil())

		})
		It(`GetPitRdata request example`, func() {
			// begin-getPITRdata

			getPitRdataOptions := cloudDatabasesService.NewGetPitRdataOptions(
				"testString",
			)

			pointInTimeRecoveryData, response, err := cloudDatabasesService.GetPitRdata(getPitRdataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pointInTimeRecoveryData, "", "  ")
			fmt.Println(string(b))

			// end-getPITRdata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pointInTimeRecoveryData).ToNot(BeNil())

		})
		It(`GetConnection request example`, func() {
			// begin-getConnection

			getConnectionOptions := cloudDatabasesService.NewGetConnectionOptions(
				"testString",
				"database",
				"testString",
				"public",
			)

			connection, response, err := cloudDatabasesService.GetConnection(getConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(connection, "", "  ")
			fmt.Println(string(b))

			// end-getConnection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connection).ToNot(BeNil())

		})
		It(`CompleteConnection request example`, func() {
			// begin-completeConnection

			completeConnectionOptions := cloudDatabasesService.NewCompleteConnectionOptions(
				"testString",
				"database",
				"testString",
				"public",
			)
			completeConnectionOptions.SetPassword("providedpassword")

			connection, response, err := cloudDatabasesService.CompleteConnection(completeConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(connection, "", "  ")
			fmt.Println(string(b))

			// end-completeConnection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connection).ToNot(BeNil())

		})
		It(`ListDeploymentScalingGroups request example`, func() {
			// begin-listDeploymentScalingGroups

			listDeploymentScalingGroupsOptions := cloudDatabasesService.NewListDeploymentScalingGroupsOptions(
				"testString",
			)

			groups, response, err := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(groups, "", "  ")
			fmt.Println(string(b))

			// end-listDeploymentScalingGroups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groups).ToNot(BeNil())

		})
		It(`GetDefaultScalingGroups request example`, func() {
			// begin-getDefaultScalingGroups

			getDefaultScalingGroupsOptions := cloudDatabasesService.NewGetDefaultScalingGroupsOptions(
				"postgresql",
			)

			groups, response, err := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(groups, "", "  ")
			fmt.Println(string(b))

			// end-getDefaultScalingGroups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groups).ToNot(BeNil())

		})
		It(`GetAutoscalingConditions request example`, func() {
			// begin-getAutoscalingConditions

			getAutoscalingConditionsOptions := cloudDatabasesService.NewGetAutoscalingConditionsOptions(
				"testString",
				"testString",
			)

			autoscalingGroup, response, err := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(autoscalingGroup, "", "  ")
			fmt.Println(string(b))

			// end-getAutoscalingConditions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(autoscalingGroup).ToNot(BeNil())

		})
		It(`GetAllowlist request example`, func() {
			// begin-getAllowlist

			getAllowlistOptions := cloudDatabasesService.NewGetAllowlistOptions(
				"testString",
			)

			allowlist, response, err := cloudDatabasesService.GetAllowlist(getAllowlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(allowlist, "", "  ")
			fmt.Println(string(b))

			// end-getAllowlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allowlist).ToNot(BeNil())

		})
	})
})
