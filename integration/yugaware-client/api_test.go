package yugaware_client_test

import (
	"context"
	"encoding/pem"
	"time"

	"github.com/go-openapi/strfmt"
	. "github.com/icza/gox/gox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/access_keys"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/certificate_info"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/cloud_providers"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/session_management"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/universe_cluster_mutations"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/client/universe_management"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
	"github.com/yugabyte/yb-tools/yugaware-client/pkg/cmdutil"
)

var _ = Describe("Yugaware API Compatibility Tests", func() {
	Context("AccessKeys API", func() {
		// TODO: This will only return access keys when a non-kubernetes provider is used
		PWhen("Listing access keys", func() {
			var (
				accessKeys []*models.AccessKey
			)
			BeforeEach(func() {
				universe := CreateTestUniverseIfNotExists()

				params := access_keys.NewListParams().
					WithCUUID(ywclient.CustomerUUID()).
					WithPUUID(strfmt.UUID(universe.UniverseDetails.Clusters[0].UserIntent.Provider))

				keysResponse, err := ywclient.PlatformAPIs.AccessKeys.List(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())
				accessKeys = keysResponse.GetPayload()
			})

			It("Returns a list of access keys", func() {
				Expect(len(accessKeys)).NotTo(BeZero())
			})
		})
	})

	Context("CertificateInfo", func() {
		When("Listing Certificates", func() {
			var (
				certificates []*models.CertificateInfo
				certificate  *models.CertificateRoot
			)

			BeforeEach(func() {
				_ = CreateTestUniverseIfNotExists()

				params := certificate_info.NewGetListOfCertificateParams().WithCUUID(ywclient.CustomerUUID())

				certificatesResponse, err := ywclient.PlatformAPIs.CertificateInfo.GetListOfCertificate(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())
				certificates = certificatesResponse.GetPayload()
			})
			It("Returns a list of cloud certificates", func() {
				Expect(len(certificates)).NotTo(BeZero())
			})

			When("Getting a root certificate", func() {
				BeforeEach(func() {
					params := certificate_info.NewGetRootCertParams().
						WithCUUID(ywclient.CustomerUUID()).
						WithRUUID(certificates[0].UUID)

					certificateResponse, err := ywclient.PlatformAPIs.CertificateInfo.GetRootCert(params, ywclient.SwaggerAuth)
					Expect(err).NotTo(HaveOccurred())
					certificate = certificateResponse.GetPayload()
				})

				It("Returns a certificate", func() {
					block, _ := pem.Decode([]byte(certificate.RootCrt))
					Expect(block.Type).To(Equal("CERTIFICATE"))
				})
			})
		})
	})

	Context("CloudProviders API", func() {
		When("Listing cloud providers", func() {
			var (
				providers []*models.Provider
			)

			BeforeEach(func() {
				params := cloud_providers.NewGetListOfProvidersParams().
					WithCUUID(ywclient.CustomerUUID())
				response, err := ywclient.PlatformAPIs.CloudProviders.GetListOfProviders(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())
				providers = response.GetPayload()
			})
			It("Returns a list of cloud providers", func() {
				Expect(len(providers)).NotTo(BeZero())
			})
		})
	})

	Context("SessionManagement API", func() {
		When("Determining a server version", func() {
			var (
				appVersionReturn map[string]string
			)
			BeforeEach(func() {
				params := session_management.NewAppVersionParams().WithDefaults()
				response, err := ywclient.PlatformAPIs.SessionManagement.AppVersion(params)
				Expect(err).NotTo(HaveOccurred())

				appVersionReturn = response.GetPayload()
			})
			It("Returns a valid version string", func() {
				Expect(appVersionReturn).To(gstruct.MatchKeys(gstruct.IgnoreExtras, gstruct.Keys{"version": MatchRegexp(`\d+\.\d+\.\d-b\d+`)}))
			})
		})

		When("Requesting the customer count", func() {
			var (
				customerCount *int32
			)
			BeforeEach(func() {
				params := session_management.NewCustomerCountParams().WithDefaults()
				response, err := ywclient.PlatformAPIs.SessionManagement.CustomerCount(params)
				Expect(err).NotTo(HaveOccurred())

				customerCount = response.GetPayload().Count
			})
			It("Returns a valid version string", func() {
				Expect(*customerCount).To(Equal(int32(1)))
			})
		})

		// TODO: investigate why this API call fails on yugaware 2.11.2
		//------------------------------
		//• Failure in Spec Setup (BeforeEach) [98.760 seconds]
		//Yugaware API Compatibility Tests
		///tmp/build/be563b6d/yb-tools/integration/yugaware-client/api_test.go:16
		//  SessionManagement API
		//  /tmp/build/be563b6d/yb-tools/integration/yugaware-client/api_test.go:17
		//    when Requesting filtered yugaware logs
		//    /tmp/build/be563b6d/yb-tools/integration/yugaware-client/api_test.go:50
		//      Returns valid logs [BeforeEach]
		//      /tmp/build/be563b6d/yb-tools/integration/yugaware-client/api_test.go:68
		//
		//      Unexpected error:
		//          <*errors.errorString | 0xc0000a5db0>: {
		//              s: "&{[]} (*models.LogData) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface",
		//          }
		//          &{[]} (*models.LogData) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface
		//      occurred
		//
		//      /tmp/build/be563b6d/yb-tools/integration/yugaware-client/api_test.go:64
		XWhen("Requesting filtered yugaware logs", func() {
			var (
				universe *models.UniverseResp
				logLines []string
			)
			BeforeEach(func() {
				universe = CreateTestUniverseIfNotExists()

				params := (session_management.NewGetFilteredLogsParams().
					WithMaxLines(NewInt32(5))).
					WithUniverseName(NewString(universe.Name)).
					WithTimeout(time.Duration(options.DialTimeout) * time.Second)

				response, err := ywclient.PlatformAPIs.SessionManagement.GetFilteredLogs(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())

				logLines = response.GetPayload().Lines
			})
			It("Returns valid logs", func() {
				Expect(logLines).To(ContainElement(ContainSubstring(string(universe.UniverseUUID))))
			})
		})

		When("Requesting yugaware logs", func() {
			var (
				logLines []string
			)
			BeforeEach(func() {
				params := session_management.NewGetLogsParams().
					WithMaxLines(int32(10))

				response, err := ywclient.PlatformAPIs.SessionManagement.GetLogs(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())

				logLines = response.GetPayload().Lines
			})
			It("Returns valid logs", func() {
				Expect(logLines).To(HaveLen(10))
				Expect(logLines).To(ContainElement(MatchRegexp(`^\d{4,4}-\d{2,2}-\d{2,2} \d{2,2}:\d{2,2}:\d{2,2},\d{3,3}`)))
			})
		})

		// TODO: This returns a value similar to:
		//          {"minLength":8,"minUppercase":1,"minLowercase":1,"minDigits":1,"minSpecialCharacters":1}
		//       The default swagger.json does not mark these values as required
		XWhen("Requesting the password policy", func() {
			var (
				err error
			)
			BeforeEach(func() {
				params := session_management.NewGetPasswordPolicyParams().
					WithCUUID(ywclient.CustomerUUID())

				err = ywclient.PlatformAPIs.SessionManagement.GetPasswordPolicy(params)

			})
			It("Returns a valid password policy", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("Requesting the session info", func() {
			var (
				sessionInfo *models.SessionInfo
			)
			BeforeEach(func() {
				params := session_management.NewGetSessionInfoParams().
					WithDefaults()

				response, err := ywclient.PlatformAPIs.SessionManagement.GetSessionInfo(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())

				sessionInfo = response.GetPayload()
			})
			It("Returns a valid password policy", func() {
				Expect(sessionInfo.CustomerUUID).To(Equal(ywclient.CustomerUUID()))
			})
		})
	})

	Context("UniverseClusterMutations API", func() {
		When("Updating an existing universe", func() {
			var (
				newNodeCount int32
				universe     *models.UniverseResp
			)
			BeforeEach(func() {
				originalUniverse := CreateTestUniverseIfNotExists()

				clusters := originalUniverse.UniverseDetails.Clusters
				newNodeCount = clusters[0].UserIntent.NumNodes + 1
				clusters[0].UserIntent.NumNodes = newNodeCount

				params := universe_cluster_mutations.NewUpdatePrimaryClusterParams().
					WithCUUID(ywclient.CustomerUUID()).
					WithUniUUID(originalUniverse.UniverseUUID).
					WithUniverseConfigureTaskParams(&models.UniverseConfigureTaskParams{
						Clusters:     clusters,
						UniverseUUID: originalUniverse.UniverseUUID,
					})

				response, err := ywclient.PlatformAPIs.UniverseClusterMutations.UpdatePrimaryCluster(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())
				err = cmdutil.WaitForTaskCompletion(context.Background(), ywclient, response.GetPayload())
				Expect(err).NotTo(HaveOccurred())

				universe = GetTestUniverse()
			})
			It("The universe increases its node count", func() {
				Expect(universe.UniverseDetails.Clusters[0].UserIntent.NumNodes).To(BeNumerically("==", newNodeCount))

				var tservers []*models.NodeDetailsResp
				for _, node := range universe.UniverseDetails.NodeDetailsSet {
					if node.IsTserver {
						tservers = append(tservers, node)
					}
				}

				Expect(tservers).To(HaveLen(int(newNodeCount)))
			})
		})
	})

	Context("UniverseManagement API", func() {
		When("Resetting the universe version", func() {
			var resetResponse *models.YBPSuccess
			BeforeEach(func() {
				universe := CreateTestUniverseIfNotExists()

				// This requires a dummy value to work around https://yugabyte.atlassian.net/browse/PLAT-2076
				params := universe_management.NewResetUniverseVersionParams().
					WithCUUID(ywclient.CustomerUUID()).
					WithUniUUID(universe.UniverseUUID).
					WithDummy(&models.DummyBody{})

				response, err := ywclient.PlatformAPIs.UniverseManagement.ResetUniverseVersion(params, ywclient.SwaggerAuth)
				Expect(err).NotTo(HaveOccurred())

				resetResponse = response.GetPayload()
			})
			It("The universe version is properly reset", func() {
				Expect(*resetResponse.Success).To(BeTrue())
			})
		})

	})
})
