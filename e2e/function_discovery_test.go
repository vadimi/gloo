package e2e

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	"github.com/solo-io/gloo-plugins/grpc"
	"github.com/solo-io/gloo-plugins/nats-streaming"
	"github.com/solo-io/gloo-plugins/rest"
	"github.com/solo-io/gloo/pkg/log"
)

var _ = Describe("Function Discovery Service Detection", func() {
	for _, test := range []struct {
		description         string
		upstreamName        string
		expectedServiceInfo *v1.ServiceInfo
	}{
		{
			description:         "grpc",
			upstreamName:        namespace + "-grpc-test-service-8080",
			expectedServiceInfo: &v1.ServiceInfo{Type: grpc.ServiceTypeGRPC},
		},
		{
			description:         "nats",
			upstreamName:        namespace + "-nats-streaming-4222",
			expectedServiceInfo: &v1.ServiceInfo{Type: natsstreaming.ServiceTypeNatsStreaming},
		},
		{
			description:         "swagger",
			upstreamName:        namespace + "-petstore-8080",
			expectedServiceInfo: &v1.ServiceInfo{Type: rest.ServiceTypeREST},
		},
	} {
		Context("discovery for "+test.description+" upstreams", func() {
			It("should detect the upstream service info", func() {
				Eventually(func() (*v1.ServiceInfo, error) {
					list, err := gloo.V1().Upstreams().List()
					if err != nil {
						return nil, err
					}
					var upstreamToTest *v1.Upstream
					for _, us := range list {
						if us.Name == test.upstreamName {
							upstreamToTest = us
							break
						}
					}

					if upstreamToTest == nil {
						return nil, errors.New("could not find upstream " + test.upstreamName)
					}
					log.Printf("got upstream: %v", upstreamToTest)
					return upstreamToTest.ServiceInfo, nil
				}, time.Minute*2).Should(Equal(test.expectedServiceInfo))
			})
		})
	}
})
