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

			addAllowlistEntryOptions := cloudDatabasesService.NewAddAllowlistEntryOptions(
				"testString",
			)

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
		It(`CreateDatabaseUser request example`, func() {
			// begin-createDatabaseUser

			createDatabaseUserOptions := cloudDatabasesService.NewCreateDatabaseUserOptions(
				"testString",
				"testString",
			)

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
				"testString",
				"testString",
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
		It(`ReplaceAllowlist request example`, func() {
			// begin-replaceAllowlist

			replaceAllowlistOptions := cloudDatabasesService.NewReplaceAllowlistOptions(
				"testString",
			)

			replaceAllowlistResponse, response, err := cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(replaceAllowlistResponse, "", "  ")
			fmt.Println(string(b))

			// end-replaceAllowlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(replaceAllowlistResponse).ToNot(BeNil())

			taskIDLink = *replaceAllowlistResponse.Task.ID

		})
		It(`SetDeploymentScalingGroup request example`, func() {
			// begin-setDeploymentScalingGroup

			setDeploymentScalingGroupRequestModel := &clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup{}

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
		It(`ChangeUserPassword request example`, func() {
			// begin-changeUserPassword

			changeUserPasswordOptions := cloudDatabasesService.NewChangeUserPasswordOptions(
				"testString",
				"testString",
				"testString",
			)

			changeUserPasswordResponse, response, err := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(changeUserPasswordResponse, "", "  ")
			fmt.Println(string(b))

			// end-changeUserPassword

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(changeUserPasswordResponse).ToNot(BeNil())

		})
		It(`GetUser request example`, func() {
			// begin-getUser

			getUserOptions := cloudDatabasesService.NewGetUserOptions(
				"testString",
				"testString",
			)

			task, response, err := cloudDatabasesService.GetUser(getUserOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(task, "", "  ")
			fmt.Println(string(b))

			// end-getUser

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())

		})
		It(`SetDatabaseConfiguration request example`, func() {
			// begin-setDatabaseConfiguration

			setConfigurationConfigurationModel := &clouddatabasesv5.SetConfigurationConfigurationPgConfiguration{}

			setDatabaseConfigurationOptions := cloudDatabasesService.NewSetDatabaseConfigurationOptions(
				"testString",
				setConfigurationConfigurationModel,
			)

			setDatabaseConfigurationResponse, response, err := cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(setDatabaseConfigurationResponse, "", "  ")
			fmt.Println(string(b))

			// end-setDatabaseConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setDatabaseConfigurationResponse).ToNot(BeNil())

		})
		It(`GetDatabaseConfigurationSchema request example`, func() {
			// begin-getDatabaseConfigurationSchema

			getDatabaseConfigurationSchemaOptions := cloudDatabasesService.NewGetDatabaseConfigurationSchemaOptions(
				"testString",
			)

			configurationSchema, response, err := cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configurationSchema, "", "  ")
			fmt.Println(string(b))

			// end-getDatabaseConfigurationSchema

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configurationSchema).ToNot(BeNil())

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
		It(`GetRemotesSchema request example`, func() {
			// begin-getRemotesSchema

			getRemotesSchemaOptions := cloudDatabasesService.NewGetRemotesSchemaOptions(
				"testString",
			)

			getRemotesSchemaResponse, response, err := cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getRemotesSchemaResponse, "", "  ")
			fmt.Println(string(b))

			// end-getRemotesSchema

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getRemotesSchemaResponse).ToNot(BeNil())

		})
		It(`SetPromotion request example`, func() {
			// begin-setPromotion

			setPromotionPromotionModel := &clouddatabasesv5.SetPromotionPromotionPromote{}

			setPromotionOptions := cloudDatabasesService.NewSetPromotionOptions(
				"testString",
				setPromotionPromotionModel,
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
				"testString",
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
				"testString",
				"testString",
				"public",
			)

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
		It(`SetAutoscalingConditions request example`, func() {
			// begin-setAutoscalingConditions

			autoscalingSetGroupAutoscalingModel := &clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup{}

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
			Expect(response.StatusCode).To(Equal(200))
			Expect(setAutoscalingConditionsResponse).ToNot(BeNil())

		})
		It(`FileSync request example`, func() {
			// begin-fileSync

			fileSyncOptions := cloudDatabasesService.NewFileSyncOptions(
				"testString",
			)

			fileSyncResponse, response, err := cloudDatabasesService.FileSync(fileSyncOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(fileSyncResponse, "", "  ")
			fmt.Println(string(b))

			// end-fileSync

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(fileSyncResponse).ToNot(BeNil())

		})
		It(`CreateLogicalReplicationSlot request example`, func() {
			// begin-createLogicalReplicationSlot

			createLogicalReplicationSlotOptions := cloudDatabasesService.NewCreateLogicalReplicationSlotOptions(
				"testString",
			)

			createLogicalReplicationSlotResponse, response, err := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createLogicalReplicationSlotResponse, "", "  ")
			fmt.Println(string(b))

			// end-createLogicalReplicationSlot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createLogicalReplicationSlotResponse).ToNot(BeNil())

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
			Expect(response.StatusCode).To(Equal(200))
			Expect(killConnectionsResponse).ToNot(BeNil())

		})
		It(`DeleteLogicalReplicationSlot request example`, func() {
			// begin-deleteLogicalReplicationSlot

			deleteLogicalReplicationSlotOptions := cloudDatabasesService.NewDeleteLogicalReplicationSlotOptions(
				"testString",
				"testString",
			)

			deleteLogicalReplicationSlotResponse, response, err := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteLogicalReplicationSlotResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteLogicalReplicationSlot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteLogicalReplicationSlotResponse).ToNot(BeNil())

		})
		It(`DeleteDatabaseUser request example`, func() {
			// begin-deleteDatabaseUser

			deleteDatabaseUserOptions := cloudDatabasesService.NewDeleteDatabaseUserOptions(
				"testString",
				"testString",
				"testString",
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
	})
})
