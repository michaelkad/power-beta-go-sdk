// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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

package powervsv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/michaelkad/power-beta-go-sdk/powervsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the powervsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PowervsV1 Integration Tests`, func() {
	const externalConfigFile = "../powervs_v1.env"

	var (
		err          error
		powervsService *powervsv1.PowervsV1
		serviceURL   string
		config       map[string]string
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
			config, err = core.GetServiceProperties(powervsv1.DefaultServiceName)
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
			powervsServiceOptions := &powervsv1.PowervsV1Options{}

			powervsService, err = powervsv1.NewPowervsV1UsingExternalConfig(powervsServiceOptions)
			Expect(err).To(BeNil())
			Expect(powervsService).ToNot(BeNil())
			Expect(powervsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			powervsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ServiceBrokerAuthCallback - Returns an accessToken (and set cookie)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthCallback(serviceBrokerAuthCallbackOptions *ServiceBrokerAuthCallbackOptions)`, func() {
			serviceBrokerAuthCallbackOptions := &powervsv1.ServiceBrokerAuthCallbackOptions{
			}

			accessToken, response, err := powervsService.ServiceBrokerAuthCallback(serviceBrokerAuthCallbackOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessToken).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthRegistrationCallback - Associates the user with a tenant and returns an accessToken`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthRegistrationCallback(serviceBrokerAuthRegistrationCallbackOptions *ServiceBrokerAuthRegistrationCallbackOptions)`, func() {
			serviceBrokerAuthRegistrationCallbackOptions := &powervsv1.ServiceBrokerAuthRegistrationCallbackOptions{
			}

			accessToken, response, err := powervsService.ServiceBrokerAuthRegistrationCallback(serviceBrokerAuthRegistrationCallbackOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessToken).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthDeviceCodePost - Request a authorization device code`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthDeviceCodePost(serviceBrokerAuthDeviceCodePostOptions *ServiceBrokerAuthDeviceCodePostOptions)`, func() {
			serviceBrokerAuthDeviceCodePostOptions := &powervsv1.ServiceBrokerAuthDeviceCodePostOptions{
			}

			deviceCode, response, err := powervsService.ServiceBrokerAuthDeviceCodePost(serviceBrokerAuthDeviceCodePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deviceCode).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthDeviceTokenPost - Poll for authorization device token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthDeviceTokenPost(serviceBrokerAuthDeviceTokenPostOptions *ServiceBrokerAuthDeviceTokenPostOptions)`, func() {
			serviceBrokerAuthDeviceTokenPostOptions := &powervsv1.ServiceBrokerAuthDeviceTokenPostOptions{
				DeviceCode: core.StringPtr("testString"),
			}

			token, response, err := powervsService.ServiceBrokerAuthDeviceTokenPost(serviceBrokerAuthDeviceTokenPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(token).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthInfoToken - Information about current access token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthInfoToken(serviceBrokerAuthInfoTokenOptions *ServiceBrokerAuthInfoTokenOptions)`, func() {
			serviceBrokerAuthInfoTokenOptions := &powervsv1.ServiceBrokerAuthInfoTokenOptions{
			}

			tokenExtra, response, err := powervsService.ServiceBrokerAuthInfoToken(serviceBrokerAuthInfoTokenOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tokenExtra).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthInfoUser - Information about current user`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthInfoUser(serviceBrokerAuthInfoUserOptions *ServiceBrokerAuthInfoUserOptions)`, func() {
			serviceBrokerAuthInfoUserOptions := &powervsv1.ServiceBrokerAuthInfoUserOptions{
			}

			userInfo, response, err := powervsService.ServiceBrokerAuthInfoUser(serviceBrokerAuthInfoUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userInfo).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthLogin - Login`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthLogin(serviceBrokerAuthLoginOptions *ServiceBrokerAuthLoginOptions)`, func() {
			serviceBrokerAuthLoginOptions := &powervsv1.ServiceBrokerAuthLoginOptions{
				UserID: core.StringPtr("testString"),
				RedirectURL: core.StringPtr("testString"),
				AccessType: core.StringPtr("online"),
			}

			accessToken, response, err := powervsService.ServiceBrokerAuthLogin(serviceBrokerAuthLoginOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessToken).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthLogout - Logout`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthLogout(serviceBrokerAuthLogoutOptions *ServiceBrokerAuthLogoutOptions)`, func() {
			serviceBrokerAuthLogoutOptions := &powervsv1.ServiceBrokerAuthLogoutOptions{
			}

			object, response, err := powervsService.ServiceBrokerAuthLogout(serviceBrokerAuthLogoutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthRegistration - Registration of a new Tenant and Login`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthRegistration(serviceBrokerAuthRegistrationOptions *ServiceBrokerAuthRegistrationOptions)`, func() {
			serviceBrokerAuthRegistrationOptions := &powervsv1.ServiceBrokerAuthRegistrationOptions{
				TenantID: core.StringPtr("testString"),
				EntitlementID: core.StringPtr("testString"),
				Plan: core.StringPtr("testString"),
				Icn: core.StringPtr("testString"),
				Regions: []string{"testString"},
				RedirectURL: core.StringPtr("testString"),
			}

			accessToken, response, err := powervsService.ServiceBrokerAuthRegistration(serviceBrokerAuthRegistrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accessToken).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerAuthTokenPost - Request a new token from a refresh token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerAuthTokenPost(serviceBrokerAuthTokenPostOptions *ServiceBrokerAuthTokenPostOptions)`, func() {
			serviceBrokerAuthTokenPostOptions := &powervsv1.ServiceBrokerAuthTokenPostOptions{
				RefreshToken: core.StringPtr("testString"),
				Source: core.StringPtr("cli"),
			}

			token, response, err := powervsService.ServiceBrokerAuthTokenPost(serviceBrokerAuthTokenPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(token).ToNot(BeNil())
		})
	})

	Describe(`BluemixServiceInstanceGet - Get the current state information associated with the service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`BluemixServiceInstanceGet(bluemixServiceInstanceGetOptions *BluemixServiceInstanceGetOptions)`, func() {
			bluemixServiceInstanceGetOptions := &powervsv1.BluemixServiceInstanceGetOptions{
				InstanceID: core.StringPtr("testString"),
			}

			serviceInstance, response, err := powervsService.BluemixServiceInstanceGet(bluemixServiceInstanceGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstance).ToNot(BeNil())
		})
	})

	Describe(`BluemixServiceInstancePut - Update (disable or enable) the state of a provisioned service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`BluemixServiceInstancePut(bluemixServiceInstancePutOptions *BluemixServiceInstancePutOptions)`, func() {
			bluemixServiceInstancePutOptions := &powervsv1.BluemixServiceInstancePutOptions{
				InstanceID: core.StringPtr("testString"),
				Enabled: core.BoolPtr(true),
				InitiatorID: core.StringPtr("testString"),
				ReasonCode: core.StringPtr("testString"),
			}

			serviceInstance, response, err := powervsService.BluemixServiceInstancePut(bluemixServiceInstancePutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstance).ToNot(BeNil())
		})
	})

	Describe(`CatalogGet - get the catalog of services that the service broker offers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CatalogGet(catalogGetOptions *CatalogGetOptions)`, func() {
			catalogGetOptions := &powervsv1.CatalogGetOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
			}

			catalog, response, err := powervsService.CatalogGet(catalogGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalog).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerHardwareplatformsGet - Available hardware statistics and limits`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerHardwareplatformsGet(serviceBrokerHardwareplatformsGetOptions *ServiceBrokerHardwareplatformsGetOptions)`, func() {
			serviceBrokerHardwareplatformsGetOptions := &powervsv1.ServiceBrokerHardwareplatformsGetOptions{
				RegionZone: core.StringPtr("us-south"),
			}

			mapStringHardwarePlatform, response, err := powervsService.ServiceBrokerHardwareplatformsGet(serviceBrokerHardwareplatformsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringHardwarePlatform).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerHealthHead - Get current server health`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerHealthHead(serviceBrokerHealthHeadOptions *ServiceBrokerHealthHeadOptions)`, func() {
			serviceBrokerHealthHeadOptions := &powervsv1.ServiceBrokerHealthHeadOptions{
			}

			response, err := powervsService.ServiceBrokerHealthHead(serviceBrokerHealthHeadOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})

	Describe(`ServiceBrokerHealth - Get current server health`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerHealth(serviceBrokerHealthOptions *ServiceBrokerHealthOptions)`, func() {
			serviceBrokerHealthOptions := &powervsv1.ServiceBrokerHealthOptions{
			}

			health, response, err := powervsService.ServiceBrokerHealth(serviceBrokerHealthOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(health).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerTestTimeout - Get current server version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerTestTimeout(serviceBrokerTestTimeoutOptions *ServiceBrokerTestTimeoutOptions)`, func() {
			serviceBrokerTestTimeoutOptions := &powervsv1.ServiceBrokerTestTimeoutOptions{
				T: core.Int64Ptr(int64(38)),
			}

			object, response, err := powervsService.ServiceBrokerTestTimeout(serviceBrokerTestTimeoutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerVersion - Get current server version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerVersion(serviceBrokerVersionOptions *ServiceBrokerVersionOptions)`, func() {
			serviceBrokerVersionOptions := &powervsv1.ServiceBrokerVersionOptions{
			}

			version, response, err := powervsService.ServiceBrokerVersion(serviceBrokerVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(version).ToNot(BeNil())
		})
	})

	Describe(`InternalV1PowervsInstancesGet - Get List of PowerVS Cloud Instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1PowervsInstancesGet(internalV1PowervsInstancesGetOptions *InternalV1PowervsInstancesGetOptions)`, func() {
			internalV1PowervsInstancesGetOptions := &powervsv1.InternalV1PowervsInstancesGetOptions{
				PowervsLocation: core.StringPtr("testString"),
			}

			powerVsInstances, response, err := powervsService.InternalV1PowervsInstancesGet(internalV1PowervsInstancesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(powerVsInstances).ToNot(BeNil())
		})
	})

	Describe(`InternalV1PowervsLocationsTransitgatewayGet - Get List of PER enabled PowerVS Locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1PowervsLocationsTransitgatewayGet(internalV1PowervsLocationsTransitgatewayGetOptions *InternalV1PowervsLocationsTransitgatewayGetOptions)`, func() {
			internalV1PowervsLocationsTransitgatewayGetOptions := &powervsv1.InternalV1PowervsLocationsTransitgatewayGetOptions{
			}

			transitGatewayLocations, response, err := powervsService.InternalV1PowervsLocationsTransitgatewayGet(internalV1PowervsLocationsTransitgatewayGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(transitGatewayLocations).ToNot(BeNil())
		})
	})

	Describe(`InternalV1StorageRegionsStoragePoolsGetall - Get the current storage pools settings for a region-zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1StorageRegionsStoragePoolsGetall(internalV1StorageRegionsStoragePoolsGetallOptions *InternalV1StorageRegionsStoragePoolsGetallOptions)`, func() {
			internalV1StorageRegionsStoragePoolsGetallOptions := &powervsv1.InternalV1StorageRegionsStoragePoolsGetallOptions{
				RegionZoneID: core.StringPtr("testString"),
			}

			storagePool, response, err := powervsService.InternalV1StorageRegionsStoragePoolsGetall(internalV1StorageRegionsStoragePoolsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storagePool).ToNot(BeNil())
		})
	})

	Describe(`InternalV1StorageRegionsStoragePoolsGet - Get the settings for given pool name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1StorageRegionsStoragePoolsGet(internalV1StorageRegionsStoragePoolsGetOptions *InternalV1StorageRegionsStoragePoolsGetOptions)`, func() {
			internalV1StorageRegionsStoragePoolsGetOptions := &powervsv1.InternalV1StorageRegionsStoragePoolsGetOptions{
				RegionZoneID: core.StringPtr("testString"),
				StoragePoolName: core.StringPtr("testString"),
			}

			storagePool, response, err := powervsService.InternalV1StorageRegionsStoragePoolsGet(internalV1StorageRegionsStoragePoolsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storagePool).ToNot(BeNil())
		})
	})

	Describe(`InternalV1StorageRegionsStoragePoolsPut - Update the settings for given pool name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1StorageRegionsStoragePoolsPut(internalV1StorageRegionsStoragePoolsPutOptions *InternalV1StorageRegionsStoragePoolsPutOptions)`, func() {
			storageEntitiesModel := &powervsv1.StorageEntities{
				ExistingEntity: core.Int64Ptr(int64(38)),
				NewEntity: core.Int64Ptr(int64(38)),
			}

			thresholdsModel := &powervsv1.Thresholds{
				Capacity: storageEntitiesModel,
				Overcommit: storageEntitiesModel,
				PhysicalCapacity: storageEntitiesModel,
				VdiskCapacity: storageEntitiesModel,
				VdiskLimit: storageEntitiesModel,
			}

			internalV1StorageRegionsStoragePoolsPutOptions := &powervsv1.InternalV1StorageRegionsStoragePoolsPutOptions{
				RegionZoneID: core.StringPtr("testString"),
				StoragePoolName: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				DrEnabled: core.BoolPtr(true),
				OverrideThresholds: thresholdsModel,
				State: core.StringPtr("closed"),
			}

			storagePool, response, err := powervsService.InternalV1StorageRegionsStoragePoolsPut(internalV1StorageRegionsStoragePoolsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storagePool).ToNot(BeNil())
		})
	})

	Describe(`InternalV1StorageRegionsThresholdsGet - Get the current default threshold settings for a region-zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1StorageRegionsThresholdsGet(internalV1StorageRegionsThresholdsGetOptions *InternalV1StorageRegionsThresholdsGetOptions)`, func() {
			internalV1StorageRegionsThresholdsGetOptions := &powervsv1.InternalV1StorageRegionsThresholdsGetOptions{
				RegionZoneID: core.StringPtr("testString"),
			}

			thresholds, response, err := powervsService.InternalV1StorageRegionsThresholdsGet(internalV1StorageRegionsThresholdsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(thresholds).ToNot(BeNil())
		})
	})

	Describe(`InternalV1StorageRegionsThresholdsPut - Update a default threshold setting for a region-zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1StorageRegionsThresholdsPut(internalV1StorageRegionsThresholdsPutOptions *InternalV1StorageRegionsThresholdsPutOptions)`, func() {
			storageEntitiesModel := &powervsv1.StorageEntities{
				ExistingEntity: core.Int64Ptr(int64(38)),
				NewEntity: core.Int64Ptr(int64(38)),
			}

			internalV1StorageRegionsThresholdsPutOptions := &powervsv1.InternalV1StorageRegionsThresholdsPutOptions{
				RegionZoneID: core.StringPtr("testString"),
				Capacity: storageEntitiesModel,
				Overcommit: storageEntitiesModel,
				PhysicalCapacity: storageEntitiesModel,
				VdiskCapacity: storageEntitiesModel,
				VdiskLimit: storageEntitiesModel,
			}

			thresholds, response, err := powervsService.InternalV1StorageRegionsThresholdsPut(internalV1StorageRegionsThresholdsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(thresholds).ToNot(BeNil())
		})
	})

	Describe(`InternalV1TransitgatewayGet - Get the Cloud Instance Transit Gateway information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InternalV1TransitgatewayGet(internalV1TransitgatewayGetOptions *InternalV1TransitgatewayGetOptions)`, func() {
			internalV1TransitgatewayGetOptions := &powervsv1.InternalV1TransitgatewayGetOptions{
				PowervsServiceCRN: core.StringPtr("testString"),
				IBMUserAuthorization: core.StringPtr("testString"),
			}

			transitGatewayInstance, response, err := powervsService.InternalV1TransitgatewayGet(internalV1TransitgatewayGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(transitGatewayInstance).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerOpenstacksGet - List all OpenStack instances being managed`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerOpenstacksGet(serviceBrokerOpenstacksGetOptions *ServiceBrokerOpenstacksGetOptions)`, func() {
			serviceBrokerOpenstacksGetOptions := &powervsv1.ServiceBrokerOpenstacksGetOptions{
			}

			openStacks, response, err := powervsService.ServiceBrokerOpenstacksGet(serviceBrokerOpenstacksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(openStacks).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerOpenstacksPost - Create a new OpenStack instance to be managed`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerOpenstacksPost(serviceBrokerOpenstacksPostOptions *ServiceBrokerOpenstacksPostOptions)`, func() {
			serviceBrokerOpenstacksPostOptions := &powervsv1.ServiceBrokerOpenstacksPostOptions{
				IPAddress: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			openStack, response, err := powervsService.ServiceBrokerOpenstacksPost(serviceBrokerOpenstacksPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(openStack).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerOpenstacksOpenstackGet - List account information for all pvm instances on hostname`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerOpenstacksOpenstackGet(serviceBrokerOpenstacksOpenstackGetOptions *ServiceBrokerOpenstacksOpenstackGetOptions)`, func() {
			serviceBrokerOpenstacksOpenstackGetOptions := &powervsv1.ServiceBrokerOpenstacksOpenstackGetOptions{
				OpenstackID: core.StringPtr("testString"),
			}

			openStackInfo, response, err := powervsService.ServiceBrokerOpenstacksOpenstackGet(serviceBrokerOpenstacksOpenstackGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(openStackInfo).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerOpenstacksHostsGet - List account information for all pvm instances on hostname`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerOpenstacksHostsGet(serviceBrokerOpenstacksHostsGetOptions *ServiceBrokerOpenstacksHostsGetOptions)`, func() {
			serviceBrokerOpenstacksHostsGetOptions := &powervsv1.ServiceBrokerOpenstacksHostsGetOptions{
				Hostname: core.StringPtr("testString"),
				OpenstackID: core.StringPtr("testString"),
			}

			hostInfo, response, err := powervsService.ServiceBrokerOpenstacksHostsGet(serviceBrokerOpenstacksHostsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hostInfo).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerOpenstacksServersGet - List account information for a pvm instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerOpenstacksServersGet(serviceBrokerOpenstacksServersGetOptions *ServiceBrokerOpenstacksServersGetOptions)`, func() {
			serviceBrokerOpenstacksServersGetOptions := &powervsv1.ServiceBrokerOpenstacksServersGetOptions{
				OpenstackID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			hostPvmInstance, response, err := powervsService.ServiceBrokerOpenstacksServersGet(serviceBrokerOpenstacksServersGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(hostPvmInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsGetall - Get all cloud connections in this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsGetall(pcloudCloudconnectionsGetallOptions *PcloudCloudconnectionsGetallOptions)`, func() {
			pcloudCloudconnectionsGetallOptions := &powervsv1.PcloudCloudconnectionsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			cloudConnections, response, err := powervsService.PcloudCloudconnectionsGetall(pcloudCloudconnectionsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudConnections).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsPost - Create a new cloud connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsPost(pcloudCloudconnectionsPostOptions *PcloudCloudconnectionsPostOptions)`, func() {
			cloudConnectionGreTunnelCreateModel := &powervsv1.CloudConnectionGreTunnelCreate{
				CIDR: core.StringPtr("testString"),
				DestIPAddress: core.StringPtr("testString"),
			}

			cloudConnectionEndpointClassicUpdateModel := &powervsv1.CloudConnectionEndpointClassicUpdate{
				Enabled: core.BoolPtr(true),
				Gre: cloudConnectionGreTunnelCreateModel,
			}

			cloudConnectionVPCModel := &powervsv1.CloudConnectionVPC{
				Name: core.StringPtr("testString"),
				VPCID: core.StringPtr("testString"),
			}

			cloudConnectionEndpointVPCModel := &powervsv1.CloudConnectionEndpointVPC{
				Enabled: core.BoolPtr(true),
				Vpcs: []powervsv1.CloudConnectionVPC{*cloudConnectionVPCModel},
			}

			pcloudCloudconnectionsPostOptions := &powervsv1.PcloudCloudconnectionsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Speed: core.Int64Ptr(int64(50)),
				Classic: cloudConnectionEndpointClassicUpdateModel,
				GlobalRouting: core.BoolPtr(true),
				Metered: core.BoolPtr(true),
				Subnets: []string{"testString"},
				TransitEnabled: core.BoolPtr(true),
				VPC: cloudConnectionEndpointVPCModel,
			}

			cloudConnection, response, err := powervsService.PcloudCloudconnectionsPost(pcloudCloudconnectionsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudConnection).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsVirtualprivatecloudsGetall - Get all virtual private cloud connections in this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsVirtualprivatecloudsGetall(pcloudCloudconnectionsVirtualprivatecloudsGetallOptions *PcloudCloudconnectionsVirtualprivatecloudsGetallOptions)`, func() {
			pcloudCloudconnectionsVirtualprivatecloudsGetallOptions := &powervsv1.PcloudCloudconnectionsVirtualprivatecloudsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			cloudConnectionVirtualPrivateClouds, response, err := powervsService.PcloudCloudconnectionsVirtualprivatecloudsGetall(pcloudCloudconnectionsVirtualprivatecloudsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudConnectionVirtualPrivateClouds).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsGet - Get a cloud connection's state/information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsGet(pcloudCloudconnectionsGetOptions *PcloudCloudconnectionsGetOptions)`, func() {
			pcloudCloudconnectionsGetOptions := &powervsv1.PcloudCloudconnectionsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
			}

			cloudConnection, response, err := powervsService.PcloudCloudconnectionsGet(pcloudCloudconnectionsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudConnection).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsPut - Update a Cloud Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsPut(pcloudCloudconnectionsPutOptions *PcloudCloudconnectionsPutOptions)`, func() {
			cloudConnectionGreTunnelCreateModel := &powervsv1.CloudConnectionGreTunnelCreate{
				CIDR: core.StringPtr("testString"),
				DestIPAddress: core.StringPtr("testString"),
			}

			cloudConnectionEndpointClassicUpdateModel := &powervsv1.CloudConnectionEndpointClassicUpdate{
				Enabled: core.BoolPtr(true),
				Gre: cloudConnectionGreTunnelCreateModel,
			}

			cloudConnectionVPCModel := &powervsv1.CloudConnectionVPC{
				Name: core.StringPtr("testString"),
				VPCID: core.StringPtr("testString"),
			}

			cloudConnectionEndpointVPCModel := &powervsv1.CloudConnectionEndpointVPC{
				Enabled: core.BoolPtr(true),
				Vpcs: []powervsv1.CloudConnectionVPC{*cloudConnectionVPCModel},
			}

			pcloudCloudconnectionsPutOptions := &powervsv1.PcloudCloudconnectionsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
				Classic: cloudConnectionEndpointClassicUpdateModel,
				GlobalRouting: core.BoolPtr(true),
				Metered: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
				Speed: core.Int64Ptr(int64(50)),
				VPC: cloudConnectionEndpointVPCModel,
			}

			cloudConnection, response, err := powervsService.PcloudCloudconnectionsPut(pcloudCloudconnectionsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudConnection).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsNetworksPut - Attach a network to the cloud connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsNetworksPut(pcloudCloudconnectionsNetworksPutOptions *PcloudCloudconnectionsNetworksPutOptions)`, func() {
			pcloudCloudconnectionsNetworksPutOptions := &powervsv1.PcloudCloudconnectionsNetworksPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudconnectionsNetworksPut(pcloudCloudconnectionsNetworksPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudLocationsDisasterrecoveryGet - Get the disaster recovery site details for the current location`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudLocationsDisasterrecoveryGet(pcloudLocationsDisasterrecoveryGetOptions *PcloudLocationsDisasterrecoveryGetOptions)`, func() {
			pcloudLocationsDisasterrecoveryGetOptions := &powervsv1.PcloudLocationsDisasterrecoveryGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			disasterRecoveryLocation, response, err := powervsService.PcloudLocationsDisasterrecoveryGet(pcloudLocationsDisasterrecoveryGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(disasterRecoveryLocation).ToNot(BeNil())
		})
	})

	Describe(`PcloudLocationsDisasterrecoveryGetall - Get all disaster recovery locations supported by Power Virtual Server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudLocationsDisasterrecoveryGetall(pcloudLocationsDisasterrecoveryGetallOptions *PcloudLocationsDisasterrecoveryGetallOptions)`, func() {
			pcloudLocationsDisasterrecoveryGetallOptions := &powervsv1.PcloudLocationsDisasterrecoveryGetallOptions{
			}

			disasterRecoveryLocations, response, err := powervsService.PcloudLocationsDisasterrecoveryGetall(pcloudLocationsDisasterrecoveryGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(disasterRecoveryLocations).ToNot(BeNil())
		})
	})

	Describe(`PcloudEventsGetquery - Get events from this cloud instance since a specific timestamp`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudEventsGetquery(pcloudEventsGetqueryOptions *PcloudEventsGetqueryOptions)`, func() {
			pcloudEventsGetqueryOptions := &powervsv1.PcloudEventsGetqueryOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Time: core.StringPtr("testString"),
				FromTime: core.StringPtr("testString"),
				ToTime: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			events, response, err := powervsService.PcloudEventsGetquery(pcloudEventsGetqueryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(events).ToNot(BeNil())
		})
	})

	Describe(`PcloudEventsGet - Get a single event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudEventsGet(pcloudEventsGetOptions *PcloudEventsGetOptions)`, func() {
			pcloudEventsGetOptions := &powervsv1.PcloudEventsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				EventID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			event, response, err := powervsService.PcloudEventsGet(pcloudEventsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(event).ToNot(BeNil())
		})
	})

	Describe(`PcloudV1CloudinstancesCosimagesGet - Get detail of last cos-image import job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV1CloudinstancesCosimagesGet(pcloudV1CloudinstancesCosimagesGetOptions *PcloudV1CloudinstancesCosimagesGetOptions)`, func() {
			pcloudV1CloudinstancesCosimagesGetOptions := &powervsv1.PcloudV1CloudinstancesCosimagesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			job, response, err := powervsService.PcloudV1CloudinstancesCosimagesGet(pcloudV1CloudinstancesCosimagesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`PcloudV1CloudinstancesCosimagesPost - Create an cos-image import job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV1CloudinstancesCosimagesPost(pcloudV1CloudinstancesCosimagesPostOptions *PcloudV1CloudinstancesCosimagesPostOptions)`, func() {
			storageAffinityModel := &powervsv1.StorageAffinity{
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
			}

			pcloudV1CloudinstancesCosimagesPostOptions := &powervsv1.PcloudV1CloudinstancesCosimagesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				BucketName: core.StringPtr("testString"),
				ImageFilename: core.StringPtr("testString"),
				ImageName: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				AccessKey: core.StringPtr("testString"),
				BucketAccess: core.StringPtr("private"),
				OsType: core.StringPtr("aix"),
				SecretKey: core.StringPtr("testString"),
				StorageAffinity: storageAffinityModel,
				StoragePool: core.StringPtr("testString"),
				StorageType: core.StringPtr("testString"),
			}

			jobReference, response, err := powervsService.PcloudV1CloudinstancesCosimagesPost(pcloudV1CloudinstancesCosimagesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesImagesGetall - List all images for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesImagesGetall(pcloudCloudinstancesImagesGetallOptions *PcloudCloudinstancesImagesGetallOptions)`, func() {
			pcloudCloudinstancesImagesGetallOptions := &powervsv1.PcloudCloudinstancesImagesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			images, response, err := powervsService.PcloudCloudinstancesImagesGetall(pcloudCloudinstancesImagesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(images).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesImagesPost - Create a new Image (from available images)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesImagesPost(pcloudCloudinstancesImagesPostOptions *PcloudCloudinstancesImagesPostOptions)`, func() {
			storageAffinityModel := &powervsv1.StorageAffinity{
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
			}

			pcloudCloudinstancesImagesPostOptions := &powervsv1.PcloudCloudinstancesImagesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Source: core.StringPtr("root-project"),
				AccessKey: core.StringPtr("testString"),
				BucketName: core.StringPtr("testString"),
				DiskType: core.StringPtr("testString"),
				ImageFilename: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
				ImageName: core.StringPtr("testString"),
				ImagePath: core.StringPtr("testString"),
				OsType: core.StringPtr("aix"),
				Region: core.StringPtr("testString"),
				SecretKey: core.StringPtr("testString"),
				Source2: core.StringPtr("testString"),
				StorageAffinity: storageAffinityModel,
				StoragePool: core.StringPtr("testString"),
			}

			image, response, err := powervsService.PcloudCloudinstancesImagesPost(pcloudCloudinstancesImagesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesImagesGet - Detailed info of an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesImagesGet(pcloudCloudinstancesImagesGetOptions *PcloudCloudinstancesImagesGetOptions)`, func() {
			pcloudCloudinstancesImagesGetOptions := &powervsv1.PcloudCloudinstancesImagesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
			}

			image, response, err := powervsService.PcloudCloudinstancesImagesGet(pcloudCloudinstancesImagesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesImagesExportPost - Export an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesImagesExportPost(pcloudCloudinstancesImagesExportPostOptions *PcloudCloudinstancesImagesExportPostOptions)`, func() {
			pcloudCloudinstancesImagesExportPostOptions := &powervsv1.PcloudCloudinstancesImagesExportPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
				AccessKey: core.StringPtr("testString"),
				BucketName: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				SecretKey: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesImagesExportPost(pcloudCloudinstancesImagesExportPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesStockimagesGetall - List all available stock images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesStockimagesGetall(pcloudCloudinstancesStockimagesGetallOptions *PcloudCloudinstancesStockimagesGetallOptions)`, func() {
			pcloudCloudinstancesStockimagesGetallOptions := &powervsv1.PcloudCloudinstancesStockimagesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Sap: core.BoolPtr(true),
				Vtl: core.BoolPtr(true),
			}

			images, response, err := powervsService.PcloudCloudinstancesStockimagesGetall(pcloudCloudinstancesStockimagesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(images).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesStockimagesGet - Detailed info of an available stock image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesStockimagesGet(pcloudCloudinstancesStockimagesGetOptions *PcloudCloudinstancesStockimagesGetOptions)`, func() {
			pcloudCloudinstancesStockimagesGetOptions := &powervsv1.PcloudCloudinstancesStockimagesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
			}

			image, response, err := powervsService.PcloudCloudinstancesStockimagesGet(pcloudCloudinstancesStockimagesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())
		})
	})

	Describe(`PcloudImagesGetall - List all the images in the image-catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudImagesGetall(pcloudImagesGetallOptions *PcloudImagesGetallOptions)`, func() {
			pcloudImagesGetallOptions := &powervsv1.PcloudImagesGetallOptions{
				Sap: core.BoolPtr(true),
				Vtl: core.BoolPtr(true),
			}

			images, response, err := powervsService.PcloudImagesGetall(pcloudImagesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(images).ToNot(BeNil())
		})
	})

	Describe(`PcloudImagesGet - Detailed info of an image in the image-catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudImagesGet(pcloudImagesGetOptions *PcloudImagesGetOptions)`, func() {
			pcloudImagesGetOptions := &powervsv1.PcloudImagesGetOptions{
				ImageID: core.StringPtr("testString"),
			}

			image, response, err := powervsService.PcloudImagesGet(pcloudImagesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2ImagesExportGet - Get detail of last image export job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2ImagesExportGet(pcloudV2ImagesExportGetOptions *PcloudV2ImagesExportGetOptions)`, func() {
			pcloudV2ImagesExportGetOptions := &powervsv1.PcloudV2ImagesExportGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
			}

			job, response, err := powervsService.PcloudV2ImagesExportGet(pcloudV2ImagesExportGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2ImagesExportPost - Add image export job to the jobs queue`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2ImagesExportPost(pcloudV2ImagesExportPostOptions *PcloudV2ImagesExportPostOptions)`, func() {
			pcloudV2ImagesExportPostOptions := &powervsv1.PcloudV2ImagesExportPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
				AccessKey: core.StringPtr("testString"),
				BucketName: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
				SecretKey: core.StringPtr("testString"),
			}

			jobReference, response, err := powervsService.PcloudV2ImagesExportPost(pcloudV2ImagesExportPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesGet - Get a Cloud Instance's current state/information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesGet(pcloudCloudinstancesGetOptions *PcloudCloudinstancesGetOptions)`, func() {
			pcloudCloudinstancesGetOptions := &powervsv1.PcloudCloudinstancesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			cloudInstance, response, err := powervsService.PcloudCloudinstancesGet(pcloudCloudinstancesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesPut - Update / Upgrade a Cloud Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesPut(pcloudCloudinstancesPutOptions *PcloudCloudinstancesPutOptions)`, func() {
			pcloudCloudinstancesPutOptions := &powervsv1.PcloudCloudinstancesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Instances: core.Float64Ptr(float64(72.5)),
				Memory: core.Float64Ptr(float64(72.5)),
				ProcUnits: core.Float64Ptr(float64(72.5)),
				Processors: core.Float64Ptr(float64(72.5)),
				Storage: core.Float64Ptr(float64(72.5)),
			}

			cloudInstance, response, err := powervsService.PcloudCloudinstancesPut(pcloudCloudinstancesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesJobsGetall - List up to the last 5 jobs initiated by the cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesJobsGetall(pcloudCloudinstancesJobsGetallOptions *PcloudCloudinstancesJobsGetallOptions)`, func() {
			pcloudCloudinstancesJobsGetallOptions := &powervsv1.PcloudCloudinstancesJobsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				OperationID: core.StringPtr("testString"),
				OperationTarget: core.StringPtr("cloudConnection"),
				OperationAction: core.StringPtr("vmCapture"),
			}

			jobs, response, err := powervsService.PcloudCloudinstancesJobsGetall(pcloudCloudinstancesJobsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobs).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesJobsGet - List the detail of a job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesJobsGet(pcloudCloudinstancesJobsGetOptions *PcloudCloudinstancesJobsGetOptions)`, func() {
			pcloudCloudinstancesJobsGetOptions := &powervsv1.PcloudCloudinstancesJobsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
			}

			job, response, err := powervsService.PcloudCloudinstancesJobsGet(pcloudCloudinstancesJobsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksGetall - Get all networks in this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksGetall(pcloudNetworksGetallOptions *PcloudNetworksGetallOptions)`, func() {
			pcloudNetworksGetallOptions := &powervsv1.PcloudNetworksGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Filter: core.StringPtr("testString"),
			}

			networks, response, err := powervsService.PcloudNetworksGetall(pcloudNetworksGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networks).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPost - Create a new Network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPost(pcloudNetworksPostOptions *PcloudNetworksPostOptions)`, func() {
			ipAddressRangeModel := &powervsv1.IPAddressRange{
				EndingIPAddress: core.StringPtr("testString"),
				StartingIPAddress: core.StringPtr("testString"),
			}

			pcloudNetworksPostOptions := &powervsv1.PcloudNetworksPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Type: core.StringPtr("vlan"),
				CIDR: core.StringPtr("testString"),
				DnsServers: []string{"testString"},
				Gateway: core.StringPtr("testString"),
				IPAddressRanges: []powervsv1.IPAddressRange{*ipAddressRangeModel},
				Jumbo: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
			}

			network, response, err := powervsService.PcloudNetworksPost(pcloudNetworksPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(network).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksGet - Get a network's current state/information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksGet(pcloudNetworksGetOptions *PcloudNetworksGetOptions)`, func() {
			pcloudNetworksGetOptions := &powervsv1.PcloudNetworksGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			network, response, err := powervsService.PcloudNetworksGet(pcloudNetworksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(network).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPut - Update a Network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPut(pcloudNetworksPutOptions *PcloudNetworksPutOptions)`, func() {
			ipAddressRangeModel := &powervsv1.IPAddressRange{
				EndingIPAddress: core.StringPtr("testString"),
				StartingIPAddress: core.StringPtr("testString"),
			}

			pcloudNetworksPutOptions := &powervsv1.PcloudNetworksPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				DnsServers: []string{"testString"},
				Gateway: core.StringPtr("testString"),
				IPAddressRanges: []powervsv1.IPAddressRange{*ipAddressRangeModel},
				Name: core.StringPtr("testString"),
			}

			network, response, err := powervsService.PcloudNetworksPut(pcloudNetworksPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(network).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPortsGetall - Get all ports for this network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPortsGetall(pcloudNetworksPortsGetallOptions *PcloudNetworksPortsGetallOptions)`, func() {
			pcloudNetworksPortsGetallOptions := &powervsv1.PcloudNetworksPortsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			networkPorts, response, err := powervsService.PcloudNetworksPortsGetall(pcloudNetworksPortsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkPorts).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPortsPost - Perform port addition, deletion, and listing`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPortsPost(pcloudNetworksPortsPostOptions *PcloudNetworksPortsPostOptions)`, func() {
			pcloudNetworksPortsPostOptions := &powervsv1.PcloudNetworksPortsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				IPAddress: core.StringPtr("testString"),
			}

			networkPort, response, err := powervsService.PcloudNetworksPortsPost(pcloudNetworksPortsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkPort).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPortsGet - Get a port's information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPortsGet(pcloudNetworksPortsGetOptions *PcloudNetworksPortsGetOptions)`, func() {
			pcloudNetworksPortsGetOptions := &powervsv1.PcloudNetworksPortsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				PortID: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
			}

			networkPort, response, err := powervsService.PcloudNetworksPortsGet(pcloudNetworksPortsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkPort).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPortsPut - Update a port's information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPortsPut(pcloudNetworksPortsPutOptions *PcloudNetworksPortsPutOptions)`, func() {
			pcloudNetworksPortsPutOptions := &powervsv1.PcloudNetworksPortsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				PortID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			networkPort, response, err := powervsService.PcloudNetworksPortsPut(pcloudNetworksPortsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkPort).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesGetall - Get all the pvm instances for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesGetall(pcloudPvminstancesGetallOptions *PcloudPvminstancesGetallOptions)`, func() {
			pcloudPvminstancesGetallOptions := &powervsv1.PcloudPvminstancesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			pvmInstances, response, err := powervsService.PcloudPvminstancesGetall(pcloudPvminstancesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstances).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesPost - Create a new Power VM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesPost(pcloudPvminstancesPostOptions *PcloudPvminstancesPostOptions)`, func() {
			pvmInstanceAddNetworkModel := &powervsv1.PvmInstanceAddNetwork{
				IPAddress: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			softwareLicensesModel := &powervsv1.SoftwareLicenses{
				IbmiCss: core.BoolPtr(false),
				IbmiDbq: core.BoolPtr(false),
				IbmiPha: core.BoolPtr(false),
				IbmiRds: core.BoolPtr(false),
				IbmiRdsUsers: core.Int64Ptr(int64(38)),
			}

			storageAffinityModel := &powervsv1.StorageAffinity{
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
			}

			virtualCoresModel := &powervsv1.VirtualCores{
				Assigned: core.Int64Ptr(int64(38)),
				Max: core.Int64Ptr(int64(38)),
				Min: core.Int64Ptr(int64(38)),
			}

			pcloudPvminstancesPostOptions := &powervsv1.PcloudPvminstancesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
				Memory: core.Float64Ptr(float64(72.5)),
				ProcType: core.StringPtr("dedicated"),
				Processors: core.Float64Ptr(float64(72.5)),
				ServerName: core.StringPtr("testString"),
				DeploymentType: core.StringPtr("testString"),
				KeyPairName: core.StringPtr("testString"),
				LicenseRepositoryCapacity: core.Int64Ptr(int64(38)),
				Migratable: core.BoolPtr(true),
				NetworkIDs: []string{"testString"},
				Networks: []powervsv1.PvmInstanceAddNetwork{*pvmInstanceAddNetworkModel},
				PinPolicy: core.StringPtr("none"),
				PlacementGroup: core.StringPtr("testString"),
				ReplicantAffinityPolicy: core.StringPtr("none"),
				ReplicantNamingScheme: core.StringPtr("suffix"),
				Replicants: core.Float64Ptr(float64(72.5)),
				SharedProcessorPool: core.StringPtr("testString"),
				SoftwareLicenses: softwareLicensesModel,
				StorageAffinity: storageAffinityModel,
				StorageConnection: core.StringPtr("vSCSI"),
				StoragePool: core.StringPtr("testString"),
				StorageType: core.StringPtr("testString"),
				SysType: core.StringPtr("testString"),
				UserData: core.StringPtr("testString"),
				VirtualCores: virtualCoresModel,
				VolumeIDs: []string{"testString"},
				SkipHostValidation: core.BoolPtr(true),
			}

			pvmInstance, response, err := powervsService.PcloudPvminstancesPost(pcloudPvminstancesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesGet - Get a PVM Instance's current state or information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesGet(pcloudPvminstancesGetOptions *PcloudPvminstancesGetOptions)`, func() {
			pcloudPvminstancesGetOptions := &powervsv1.PcloudPvminstancesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			pvmInstance, response, err := powervsService.PcloudPvminstancesGet(pcloudPvminstancesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesPut - Update a PCloud PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesPut(pcloudPvminstancesPutOptions *PcloudPvminstancesPutOptions)`, func() {
			softwareLicensesModel := &powervsv1.SoftwareLicenses{
				IbmiCss: core.BoolPtr(false),
				IbmiDbq: core.BoolPtr(false),
				IbmiPha: core.BoolPtr(false),
				IbmiRds: core.BoolPtr(false),
				IbmiRdsUsers: core.Int64Ptr(int64(38)),
			}

			virtualCoresModel := &powervsv1.VirtualCores{
				Assigned: core.Int64Ptr(int64(38)),
				Max: core.Int64Ptr(int64(38)),
				Min: core.Int64Ptr(int64(38)),
			}

			pcloudPvminstancesPutOptions := &powervsv1.PcloudPvminstancesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				LicenseRepositoryCapacity: core.Int64Ptr(int64(38)),
				Memory: core.Float64Ptr(float64(72.5)),
				Migratable: core.BoolPtr(true),
				PinPolicy: core.StringPtr("none"),
				ProcType: core.StringPtr("dedicated"),
				Processors: core.Float64Ptr(float64(72.5)),
				SapProfileID: core.StringPtr("testString"),
				ServerName: core.StringPtr("testString"),
				SoftwareLicenses: softwareLicensesModel,
				StoragePoolAffinity: core.BoolPtr(true),
				VirtualCores: virtualCoresModel,
			}

			pvmInstanceUpdateResponse, response, err := powervsService.PcloudPvminstancesPut(pcloudPvminstancesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvmInstanceUpdateResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesActionPost - Perform an action (start stop reboot immediate-shutdown reset) on a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesActionPost(pcloudPvminstancesActionPostOptions *PcloudPvminstancesActionPostOptions)`, func() {
			pcloudPvminstancesActionPostOptions := &powervsv1.PcloudPvminstancesActionPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				Action: core.StringPtr("start"),
			}

			object, response, err := powervsService.PcloudPvminstancesActionPost(pcloudPvminstancesActionPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesCapturePost - Capture a PVMInstance and create a deployable image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesCapturePost(pcloudPvminstancesCapturePostOptions *PcloudPvminstancesCapturePostOptions)`, func() {
			pcloudPvminstancesCapturePostOptions := &powervsv1.PcloudPvminstancesCapturePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				CaptureDestination: core.StringPtr("cloud-storage"),
				CaptureName: core.StringPtr("testString"),
				CaptureVolumeIDs: []string{"testString"},
				CloudStorageAccessKey: core.StringPtr("testString"),
				CloudStorageImagePath: core.StringPtr("testString"),
				CloudStorageRegion: core.StringPtr("testString"),
				CloudStorageSecretKey: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPvminstancesCapturePost(pcloudPvminstancesCapturePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesClonePost - Clone a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesClonePost(pcloudPvminstancesClonePostOptions *PcloudPvminstancesClonePostOptions)`, func() {
			pvmInstanceAddNetworkModel := &powervsv1.PvmInstanceAddNetwork{
				IPAddress: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			softwareLicensesModel := &powervsv1.SoftwareLicenses{
				IbmiCss: core.BoolPtr(false),
				IbmiDbq: core.BoolPtr(false),
				IbmiPha: core.BoolPtr(false),
				IbmiRds: core.BoolPtr(false),
				IbmiRdsUsers: core.Int64Ptr(int64(38)),
			}

			pcloudPvminstancesClonePostOptions := &powervsv1.PcloudPvminstancesClonePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Networks: []powervsv1.PvmInstanceAddNetwork{*pvmInstanceAddNetworkModel},
				KeyPairName: core.StringPtr("testString"),
				Memory: core.Float64Ptr(float64(72.5)),
				ProcType: core.StringPtr("dedicated"),
				Processors: core.Float64Ptr(float64(72.5)),
				SoftwareLicenses: softwareLicensesModel,
				VolumeIDs: []string{"testString"},
			}

			pvmInstance, response, err := powervsService.PcloudPvminstancesClonePost(pcloudPvminstancesClonePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvmInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesConsoleGet - List all console languages`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesConsoleGet(pcloudPvminstancesConsoleGetOptions *PcloudPvminstancesConsoleGetOptions)`, func() {
			pcloudPvminstancesConsoleGetOptions := &powervsv1.PcloudPvminstancesConsoleGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			consoleLanguages, response, err := powervsService.PcloudPvminstancesConsoleGet(pcloudPvminstancesConsoleGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(consoleLanguages).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesConsolePost - Generate the noVNC Console URL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesConsolePost(pcloudPvminstancesConsolePostOptions *PcloudPvminstancesConsolePostOptions)`, func() {
			pcloudPvminstancesConsolePostOptions := &powervsv1.PcloudPvminstancesConsolePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			pvmInstanceConsole, response, err := powervsService.PcloudPvminstancesConsolePost(pcloudPvminstancesConsolePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pvmInstanceConsole).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesConsolePut - Update PVMInstance console laguage code`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesConsolePut(pcloudPvminstancesConsolePutOptions *PcloudPvminstancesConsolePutOptions)`, func() {
			pcloudPvminstancesConsolePutOptions := &powervsv1.PcloudPvminstancesConsolePutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				Code: core.StringPtr("testString"),
				Language: core.StringPtr("testString"),
			}

			consoleLanguage, response, err := powervsService.PcloudPvminstancesConsolePut(pcloudPvminstancesConsolePutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(consoleLanguage).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesNetworksGetall - Get all networks for this PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesNetworksGetall(pcloudPvminstancesNetworksGetallOptions *PcloudPvminstancesNetworksGetallOptions)`, func() {
			pcloudPvminstancesNetworksGetallOptions := &powervsv1.PcloudPvminstancesNetworksGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			pvmInstanceNetworks, response, err := powervsService.PcloudPvminstancesNetworksGetall(pcloudPvminstancesNetworksGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstanceNetworks).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesNetworksPost - Perform network addition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesNetworksPost(pcloudPvminstancesNetworksPostOptions *PcloudPvminstancesNetworksPostOptions)`, func() {
			pcloudPvminstancesNetworksPostOptions := &powervsv1.PcloudPvminstancesNetworksPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				IPAddress: core.StringPtr("testString"),
			}

			pvmInstanceNetwork, response, err := powervsService.PcloudPvminstancesNetworksPost(pcloudPvminstancesNetworksPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pvmInstanceNetwork).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesNetworksGet - Get a PVM Instance's network information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesNetworksGet(pcloudPvminstancesNetworksGetOptions *PcloudPvminstancesNetworksGetOptions)`, func() {
			pcloudPvminstancesNetworksGetOptions := &powervsv1.PcloudPvminstancesNetworksGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			pvmInstanceNetworks, response, err := powervsService.PcloudPvminstancesNetworksGet(pcloudPvminstancesNetworksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstanceNetworks).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesOperationsPost - Perform an operation on a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesOperationsPost(pcloudPvminstancesOperationsPostOptions *PcloudPvminstancesOperationsPostOptions)`, func() {
			operationsModel := &powervsv1.Operations{
				BootMode: core.StringPtr("a"),
				OperatingMode: core.StringPtr("normal"),
				Task: core.StringPtr("dston"),
			}

			pcloudPvminstancesOperationsPostOptions := &powervsv1.PcloudPvminstancesOperationsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				Operation: operationsModel,
				OperationType: core.StringPtr("job"),
			}

			object, response, err := powervsService.PcloudPvminstancesOperationsPost(pcloudPvminstancesOperationsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesSnapshotsGetall - Get all snapshots for this PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesSnapshotsGetall(pcloudPvminstancesSnapshotsGetallOptions *PcloudPvminstancesSnapshotsGetallOptions)`, func() {
			pcloudPvminstancesSnapshotsGetallOptions := &powervsv1.PcloudPvminstancesSnapshotsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			snapshots, response, err := powervsService.PcloudPvminstancesSnapshotsGetall(pcloudPvminstancesSnapshotsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshots).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesSnapshotsPost - Create a PVM Instance snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesSnapshotsPost(pcloudPvminstancesSnapshotsPostOptions *PcloudPvminstancesSnapshotsPostOptions)`, func() {
			pcloudPvminstancesSnapshotsPostOptions := &powervsv1.PcloudPvminstancesSnapshotsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				PerformancePath: core.BoolPtr(true),
				VolumeIDs: []string{"testString"},
			}

			snapshotCreateResponse, response, err := powervsService.PcloudPvminstancesSnapshotsPost(pcloudPvminstancesSnapshotsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(snapshotCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesSnapshotsRestorePost - Restore a PVM Instance snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesSnapshotsRestorePost(pcloudPvminstancesSnapshotsRestorePostOptions *PcloudPvminstancesSnapshotsRestorePostOptions)`, func() {
			pcloudPvminstancesSnapshotsRestorePostOptions := &powervsv1.PcloudPvminstancesSnapshotsRestorePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
				Force: core.BoolPtr(false),
				RestoreFailAction: core.StringPtr("retry"),
			}

			snapshot, response, err := powervsService.PcloudPvminstancesSnapshotsRestorePost(pcloudPvminstancesSnapshotsRestorePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2PvminstancesGetall - Get all the pvm instances for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2PvminstancesGetall(pcloudV2PvminstancesGetallOptions *PcloudV2PvminstancesGetallOptions)`, func() {
			pcloudV2PvminstancesGetallOptions := &powervsv1.PcloudV2PvminstancesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			pvmInstancesV2, response, err := powervsService.PcloudV2PvminstancesGetall(pcloudV2PvminstancesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstancesV2).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2PvminstancesCaptureGet - Get detail of last capture job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2PvminstancesCaptureGet(pcloudV2PvminstancesCaptureGetOptions *PcloudV2PvminstancesCaptureGetOptions)`, func() {
			pcloudV2PvminstancesCaptureGetOptions := &powervsv1.PcloudV2PvminstancesCaptureGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			job, response, err := powervsService.PcloudV2PvminstancesCaptureGet(pcloudV2PvminstancesCaptureGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2PvminstancesCapturePost - Add a capture pvm-instance to the jobs queue`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2PvminstancesCapturePost(pcloudV2PvminstancesCapturePostOptions *PcloudV2PvminstancesCapturePostOptions)`, func() {
			pcloudV2PvminstancesCapturePostOptions := &powervsv1.PcloudV2PvminstancesCapturePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				CaptureDestination: core.StringPtr("cloud-storage"),
				CaptureName: core.StringPtr("testString"),
				CaptureVolumeIDs: []string{"testString"},
				CloudStorageAccessKey: core.StringPtr("testString"),
				CloudStorageImagePath: core.StringPtr("testString"),
				CloudStorageRegion: core.StringPtr("testString"),
				CloudStorageSecretKey: core.StringPtr("testString"),
			}

			jobReference, response, err := powervsService.PcloudV2PvminstancesCapturePost(pcloudV2PvminstancesCapturePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsGetall - Get all Server Placement Groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsGetall(pcloudPlacementgroupsGetallOptions *PcloudPlacementgroupsGetallOptions)`, func() {
			pcloudPlacementgroupsGetallOptions := &powervsv1.PcloudPlacementgroupsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			placementGroups, response, err := powervsService.PcloudPlacementgroupsGetall(pcloudPlacementgroupsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroups).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsPost - Create a new Server Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsPost(pcloudPlacementgroupsPostOptions *PcloudPlacementgroupsPostOptions)`, func() {
			pcloudPlacementgroupsPostOptions := &powervsv1.PcloudPlacementgroupsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Policy: core.StringPtr("affinity"),
			}

			placementGroup, response, err := powervsService.PcloudPlacementgroupsPost(pcloudPlacementgroupsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsGet - Get Server Placement Group detail`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsGet(pcloudPlacementgroupsGetOptions *PcloudPlacementgroupsGetOptions)`, func() {
			pcloudPlacementgroupsGetOptions := &powervsv1.PcloudPlacementgroupsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PlacementGroupID: core.StringPtr("testString"),
			}

			placementGroup, response, err := powervsService.PcloudPlacementgroupsGet(pcloudPlacementgroupsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsMembersPost - Add Server to Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsMembersPost(pcloudPlacementgroupsMembersPostOptions *PcloudPlacementgroupsMembersPostOptions)`, func() {
			pcloudPlacementgroupsMembersPostOptions := &powervsv1.PcloudPlacementgroupsMembersPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PlacementGroupID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			placementGroup, response, err := powervsService.PcloudPlacementgroupsMembersPost(pcloudPlacementgroupsMembersPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudSapGetall - Get list of SAP profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSapGetall(pcloudSapGetallOptions *PcloudSapGetallOptions)`, func() {
			pcloudSapGetallOptions := &powervsv1.PcloudSapGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			sapProfiles, response, err := powervsService.PcloudSapGetall(pcloudSapGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sapProfiles).ToNot(BeNil())
		})
	})

	Describe(`PcloudSapPost - Create a new SAP PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSapPost(pcloudSapPostOptions *PcloudSapPostOptions)`, func() {
			pvmInstanceAddNetworkModel := &powervsv1.PvmInstanceAddNetwork{
				IPAddress: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			pvmInstanceMultiCreateModel := &powervsv1.PvmInstanceMultiCreate{
				AffinityPolicy: core.StringPtr("none"),
				Count: core.Int64Ptr(int64(38)),
				Numerical: core.StringPtr("suffix"),
			}

			storageAffinityModel := &powervsv1.StorageAffinity{
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
			}

			pcloudSapPostOptions := &powervsv1.PcloudSapPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Networks: []powervsv1.PvmInstanceAddNetwork{*pvmInstanceAddNetworkModel},
				ProfileID: core.StringPtr("testString"),
				DeploymentType: core.StringPtr("testString"),
				Instances: pvmInstanceMultiCreateModel,
				PinPolicy: core.StringPtr("none"),
				PlacementGroup: core.StringPtr("testString"),
				SshKeyName: core.StringPtr("testString"),
				StorageAffinity: storageAffinityModel,
				StoragePool: core.StringPtr("testString"),
				StorageType: core.StringPtr("testString"),
				SysType: core.StringPtr("testString"),
				UserData: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
			}

			pvmInstance, response, err := powervsService.PcloudSapPost(pcloudSapPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvmInstance).ToNot(BeNil())
		})
	})

	Describe(`PcloudSapGet - Get the information on an SAP profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSapGet(pcloudSapGetOptions *PcloudSapGetOptions)`, func() {
			pcloudSapGetOptions := &powervsv1.PcloudSapGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SapProfileID: core.StringPtr("testString"),
			}

			sapProfile, response, err := powervsService.PcloudSapGet(pcloudSapGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sapProfile).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsGetall - Get the list of Shared Processor Pool Placement Groups for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsGetall(pcloudSppplacementgroupsGetallOptions *PcloudSppplacementgroupsGetallOptions)`, func() {
			pcloudSppplacementgroupsGetallOptions := &powervsv1.PcloudSppplacementgroupsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			sppPlacementGroups, response, err := powervsService.PcloudSppplacementgroupsGetall(pcloudSppplacementgroupsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sppPlacementGroups).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsPost - Create a new Shared Processor Pool Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsPost(pcloudSppplacementgroupsPostOptions *PcloudSppplacementgroupsPostOptions)`, func() {
			pcloudSppplacementgroupsPostOptions := &powervsv1.PcloudSppplacementgroupsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Policy: core.StringPtr("affinity"),
			}

			sppPlacementGroup, response, err := powervsService.PcloudSppplacementgroupsPost(pcloudSppplacementgroupsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sppPlacementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsGet - Get the detail of a Shared Processor Pool Placement Group for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsGet(pcloudSppplacementgroupsGetOptions *PcloudSppplacementgroupsGetOptions)`, func() {
			pcloudSppplacementgroupsGetOptions := &powervsv1.PcloudSppplacementgroupsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SppPlacementGroupID: core.StringPtr("testString"),
			}

			sppPlacementGroup, response, err := powervsService.PcloudSppplacementgroupsGet(pcloudSppplacementgroupsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sppPlacementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsMembersPost - Add Shared Processor Pool as a member of a Shared Processor Pool Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsMembersPost(pcloudSppplacementgroupsMembersPostOptions *PcloudSppplacementgroupsMembersPostOptions)`, func() {
			pcloudSppplacementgroupsMembersPostOptions := &powervsv1.PcloudSppplacementgroupsMembersPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SppPlacementGroupID: core.StringPtr("testString"),
				SharedProcessorPoolID: core.StringPtr("testString"),
			}

			sppPlacementGroup, response, err := powervsService.PcloudSppplacementgroupsMembersPost(pcloudSppplacementgroupsMembersPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sppPlacementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudDhcpGetall - Get all DHCP Servers information (OpenShift Internal Use Only)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudDhcpGetall(pcloudDhcpGetallOptions *PcloudDhcpGetallOptions)`, func() {
			pcloudDhcpGetallOptions := &powervsv1.PcloudDhcpGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			dhcpServer, response, err := powervsService.PcloudDhcpGetall(pcloudDhcpGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dhcpServer).ToNot(BeNil())
		})
	})

	Describe(`PcloudDhcpPost - Create a DHCP Server (OpenShift Internal Use Only)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudDhcpPost(pcloudDhcpPostOptions *PcloudDhcpPostOptions)`, func() {
			pcloudDhcpPostOptions := &powervsv1.PcloudDhcpPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CIDR: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
				DnsServer: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				SnatEnabled: core.BoolPtr(true),
			}

			dhcpServer, response, err := powervsService.PcloudDhcpPost(pcloudDhcpPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(dhcpServer).ToNot(BeNil())
		})
	})

	Describe(`PcloudDhcpGet - Get DHCP Server information (OpenShift Internal Use Only)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudDhcpGet(pcloudDhcpGetOptions *PcloudDhcpGetOptions)`, func() {
			pcloudDhcpGetOptions := &powervsv1.PcloudDhcpGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				DhcpID: core.StringPtr("testString"),
			}

			dhcpServerDetail, response, err := powervsService.PcloudDhcpGet(pcloudDhcpGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dhcpServerDetail).ToNot(BeNil())
		})
	})

	Describe(`PcloudSharedprocessorpoolsGetall - Get the list of Shared Processor Pools for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSharedprocessorpoolsGetall(pcloudSharedprocessorpoolsGetallOptions *PcloudSharedprocessorpoolsGetallOptions)`, func() {
			pcloudSharedprocessorpoolsGetallOptions := &powervsv1.PcloudSharedprocessorpoolsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			sharedProcessorPools, response, err := powervsService.PcloudSharedprocessorpoolsGetall(pcloudSharedprocessorpoolsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sharedProcessorPools).ToNot(BeNil())
		})
	})

	Describe(`PcloudSharedprocessorpoolsPost - Create a new Shared Processor Pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSharedprocessorpoolsPost(pcloudSharedprocessorpoolsPostOptions *PcloudSharedprocessorpoolsPostOptions)`, func() {
			pcloudSharedprocessorpoolsPostOptions := &powervsv1.PcloudSharedprocessorpoolsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				HostGroup: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ReservedCores: core.Int64Ptr(int64(38)),
				PlacementGroupID: core.StringPtr("testString"),
			}

			sharedProcessorPool, response, err := powervsService.PcloudSharedprocessorpoolsPost(pcloudSharedprocessorpoolsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(sharedProcessorPool).ToNot(BeNil())
		})
	})

	Describe(`PcloudSharedprocessorpoolsGet - Get the detail of a Shared Processor Pool for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSharedprocessorpoolsGet(pcloudSharedprocessorpoolsGetOptions *PcloudSharedprocessorpoolsGetOptions)`, func() {
			pcloudSharedprocessorpoolsGetOptions := &powervsv1.PcloudSharedprocessorpoolsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SharedProcessorPoolID: core.StringPtr("testString"),
			}

			sharedProcessorPoolDetail, response, err := powervsService.PcloudSharedprocessorpoolsGet(pcloudSharedprocessorpoolsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sharedProcessorPoolDetail).ToNot(BeNil())
		})
	})

	Describe(`PcloudSharedprocessorpoolsPut - Update a Shared Processor Pool for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSharedprocessorpoolsPut(pcloudSharedprocessorpoolsPutOptions *PcloudSharedprocessorpoolsPutOptions)`, func() {
			pcloudSharedprocessorpoolsPutOptions := &powervsv1.PcloudSharedprocessorpoolsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SharedProcessorPoolID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				ReservedCores: core.Int64Ptr(int64(38)),
			}

			sharedProcessorPool, response, err := powervsService.PcloudSharedprocessorpoolsPut(pcloudSharedprocessorpoolsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sharedProcessorPool).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesSnapshotsGetall - List all PVM instance snapshots for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesSnapshotsGetall(pcloudCloudinstancesSnapshotsGetallOptions *PcloudCloudinstancesSnapshotsGetallOptions)`, func() {
			pcloudCloudinstancesSnapshotsGetallOptions := &powervsv1.PcloudCloudinstancesSnapshotsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			snapshots, response, err := powervsService.PcloudCloudinstancesSnapshotsGetall(pcloudCloudinstancesSnapshotsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshots).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesSnapshotsGet - Get the detail of a snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesSnapshotsGet(pcloudCloudinstancesSnapshotsGetOptions *PcloudCloudinstancesSnapshotsGetOptions)`, func() {
			pcloudCloudinstancesSnapshotsGetOptions := &powervsv1.PcloudCloudinstancesSnapshotsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
			}

			snapshot, response, err := powervsService.PcloudCloudinstancesSnapshotsGet(pcloudCloudinstancesSnapshotsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesSnapshotsPut - Update a PVM instance snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesSnapshotsPut(pcloudCloudinstancesSnapshotsPutOptions *PcloudCloudinstancesSnapshotsPutOptions)`, func() {
			pcloudCloudinstancesSnapshotsPutOptions := &powervsv1.PcloudCloudinstancesSnapshotsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesSnapshotsPut(pcloudCloudinstancesSnapshotsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudStoragecapacityPoolsGetall - Storage capacity for all available storage pools in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudStoragecapacityPoolsGetall(pcloudStoragecapacityPoolsGetallOptions *PcloudStoragecapacityPoolsGetallOptions)`, func() {
			pcloudStoragecapacityPoolsGetallOptions := &powervsv1.PcloudStoragecapacityPoolsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			storagePoolsCapacity, response, err := powervsService.PcloudStoragecapacityPoolsGetall(pcloudStoragecapacityPoolsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storagePoolsCapacity).ToNot(BeNil())
		})
	})

	Describe(`PcloudStoragecapacityPoolsGet - Storage capacity for a storage pool in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudStoragecapacityPoolsGet(pcloudStoragecapacityPoolsGetOptions *PcloudStoragecapacityPoolsGetOptions)`, func() {
			pcloudStoragecapacityPoolsGetOptions := &powervsv1.PcloudStoragecapacityPoolsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				StoragePoolName: core.StringPtr("testString"),
			}

			storagePoolCapacity, response, err := powervsService.PcloudStoragecapacityPoolsGet(pcloudStoragecapacityPoolsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storagePoolCapacity).ToNot(BeNil())
		})
	})

	Describe(`PcloudStoragecapacityTypesGetall - Storage capacity for all available storage types in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudStoragecapacityTypesGetall(pcloudStoragecapacityTypesGetallOptions *PcloudStoragecapacityTypesGetallOptions)`, func() {
			pcloudStoragecapacityTypesGetallOptions := &powervsv1.PcloudStoragecapacityTypesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			storageTypesCapacity, response, err := powervsService.PcloudStoragecapacityTypesGetall(pcloudStoragecapacityTypesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageTypesCapacity).ToNot(BeNil())
		})
	})

	Describe(`PcloudStoragecapacityTypesGet - Storage capacity for a storage type in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudStoragecapacityTypesGet(pcloudStoragecapacityTypesGetOptions *PcloudStoragecapacityTypesGetOptions)`, func() {
			pcloudStoragecapacityTypesGetOptions := &powervsv1.PcloudStoragecapacityTypesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				StorageTypeName: core.StringPtr("testString"),
			}

			storageTypeCapacity, response, err := powervsService.PcloudStoragecapacityTypesGet(pcloudStoragecapacityTypesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(storageTypeCapacity).ToNot(BeNil())
		})
	})

	Describe(`PcloudSystempoolsGet - List of available system pools within a particular DataCenter`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSystempoolsGet(pcloudSystempoolsGetOptions *PcloudSystempoolsGetOptions)`, func() {
			pcloudSystempoolsGetOptions := &powervsv1.PcloudSystempoolsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			mapStringSystemPool, response, err := powervsService.PcloudSystempoolsGet(pcloudSystempoolsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringSystemPool).ToNot(BeNil())
		})
	})

	Describe(`PcloudTasksGet - Get a Task`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTasksGet(pcloudTasksGetOptions *PcloudTasksGetOptions)`, func() {
			pcloudTasksGetOptions := &powervsv1.PcloudTasksGetOptions{
				TaskID: core.StringPtr("testString"),
			}

			task, response, err := powervsService.PcloudTasksGet(pcloudTasksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsGet - Get a Tenant's current state/information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsGet(pcloudTenantsGetOptions *PcloudTenantsGetOptions)`, func() {
			pcloudTenantsGetOptions := &powervsv1.PcloudTenantsGetOptions{
				TenantID: core.StringPtr("testString"),
			}

			tenant, response, err := powervsService.PcloudTenantsGet(pcloudTenantsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenant).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsPut - Update a tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsPut(pcloudTenantsPutOptions *PcloudTenantsPutOptions)`, func() {
			peeringNetworkModel := &powervsv1.PeeringNetwork{
				CIDR: core.StringPtr("testString"),
				DnsServers: []string{"testString"},
				ProjectName: core.StringPtr("testString"),
			}

			pcloudTenantsPutOptions := &powervsv1.PcloudTenantsPutOptions{
				TenantID: core.StringPtr("testString"),
				Icn: core.StringPtr("testString"),
				PeeringNetworks: []powervsv1.PeeringNetwork{*peeringNetworkModel},
			}

			tenant, response, err := powervsService.PcloudTenantsPut(pcloudTenantsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenant).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsSshkeysGetall - List a Tenant's SSH Keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsSshkeysGetall(pcloudTenantsSshkeysGetallOptions *PcloudTenantsSshkeysGetallOptions)`, func() {
			pcloudTenantsSshkeysGetallOptions := &powervsv1.PcloudTenantsSshkeysGetallOptions{
				TenantID: core.StringPtr("testString"),
			}

			sshKeys, response, err := powervsService.PcloudTenantsSshkeysGetall(pcloudTenantsSshkeysGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sshKeys).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsSshkeysPost - Add a new SSH key to the Tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsSshkeysPost(pcloudTenantsSshkeysPostOptions *PcloudTenantsSshkeysPostOptions)`, func() {
			pcloudTenantsSshkeysPostOptions := &powervsv1.PcloudTenantsSshkeysPostOptions{
				TenantID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				SshKey: core.StringPtr("testString"),
				CreationDate: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			sshKey, response, err := powervsService.PcloudTenantsSshkeysPost(pcloudTenantsSshkeysPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sshKey).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsSshkeysGet - Get a Tenant's SSH Key by name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsSshkeysGet(pcloudTenantsSshkeysGetOptions *PcloudTenantsSshkeysGetOptions)`, func() {
			pcloudTenantsSshkeysGetOptions := &powervsv1.PcloudTenantsSshkeysGetOptions{
				TenantID: core.StringPtr("testString"),
				SshkeyName: core.StringPtr("testString"),
			}

			sshKey, response, err := powervsService.PcloudTenantsSshkeysGet(pcloudTenantsSshkeysGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sshKey).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsSshkeysPut - Update an SSH Key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsSshkeysPut(pcloudTenantsSshkeysPutOptions *PcloudTenantsSshkeysPutOptions)`, func() {
			pcloudTenantsSshkeysPutOptions := &powervsv1.PcloudTenantsSshkeysPutOptions{
				TenantID: core.StringPtr("testString"),
				SshkeyName: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				SshKey: core.StringPtr("testString"),
				CreationDate: CreateMockDateTime("2019-01-01T12:00:00.000Z"),
			}

			sshKey, response, err := powervsService.PcloudTenantsSshkeysPut(pcloudTenantsSshkeysPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sshKey).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsGetall - Get all VPN Connections`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsGetall(pcloudVpnconnectionsGetallOptions *PcloudVpnconnectionsGetallOptions)`, func() {
			pcloudVpnconnectionsGetallOptions := &powervsv1.PcloudVpnconnectionsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			vpnConnections, response, err := powervsService.PcloudVpnconnectionsGetall(pcloudVpnconnectionsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnConnections).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsPost - Create VPN Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsPost(pcloudVpnconnectionsPostOptions *PcloudVpnconnectionsPostOptions)`, func() {
			pcloudVpnconnectionsPostOptions := &powervsv1.PcloudVpnconnectionsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IkePolicy: core.StringPtr("c36723ec-8593-11eb-8dcd-0242ac133853"),
				IPSecPolicy: core.StringPtr("c12345d-8593-11eb-8dcd-0242ac134573"),
				Mode: core.StringPtr("policy"),
				Name: core.StringPtr("VPN-Connection-1"),
				Networks: []string{"7f950c76-8582-11veb-8dcd-0242ac153", "7f950c76-8582-11veb-8dcd-0242ac144", "7f950c76-8582-11veb-8dcd-0242ac199"},
				PeerGatewayAddress: core.StringPtr("192.168.1.1"),
				PeerSubnets: []string{"128.170.1.0/20", "128.169.1.0/24", "128.168.1.0/27", "128.170.1.0/32"},
			}

			vpnConnectionCreateResponse, response, err := powervsService.PcloudVpnconnectionsPost(pcloudVpnconnectionsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vpnConnectionCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsGet - Get VPN Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsGet(pcloudVpnconnectionsGetOptions *PcloudVpnconnectionsGetOptions)`, func() {
			pcloudVpnconnectionsGetOptions := &powervsv1.PcloudVpnconnectionsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
			}

			vpnConnection, response, err := powervsService.PcloudVpnconnectionsGet(pcloudVpnconnectionsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnConnection).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsPut - Update VPN Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsPut(pcloudVpnconnectionsPutOptions *PcloudVpnconnectionsPutOptions)`, func() {
			vpnConnectionUpdateModel := &powervsv1.VPNConnectionUpdate{
				IkePolicy: core.StringPtr("c36723ec-8593-11eb-8dcd-0242ac133853"),
				IPSecPolicy: core.StringPtr("c12345d-8593-11eb-8dcd-0242ac134573"),
				Name: core.StringPtr("VPN-Connection-1"),
				PeerGatewayAddress: core.StringPtr("192.168.1.1"),
			}
			vpnConnectionUpdateModel.SetProperty("foo", core.StringPtr("testString"))

			pcloudVpnconnectionsPutOptions := &powervsv1.PcloudVpnconnectionsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
				VPNConnectionUpdate: vpnConnectionUpdateModel,
			}

			vpnConnection, response, err := powervsService.PcloudVpnconnectionsPut(pcloudVpnconnectionsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnConnection).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsNetworksGet - Get attached networks`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsNetworksGet(pcloudVpnconnectionsNetworksGetOptions *PcloudVpnconnectionsNetworksGetOptions)`, func() {
			pcloudVpnconnectionsNetworksGetOptions := &powervsv1.PcloudVpnconnectionsNetworksGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
			}

			networkIDs, response, err := powervsService.PcloudVpnconnectionsNetworksGet(pcloudVpnconnectionsNetworksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkIDs).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsNetworksPut - Attach network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsNetworksPut(pcloudVpnconnectionsNetworksPutOptions *PcloudVpnconnectionsNetworksPutOptions)`, func() {
			pcloudVpnconnectionsNetworksPutOptions := &powervsv1.PcloudVpnconnectionsNetworksPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("7f950c76-8582-11qeb-8dcd-0242ac172"),
			}

			jobReference, response, err := powervsService.PcloudVpnconnectionsNetworksPut(pcloudVpnconnectionsNetworksPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsPeersubnetsGet - Get Peer Subnets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsPeersubnetsGet(pcloudVpnconnectionsPeersubnetsGetOptions *PcloudVpnconnectionsPeersubnetsGetOptions)`, func() {
			pcloudVpnconnectionsPeersubnetsGetOptions := &powervsv1.PcloudVpnconnectionsPeersubnetsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
			}

			peerSubnets, response, err := powervsService.PcloudVpnconnectionsPeersubnetsGet(pcloudVpnconnectionsPeersubnetsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(peerSubnets).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsPeersubnetsPut - Attach Peer Subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsPeersubnetsPut(pcloudVpnconnectionsPeersubnetsPutOptions *PcloudVpnconnectionsPeersubnetsPutOptions)`, func() {
			pcloudVpnconnectionsPeersubnetsPutOptions := &powervsv1.PcloudVpnconnectionsPeersubnetsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
				CIDR: core.StringPtr("128.170.1.0/32"),
			}

			peerSubnets, response, err := powervsService.PcloudVpnconnectionsPeersubnetsPut(pcloudVpnconnectionsPeersubnetsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(peerSubnets).ToNot(BeNil())
		})
	})

	Describe(`PcloudIkepoliciesGetall - Get all IKE Policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIkepoliciesGetall(pcloudIkepoliciesGetallOptions *PcloudIkepoliciesGetallOptions)`, func() {
			pcloudIkepoliciesGetallOptions := &powervsv1.PcloudIkepoliciesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			ikePolicies, response, err := powervsService.PcloudIkepoliciesGetall(pcloudIkepoliciesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicies).ToNot(BeNil())
		})
	})

	Describe(`PcloudIkepoliciesPost - Add IKE Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIkepoliciesPost(pcloudIkepoliciesPostOptions *PcloudIkepoliciesPostOptions)`, func() {
			pcloudIkepoliciesPostOptions := &powervsv1.PcloudIkepoliciesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				DhGroup: core.Int64Ptr(int64(2)),
				Encryption: core.StringPtr("aes-256-cbc"),
				KeyLifetime: core.Int64Ptr(int64(28800)),
				Name: core.StringPtr("ikePolicy1"),
				PresharedKey: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(2)),
				Authentication: core.StringPtr("sha-256"),
			}

			ikePolicy, response, err := powervsService.PcloudIkepoliciesPost(pcloudIkepoliciesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudIkepoliciesGet - Get the specified IKE Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIkepoliciesGet(pcloudIkepoliciesGetOptions *PcloudIkepoliciesGetOptions)`, func() {
			pcloudIkepoliciesGetOptions := &powervsv1.PcloudIkepoliciesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IkePolicyID: core.StringPtr("testString"),
			}

			ikePolicy, response, err := powervsService.PcloudIkepoliciesGet(pcloudIkepoliciesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudIkepoliciesPut - Update IKE Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIkepoliciesPut(pcloudIkepoliciesPutOptions *PcloudIkepoliciesPutOptions)`, func() {
			ikePolicyUpdateModel := &powervsv1.IkePolicyUpdate{
				Authentication: core.StringPtr("sha-256"),
				DhGroup: core.Int64Ptr(int64(2)),
				Encryption: core.StringPtr("aes-256-cbc"),
				KeyLifetime: core.Int64Ptr(int64(28800)),
				Name: core.StringPtr("ikePolicy1"),
				PresharedKey: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(2)),
			}
			ikePolicyUpdateModel.SetProperty("foo", core.StringPtr("testString"))

			pcloudIkepoliciesPutOptions := &powervsv1.PcloudIkepoliciesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IkePolicyID: core.StringPtr("testString"),
				IkePolicyUpdate: ikePolicyUpdateModel,
			}

			ikePolicy, response, err := powervsService.PcloudIkepoliciesPut(pcloudIkepoliciesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudIpsecpoliciesGetall - Get all IPSec Policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIpsecpoliciesGetall(pcloudIpsecpoliciesGetallOptions *PcloudIpsecpoliciesGetallOptions)`, func() {
			pcloudIpsecpoliciesGetallOptions := &powervsv1.PcloudIpsecpoliciesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			ipSecPolicies, response, err := powervsService.PcloudIpsecpoliciesGetall(pcloudIpsecpoliciesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipSecPolicies).ToNot(BeNil())
		})
	})

	Describe(`PcloudIpsecpoliciesPost - Add IPSec Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIpsecpoliciesPost(pcloudIpsecpoliciesPostOptions *PcloudIpsecpoliciesPostOptions)`, func() {
			pcloudIpsecpoliciesPostOptions := &powervsv1.PcloudIpsecpoliciesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				DhGroup: core.Int64Ptr(int64(2)),
				Encryption: core.StringPtr("aes-256-cbc"),
				KeyLifetime: core.Int64Ptr(int64(28800)),
				Name: core.StringPtr("ipSecPolicy2"),
				Pfs: core.BoolPtr(true),
				Authentication: core.StringPtr("hmac-sha-256-128"),
			}

			ipSecPolicy, response, err := powervsService.PcloudIpsecpoliciesPost(pcloudIpsecpoliciesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipSecPolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudIpsecpoliciesGet - Get the specified IPSec Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIpsecpoliciesGet(pcloudIpsecpoliciesGetOptions *PcloudIpsecpoliciesGetOptions)`, func() {
			pcloudIpsecpoliciesGetOptions := &powervsv1.PcloudIpsecpoliciesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IpsecPolicyID: core.StringPtr("testString"),
			}

			ipSecPolicy, response, err := powervsService.PcloudIpsecpoliciesGet(pcloudIpsecpoliciesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipSecPolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudIpsecpoliciesPut - Update IPSec Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIpsecpoliciesPut(pcloudIpsecpoliciesPutOptions *PcloudIpsecpoliciesPutOptions)`, func() {
			ipSecPolicyUpdateModel := &powervsv1.IPSecPolicyUpdate{
				Authentication: core.StringPtr("hmac-sha-256-128"),
				DhGroup: core.Int64Ptr(int64(2)),
				Encryption: core.StringPtr("aes-256-cbc"),
				KeyLifetime: core.Int64Ptr(int64(28800)),
				Name: core.StringPtr("ipSecPolicy2"),
				Pfs: core.BoolPtr(true),
			}
			ipSecPolicyUpdateModel.SetProperty("foo", core.StringPtr("testString"))

			pcloudIpsecpoliciesPutOptions := &powervsv1.PcloudIpsecpoliciesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IpsecPolicyID: core.StringPtr("testString"),
				IPSecPolicyUpdate: ipSecPolicyUpdateModel,
			}

			ipSecPolicy, response, err := powervsService.PcloudIpsecpoliciesPut(pcloudIpsecpoliciesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipSecPolicy).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsGetall - Get all volume groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsGetall(pcloudVolumegroupsGetallOptions *PcloudVolumegroupsGetallOptions)`, func() {
			pcloudVolumegroupsGetallOptions := &powervsv1.PcloudVolumegroupsGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			volumeGroups, response, err := powervsService.PcloudVolumegroupsGetall(pcloudVolumegroupsGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroups).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsPost - Create a new volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsPost(pcloudVolumegroupsPostOptions *PcloudVolumegroupsPostOptions)`, func() {
			pcloudVolumegroupsPostOptions := &powervsv1.PcloudVolumegroupsPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
				ConsistencyGroupName: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			volumeGroupCreateResponse, response, err := powervsService.PcloudVolumegroupsPost(pcloudVolumegroupsPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeGroupCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsGetallDetails - Get all volume groups with details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsGetallDetails(pcloudVolumegroupsGetallDetailsOptions *PcloudVolumegroupsGetallDetailsOptions)`, func() {
			pcloudVolumegroupsGetallDetailsOptions := &powervsv1.PcloudVolumegroupsGetallDetailsOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			volumeGroupsDetails, response, err := powervsService.PcloudVolumegroupsGetallDetails(pcloudVolumegroupsGetallDetailsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroupsDetails).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsGet - Get volume Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsGet(pcloudVolumegroupsGetOptions *PcloudVolumegroupsGetOptions)`, func() {
			pcloudVolumegroupsGetOptions := &powervsv1.PcloudVolumegroupsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
			}

			volumeGroup, response, err := powervsService.PcloudVolumegroupsGet(pcloudVolumegroupsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsPut - updates the volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsPut(pcloudVolumegroupsPutOptions *PcloudVolumegroupsPutOptions)`, func() {
			pcloudVolumegroupsPutOptions := &powervsv1.PcloudVolumegroupsPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
				AddVolumes: []string{"testString"},
				RemoveVolumes: []string{"testString"},
			}

			object, response, err := powervsService.PcloudVolumegroupsPut(pcloudVolumegroupsPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsActionPost - Perform an action (start stop reset ) on a volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsActionPost(pcloudVolumegroupsActionPostOptions *PcloudVolumegroupsActionPostOptions)`, func() {
			volumeGroupActionResetModel := &powervsv1.VolumeGroupActionReset{
				Status: core.StringPtr("available"),
			}

			volumeGroupActionStartModel := &powervsv1.VolumeGroupActionStart{
				Source: core.StringPtr("master"),
			}

			volumeGroupActionStopModel := &powervsv1.VolumeGroupActionStop{
				Access: core.BoolPtr(true),
			}

			volumeGroupActionModel := &powervsv1.VolumeGroupAction{
				Reset: volumeGroupActionResetModel,
				Start: volumeGroupActionStartModel,
				Stop: volumeGroupActionStopModel,
			}
			volumeGroupActionModel.SetProperty("foo", core.StringPtr("testString"))

			pcloudVolumegroupsActionPostOptions := &powervsv1.PcloudVolumegroupsActionPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
				VolumeGroupAction: volumeGroupActionModel,
			}

			object, response, err := powervsService.PcloudVolumegroupsActionPost(pcloudVolumegroupsActionPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsGetDetails - Get volume Group details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsGetDetails(pcloudVolumegroupsGetDetailsOptions *PcloudVolumegroupsGetDetailsOptions)`, func() {
			pcloudVolumegroupsGetDetailsOptions := &powervsv1.PcloudVolumegroupsGetDetailsOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
			}

			volumeGroupDetails, response, err := powervsService.PcloudVolumegroupsGetDetails(pcloudVolumegroupsGetDetailsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroupDetails).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsRemoteCopyRelationshipsGet - Get remote copy relationships of the volume belonging to volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsRemoteCopyRelationshipsGet(pcloudVolumegroupsRemoteCopyRelationshipsGetOptions *PcloudVolumegroupsRemoteCopyRelationshipsGetOptions)`, func() {
			pcloudVolumegroupsRemoteCopyRelationshipsGetOptions := &powervsv1.PcloudVolumegroupsRemoteCopyRelationshipsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
			}

			volumeGroupRemoteCopyRelationships, response, err := powervsService.PcloudVolumegroupsRemoteCopyRelationshipsGet(pcloudVolumegroupsRemoteCopyRelationshipsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroupRemoteCopyRelationships).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsStorageDetailsGet - Get storage details of volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsStorageDetailsGet(pcloudVolumegroupsStorageDetailsGetOptions *PcloudVolumegroupsStorageDetailsGetOptions)`, func() {
			pcloudVolumegroupsStorageDetailsGetOptions := &powervsv1.PcloudVolumegroupsStorageDetailsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
			}

			volumeGroupStorageDetails, response, err := powervsService.PcloudVolumegroupsStorageDetailsGet(pcloudVolumegroupsStorageDetailsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeGroupStorageDetails).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumeOnboardingGetall - List all volume onboardings for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumeOnboardingGetall(pcloudVolumeOnboardingGetallOptions *PcloudVolumeOnboardingGetallOptions)`, func() {
			pcloudVolumeOnboardingGetallOptions := &powervsv1.PcloudVolumeOnboardingGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			volumeOnboardings, response, err := powervsService.PcloudVolumeOnboardingGetall(pcloudVolumeOnboardingGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeOnboardings).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumeOnboardingPost - Onboard auxiliary volumes to target site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumeOnboardingPost(pcloudVolumeOnboardingPostOptions *PcloudVolumeOnboardingPostOptions)`, func() {
			auxiliaryVolumeForOnboardingModel := &powervsv1.AuxiliaryVolumeForOnboarding{
				AuxVolumeName: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			auxiliaryVolumesForOnboardingModel := &powervsv1.AuxiliaryVolumesForOnboarding{
				AuxiliaryVolumes: []powervsv1.AuxiliaryVolumeForOnboarding{*auxiliaryVolumeForOnboardingModel},
				SourceCRN: core.StringPtr("testString"),
			}

			pcloudVolumeOnboardingPostOptions := &powervsv1.PcloudVolumeOnboardingPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Volumes: []powervsv1.AuxiliaryVolumesForOnboarding{*auxiliaryVolumesForOnboardingModel},
				Description: core.StringPtr("testString"),
			}

			volumeOnboardingCreateResponse, response, err := powervsService.PcloudVolumeOnboardingPost(pcloudVolumeOnboardingPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumeOnboardingCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumeOnboardingGet - Get the information of volume onboarding operation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumeOnboardingGet(pcloudVolumeOnboardingGetOptions *PcloudVolumeOnboardingGetOptions)`, func() {
			pcloudVolumeOnboardingGetOptions := &powervsv1.PcloudVolumeOnboardingGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeOnboardingID: core.StringPtr("testString"),
			}

			volumeOnboarding, response, err := powervsService.PcloudVolumeOnboardingGet(pcloudVolumeOnboardingGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeOnboarding).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesGetall - List all volumes attached to a PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesGetall(pcloudPvminstancesVolumesGetallOptions *PcloudPvminstancesVolumesGetallOptions)`, func() {
			pcloudPvminstancesVolumesGetallOptions := &powervsv1.PcloudPvminstancesVolumesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
			}

			volumes, response, err := powervsService.PcloudPvminstancesVolumesGetall(pcloudPvminstancesVolumesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumes).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesGet - Detailed info of a volume attached to a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesGet(pcloudPvminstancesVolumesGetOptions *PcloudPvminstancesVolumesGetOptions)`, func() {
			pcloudPvminstancesVolumesGetOptions := &powervsv1.PcloudPvminstancesVolumesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			volume, response, err := powervsService.PcloudPvminstancesVolumesGet(pcloudPvminstancesVolumesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesPost - Attach a volume to a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesPost(pcloudPvminstancesVolumesPostOptions *PcloudPvminstancesVolumesPostOptions)`, func() {
			pcloudPvminstancesVolumesPostOptions := &powervsv1.PcloudPvminstancesVolumesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPvminstancesVolumesPost(pcloudPvminstancesVolumesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesPut - Update a volume attached to a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesPut(pcloudPvminstancesVolumesPutOptions *PcloudPvminstancesVolumesPutOptions)`, func() {
			pcloudPvminstancesVolumesPutOptions := &powervsv1.PcloudPvminstancesVolumesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
				DeleteOnTermination: core.BoolPtr(true),
			}

			object, response, err := powervsService.PcloudPvminstancesVolumesPut(pcloudPvminstancesVolumesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesSetbootPut - Set the PVMInstance volume as the boot volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesSetbootPut(pcloudPvminstancesVolumesSetbootPutOptions *PcloudPvminstancesVolumesSetbootPutOptions)`, func() {
			pcloudPvminstancesVolumesSetbootPutOptions := &powervsv1.PcloudPvminstancesVolumesSetbootPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPvminstancesVolumesSetbootPut(pcloudPvminstancesVolumesSetbootPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesGetall - List all volumes for this cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesGetall(pcloudCloudinstancesVolumesGetallOptions *PcloudCloudinstancesVolumesGetallOptions)`, func() {
			pcloudCloudinstancesVolumesGetallOptions := &powervsv1.PcloudCloudinstancesVolumesGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ReplicationEnabled: core.BoolPtr(true),
				Affinity: core.StringPtr("testString"),
				Auxiliary: core.BoolPtr(true),
			}

			volumes, response, err := powervsService.PcloudCloudinstancesVolumesGetall(pcloudCloudinstancesVolumesGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumes).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesPost - Create a new data Volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesPost(pcloudCloudinstancesVolumesPostOptions *PcloudCloudinstancesVolumesPostOptions)`, func() {
			pcloudCloudinstancesVolumesPostOptions := &powervsv1.PcloudCloudinstancesVolumesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Size: core.Float64Ptr(float64(72.5)),
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
				DiskType: core.StringPtr("testString"),
				ReplicationEnabled: core.BoolPtr(true),
				Shareable: core.BoolPtr(true),
				VolumePool: core.StringPtr("testString"),
			}

			volume, response, err := powervsService.PcloudCloudinstancesVolumesPost(pcloudCloudinstancesVolumesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumesClonePost - Create a volume clone for specified volumes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumesClonePost(pcloudVolumesClonePostOptions *PcloudVolumesClonePostOptions)`, func() {
			pcloudVolumesClonePostOptions := &powervsv1.PcloudVolumesClonePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				DisplayName: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
			}

			volumesCloneResponse, response, err := powervsService.PcloudVolumesClonePost(pcloudVolumesClonePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumesCloneResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesGet - Detailed info of a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesGet(pcloudCloudinstancesVolumesGetOptions *PcloudCloudinstancesVolumesGetOptions)`, func() {
			pcloudCloudinstancesVolumesGetOptions := &powervsv1.PcloudCloudinstancesVolumesGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			volume, response, err := powervsService.PcloudCloudinstancesVolumesGet(pcloudCloudinstancesVolumesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesPut - Update a cloud instance volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesPut(pcloudCloudinstancesVolumesPutOptions *PcloudCloudinstancesVolumesPutOptions)`, func() {
			pcloudCloudinstancesVolumesPutOptions := &powervsv1.PcloudCloudinstancesVolumesPutOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
				Bootable: core.BoolPtr(true),
				Name: core.StringPtr("testString"),
				Shareable: core.BoolPtr(true),
				Size: core.Float64Ptr(float64(72.5)),
			}

			volume, response, err := powervsService.PcloudCloudinstancesVolumesPut(pcloudCloudinstancesVolumesPutOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesActionPost - Perform an action on a Volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesActionPost(pcloudCloudinstancesVolumesActionPostOptions *PcloudCloudinstancesVolumesActionPostOptions)`, func() {
			pcloudCloudinstancesVolumesActionPostOptions := &powervsv1.PcloudCloudinstancesVolumesActionPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
				ReplicationEnabled: core.BoolPtr(true),
			}

			object, response, err := powervsService.PcloudCloudinstancesVolumesActionPost(pcloudCloudinstancesVolumesActionPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesFlashCopyMappingsGet - Get a list of flashcopy mappings of a given volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesFlashCopyMappingsGet(pcloudCloudinstancesVolumesFlashCopyMappingsGetOptions *PcloudCloudinstancesVolumesFlashCopyMappingsGetOptions)`, func() {
			pcloudCloudinstancesVolumesFlashCopyMappingsGetOptions := &powervsv1.PcloudCloudinstancesVolumesFlashCopyMappingsGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			flashCopyMapping, response, err := powervsService.PcloudCloudinstancesVolumesFlashCopyMappingsGet(pcloudCloudinstancesVolumesFlashCopyMappingsGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flashCopyMapping).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesRemoteCopyRelationshipGet - Get remote copy relationship of a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesRemoteCopyRelationshipGet(pcloudCloudinstancesVolumesRemoteCopyRelationshipGetOptions *PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOptions)`, func() {
			pcloudCloudinstancesVolumesRemoteCopyRelationshipGetOptions := &powervsv1.PcloudCloudinstancesVolumesRemoteCopyRelationshipGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			volumeRemoteCopyRelationship, response, err := powervsService.PcloudCloudinstancesVolumesRemoteCopyRelationshipGet(pcloudCloudinstancesVolumesRemoteCopyRelationshipGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeRemoteCopyRelationship).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2PvminstancesVolumesPost - Attach all volumes to a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2PvminstancesVolumesPost(pcloudV2PvminstancesVolumesPostOptions *PcloudV2PvminstancesVolumesPostOptions)`, func() {
			pcloudV2PvminstancesVolumesPostOptions := &powervsv1.PcloudV2PvminstancesVolumesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
				PerformancePath: core.BoolPtr(true),
			}

			volumesAttachmentResponse, response, err := powervsService.PcloudV2PvminstancesVolumesPost(pcloudV2PvminstancesVolumesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumesAttachmentResponse).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumesPost - Create multiple data volumes from a single definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumesPost(pcloudV2VolumesPostOptions *PcloudV2VolumesPostOptions)`, func() {
			pcloudV2VolumesPostOptions := &powervsv1.PcloudV2VolumesPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Size: core.Int64Ptr(int64(38)),
				AffinityPvmInstance: core.StringPtr("testString"),
				AffinityPolicy: core.StringPtr("affinity"),
				AffinityVolume: core.StringPtr("testString"),
				AntiAffinityPvmInstances: []string{"testString"},
				AntiAffinityVolumes: []string{"testString"},
				Count: core.Int64Ptr(int64(38)),
				DiskType: core.StringPtr("testString"),
				ReplicationEnabled: core.BoolPtr(true),
				Shareable: core.BoolPtr(true),
				VolumePool: core.StringPtr("testString"),
			}

			volumes, response, err := powervsService.PcloudV2VolumesPost(pcloudV2VolumesPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumes).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneGetall - Get the list of volumes-clone request for a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneGetall(pcloudV2VolumescloneGetallOptions *PcloudV2VolumescloneGetallOptions)`, func() {
			pcloudV2VolumescloneGetallOptions := &powervsv1.PcloudV2VolumescloneGetallOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Filter: core.StringPtr("prepare"),
			}

			volumesClones, response, err := powervsService.PcloudV2VolumescloneGetall(pcloudV2VolumescloneGetallOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumesClones).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumesclonePost - Create a new volumes clone request and initiates the Prepare action   Requires a minimum of two volumes   Requires a minimum of one volume to be in the 'in-use' state   Requires a unique volumes clone name   Prepare action does the preparatory work for creating the snapshot volumes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumesclonePost(pcloudV2VolumesclonePostOptions *PcloudV2VolumesclonePostOptions)`, func() {
			pcloudV2VolumesclonePostOptions := &powervsv1.PcloudV2VolumesclonePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
				PerformancePath: core.BoolPtr(true),
			}

			volumesClone, response, err := powervsService.PcloudV2VolumesclonePost(pcloudV2VolumesclonePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumesClone).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneGet - Get the details for a volumes-clone request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneGet(pcloudV2VolumescloneGetOptions *PcloudV2VolumescloneGetOptions)`, func() {
			pcloudV2VolumescloneGetOptions := &powervsv1.PcloudV2VolumescloneGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumesCloneID: core.StringPtr("testString"),
			}

			volumesCloneDetail, response, err := powervsService.PcloudV2VolumescloneGet(pcloudV2VolumescloneGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumesCloneDetail).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneCancelPost - Cancel a volumes-clone request, initiates the Cleanup action Cleanup action performs the cleanup of the preparatory clones and snapshot volumes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneCancelPost(pcloudV2VolumescloneCancelPostOptions *PcloudV2VolumescloneCancelPostOptions)`, func() {
			pcloudV2VolumescloneCancelPostOptions := &powervsv1.PcloudV2VolumescloneCancelPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumesCloneID: core.StringPtr("testString"),
				Force: core.BoolPtr(true),
			}

			volumesClone, response, err := powervsService.PcloudV2VolumescloneCancelPost(pcloudV2VolumescloneCancelPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumesClone).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneExecutePost - Initiate the Execute action for a volumes-clone request Execute action creates the cloned volumes using the volume snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneExecutePost(pcloudV2VolumescloneExecutePostOptions *PcloudV2VolumescloneExecutePostOptions)`, func() {
			pcloudV2VolumescloneExecutePostOptions := &powervsv1.PcloudV2VolumescloneExecutePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumesCloneID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				RollbackPrepare: core.BoolPtr(true),
			}

			volumesClone, response, err := powervsService.PcloudV2VolumescloneExecutePost(pcloudV2VolumescloneExecutePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(volumesClone).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneStartPost - Initiate the Start action for a volumes-clone request Start action starts the consistency group to initiate the flash copy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneStartPost(pcloudV2VolumescloneStartPostOptions *PcloudV2VolumescloneStartPostOptions)`, func() {
			pcloudV2VolumescloneStartPostOptions := &powervsv1.PcloudV2VolumescloneStartPostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumesCloneID: core.StringPtr("testString"),
			}

			volumesClone, response, err := powervsService.PcloudV2VolumescloneStartPost(pcloudV2VolumescloneStartPostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumesClone).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumesClonePost - Create a volume clone for specified volumes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumesClonePost(pcloudV2VolumesClonePostOptions *PcloudV2VolumesClonePostOptions)`, func() {
			pcloudV2VolumesClonePostOptions := &powervsv1.PcloudV2VolumesClonePostOptions{
				CloudInstanceID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				VolumeIDs: []string{"testString"},
			}

			cloneTaskReference, response, err := powervsService.PcloudV2VolumesClonePost(pcloudV2VolumesClonePostOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(cloneTaskReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumesClonetasksGet - Get the status of a volumes clone request for the specified clone task ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumesClonetasksGet(pcloudV2VolumesClonetasksGetOptions *PcloudV2VolumesClonetasksGetOptions)`, func() {
			pcloudV2VolumesClonetasksGetOptions := &powervsv1.PcloudV2VolumesClonetasksGetOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloneTaskID: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
			}

			cloneTaskStatus, response, err := powervsService.PcloudV2VolumesClonetasksGet(pcloudV2VolumesClonetasksGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloneTaskStatus).ToNot(BeNil())
		})
	})

	Describe(`ServiceBindingGet - gets a service binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBindingGet(serviceBindingGetOptions *ServiceBindingGetOptions)`, func() {
			serviceBindingGetOptions := &powervsv1.ServiceBindingGetOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				BindingID: core.StringPtr("testString"),
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
			}

			serviceBindingResource, response, err := powervsService.ServiceBindingGet(serviceBindingGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceBindingResource).ToNot(BeNil())
		})
	})

	Describe(`ServiceBindingBinding - generation of a service binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBindingBinding(serviceBindingBindingOptions *ServiceBindingBindingOptions)`, func() {
			serviceBindingResourceObjectModel := &powervsv1.ServiceBindingResourceObject{
				AppGuid: core.StringPtr("testString"),
				Route: core.StringPtr("testString"),
			}

			contextModel := &powervsv1.Context{
			}
			contextModel.SetProperty("foo", core.StringPtr("testString"))

			objectModel := &powervsv1.Object{
			}
			objectModel.SetProperty("foo", core.StringPtr("testString"))

			serviceBindingBindingOptions := &powervsv1.ServiceBindingBindingOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				BindingID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				AppGuid: core.StringPtr("testString"),
				BindResource: serviceBindingResourceObjectModel,
				Context: contextModel,
				Parameters: objectModel,
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
				AcceptsIncomplete: core.BoolPtr(true),
			}

			serviceBinding, response, err := powervsService.ServiceBindingBinding(serviceBindingBindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceBinding).ToNot(BeNil())
		})
	})

	Describe(`ServiceBindingLastOperationGet - last requested operation state for service binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBindingLastOperationGet(serviceBindingLastOperationGetOptions *ServiceBindingLastOperationGetOptions)`, func() {
			serviceBindingLastOperationGetOptions := &powervsv1.ServiceBindingLastOperationGetOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				BindingID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Operation: core.StringPtr("testString"),
			}

			lastOperationResource, response, err := powervsService.ServiceBindingLastOperationGet(serviceBindingLastOperationGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(lastOperationResource).ToNot(BeNil())
		})
	})

	Describe(`ServiceInstanceGet - gets a service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceInstanceGet(serviceInstanceGetOptions *ServiceInstanceGetOptions)`, func() {
			serviceInstanceGetOptions := &powervsv1.ServiceInstanceGetOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
			}

			serviceInstanceResource, response, err := powervsService.ServiceInstanceGet(serviceInstanceGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceResource).ToNot(BeNil())
		})
	})

	Describe(`ServiceInstanceUpdate - update a service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceInstanceUpdate(serviceInstanceUpdateOptions *ServiceInstanceUpdateOptions)`, func() {
			contextModel := &powervsv1.Context{
			}
			contextModel.SetProperty("foo", core.StringPtr("testString"))

			objectModel := &powervsv1.Object{
			}
			objectModel.SetProperty("foo", core.StringPtr("testString"))

			serviceInstancePreviousValuesModel := &powervsv1.ServiceInstancePreviousValues{
				OrganizationID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				SpaceID: core.StringPtr("testString"),
			}

			serviceInstanceUpdateOptions := &powervsv1.ServiceInstanceUpdateOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				Context: contextModel,
				Parameters: objectModel,
				PlanID: core.StringPtr("testString"),
				PreviousValues: serviceInstancePreviousValuesModel,
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
				AcceptsIncomplete: core.BoolPtr(true),
			}

			object, response, err := powervsService.ServiceInstanceUpdate(serviceInstanceUpdateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`ServiceInstanceProvision - provision a service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceInstanceProvision(serviceInstanceProvisionOptions *ServiceInstanceProvisionOptions)`, func() {
			contextModel := &powervsv1.Context{
			}
			contextModel.SetProperty("foo", core.StringPtr("testString"))

			objectModel := &powervsv1.Object{
			}
			objectModel.SetProperty("foo", core.StringPtr("testString"))

			serviceInstanceProvisionOptions := &powervsv1.ServiceInstanceProvisionOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				Context: contextModel,
				OrganizationGuid: core.StringPtr("testString"),
				Parameters: objectModel,
				SpaceGuid: core.StringPtr("testString"),
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
				AcceptsIncomplete: core.BoolPtr(true),
			}

			serviceInstanceProvision, response, err := powervsService.ServiceInstanceProvision(serviceInstanceProvisionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceProvision).ToNot(BeNil())
		})
	})

	Describe(`ServiceInstanceLastOperationGet - last requested operation state for service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceInstanceLastOperationGet(serviceInstanceLastOperationGetOptions *ServiceInstanceLastOperationGetOptions)`, func() {
			serviceInstanceLastOperationGetOptions := &powervsv1.ServiceInstanceLastOperationGetOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Operation: core.StringPtr("testString"),
			}

			lastOperationResource, response, err := powervsService.ServiceInstanceLastOperationGet(serviceInstanceLastOperationGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(lastOperationResource).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerStoragetypesGet - Available storage types in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerStoragetypesGet(serviceBrokerStoragetypesGetOptions *ServiceBrokerStoragetypesGetOptions)`, func() {
			serviceBrokerStoragetypesGetOptions := &powervsv1.ServiceBrokerStoragetypesGetOptions{
			}

			mapStringStorageType, response, err := powervsService.ServiceBrokerStoragetypesGet(serviceBrokerStoragetypesGetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringStorageType).ToNot(BeNil())
		})
	})

	Describe(`ServiceBrokerSwaggerspec - Get swagger json spec`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBrokerSwaggerspec(serviceBrokerSwaggerspecOptions *ServiceBrokerSwaggerspecOptions)`, func() {
			serviceBrokerSwaggerspecOptions := &powervsv1.ServiceBrokerSwaggerspecOptions{
			}

			object, response, err := powervsService.ServiceBrokerSwaggerspec(serviceBrokerSwaggerspecOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsDelete - Delete a Cloud Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsDelete(pcloudCloudconnectionsDeleteOptions *PcloudCloudconnectionsDeleteOptions)`, func() {
			pcloudCloudconnectionsDeleteOptions := &powervsv1.PcloudCloudconnectionsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudconnectionsDelete(pcloudCloudconnectionsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudconnectionsNetworksDelete - Detach a network from a Cloud Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudconnectionsNetworksDelete(pcloudCloudconnectionsNetworksDeleteOptions *PcloudCloudconnectionsNetworksDeleteOptions)`, func() {
			pcloudCloudconnectionsNetworksDeleteOptions := &powervsv1.PcloudCloudconnectionsNetworksDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				CloudConnectionID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudconnectionsNetworksDelete(pcloudCloudconnectionsNetworksDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesImagesDelete - Delete an Image from a Cloud Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesImagesDelete(pcloudCloudinstancesImagesDeleteOptions *PcloudCloudinstancesImagesDeleteOptions)`, func() {
			pcloudCloudinstancesImagesDeleteOptions := &powervsv1.PcloudCloudinstancesImagesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				ImageID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesImagesDelete(pcloudCloudinstancesImagesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesDelete - Delete a Power Cloud Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesDelete(pcloudCloudinstancesDeleteOptions *PcloudCloudinstancesDeleteOptions)`, func() {
			pcloudCloudinstancesDeleteOptions := &powervsv1.PcloudCloudinstancesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesDelete(pcloudCloudinstancesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesJobsDelete - Delete a cloud instance job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesJobsDelete(pcloudCloudinstancesJobsDeleteOptions *PcloudCloudinstancesJobsDeleteOptions)`, func() {
			pcloudCloudinstancesJobsDeleteOptions := &powervsv1.PcloudCloudinstancesJobsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				JobID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesJobsDelete(pcloudCloudinstancesJobsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksDelete - Delete a Network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksDelete(pcloudNetworksDeleteOptions *PcloudNetworksDeleteOptions)`, func() {
			pcloudNetworksDeleteOptions := &powervsv1.PcloudNetworksDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudNetworksDelete(pcloudNetworksDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudNetworksPortsDelete - Delete a Network Port`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudNetworksPortsDelete(pcloudNetworksPortsDeleteOptions *PcloudNetworksPortsDeleteOptions)`, func() {
			pcloudNetworksPortsDeleteOptions := &powervsv1.PcloudNetworksPortsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				PortID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudNetworksPortsDelete(pcloudNetworksPortsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesDelete - Delete a PCloud PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesDelete(pcloudPvminstancesDeleteOptions *PcloudPvminstancesDeleteOptions)`, func() {
			pcloudPvminstancesDeleteOptions := &powervsv1.PcloudPvminstancesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				DeleteDataVolumes: core.BoolPtr(true),
			}

			object, response, err := powervsService.PcloudPvminstancesDelete(pcloudPvminstancesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesNetworksDelete - Remove all Address of Network from a PVM Instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesNetworksDelete(pcloudPvminstancesNetworksDeleteOptions *PcloudPvminstancesNetworksDeleteOptions)`, func() {
			pcloudPvminstancesNetworksDeleteOptions := &powervsv1.PcloudPvminstancesNetworksDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("testString"),
				MacAddress: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPvminstancesNetworksDelete(pcloudPvminstancesNetworksDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsDelete - Delete Server Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsDelete(pcloudPlacementgroupsDeleteOptions *PcloudPlacementgroupsDeleteOptions)`, func() {
			pcloudPlacementgroupsDeleteOptions := &powervsv1.PcloudPlacementgroupsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PlacementGroupID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPlacementgroupsDelete(pcloudPlacementgroupsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPlacementgroupsMembersDelete - Remove Server from Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPlacementgroupsMembersDelete(pcloudPlacementgroupsMembersDeleteOptions *PcloudPlacementgroupsMembersDeleteOptions)`, func() {
			pcloudPlacementgroupsMembersDeleteOptions := &powervsv1.PcloudPlacementgroupsMembersDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PlacementGroupID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			placementGroup, response, err := powervsService.PcloudPlacementgroupsMembersDelete(pcloudPlacementgroupsMembersDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsDelete - Delete a Shared Processor Pool Placement Group from a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsDelete(pcloudSppplacementgroupsDeleteOptions *PcloudSppplacementgroupsDeleteOptions)`, func() {
			pcloudSppplacementgroupsDeleteOptions := &powervsv1.PcloudSppplacementgroupsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SppPlacementGroupID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudSppplacementgroupsDelete(pcloudSppplacementgroupsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudSppplacementgroupsMembersDelete - Delete Shared Processor Pool member from a Shared Processor Pool Placement Group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSppplacementgroupsMembersDelete(pcloudSppplacementgroupsMembersDeleteOptions *PcloudSppplacementgroupsMembersDeleteOptions)`, func() {
			pcloudSppplacementgroupsMembersDeleteOptions := &powervsv1.PcloudSppplacementgroupsMembersDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SppPlacementGroupID: core.StringPtr("testString"),
				SharedProcessorPoolID: core.StringPtr("testString"),
			}

			sppPlacementGroup, response, err := powervsService.PcloudSppplacementgroupsMembersDelete(pcloudSppplacementgroupsMembersDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sppPlacementGroup).ToNot(BeNil())
		})
	})

	Describe(`PcloudDhcpDelete - Delete DHCP Server (OpenShift Internal Use Only)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudDhcpDelete(pcloudDhcpDeleteOptions *PcloudDhcpDeleteOptions)`, func() {
			pcloudDhcpDeleteOptions := &powervsv1.PcloudDhcpDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				DhcpID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudDhcpDelete(pcloudDhcpDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudSharedprocessorpoolsDelete - Delete a Shared Processor Pool from a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudSharedprocessorpoolsDelete(pcloudSharedprocessorpoolsDeleteOptions *PcloudSharedprocessorpoolsDeleteOptions)`, func() {
			pcloudSharedprocessorpoolsDeleteOptions := &powervsv1.PcloudSharedprocessorpoolsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SharedProcessorPoolID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudSharedprocessorpoolsDelete(pcloudSharedprocessorpoolsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesSnapshotsDelete - Delete a PVM instance snapshot of a cloud instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesSnapshotsDelete(pcloudCloudinstancesSnapshotsDeleteOptions *PcloudCloudinstancesSnapshotsDeleteOptions)`, func() {
			pcloudCloudinstancesSnapshotsDeleteOptions := &powervsv1.PcloudCloudinstancesSnapshotsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				SnapshotID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesSnapshotsDelete(pcloudCloudinstancesSnapshotsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudTasksDelete - Delete a Task`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTasksDelete(pcloudTasksDeleteOptions *PcloudTasksDeleteOptions)`, func() {
			pcloudTasksDeleteOptions := &powervsv1.PcloudTasksDeleteOptions{
				TaskID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudTasksDelete(pcloudTasksDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudTenantsSshkeysDelete - Delete a Tenant's SSH key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudTenantsSshkeysDelete(pcloudTenantsSshkeysDeleteOptions *PcloudTenantsSshkeysDeleteOptions)`, func() {
			pcloudTenantsSshkeysDeleteOptions := &powervsv1.PcloudTenantsSshkeysDeleteOptions{
				TenantID: core.StringPtr("testString"),
				SshkeyName: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudTenantsSshkeysDelete(pcloudTenantsSshkeysDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsDelete - Delete VPN Connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsDelete(pcloudVpnconnectionsDeleteOptions *PcloudVpnconnectionsDeleteOptions)`, func() {
			pcloudVpnconnectionsDeleteOptions := &powervsv1.PcloudVpnconnectionsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
			}

			jobReference, response, err := powervsService.PcloudVpnconnectionsDelete(pcloudVpnconnectionsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsNetworksDelete - Detach network`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsNetworksDelete(pcloudVpnconnectionsNetworksDeleteOptions *PcloudVpnconnectionsNetworksDeleteOptions)`, func() {
			pcloudVpnconnectionsNetworksDeleteOptions := &powervsv1.PcloudVpnconnectionsNetworksDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
				NetworkID: core.StringPtr("7f950c76-8582-11qeb-8dcd-0242ac172"),
			}

			jobReference, response, err := powervsService.PcloudVpnconnectionsNetworksDelete(pcloudVpnconnectionsNetworksDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(jobReference).ToNot(BeNil())
		})
	})

	Describe(`PcloudVpnconnectionsPeersubnetsDelete - Detach Peer Subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVpnconnectionsPeersubnetsDelete(pcloudVpnconnectionsPeersubnetsDeleteOptions *PcloudVpnconnectionsPeersubnetsDeleteOptions)`, func() {
			pcloudVpnconnectionsPeersubnetsDeleteOptions := &powervsv1.PcloudVpnconnectionsPeersubnetsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VPNConnectionID: core.StringPtr("testString"),
				CIDR: core.StringPtr("128.170.1.0/32"),
			}

			peerSubnets, response, err := powervsService.PcloudVpnconnectionsPeersubnetsDelete(pcloudVpnconnectionsPeersubnetsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(peerSubnets).ToNot(BeNil())
		})
	})

	Describe(`PcloudIkepoliciesDelete - Delete IKE Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIkepoliciesDelete(pcloudIkepoliciesDeleteOptions *PcloudIkepoliciesDeleteOptions)`, func() {
			pcloudIkepoliciesDeleteOptions := &powervsv1.PcloudIkepoliciesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IkePolicyID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudIkepoliciesDelete(pcloudIkepoliciesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudIpsecpoliciesDelete - Delete IPSec Policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudIpsecpoliciesDelete(pcloudIpsecpoliciesDeleteOptions *PcloudIpsecpoliciesDeleteOptions)`, func() {
			pcloudIpsecpoliciesDeleteOptions := &powervsv1.PcloudIpsecpoliciesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				IpsecPolicyID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudIpsecpoliciesDelete(pcloudIpsecpoliciesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudVolumegroupsDelete - Delete a cloud instance volume group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudVolumegroupsDelete(pcloudVolumegroupsDeleteOptions *PcloudVolumegroupsDeleteOptions)`, func() {
			pcloudVolumegroupsDeleteOptions := &powervsv1.PcloudVolumegroupsDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeGroupID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudVolumegroupsDelete(pcloudVolumegroupsDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudPvminstancesVolumesDelete - Detach a volume from a PVMInstance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudPvminstancesVolumesDelete(pcloudPvminstancesVolumesDeleteOptions *PcloudPvminstancesVolumesDeleteOptions)`, func() {
			pcloudPvminstancesVolumesDeleteOptions := &powervsv1.PcloudPvminstancesVolumesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				PvmInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudPvminstancesVolumesDelete(pcloudPvminstancesVolumesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudCloudinstancesVolumesDelete - Delete a cloud instance volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudCloudinstancesVolumesDelete(pcloudCloudinstancesVolumesDeleteOptions *PcloudCloudinstancesVolumesDeleteOptions)`, func() {
			pcloudCloudinstancesVolumesDeleteOptions := &powervsv1.PcloudCloudinstancesVolumesDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumeID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudCloudinstancesVolumesDelete(pcloudCloudinstancesVolumesDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`PcloudV2VolumescloneDelete - Delete a volumes-clone request`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PcloudV2VolumescloneDelete(pcloudV2VolumescloneDeleteOptions *PcloudV2VolumescloneDeleteOptions)`, func() {
			pcloudV2VolumescloneDeleteOptions := &powervsv1.PcloudV2VolumescloneDeleteOptions{
				CloudInstanceID: core.StringPtr("testString"),
				VolumesCloneID: core.StringPtr("testString"),
			}

			object, response, err := powervsService.PcloudV2VolumescloneDelete(pcloudV2VolumescloneDeleteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`ServiceBindingUnbinding - deprovision of a service binding`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceBindingUnbinding(serviceBindingUnbindingOptions *ServiceBindingUnbindingOptions)`, func() {
			serviceBindingUnbindingOptions := &powervsv1.ServiceBindingUnbindingOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				BindingID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
				AcceptsIncomplete: core.BoolPtr(true),
			}

			object, response, err := powervsService.ServiceBindingUnbinding(serviceBindingUnbindingOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})

	Describe(`ServiceInstanceDeprovision - deprovision a service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ServiceInstanceDeprovision(serviceInstanceDeprovisionOptions *ServiceInstanceDeprovisionOptions)`, func() {
			serviceInstanceDeprovisionOptions := &powervsv1.ServiceInstanceDeprovisionOptions{
				XBrokerApiVersion: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				ServiceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				XBrokerApiOriginatingIdentity: core.StringPtr("testString"),
				AcceptsIncomplete: core.BoolPtr(true),
			}

			object, response, err := powervsService.ServiceInstanceDeprovision(serviceInstanceDeprovisionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(object).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
