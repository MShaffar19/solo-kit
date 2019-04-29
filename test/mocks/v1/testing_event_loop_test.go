// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"sync"
	"time"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("TestingEventLoop", func() {
	var (
		namespace                 string
		err                       error
		emitter                   TestingEmitter
		mockResourceClient        MockResourceClient
		fakeResourceClient        FakeResourceClient
		anotherMockResourceClient AnotherMockResourceClient
		clusterResourceClient     ClusterResourceClient
		mockCustomTypeClient      MockCustomTypeClient
		podClient                 github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient
	)

	BeforeEach(func() {

		mockResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		mockResourceClient, err = NewMockResourceClient(mockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		fakeResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		fakeResourceClient, err = NewFakeResourceClient(fakeResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		anotherMockResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		anotherMockResourceClient, err = NewAnotherMockResourceClient(anotherMockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		clusterResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		clusterResourceClient, err = NewClusterResourceClient(clusterResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())

		mockCustomTypeClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		mockCustomTypeClient, err = NewMockCustomTypeClient(mockCustomTypeClientFactory)
		Expect(err).NotTo(HaveOccurred())

		podClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		podClient, err = github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPodClient(podClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewTestingEmitter(mockResourceClient, fakeResourceClient, anotherMockResourceClient, clusterResourceClient, mockCustomTypeClient, podClient)
	})
	It("runs sync function on a new snapshot", func() {
		mockResourceClient.Write(NewMockResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		fakeResourceClient.Write(NewFakeResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		anotherMockResourceClient.Write(NewAnotherMockResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		clusterResourceClient.Write(NewClusterResource(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		mockCustomTypeClient.Write(NewMockCustomType(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockTestingSyncer{}
		el := NewTestingEventLoop(emitter, sync)
		_, err := el.Run(nil, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockTestingSyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockTestingSyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockTestingSyncer) Sync(ctx context.Context, snap *TestingSnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}
