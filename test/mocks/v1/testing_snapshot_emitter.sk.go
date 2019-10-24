// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mTestingResourcesIn
	mTestingSnapshotIn = stats.Int64("testing.solo.io/emitter/snap_in", "Deprecated. Use testing.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mTestingResourcesIn    = stats.Int64("testing.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mTestingSnapshotOut    = stats.Int64("testing.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mTestingSnapshotMissed = stats.Int64("testing.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see testingResourcesInView
	testingsnapshotInView = &view.View{
		Name:        "testing.solo.io/emitter/snap_in",
		Measure:     mTestingSnapshotIn,
		Description: "Deprecated. Use testing.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	testingResourcesInView = &view.View{
		Name:        "testing.solo.io/emitter/resources_in",
		Measure:     mTestingResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	testingsnapshotOutView = &view.View{
		Name:        "testing.solo.io/emitter/snap_out",
		Measure:     mTestingSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	testingsnapshotMissedView = &view.View{
		Name:        "testing.solo.io/emitter/snap_missed",
		Measure:     mTestingSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		testingsnapshotInView,
		testingsnapshotOutView,
		testingsnapshotMissedView,
		testingResourcesInView,
	)
}

type TestingSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TestingSnapshot, <-chan error, error)
}

type TestingEmitter interface {
	TestingSnapshotEmitter
	Register() error
	MockResource() MockResourceClient
	FakeResource() FakeResourceClient
	AnotherMockResource() AnotherMockResourceClient
	ClusterResource() ClusterResourceClient
	MockCustomType() MockCustomTypeClient
	Pod() github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient
}

func NewTestingEmitter(mockResourceClient MockResourceClient, fakeResourceClient FakeResourceClient, anotherMockResourceClient AnotherMockResourceClient, clusterResourceClient ClusterResourceClient, mockCustomTypeClient MockCustomTypeClient, podClient github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient) TestingEmitter {
	return NewTestingEmitterWithEmit(mockResourceClient, fakeResourceClient, anotherMockResourceClient, clusterResourceClient, mockCustomTypeClient, podClient, make(chan struct{}))
}

func NewTestingEmitterWithEmit(mockResourceClient MockResourceClient, fakeResourceClient FakeResourceClient, anotherMockResourceClient AnotherMockResourceClient, clusterResourceClient ClusterResourceClient, mockCustomTypeClient MockCustomTypeClient, podClient github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient, emit <-chan struct{}) TestingEmitter {
	return &testingEmitter{
		mockResource:        mockResourceClient,
		fakeResource:        fakeResourceClient,
		anotherMockResource: anotherMockResourceClient,
		clusterResource:     clusterResourceClient,
		mockCustomType:      mockCustomTypeClient,
		pod:                 podClient,
		forceEmit:           emit,
	}
}

type testingEmitter struct {
	forceEmit           <-chan struct{}
	mockResource        MockResourceClient
	fakeResource        FakeResourceClient
	anotherMockResource AnotherMockResourceClient
	clusterResource     ClusterResourceClient
	mockCustomType      MockCustomTypeClient
	pod                 github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient
}

func (c *testingEmitter) Register() error {
	if err := c.mockResource.Register(); err != nil {
		return err
	}
	if err := c.fakeResource.Register(); err != nil {
		return err
	}
	if err := c.anotherMockResource.Register(); err != nil {
		return err
	}
	if err := c.clusterResource.Register(); err != nil {
		return err
	}
	if err := c.mockCustomType.Register(); err != nil {
		return err
	}
	if err := c.pod.Register(); err != nil {
		return err
	}
	return nil
}

func (c *testingEmitter) MockResource() MockResourceClient {
	return c.mockResource
}

func (c *testingEmitter) FakeResource() FakeResourceClient {
	return c.fakeResource
}

func (c *testingEmitter) AnotherMockResource() AnotherMockResourceClient {
	return c.anotherMockResource
}

func (c *testingEmitter) ClusterResource() ClusterResourceClient {
	return c.clusterResource
}

func (c *testingEmitter) MockCustomType() MockCustomTypeClient {
	return c.mockCustomType
}

func (c *testingEmitter) Pod() github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient {
	return c.pod
}

func (c *testingEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TestingSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for MockResource */
	type mockResourceListWithNamespace struct {
		list      MockResourceList
		namespace string
	}
	mockResourceChan := make(chan mockResourceListWithNamespace)

	var initialMockResourceList MockResourceList
	/* Create channel for FakeResource */
	type fakeResourceListWithNamespace struct {
		list      FakeResourceList
		namespace string
	}
	fakeResourceChan := make(chan fakeResourceListWithNamespace)

	var initialFakeResourceList FakeResourceList
	/* Create channel for AnotherMockResource */
	type anotherMockResourceListWithNamespace struct {
		list      AnotherMockResourceList
		namespace string
	}
	anotherMockResourceChan := make(chan anotherMockResourceListWithNamespace)

	var initialAnotherMockResourceList AnotherMockResourceList
	/* Create channel for ClusterResource */
	/* Create channel for MockCustomType */
	type mockCustomTypeListWithNamespace struct {
		list      MockCustomTypeList
		namespace string
	}
	mockCustomTypeChan := make(chan mockCustomTypeListWithNamespace)

	var initialMockCustomTypeList MockCustomTypeList
	/* Create channel for Pod */
	type podListWithNamespace struct {
		list      github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList
		namespace string
	}
	podChan := make(chan podListWithNamespace)

	var initialPodList github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList

	currentSnapshot := TestingSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for MockResource */
		{
			mocks, err := c.mockResource.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial MockResource list")
			}
			initialMockResourceList = append(initialMockResourceList, mocks...)
		}
		mockResourceNamespacesChan, mockResourceErrs, err := c.mockResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MockResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, mockResourceErrs, namespace+"-mocks")
		}(namespace)
		/* Setup namespaced watch for FakeResource */
		{
			fakes, err := c.fakeResource.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial FakeResource list")
			}
			initialFakeResourceList = append(initialFakeResourceList, fakes...)
		}
		fakeResourceNamespacesChan, fakeResourceErrs, err := c.fakeResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting FakeResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, fakeResourceErrs, namespace+"-fakes")
		}(namespace)
		/* Setup namespaced watch for AnotherMockResource */
		{
			anothermockresources, err := c.anotherMockResource.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial AnotherMockResource list")
			}
			initialAnotherMockResourceList = append(initialAnotherMockResourceList, anothermockresources...)
		}
		anotherMockResourceNamespacesChan, anotherMockResourceErrs, err := c.anotherMockResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting AnotherMockResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, anotherMockResourceErrs, namespace+"-anothermockresources")
		}(namespace)
		/* Setup namespaced watch for MockCustomType */
		{
			mcts, err := c.mockCustomType.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial MockCustomType list")
			}
			initialMockCustomTypeList = append(initialMockCustomTypeList, mcts...)
		}
		mockCustomTypeNamespacesChan, mockCustomTypeErrs, err := c.mockCustomType.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MockCustomType watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, mockCustomTypeErrs, namespace+"-mcts")
		}(namespace)
		/* Setup namespaced watch for Pod */
		{
			pods, err := c.pod.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Pod list")
			}
			initialPodList = append(initialPodList, pods...)
		}
		podNamespacesChan, podErrs, err := c.pod.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Pod watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, podErrs, namespace+"-pods")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case mockResourceList := <-mockResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case mockResourceChan <- mockResourceListWithNamespace{list: mockResourceList, namespace: namespace}:
					}
				case fakeResourceList := <-fakeResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case fakeResourceChan <- fakeResourceListWithNamespace{list: fakeResourceList, namespace: namespace}:
					}
				case anotherMockResourceList := <-anotherMockResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case anotherMockResourceChan <- anotherMockResourceListWithNamespace{list: anotherMockResourceList, namespace: namespace}:
					}
				case mockCustomTypeList := <-mockCustomTypeNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case mockCustomTypeChan <- mockCustomTypeListWithNamespace{list: mockCustomTypeList, namespace: namespace}:
					}
				case podList := <-podNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case podChan <- podListWithNamespace{list: podList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Mocks */
	currentSnapshot.Mocks = initialMockResourceList.Sort()
	/* Initialize snapshot for Fakes */
	currentSnapshot.Fakes = initialFakeResourceList.Sort()
	/* Initialize snapshot for Anothermockresources */
	currentSnapshot.Anothermockresources = initialAnotherMockResourceList.Sort()
	/* Setup cluster-wide watch for ClusterResource */
	var err error
	currentSnapshot.Clusterresources, err = c.clusterResource.List(clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "initial ClusterResource list")
	}
	clusterResourceChan, clusterResourceErrs, err := c.clusterResource.Watch(opts)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "starting ClusterResource watch")
	}
	done.Add(1)
	go func() {
		defer done.Done()
		errutils.AggregateErrs(ctx, errs, clusterResourceErrs, "clusterresources")
	}()
	/* Initialize snapshot for Mcts */
	currentSnapshot.Mcts = initialMockCustomTypeList.Sort()
	/* Initialize snapshot for Pods */
	currentSnapshot.Pods = initialPodList.Sort()

	snapshots := make(chan *TestingSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot
		stats.Record(ctx, mTestingSnapshotOut.M(1))

		timer := time.NewTicker(time.Second * 1)
		previousHash := currentSnapshot.Hash()
		sync := func() {
			currentHash := currentSnapshot.Hash()
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mTestingSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mTestingSnapshotMissed.M(1))
			}
		}
		mocksByNamespace := make(map[string]MockResourceList)
		fakesByNamespace := make(map[string]FakeResourceList)
		anothermockresourcesByNamespace := make(map[string]AnotherMockResourceList)
		mctsByNamespace := make(map[string]MockCustomTypeList)
		podsByNamespace := make(map[string]github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList)

		for {
			record := func() { stats.Record(ctx, mTestingSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case mockResourceNamespacedList := <-mockResourceChan:
				record()

				namespace := mockResourceNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"mock_resource",
					mTestingResourcesIn,
				)

				// merge lists by namespace
				mocksByNamespace[namespace] = mockResourceNamespacedList.list
				var mockResourceList MockResourceList
				for _, mocks := range mocksByNamespace {
					mockResourceList = append(mockResourceList, mocks...)
				}
				currentSnapshot.Mocks = mockResourceList.Sort()
			case fakeResourceNamespacedList := <-fakeResourceChan:
				record()

				namespace := fakeResourceNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"fake_resource",
					mTestingResourcesIn,
				)

				// merge lists by namespace
				fakesByNamespace[namespace] = fakeResourceNamespacedList.list
				var fakeResourceList FakeResourceList
				for _, fakes := range fakesByNamespace {
					fakeResourceList = append(fakeResourceList, fakes...)
				}
				currentSnapshot.Fakes = fakeResourceList.Sort()
			case anotherMockResourceNamespacedList := <-anotherMockResourceChan:
				record()

				namespace := anotherMockResourceNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"another_mock_resource",
					mTestingResourcesIn,
				)

				// merge lists by namespace
				anothermockresourcesByNamespace[namespace] = anotherMockResourceNamespacedList.list
				var anotherMockResourceList AnotherMockResourceList
				for _, anothermockresources := range anothermockresourcesByNamespace {
					anotherMockResourceList = append(anotherMockResourceList, anothermockresources...)
				}
				currentSnapshot.Anothermockresources = anotherMockResourceList.Sort()
			case clusterResourceList := <-clusterResourceChan:
				record()

				skstats.IncrementResourceCount(
					ctx,
					"<all>",
					"cluster_resource",
					mTestingResourcesIn,
				)

				currentSnapshot.Clusterresources = clusterResourceList
			case mockCustomTypeNamespacedList := <-mockCustomTypeChan:
				record()

				namespace := mockCustomTypeNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"mock_custom_type",
					mTestingResourcesIn,
				)

				// merge lists by namespace
				mctsByNamespace[namespace] = mockCustomTypeNamespacedList.list
				var mockCustomTypeList MockCustomTypeList
				for _, mcts := range mctsByNamespace {
					mockCustomTypeList = append(mockCustomTypeList, mcts...)
				}
				currentSnapshot.Mcts = mockCustomTypeList.Sort()
			case podNamespacedList := <-podChan:
				record()

				namespace := podNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"pod",
					mTestingResourcesIn,
				)

				// merge lists by namespace
				podsByNamespace[namespace] = podNamespacedList.list
				var podList github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList
				for _, pods := range podsByNamespace {
					podList = append(podList, pods...)
				}
				currentSnapshot.Pods = podList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
