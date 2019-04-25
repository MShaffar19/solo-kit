// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1       string
		namespace2       string
		name1, name2     = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		emitter          KubeconfigsEmitter
		kubeConfigClient KubeConfigClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		var err error
		err = setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		Expect(err).NotTo(HaveOccurred())
		// KubeConfig Constructor
		kubeConfigClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		kubeConfigClient, err = NewKubeConfigClient(kubeConfigClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewKubeconfigsEmitter(kubeConfigClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
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

		var snap *KubeconfigsSnapshot

		/*
			KubeConfig
		*/

		assertSnapshotkubeconfigs := func(expectkubeconfigs KubeConfigList, unexpectkubeconfigs KubeConfigList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectkubeconfigs {
						if _, err := snap.Kubeconfigs.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectkubeconfigs {
						if _, err := snap.Kubeconfigs.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := kubeConfigClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := kubeConfigClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		kubeConfig1a, err := kubeConfigClient.Write(NewKubeConfig(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		kubeConfig1b, err := kubeConfigClient.Write(NewKubeConfig(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b}, nil)
		kubeConfig2a, err := kubeConfigClient.Write(NewKubeConfig(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		kubeConfig2b, err := kubeConfigClient.Write(NewKubeConfig(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b, kubeConfig2a, kubeConfig2b}, nil)

		err = kubeConfigClient.Delete(kubeConfig2a.GetMetadata().Namespace, kubeConfig2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = kubeConfigClient.Delete(kubeConfig2b.GetMetadata().Namespace, kubeConfig2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b}, KubeConfigList{kubeConfig2a, kubeConfig2b})

		err = kubeConfigClient.Delete(kubeConfig1a.GetMetadata().Namespace, kubeConfig1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = kubeConfigClient.Delete(kubeConfig1b.GetMetadata().Namespace, kubeConfig1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(nil, KubeConfigList{kubeConfig1a, kubeConfig1b, kubeConfig2a, kubeConfig2b})
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

		var snap *KubeconfigsSnapshot

		/*
			KubeConfig
		*/

		assertSnapshotkubeconfigs := func(expectkubeconfigs KubeConfigList, unexpectkubeconfigs KubeConfigList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectkubeconfigs {
						if _, err := snap.Kubeconfigs.List().Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectkubeconfigs {
						if _, err := snap.Kubeconfigs.List().Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := kubeConfigClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := kubeConfigClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		kubeConfig1a, err := kubeConfigClient.Write(NewKubeConfig(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		kubeConfig1b, err := kubeConfigClient.Write(NewKubeConfig(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b}, nil)
		kubeConfig2a, err := kubeConfigClient.Write(NewKubeConfig(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		kubeConfig2b, err := kubeConfigClient.Write(NewKubeConfig(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b, kubeConfig2a, kubeConfig2b}, nil)

		err = kubeConfigClient.Delete(kubeConfig2a.GetMetadata().Namespace, kubeConfig2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = kubeConfigClient.Delete(kubeConfig2b.GetMetadata().Namespace, kubeConfig2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(KubeConfigList{kubeConfig1a, kubeConfig1b}, KubeConfigList{kubeConfig2a, kubeConfig2b})

		err = kubeConfigClient.Delete(kubeConfig1a.GetMetadata().Namespace, kubeConfig1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = kubeConfigClient.Delete(kubeConfig1b.GetMetadata().Namespace, kubeConfig1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotkubeconfigs(nil, KubeConfigList{kubeConfig1a, kubeConfig1b, kubeConfig2a, kubeConfig2b})
	})
})
