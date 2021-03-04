// +build integration

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
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/icd-go-sdk/clouddatabasesv5"
)

/**
 * This file contains an integration test for the clouddatabasesv5 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CloudDatabasesV5 Integration Tests`, func() {

	const externalConfigFile = "../cloud_databases.env"

	var (
		err                   error
		cloudDatabasesService *clouddatabasesv5.CloudDatabasesV5
		serviceURL            string
		config                map[string]string
	)

	// Globlal variables to hold link values
	var (
		taskIDLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(clouddatabasesv5.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			cloudDatabasesServiceOptions := &clouddatabasesv5.CloudDatabasesV5Options{}

			cloudDatabasesService, err = clouddatabasesv5.NewCloudDatabasesV5UsingExternalConfig(cloudDatabasesServiceOptions)

			Expect(err).To(BeNil())
			Expect(cloudDatabasesService).ToNot(BeNil())
			Expect(cloudDatabasesService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`AddAllowlistEntry - Add an address or range to the allowlist for a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions)`, func() {

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			addAllowlistEntryOptions := &clouddatabasesv5.AddAllowlistEntryOptions{
				ID:        core.StringPtr("testString"),
				IPAddress: allowlistEntryModel,
			}

			addAllowlistEntryResponse, response, err := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(addAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *addAllowlistEntryResponse.Task.ID

		})
	})

	Describe(`CreateDatabaseUser - Creates a user based on user type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions)`, func() {

			createDatabaseUserRequestUserModel := &clouddatabasesv5.CreateDatabaseUserRequestUser{
				UserType: core.StringPtr("database"),
				Username: core.StringPtr("james"),
				Password: core.StringPtr("kickoutthe"),
			}

			createDatabaseUserOptions := &clouddatabasesv5.CreateDatabaseUserOptions{
				ID:       core.StringPtr("testString"),
				UserType: core.StringPtr("testString"),
				User:     createDatabaseUserRequestUserModel,
			}

			createDatabaseUserResponse, response, err := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *createDatabaseUserResponse.Task.ID

		})
	})

	Describe(`DeleteAllowlistEntry - Delete an address or range from the allowlist of a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions)`, func() {

			deleteAllowlistEntryOptions := &clouddatabasesv5.DeleteAllowlistEntryOptions{
				ID:        core.StringPtr("testString"),
				Ipaddress: core.StringPtr("testString"),
			}

			deleteAllowlistEntryResponse, response, err := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *deleteAllowlistEntryResponse.Task.ID

		})
	})

	Describe(`DeleteDatabaseUser - Deletes a user based on user type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions)`, func() {

			deleteDatabaseUserOptions := &clouddatabasesv5.DeleteDatabaseUserOptions{
				ID:       core.StringPtr("testString"),
				UserType: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
			}

			deleteDatabaseUserResponse, response, err := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *deleteDatabaseUserResponse.Task.ID

		})
	})

	Describe(`ReplaceAllowlist - Replace the allowlist for a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceAllowlist(replaceAllowlistOptions *ReplaceAllowlistOptions)`, func() {

			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address:     core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
			}

			replaceAllowlistOptions := &clouddatabasesv5.ReplaceAllowlistOptions{
				ID:          core.StringPtr("testString"),
				IPAddresses: []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel},
				IfMatch:     core.StringPtr("testString"),
			}

			replaceAllowlistResponse, response, err := cloudDatabasesService.ReplaceAllowlist(replaceAllowlistOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(replaceAllowlistResponse).ToNot(BeNil())

			taskIDLink = *replaceAllowlistResponse.Task.ID

		})
	})

	Describe(`SetDeploymentScalingGroup - Set scaling values on a specified group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions)`, func() {

			setMembersGroupMembersModel := &clouddatabasesv5.SetMembersGroupMembers{
				AllocationCount: core.Int64Ptr(int64(4)),
			}

			setDeploymentScalingGroupRequestModel := &clouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup{
				Members: setMembersGroupMembersModel,
			}

			setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
				ID:                               core.StringPtr("testString"),
				GroupID:                          core.StringPtr("testString"),
				SetDeploymentScalingGroupRequest: setDeploymentScalingGroupRequestModel,
			}

			setDeploymentScalingGroupResponse, response, err := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setDeploymentScalingGroupResponse).ToNot(BeNil())

			taskIDLink = *setDeploymentScalingGroupResponse.Task.ID

		})
	})

	Describe(`ListDeployables - List all deployable databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeployables(listDeployablesOptions *ListDeployablesOptions)`, func() {

			listDeployablesOptions := &clouddatabasesv5.ListDeployablesOptions{}

			listDeployablesResponse, response, err := cloudDatabasesService.ListDeployables(listDeployablesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDeployablesResponse).ToNot(BeNil())

		})
	})

	Describe(`ListRegions - List all deployable regions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRegions(listRegionsOptions *ListRegionsOptions)`, func() {

			listRegionsOptions := &clouddatabasesv5.ListRegionsOptions{}

			listRegionsResponse, response, err := cloudDatabasesService.ListRegions(listRegionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRegionsResponse).ToNot(BeNil())

		})
	})

	Describe(`GetDeploymentInfo - Get deployment information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions)`, func() {

			getDeploymentInfoOptions := &clouddatabasesv5.GetDeploymentInfoOptions{
				ID: core.StringPtr("testString"),
			}

			getDeploymentInfoResponse, response, err := cloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDeploymentInfoResponse).ToNot(BeNil())

		})
	})

	Describe(`ChangeUserPassword - Set specified user's password`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions)`, func() {

			aPasswordSettingUserModel := &clouddatabasesv5.APasswordSettingUser{
				Password: core.StringPtr("xyzzy"),
			}

			changeUserPasswordOptions := &clouddatabasesv5.ChangeUserPasswordOptions{
				ID:       core.StringPtr("testString"),
				UserType: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
				User:     aPasswordSettingUserModel,
			}

			changeUserPasswordResponse, response, err := cloudDatabasesService.ChangeUserPassword(changeUserPasswordOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(changeUserPasswordResponse).ToNot(BeNil())

		})
	})

	Describe(`GetUser - Discover user name and password information for a deployment for a user with an endpoint type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUser(getUserOptions *GetUserOptions)`, func() {

			getUserOptions := &clouddatabasesv5.GetUserOptions{
				ID:     core.StringPtr("testString"),
				UserID: core.StringPtr("testString"),
			}

			task, response, err := cloudDatabasesService.GetUser(getUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())

		})
	})

	Describe(`SetDatabaseConfiguration - Change your database configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions)`, func() {

			setConfigurationConfigurationModel := &clouddatabasesv5.SetConfigurationConfigurationPgConfiguration{
				MaxConnections:          core.Int64Ptr(int64(115)),
				MaxPreparedTransactions: core.Int64Ptr(int64(0)),
				DeadlockTimeout:         core.Int64Ptr(int64(100)),
				EffectiveIoConcurrency:  core.Int64Ptr(int64(1)),
				MaxReplicationSlots:     core.Int64Ptr(int64(10)),
				MaxWalSenders:           core.Int64Ptr(int64(12)),
				SharedBuffers:           core.Int64Ptr(int64(16)),
				SynchronousCommit:       core.StringPtr("local"),
				WalLevel:                core.StringPtr("hot_standby"),
				ArchiveTimeout:          core.Int64Ptr(int64(300)),
				LogMinDurationStatement: core.Int64Ptr(int64(100)),
			}

			setDatabaseConfigurationOptions := &clouddatabasesv5.SetDatabaseConfigurationOptions{
				ID:            core.StringPtr("testString"),
				Configuration: setConfigurationConfigurationModel,
			}

			setDatabaseConfigurationResponse, response, err := cloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setDatabaseConfigurationResponse).ToNot(BeNil())

		})
	})

	Describe(`GetDatabaseConfigurationSchema - Get the schema of the database configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions)`, func() {

			getDatabaseConfigurationSchemaOptions := &clouddatabasesv5.GetDatabaseConfigurationSchemaOptions{
				ID: core.StringPtr("testString"),
			}

			configurationSchema, response, err := cloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configurationSchema).ToNot(BeNil())

		})
	})

	Describe(`ListRemotes - List read-only replica information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRemotes(listRemotesOptions *ListRemotesOptions)`, func() {

			listRemotesOptions := &clouddatabasesv5.ListRemotesOptions{
				ID: core.StringPtr("testString"),
			}

			listRemotesResponse, response, err := cloudDatabasesService.ListRemotes(listRemotesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listRemotesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetRemotesSchema - Resync read-only replica`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions)`, func() {

			getRemotesSchemaOptions := &clouddatabasesv5.GetRemotesSchemaOptions{
				ID: core.StringPtr("testString"),
			}

			getRemotesSchemaResponse, response, err := cloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getRemotesSchemaResponse).ToNot(BeNil())

		})
	})

	Describe(`SetPromotion - Promote read-only replica to a full deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetPromotion(setPromotionOptions *SetPromotionOptions)`, func() {

			setPromotionPromotionModel := &clouddatabasesv5.SetPromotionPromotionPromote{
				Promotion: make(map[string]interface{}),
			}

			setPromotionOptions := &clouddatabasesv5.SetPromotionOptions{
				ID:        core.StringPtr("testString"),
				Promotion: setPromotionPromotionModel,
			}

			setPromotionResponse, response, err := cloudDatabasesService.SetPromotion(setPromotionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setPromotionResponse).ToNot(BeNil())

		})
	})

	Describe(`ListDeploymentTasks - List currently running tasks on a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeploymentTasks(listDeploymentTasksOptions *ListDeploymentTasksOptions)`, func() {

			listDeploymentTasksOptions := &clouddatabasesv5.ListDeploymentTasksOptions{
				ID: core.StringPtr("testString"),
			}

			tasks, response, err := cloudDatabasesService.ListDeploymentTasks(listDeploymentTasksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tasks).ToNot(BeNil())

		})
	})

	Describe(`GetTask - Get information about a task`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTask(getTaskOptions *GetTaskOptions)`, func() {

			getTaskOptions := &clouddatabasesv5.GetTaskOptions{
				ID: &taskIDLink,
			}

			getTaskResponse, response, err := cloudDatabasesService.GetTask(getTaskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTaskResponse).ToNot(BeNil())

		})
	})

	Describe(`GetBackupInfo - Get information about a backup`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions)`, func() {

			getBackupInfoOptions := &clouddatabasesv5.GetBackupInfoOptions{
				BackupID: core.StringPtr("testString"),
			}

			getBackupInfoResponse, response, err := cloudDatabasesService.GetBackupInfo(getBackupInfoOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBackupInfoResponse).ToNot(BeNil())

		})
	})

	Describe(`ListDeploymentBackups - List currently available backups from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeploymentBackups(listDeploymentBackupsOptions *ListDeploymentBackupsOptions)`, func() {

			listDeploymentBackupsOptions := &clouddatabasesv5.ListDeploymentBackupsOptions{
				ID: core.StringPtr("testString"),
			}

			backups, response, err := cloudDatabasesService.ListDeploymentBackups(listDeploymentBackupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backups).ToNot(BeNil())

		})
	})

	Describe(`StartOndemandBackup - Initiate an on-demand backup`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions)`, func() {

			startOndemandBackupOptions := &clouddatabasesv5.StartOndemandBackupOptions{
				ID: core.StringPtr("testString"),
			}

			startOndemandBackupResponse, response, err := cloudDatabasesService.StartOndemandBackup(startOndemandBackupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(startOndemandBackupResponse).ToNot(BeNil())

		})
	})

	Describe(`GetPitRdata - Get earliest point-in-time-recovery timestamp`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPitRdata(getPitRdataOptions *GetPitRdataOptions)`, func() {

			getPitRdataOptions := &clouddatabasesv5.GetPitRdataOptions{
				ID: core.StringPtr("testString"),
			}

			pointInTimeRecoveryData, response, err := cloudDatabasesService.GetPitRdata(getPitRdataOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pointInTimeRecoveryData).ToNot(BeNil())

		})
	})

	Describe(`GetConnection - Discover connection information for a deployment for a user with an endpoint type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConnection(getConnectionOptions *GetConnectionOptions)`, func() {

			getConnectionOptions := &clouddatabasesv5.GetConnectionOptions{
				ID:              core.StringPtr("testString"),
				UserType:        core.StringPtr("testString"),
				UserID:          core.StringPtr("testString"),
				EndpointType:    core.StringPtr("public"),
				CertificateRoot: core.StringPtr("testString"),
			}

			connection, response, err := cloudDatabasesService.GetConnection(getConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connection).ToNot(BeNil())

		})
	})

	Describe(`CompleteConnection - Discover connection information for a deployment for a user with substitutions and an endpoint type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions)`, func() {

			completeConnectionOptions := &clouddatabasesv5.CompleteConnectionOptions{
				ID:              core.StringPtr("testString"),
				UserType:        core.StringPtr("testString"),
				UserID:          core.StringPtr("testString"),
				EndpointType:    core.StringPtr("public"),
				Password:        core.StringPtr("testString"),
				CertificateRoot: core.StringPtr("testString"),
			}

			connection, response, err := cloudDatabasesService.CompleteConnection(completeConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connection).ToNot(BeNil())

		})
	})

	Describe(`ListDeploymentScalingGroups - List currently available scaling groups from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions *ListDeploymentScalingGroupsOptions)`, func() {

			listDeploymentScalingGroupsOptions := &clouddatabasesv5.ListDeploymentScalingGroupsOptions{
				ID: core.StringPtr("testString"),
			}

			groups, response, err := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groups).ToNot(BeNil())

		})
	})

	Describe(`GetDefaultScalingGroups - Get default scaling groups for a new deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions)`, func() {

			getDefaultScalingGroupsOptions := &clouddatabasesv5.GetDefaultScalingGroupsOptions{
				Type: core.StringPtr("postgresql"),
			}

			groups, response, err := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(groups).ToNot(BeNil())

		})
	})

	Describe(`GetAutoscalingConditions - Get the autoscaling configuration from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions)`, func() {

			getAutoscalingConditionsOptions := &clouddatabasesv5.GetAutoscalingConditionsOptions{
				ID:      core.StringPtr("testString"),
				GroupID: core.StringPtr("testString"),
			}

			autoscalingGroup, response, err := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(autoscalingGroup).ToNot(BeNil())

		})
	})

	Describe(`SetAutoscalingConditions - Set the autoscaling configuration from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions)`, func() {

			autoscalingDiskGroupDiskScalersCapacityModel := &clouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity{
				Enabled:                  core.BoolPtr(true),
				FreeSpaceLessThanPercent: core.Int64Ptr(int64(10)),
			}

			autoscalingDiskGroupDiskScalersIoUtilizationModel := &clouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization{
				Enabled:      core.BoolPtr(true),
				OverPeriod:   core.StringPtr("30m"),
				AbovePercent: core.Int64Ptr(int64(45)),
			}

			autoscalingDiskGroupDiskScalersModel := &clouddatabasesv5.AutoscalingDiskGroupDiskScalers{
				Capacity:      autoscalingDiskGroupDiskScalersCapacityModel,
				IoUtilization: autoscalingDiskGroupDiskScalersIoUtilizationModel,
			}

			autoscalingDiskGroupDiskRateModel := &clouddatabasesv5.AutoscalingDiskGroupDiskRate{
				IncreasePercent:  core.Float64Ptr(float64(20)),
				PeriodSeconds:    core.Int64Ptr(int64(900)),
				LimitMbPerMember: core.Float64Ptr(float64(3670016)),
				Units:            core.StringPtr("mb"),
			}

			autoscalingDiskGroupDiskModel := &clouddatabasesv5.AutoscalingDiskGroupDisk{
				Scalers: autoscalingDiskGroupDiskScalersModel,
				Rate:    autoscalingDiskGroupDiskRateModel,
			}

			autoscalingSetGroupAutoscalingModel := &clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingDiskGroup{
				Disk: autoscalingDiskGroupDiskModel,
			}

			setAutoscalingConditionsOptions := &clouddatabasesv5.SetAutoscalingConditionsOptions{
				ID:          core.StringPtr("testString"),
				GroupID:     core.StringPtr("testString"),
				Autoscaling: autoscalingSetGroupAutoscalingModel,
			}

			setAutoscalingConditionsResponse, response, err := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setAutoscalingConditionsResponse).ToNot(BeNil())

		})
	})

	Describe(`FileSync - Sync files uploaded to Elasticsearch deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`FileSync(fileSyncOptions *FileSyncOptions)`, func() {

			fileSyncOptions := &clouddatabasesv5.FileSyncOptions{
				ID: core.StringPtr("testString"),
			}

			fileSyncResponse, response, err := cloudDatabasesService.FileSync(fileSyncOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(fileSyncResponse).ToNot(BeNil())

		})
	})

	Describe(`CreateLogicalReplicationSlot - Create a new logical replication slot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions)`, func() {

			logicalReplicationSlotLogicalReplicationSlotModel := &clouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot{
				Name:         core.StringPtr("customer_replication"),
				DatabaseName: core.StringPtr("customers"),
				PluginType:   core.StringPtr("wal2json"),
			}

			createLogicalReplicationSlotOptions := &clouddatabasesv5.CreateLogicalReplicationSlotOptions{
				ID:                     core.StringPtr("testString"),
				LogicalReplicationSlot: logicalReplicationSlotLogicalReplicationSlotModel,
			}

			createLogicalReplicationSlotResponse, response, err := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createLogicalReplicationSlotResponse).ToNot(BeNil())

		})
	})

	Describe(`GetAllowlist - Retrieve the allowlisted addresses and ranges for a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllowlist(getAllowlistOptions *GetAllowlistOptions)`, func() {

			getAllowlistOptions := &clouddatabasesv5.GetAllowlistOptions{
				ID: core.StringPtr("testString"),
			}

			allowlist, response, err := cloudDatabasesService.GetAllowlist(getAllowlistOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(allowlist).ToNot(BeNil())

		})
	})

	Describe(`KillConnections - Kill connections to a PostgreSQL or EnterpriseDB deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`KillConnections(killConnectionsOptions *KillConnectionsOptions)`, func() {

			killConnectionsOptions := &clouddatabasesv5.KillConnectionsOptions{
				ID: core.StringPtr("testString"),
			}

			killConnectionsResponse, response, err := cloudDatabasesService.KillConnections(killConnectionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(killConnectionsResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteLogicalReplicationSlot - Delete a logical replication slot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions)`, func() {

			deleteLogicalReplicationSlotOptions := &clouddatabasesv5.DeleteLogicalReplicationSlotOptions{
				ID:   core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			deleteLogicalReplicationSlotResponse, response, err := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteLogicalReplicationSlotResponse).ToNot(BeNil())

		})
	})

	Describe(`DeleteDatabaseUser - Deletes a user based on user type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions)`, func() {

			deleteDatabaseUserOptions := &clouddatabasesv5.DeleteDatabaseUserOptions{
				ID:       core.StringPtr("testString"),
				UserType: core.StringPtr("testString"),
				Username: core.StringPtr("testString"),
			}

			deleteDatabaseUserResponse, response, err := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *deleteDatabaseUserResponse.Task.ID

		})
	})

	Describe(`DeleteAllowlistEntry - Delete an address or range from the allowlist of a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions)`, func() {

			deleteAllowlistEntryOptions := &clouddatabasesv5.DeleteAllowlistEntryOptions{
				ID:        core.StringPtr("testString"),
				Ipaddress: core.StringPtr("testString"),
			}

			deleteAllowlistEntryResponse, response, err := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *deleteAllowlistEntryResponse.Task.ID

		})
	})
})

//
// Utility functions are declared in the unit test file
//
