//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
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
	"log"
	"os"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the clouddatabasesv5 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`CloudDatabasesV5 Integration Tests`, func() {
	const externalConfigFile = "../cloud_databases_v5.env"

	var (
		err          error
		cloudDatabasesService *clouddatabasesv5.CloudDatabasesV5
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
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

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
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

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			cloudDatabasesService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateDatabaseUser - Creates a user based on user type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions)`, func() {
			userModel := &clouddatabasesv5.UserDatabaseUser{
				Username: core.StringPtr("user"),
				Password: core.StringPtr("v3ry-1-secUre-pAssword-2"),
			}

			createDatabaseUserOptions := &clouddatabasesv5.CreateDatabaseUserOptions{
				ID: core.StringPtr("testString"),
				UserType: core.StringPtr("testString"),
				User: userModel,
			}

			createDatabaseUserResponse, response, err := cloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *createDatabaseUserResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`UpdateUser - Update a user's password or role`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateUser(updateUserOptions *UpdateUserOptions)`, func() {
			userUpdateModel := &clouddatabasesv5.UserUpdatePasswordSetting{
				Password: core.StringPtr("v3ry-1-secUre-pAssword-2"),
			}

			updateUserOptions := &clouddatabasesv5.UpdateUserOptions{
				ID: core.StringPtr("testString"),
				UserType: core.StringPtr("database"),
				Username: core.StringPtr("user"),
				User: userUpdateModel,
			}

			updateUserResponse, response, err := cloudDatabasesService.UpdateUser(updateUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updateUserResponse).ToNot(BeNil())

			taskIDLink = *updateUserResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`DeleteDatabaseUser - Deletes a user based on user type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions)`, func() {
			deleteDatabaseUserOptions := &clouddatabasesv5.DeleteDatabaseUserOptions{
				ID: core.StringPtr("testString"),
				UserType: core.StringPtr("database"),
				Username: core.StringPtr("user"),
			}

			deleteDatabaseUserResponse, response, err := cloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteDatabaseUserResponse).ToNot(BeNil())

			taskIDLink = *deleteDatabaseUserResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`UpdateDatabaseConfiguration - Change your database configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions *UpdateDatabaseConfigurationOptions)`, func() {
			configurationModel := &clouddatabasesv5.ConfigurationPgConfiguration{
				ArchiveTimeout: core.Int64Ptr(int64(300)),
				DeadlockTimeout: core.Int64Ptr(int64(100)),
				EffectiveIoConcurrency: core.Int64Ptr(int64(1)),
				LogConnections: core.StringPtr("off"),
				LogDisconnections: core.StringPtr("off"),
				LogMinDurationStatement: core.Int64Ptr(int64(100)),
				MaxConnections: core.Int64Ptr(int64(200)),
				MaxLocksPerTransaction: core.Int64Ptr(int64(10)),
				MaxPreparedTransactions: core.Int64Ptr(int64(0)),
				MaxReplicationSlots: core.Int64Ptr(int64(10)),
				MaxWalSenders: core.Int64Ptr(int64(12)),
				SharedBuffers: core.Int64Ptr(int64(16)),
				SynchronousCommit: core.StringPtr("local"),
				TCPKeepalivesCount: core.Int64Ptr(int64(0)),
				TCPKeepalivesIdle: core.Int64Ptr(int64(0)),
				TCPKeepalivesInterval: core.Int64Ptr(int64(0)),
				WalLevel: core.StringPtr("replica"),
			}

			updateDatabaseConfigurationOptions := &clouddatabasesv5.UpdateDatabaseConfigurationOptions{
				ID: core.StringPtr("testString"),
				Configuration: configurationModel,
			}

			updateDatabaseConfigurationResponse, response, err := cloudDatabasesService.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateDatabaseConfigurationResponse).ToNot(BeNil())

			taskIDLink = *updateDatabaseConfigurationResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`SetDeploymentScalingGroup - Set scaling values on a specified group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions)`, func() {
			groupScalingMembersModel := &clouddatabasesv5.GroupScalingMembers{
				AllocationCount: core.Int64Ptr(int64(4)),
			}

			groupScalingMemoryModel := &clouddatabasesv5.GroupScalingMemory{
				AllocationMb: core.Int64Ptr(int64(12288)),
			}

			groupScalingCPUModel := &clouddatabasesv5.GroupScalingCPU{
				AllocationCount: core.Int64Ptr(int64(2)),
			}

			groupScalingDiskModel := &clouddatabasesv5.GroupScalingDisk{
				AllocationMb: core.Int64Ptr(int64(20480)),
			}

			groupScalingHostFlavorModel := &clouddatabasesv5.GroupScalingHostFlavor{
				ID: core.StringPtr("b3c.16x64.encrypted"),
			}

			groupScalingModel := &clouddatabasesv5.GroupScaling{
				Members: groupScalingMembersModel,
				Memory: groupScalingMemoryModel,
				CPU: groupScalingCPUModel,
				Disk: groupScalingDiskModel,
				HostFlavor: groupScalingHostFlavorModel,
			}

			setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
				ID: core.StringPtr("testString"),
				GroupID: core.StringPtr("testString"),
				Group: groupScalingModel,
			}

			setDeploymentScalingGroupResponse, response, err := cloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setDeploymentScalingGroupResponse).ToNot(BeNil())

			taskIDLink = *setDeploymentScalingGroupResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`SetAutoscalingConditions - Set the autoscaling configuration from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions)`, func() {
			autoscalingMemoryGroupMemoryScalersIoUtilizationModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryScalersIoUtilization{
				Enabled: core.BoolPtr(true),
				OverPeriod: core.StringPtr("5m"),
				AbovePercent: core.Int64Ptr(int64(90)),
			}

			autoscalingMemoryGroupMemoryScalersModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryScalers{
				IoUtilization: autoscalingMemoryGroupMemoryScalersIoUtilizationModel,
			}

			autoscalingMemoryGroupMemoryRateModel := &clouddatabasesv5.AutoscalingMemoryGroupMemoryRate{
				IncreasePercent: core.Float64Ptr(float64(10)),
				PeriodSeconds: core.Int64Ptr(int64(300)),
				LimitMbPerMember: core.Float64Ptr(float64(125952)),
				Units: core.StringPtr("mb"),
			}

			autoscalingMemoryGroupMemoryModel := &clouddatabasesv5.AutoscalingMemoryGroupMemory{
				Scalers: autoscalingMemoryGroupMemoryScalersModel,
				Rate: autoscalingMemoryGroupMemoryRateModel,
			}

			autoscalingSetGroupAutoscalingModel := &clouddatabasesv5.AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup{
				Memory: autoscalingMemoryGroupMemoryModel,
			}

			setAutoscalingConditionsOptions := &clouddatabasesv5.SetAutoscalingConditionsOptions{
				ID: core.StringPtr("testString"),
				GroupID: core.StringPtr("testString"),
				Autoscaling: autoscalingSetGroupAutoscalingModel,
			}

			setAutoscalingConditionsResponse, response, err := cloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setAutoscalingConditionsResponse).ToNot(BeNil())

			taskIDLink = *setAutoscalingConditionsResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
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
			Expect(response.StatusCode).To(Equal(202))
			Expect(killConnectionsResponse).ToNot(BeNil())

			taskIDLink = *killConnectionsResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`CreateLogicalReplicationSlot - Create a new logical replication slot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions)`, func() {
			logicalReplicationSlotModel := &clouddatabasesv5.LogicalReplicationSlot{
				Name: core.StringPtr("customer_replication"),
				DatabaseName: core.StringPtr("customers"),
				PluginType: core.StringPtr("wal2json"),
			}

			createLogicalReplicationSlotOptions := &clouddatabasesv5.CreateLogicalReplicationSlotOptions{
				ID: core.StringPtr("testString"),
				LogicalReplicationSlot: logicalReplicationSlotModel,
			}

			createLogicalReplicationSlotResponse, response, err := cloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createLogicalReplicationSlotResponse).ToNot(BeNil())

			taskIDLink = *createLogicalReplicationSlotResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`DeleteLogicalReplicationSlot - Delete a logical replication slot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions)`, func() {
			deleteLogicalReplicationSlotOptions := &clouddatabasesv5.DeleteLogicalReplicationSlotOptions{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			deleteLogicalReplicationSlotResponse, response, err := cloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteLogicalReplicationSlotResponse).ToNot(BeNil())

			taskIDLink = *deleteLogicalReplicationSlotResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`SetAllowlist - Set the allowlist for a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetAllowlist(setAllowlistOptions *SetAllowlistOptions)`, func() {
			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address: core.StringPtr("195.212.0.0/16"),
				Description: core.StringPtr("Dev IP space 1"),
			}

			setAllowlistOptions := &clouddatabasesv5.SetAllowlistOptions{
				ID: core.StringPtr("testString"),
				IPAddresses: []clouddatabasesv5.AllowlistEntry{*allowlistEntryModel},
				IfMatch: core.StringPtr("testString"),
			}

			setAllowlistResponse, response, err := cloudDatabasesService.SetAllowlist(setAllowlistOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(setAllowlistResponse).ToNot(BeNil())

			taskIDLink = *setAllowlistResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`AddAllowlistEntry - Add an address or range to the allowlist for a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddAllowlistEntry(addAllowlistEntryOptions *AddAllowlistEntryOptions)`, func() {
			allowlistEntryModel := &clouddatabasesv5.AllowlistEntry{
				Address: core.StringPtr("172.16.0.0/16"),
				Description: core.StringPtr("Dev IP space 3"),
			}

			addAllowlistEntryOptions := &clouddatabasesv5.AddAllowlistEntryOptions{
				ID: core.StringPtr("testString"),
				IPAddress: allowlistEntryModel,
			}

			addAllowlistEntryResponse, response, err := cloudDatabasesService.AddAllowlistEntry(addAllowlistEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(addAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *addAllowlistEntryResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`DeleteAllowlistEntry - Delete an address or range from the allowlist of a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAllowlistEntry(deleteAllowlistEntryOptions *DeleteAllowlistEntryOptions)`, func() {
			deleteAllowlistEntryOptions := &clouddatabasesv5.DeleteAllowlistEntryOptions{
				ID: core.StringPtr("testString"),
				Ipaddress: core.StringPtr("testString"),
			}

			deleteAllowlistEntryResponse, response, err := cloudDatabasesService.DeleteAllowlistEntry(deleteAllowlistEntryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(deleteAllowlistEntryResponse).ToNot(BeNil())

			taskIDLink = *deleteAllowlistEntryResponse.Task.ID
			fmt.Fprintf(GinkgoWriter, "Saved taskIDLink value: %v\n", taskIDLink)
		})
	})

	Describe(`ListDeployables - List all deployable databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeployables(listDeployablesOptions *ListDeployablesOptions)`, func() {
			listDeployablesOptions := &clouddatabasesv5.ListDeployablesOptions{
			}

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
			listRegionsOptions := &clouddatabasesv5.ListRegionsOptions{
			}

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

	Describe(`ResyncReplica - Resync read-only replica`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ResyncReplica(resyncReplicaOptions *ResyncReplicaOptions)`, func() {
			resyncReplicaOptions := &clouddatabasesv5.ResyncReplicaOptions{
				ID: core.StringPtr("testString"),
			}

			resyncReplicaResponse, response, err := cloudDatabasesService.ResyncReplica(resyncReplicaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resyncReplicaResponse).ToNot(BeNil())
		})
	})

	Describe(`PromoteReadOnlyReplica - Promote read-only replica to a full deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions *PromoteReadOnlyReplicaOptions)`, func() {
			promoteReadOnlyReplicaOptions := &clouddatabasesv5.PromoteReadOnlyReplicaOptions{
				ID: core.StringPtr("testString"),
				Promotion: map[string]interface{}{"anyKey": "anyValue"},
			}

			promoteReadOnlyReplicaResponse, response, err := cloudDatabasesService.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(promoteReadOnlyReplicaResponse).ToNot(BeNil())
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

	Describe(`GetPitrData - Get earliest point-in-time-recovery timestamp`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPitrData(getPitrDataOptions *GetPitrDataOptions)`, func() {
			getPitrDataOptions := &clouddatabasesv5.GetPitrDataOptions{
				ID: core.StringPtr("testString"),
			}

			getPitrDataResponse, response, err := cloudDatabasesService.GetPitrData(getPitrDataOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getPitrDataResponse).ToNot(BeNil())
		})
	})

	Describe(`GetConnection - Discover connection information for a deployment for a user with an endpoint type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConnection(getConnectionOptions *GetConnectionOptions)`, func() {
			getConnectionOptions := &clouddatabasesv5.GetConnectionOptions{
				ID: core.StringPtr("testString"),
				UserType: core.StringPtr("database"),
				UserID: core.StringPtr("testString"),
				EndpointType: core.StringPtr("public"),
				CertificateRoot: core.StringPtr("testString"),
			}

			getConnectionResponse, response, err := cloudDatabasesService.GetConnection(getConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getConnectionResponse).ToNot(BeNil())
		})
	})

	Describe(`CompleteConnection - Discover connection information for a deployment for a user with substitutions and an endpoint type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions)`, func() {
			completeConnectionOptions := &clouddatabasesv5.CompleteConnectionOptions{
				ID: core.StringPtr("testString"),
				UserType: core.StringPtr("database"),
				UserID: core.StringPtr("testString"),
				EndpointType: core.StringPtr("public"),
				Password: core.StringPtr("providedpassword"),
				CertificateRoot: core.StringPtr("testString"),
			}

			completeConnectionResponse, response, err := cloudDatabasesService.CompleteConnection(completeConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(completeConnectionResponse).ToNot(BeNil())
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

			listDeploymentScalingGroupsResponse, response, err := cloudDatabasesService.ListDeploymentScalingGroups(listDeploymentScalingGroupsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDeploymentScalingGroupsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDefaultScalingGroups - Get default scaling groups for a new deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions)`, func() {
			getDefaultScalingGroupsOptions := &clouddatabasesv5.GetDefaultScalingGroupsOptions{
				Type: core.StringPtr("postgresql"),
				HostFlavor: core.StringPtr("multitenant"),
			}

			getDefaultScalingGroupsResponse, response, err := cloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDefaultScalingGroupsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetAutoscalingConditions - Get the autoscaling configuration from a deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions)`, func() {
			getAutoscalingConditionsOptions := &clouddatabasesv5.GetAutoscalingConditionsOptions{
				ID: core.StringPtr("testString"),
				GroupID: core.StringPtr("testString"),
			}

			autoscalingGroup, response, err := cloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(autoscalingGroup).ToNot(BeNil())
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

			getAllowlistResponse, response, err := cloudDatabasesService.GetAllowlist(getAllowlistOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getAllowlistResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
