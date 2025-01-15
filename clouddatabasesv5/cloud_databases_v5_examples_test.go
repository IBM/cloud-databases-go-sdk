//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
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
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Cloud Databases service.
//
// The following configuration properties are assumed to be defined:
// CLOUD_DATABASES_URL=<service base url>
// CLOUD_DATABASES_AUTH_TYPE=iam
// CLOUD_DATABASES_APIKEY=<IAM apikey>
// CLOUD_DATABASES_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// CLOUD_DATABASES_DEPLOYMENT_ID=<ID of an example deployment>
// CLOUD_DATABASES_REPLICA_ID=<ID of an example replica>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
const externalConfigFile = "../cloud_databases.env"

var (
	cloudDatabasesService *clouddatabasesv5.CloudDatabasesV5
	config                map[string]string
	configLoaded          bool = false

	IPAddress1         string = "195.212.0.0/16"
	IPAddress3         string = "172.16.0.0/16"
	username           string = "exampleUsername"
	password           string = "examplePassword"
	newPassword        string = "exampleNewPassword"
	userType           string = "database"
	autoScalingGroupID string = "member"
	deploymentID       string
	replicaID          string
	backupIDLink       string
	scalingGroupIDLink string
	logicalRepName     string = "wj123"
	databaseName       string = "exampleDatabase"
	pluginType         string = "wal2json"
)

// Globlal variables to hold link values
var (
	taskIDLink string
)

