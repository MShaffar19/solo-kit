// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v2alpha1

import (
	"context"
	"os"
	"time"

	testing_solo_io_v1 "github.com/solo-io/solo-kit/test/mocks/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/test/helpers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V2Alpha1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1         string
		namespace2         string
		name1, name2       = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg                *rest.Config
		kube               kubernetes.Interface
		emitter            TestingEmitter
		mockResourceClient MockResourceClient
		fakeResourceClient testing_solo_io_v1.FakeResourceClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		kube = helpers.MustKubeClient()
		err := kubeutils.CreateNamespacesInParallel(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		// MockResource Constructor
		mockResourceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         MockResourceCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		mockResourceClient, err = NewMockResourceClient(mockResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// FakeResource Constructor
		fakeResourceClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		fakeResourceClient, err = testing_solo_io_v1.NewFakeResourceClient(fakeResourceClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewTestingEmitter(mockResourceClient, fakeResourceClient)
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *TestingSnapshot

		/*
			MockResource
		*/

		assertSnapshotMocks := func(expectMocks MockResourceList, unexpectMocks MockResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMocks {
						if _, err := snap.Mocks.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMocks {
						if _, err := snap.Mocks.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := mockResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := mockResourceClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mockResource1a, err := mockResourceClient.Write(NewMockResource(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource1b, err := mockResourceClient.Write(NewMockResource(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, nil)
		mockResource2a, err := mockResourceClient.Write(NewMockResource(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource2b, err := mockResourceClient.Write(NewMockResource(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b}, nil)

		err = mockResourceClient.Delete(mockResource2a.GetMetadata().Namespace, mockResource2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource2b.GetMetadata().Namespace, mockResource2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, MockResourceList{mockResource2a, mockResource2b})

		err = mockResourceClient.Delete(mockResource1a.GetMetadata().Namespace, mockResource1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource1b.GetMetadata().Namespace, mockResource1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(nil, MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b})

		/*
			FakeResource
		*/

		assertSnapshotFakes := func(expectFakes testing_solo_io_v1.FakeResourceList, unexpectFakes testing_solo_io_v1.FakeResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectFakes {
						if _, err := snap.Fakes.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectFakes {
						if _, err := snap.Fakes.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := fakeResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := fakeResourceClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		fakeResource1a, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource1b, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b}, nil)
		fakeResource2a, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource2b, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b}, nil)

		err = fakeResourceClient.Delete(fakeResource2a.GetMetadata().Namespace, fakeResource2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource2b.GetMetadata().Namespace, fakeResource2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b}, testing_solo_io_v1.FakeResourceList{fakeResource2a, fakeResource2b})

		err = fakeResourceClient.Delete(fakeResource1a.GetMetadata().Namespace, fakeResource1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource1b.GetMetadata().Namespace, fakeResource1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(nil, testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b})
	})
	It("tracks snapshots on changes to any resource using AllNamespace", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{""}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *TestingSnapshot

		/*
			MockResource
		*/

		assertSnapshotMocks := func(expectMocks MockResourceList, unexpectMocks MockResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMocks {
						if _, err := snap.Mocks.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMocks {
						if _, err := snap.Mocks.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := mockResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := mockResourceClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mockResource1a, err := mockResourceClient.Write(NewMockResource(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource1b, err := mockResourceClient.Write(NewMockResource(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, nil)
		mockResource2a, err := mockResourceClient.Write(NewMockResource(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mockResource2b, err := mockResourceClient.Write(NewMockResource(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b}, nil)

		err = mockResourceClient.Delete(mockResource2a.GetMetadata().Namespace, mockResource2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource2b.GetMetadata().Namespace, mockResource2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(MockResourceList{mockResource1a, mockResource1b}, MockResourceList{mockResource2a, mockResource2b})

		err = mockResourceClient.Delete(mockResource1a.GetMetadata().Namespace, mockResource1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = mockResourceClient.Delete(mockResource1b.GetMetadata().Namespace, mockResource1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMocks(nil, MockResourceList{mockResource1a, mockResource1b, mockResource2a, mockResource2b})

		/*
			FakeResource
		*/

		assertSnapshotFakes := func(expectFakes testing_solo_io_v1.FakeResourceList, unexpectFakes testing_solo_io_v1.FakeResourceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectFakes {
						if _, err := snap.Fakes.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectFakes {
						if _, err := snap.Fakes.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := fakeResourceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := fakeResourceClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		fakeResource1a, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource1b, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b}, nil)
		fakeResource2a, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		fakeResource2b, err := fakeResourceClient.Write(testing_solo_io_v1.NewFakeResource(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b}, nil)

		err = fakeResourceClient.Delete(fakeResource2a.GetMetadata().Namespace, fakeResource2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource2b.GetMetadata().Namespace, fakeResource2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b}, testing_solo_io_v1.FakeResourceList{fakeResource2a, fakeResource2b})

		err = fakeResourceClient.Delete(fakeResource1a.GetMetadata().Namespace, fakeResource1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = fakeResourceClient.Delete(fakeResource1b.GetMetadata().Namespace, fakeResource1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotFakes(nil, testing_solo_io_v1.FakeResourceList{fakeResource1a, fakeResource1b, fakeResource2a, fakeResource2b})
	})
})
