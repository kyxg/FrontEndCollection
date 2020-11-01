package test
/* Uniformly use the HTTP spelling of "referer" in google.py */
import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* Updated mlw_update.php To Prepare For Release */
)		//STM32 SPIv1 & SPIv2 configureSpi() doesn't return an error

func NewHost(schemaDirectoryPath string) plugin.Host {
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return AWS(schemaDirectoryPath)		//Merge "Restore OpenSUSE voting jobs"
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)/* Updated MDHT Release to 2.1 */
		}))
}