var waitForTask = func(taskID string) {
	getTaskOptions := &clouddatabasesv5.GetTaskOptions{
		ID: &taskID,
	}

	// If the task runs for more than a minute, then we'll consider it to have succeeded.
	for complete, attempts := false, 0; !complete && attempts < 30; attempts++ {
		getTaskResponse, response, err := cloudDatabasesService.GetTask(getTaskOptions)

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(200))
		Expect(getTaskResponse).ToNot(BeNil())

		if getTaskResponse.Task == nil {
			complete = true
		} else {
			switch *getTaskResponse.Task.Status {
			case "completed", "failed":
				complete = true
				Expect(*getTaskResponse.Task.Status).To(Equal("completed"))
			case "queued", "running":
				break // from switch, not from for
			default:
				fmt.Println("status is " + *getTaskResponse.Task.Status)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

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

			deploymentID = config["DEPLOYMENT_ID"]
			Expect(deploymentID).ToNot(BeEmpty())

			replicaID = config["REPLICA_ID"]
			Expect(replicaID).ToNot(BeEmpty())

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
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
			fmt.Println("\nAddAllowlistEntry() result:")
			// begin-addAllowlistEntry

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     &IPAddress3,
				Description: core.StringPtr("Dev IP space 3"),
			}

			addAllowlistEntryOptions := cloudDatabasesService.NewAddAllowlistEntryOptions(
				deploymentID,
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

			waitForTask(taskIDLink)
		})
		It(`DeleteAllowlistEntry request example`, func() {
			fmt.Println("\nDeleteAllowlistEntry() result:")
			// begin-deleteAllowlistEntry

			deleteAllowlistEntryOptions := cloudDatabasesService.NewDeleteAllowlistEntryOptions(
				deploymentID,
				IPAddress3,
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

			waitForTask(taskIDLink)
		})
		It(`CreateDatabaseUser request example`, func() {
			fmt.Println("\nCreateDatabaseUser() result:")
			// begin-createDatabaseUser

			userModel := &clouddatabasesv5.User{
				Username: &username,
				Password: &password,
			}

			createDatabaseUserOptions := cloudDatabasesService.NewCreateDatabaseUserOptions(
				deploymentID,
				userType,
			)
			createDatabaseUserOptions.SetUser(userModel)

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

			waitForTask(taskIDLink)
		})
		It(`CreateLogicalReplicationSlot request example`, func() {
			fmt.Println("\nCreateLogicalReplicationSlot() result:")
			// begin-createLogicalReplicationSlot

			logicalReplicationSlotModel := &clouddatabasesv5.LogicalReplicationSlot{
				Name:         &logicalRepName,
				DatabaseName: &databaseName,
				PluginType:   &pluginType,
			}

			createLogicalReplicationSlotOptions := cloudDatabasesService.NewCreateLogicalReplicationSlotOptions(deploymentID)
			createLogicalReplicationSlotOptions.SetLogicalReplicationSlot(logicalReplicationSlotModel)

			createLogicalReplicationResponse, response, err := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createLogicalReplicationResponse, "", "  ")
			fmt.Println(string(b))

			// end-createLogicalReplicationSlot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createLogicalReplicationResponse).ToNot(BeNil())

			taskIDLink = *createLogicalReplicationResponse.Task.ID

			waitForTask(taskIDLink)
		})
		It(`ChangeUserPassword request example`, func() {
			fmt.Println("\nChangeUserPassword() result:")
			// begin-changeUserPassword

			aPasswordSettingUserModel := &clouddatabasesv5.APasswordSettingUser{
				Password: &newPassword,
			}

			changeUserPasswordOptions := cloudDatabasesService.NewChangeUserPasswordOptions(
				deploymentID,
				userType,
				username,
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

			waitForTask(taskIDLink)
		})
		It(`DeleteDatabaseUser request example`, func() {
			fmt.Println("\nDeleteDatabaseUser() result:")
			// begin-deleteDatabaseUser

			deleteDatabaseUserOptions := cloudDatabasesService.NewDeleteDatabaseUserOptions(
				deploymentID,
				userType,
				username,
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

			waitForTask(taskIDLink)
		})
		It(`CreateLogicalReplicationSlot request example`, func() {
			fmt.Println("\nCreateLogicalReplicationSlot() result:")
			// begin-createLogicalReplicationSlot

			logicalReplicationSlot := &clouddatabasesv5.LogicalReplicationSlot{
				Name:         &logicalRepName,
				DatabaseName: &databaseName,
				PluginType:   &pluginType,
			}

			createLogicalReplicationSlotOptions := &clouddatabasesv5.CreateLogicalReplicationSlotOptions{
				ID:                     &deploymentID,
				LogicalReplicationSlot: logicalReplicationSlot,
			}

			createLogicalReplicationResponse, response, err := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createLogicalReplicationResponse, "", "  ")
			fmt.Println(string(b))

			// end-createLogicalReplicationSlot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createLogicalReplicationResponse).ToNot(BeNil())

			taskIDLink = *createLogicalReplicationResponse.Task.ID

			waitForTask(taskIDLink)
		})
		It(`DeleteLogicalReplicationSlot request example`, func() {
			fmt.Println("\nDeleteLogicalReplicationSlot() result:")
			// begin-deleteLogicalReplicationSlot

			deleteLogicalReplicationSlotOptions := &clouddatabasesv5.DeleteLogicalReplicationSlotOptions{
				ID:   deploymentID,
				Name: logicalRepName,
			}

			deleteLogicalReplicationResponse, response, err := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteLogicalReplicationResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteLogicalReplicationSlot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteLogicalReplicationResponse).ToNot(BeNil())

			taskIDLink = *deleteLogicalReplicationResponse.Task.ID

			waitForTask(taskIDLink)
		})
		It(`KillConnections request example`, func() {
			fmt.Println("\nKillConnections() result:")
			// begin-killConnections

			killConnectionsOptions := cloudDatabasesService.NewKillConnectionsOptions(
				deploymentID,
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

			waitForTask(taskIDLink)
		})
		It(`SetAllowlist request example`, func() {
			fmt.Println("\nSetAllowlist() result:")
			// begin-setAllowlist

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     &IPAddress1,
				Description: core.StringPtr("Dev IP space 1"),
			}

			setAllowlistOptions := cloudDatabasesService.NewSetAllowlistOptions(
				deploymentID,
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

			waitForTask(taskIDLink)
		})
		It(`SetAutoscalingConditions request example`, func() {
			fmt.Println("\nSetAutoscalingConditions() result:")
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
				LimitMbPerMember: core.Float64Ptr(float64(114432)),
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
				deploymentID,
				autoScalingGroupID,
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

			waitForTask(taskIDLink)
		})
		It(`UpdateDatabaseConfiguration request example`, func() {
			fmt.Println("\nUpdateDatabaseConfiguration() result:")
			// begin-updateDatabaseConfiguration

			configurationModel := &clouddatabasesv5.ConfigurationPgConfiguration{
				MaxConnections: core.Int64Ptr(int64(200)),
			}

			updateDatabaseConfigurationOptions := cloudDatabasesService.NewUpdateDatabaseConfigurationOptions(
				deploymentID,
			)
			updateDatabaseConfigurationOptions.SetConfiguration(configurationModel)

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

			waitForTask(taskIDLink)
		})
		It(`ListDeployables request example`, func() {
			fmt.Println("\nListDeployables() result:")
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
			fmt.Println("\nListRegions() result:")
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
			fmt.Println("\nGetDeploymentInfo() result:")
			// begin-getDeploymentInfo

			getDeploymentInfoOptions := cloudDatabasesService.NewGetDeploymentInfoOptions(
				deploymentID,
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
			fmt.Println("\nListRemotes() result:")
			// begin-listRemotes

			listRemotesOptions := cloudDatabasesService.NewListRemotesOptions(
				deploymentID,
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
			fmt.Println("\nResyncReplica() result:")
			// begin-resyncReplica

			resyncReplicaOptions := cloudDatabasesService.NewResyncReplicaOptions(
				replicaID,
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
		It(`PromoteReadOnlyReplica request example`, func() {
			fmt.Println("\nPromoteReadOnlyReplica() result:")
			// begin-promoteReadOnlyReplica

			promotion := map[string]interface{}{
				"skip_initial_backup": false,
			}

			promoteReadOnlyReplicaOptions := cloudDatabasesService.NewPromoteReadOnlyReplicaOptions(
				replicaID,
			)
			promoteReadOnlyReplicaOptions.SetPromotion(promotion)

			promoteReadOnlyReplicaResponse, response, err := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(promoteReadOnlyReplicaResponse, "", "  ")
			fmt.Println(string(b))

			// end-promoteReadOnlyReplica

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(promoteReadOnlyReplicaResponse).ToNot(BeNil())

			taskIDLink = *promoteReadOnlyReplicaResponse.Task.ID

			waitForTask(taskIDLink)
		})
		It(`ListDeploymentTasks request example`, func() {
			fmt.Println("\nListDeploymentTasks() result:")
			// begin-listDeploymentTasks

			listDeploymentTasksOptions := cloudDatabasesService.NewListDeploymentTasksOptions(
				deploymentID,
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
			fmt.Println("\nGetTask() result:")
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
		It(`ListDeploymentBackups request example`, func() {
			fmt.Println("\nListDeploymentBackups() result:")
			// begin-listDeploymentBackups

			listDeploymentBackupsOptions := cloudDatabasesService.NewListDeploymentBackupsOptions(
				deploymentID,
			)

			backupsResponse, response, err := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(backupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listDeploymentBackups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupsResponse.Backups).ToNot(BeNil())

			backupIDLink = *backupsResponse.Backups[0].ID
		})
		It(`GetBackupInfo request example`, func() {
			fmt.Println("\nGetBackupInfo() result:")
			// begin-getBackupInfo

			getBackupInfoOptions := cloudDatabasesService.NewGetBackupInfoOptions(
				backupIDLink,
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
		It(`StartOndemandBackup request example`, func() {
			fmt.Println("\nStartOndemandBackup() result:")
			// begin-startOndemandBackup

			startOndemandBackupOptions := cloudDatabasesService.NewStartOndemandBackupOptions(
				deploymentID,
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
		It(`GetPitrData request example`, func() {
			fmt.Println("\nGetPitrData() result:")
			// begin-getPITRData

			getPitrDataOptions := cloudDatabasesService.NewGetPitrDataOptions(
				deploymentID,
			)

			getPitrDataResponse, response, err := cloudDatabasesService.GetPitrData(getPitrDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getPitrDataResponse, "", "  ")
			fmt.Println(string(b))

			// end-getPITRData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getPitrDataResponse.PointInTimeRecoveryData).ToNot(BeNil())
		})
		It(`GetConnection request example`, func() {
			fmt.Println("\nGetConnection() result:")
			// begin-getConnection

			getConnectionOptions := cloudDatabasesService.NewGetConnectionOptions(
				deploymentID,
				userType,
				"exampleUserID",
				"public",
			)
			getConnectionOptions.SetCertificateRoot("exampleCertRoot")

			getConnectionResponse, response, err := cloudDatabasesService.GetConnection(getConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getConnectionResponse, "", "  ")
			fmt.Println(string(b))

			// end-getConnection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getConnectionResponse).ToNot(BeNil())
		})
		It(`CompleteConnection request example`, func() {
			fmt.Println("\nCompleteConnection() result:")
			// begin-completeConnection

			completeConnectionOptions := cloudDatabasesService.NewCompleteConnectionOptions(
				deploymentID,
				userType,
				"exampleUserID",
				"public",
			)
			completeConnectionOptions.SetPassword(password)
			completeConnectionOptions.SetCertificateRoot("exampleCertRoot")

			completeConnectionResponse, response, err := cloudDatabasesService.CompleteConnection(completeConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(completeConnectionResponse, "", "  ")
			fmt.Println(string(b))

			// end-completeConnection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(completeConnectionResponse).ToNot(BeNil())
		})
		It(`ListDeploymentScalingGroups request example`, func() {
			fmt.Println("\nListDeploymentScalingGroups() result:")
			// begin-listDeploymentScalingGroups

			listDeploymentScalingGroupsOptions := cloudDatabasesService.NewListDeploymentScalingGroupsOptions(
				deploymentID,
			)

			groupsResponse, response, err := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(groupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listDeploymentScalingGroups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groupsResponse).ToNot(BeNil())

			scalingGroupIDLink = *groupsResponse.Groups[0].ID
		})
		It(`SetDeploymentScalingGroup request example`, func() {
			fmt.Println("\nSetDeploymentScalingGroup() result:")
			// begin-setDeploymentScalingGroup

			var handleError = func(resp *clouddatabasesv5.SetDeploymentScalingGroupResponse, r *core.DetailedResponse, err error) {
				groupScalingMemory := &clouddatabasesv5.GroupScalingMemory{
					AllocationMb: core.Int64Ptr(int64(114432)),
				}

				groupScaling := &clouddatabasesv5.GroupScaling{
					Memory: groupScalingMemory,
				}

				setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
					ID:      &deploymentID,
					GroupID: &scalingGroupIDLink,
					Group:   groupScaling,
				}

				setDeploymentScalingGroupResponse, response, err := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)
				if err != nil {
					fmt.Printf("\nError: %s", response.Result.(map[string]interface{}))
				}
				b, _ := json.MarshalIndent(setDeploymentScalingGroupResponse, "", "  ")
				fmt.Println(string(b))

				// end-setDeploymentScalingGroup

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))
				Expect(setDeploymentScalingGroupResponse).ToNot(BeNil())

				taskIDLink = *setDeploymentScalingGroupResponse.Task.ID

				waitForTask(taskIDLink)
			}

			groupScalingMemory := &clouddatabasesv5.GroupScalingMemory{
				AllocationMb: core.Int64Ptr(int64(114688)),
			}

			groupScaling := &clouddatabasesv5.GroupScaling{
				Memory: groupScalingMemory,
			}

			setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
				ID:      &deploymentID,
				GroupID: &scalingGroupIDLink,
				Group:   groupScaling,
			}

			setDeploymentScalingGroupResponse, response, err := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)
			if err != nil {
				handleError(setDeploymentScalingGroupResponse, response, err)
			}
			b, _ := json.MarshalIndent(setDeploymentScalingGroupResponse, "", "  ")
			fmt.Println(string(b))

			// end-setDeploymentScalingGroup
		})
		It(`GetDefaultScalingGroups request example`, func() {
			fmt.Println("\nGetDefaultScalingGroups() result:")
			// begin-getDefaultScalingGroups

			getDefaultScalingGroupsOptions := cloudDatabasesService.NewGetDefaultScalingGroupsOptions(
				"postgresql",
			)

			getDefaultScalingGroupsResponse, response, err := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getDefaultScalingGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-getDefaultScalingGroups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDefaultScalingGroupsResponse).ToNot(BeNil())
		})
		It(`GetAutoscalingConditions request example`, func() {
			fmt.Println("\nGetAutoscalingConditions() result:")
			// begin-getAutoscalingConditions

			getAutoscalingConditionsOptions := cloudDatabasesService.NewGetAutoscalingConditionsOptions(
				deploymentID,
				autoScalingGroupID,
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
			fmt.Println("\nGetAllowlist() result:")
			// begin-getAllowlist

			getAllowlistOptions := cloudDatabasesService.NewGetAllowlistOptions(
				deploymentID,
			)

			getAllowlistResponse, response, err := cloudDatabasesService.GetAllowlist(getAllowlistOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getAllowlistResponse, "", "  ")
			fmt.Println(string(b))

			// end-getAllowlist

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllowlistResponse).ToNot(BeNil())
		})

		It(`Capability request example`, func() {
			fmt.Println("\nCapability() result:")
			// begin-capability

			capabilityRequestDeploymentModel := &clouddatabasesv5.CapabilityRequestDeployment{
				Type:     core.StringPtr("postgresql"),
				Version:  core.StringPtr("10"),
				Platform: core.StringPtr("classic"),
				Location: core.StringPtr("us-south"),
			}

			capabilityOptions := cloudDatabasesService.NewCapabilityOptions(
				"autoscaling",
			)
			capabilityOptions.SetDeployment(capabilityRequestDeploymentModel)

			capabilityResponse, response, err := cloudDatabasesService.Capability(capabilityOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(capabilityResponse, "", "  ")
			fmt.Println(string(b))

			// end-capability

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capabilityResponse).ToNot(BeNil())
		})
	})
})
