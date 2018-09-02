package mocks

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

type BlestingSyncer interface {
	Sync(context.Context, *BlestingSnapshot) error
}

type BlestingEventLoop interface {
	Run(namespaces []string, opts clients.WatchOpts) (<-chan error, error)
}

type blestingEventLoop struct {
	emitter BlestingEmitter
	syncer  BlestingSyncer
}

func NewBlestingEventLoop(emitter BlestingEmitter, syncer BlestingSyncer) BlestingEventLoop {
	return &blestingEventLoop{
		emitter: emitter,
		syncer:  syncer,
	}
}

func (el *blestingEventLoop) Run(namespaces []string, opts clients.WatchOpts) (<-chan error, error) {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "mocks.event_loop")
	logger := contextutils.LoggerFrom(opts.Ctx)
	logger.Infof("event loop started")

	errs := make(chan error)

	watch, emitterErrs, err := el.emitter.Snapshots(namespaces, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "starting snapshot watch")
	}
	go errutils.AggregateErrs(opts.Ctx, errs, emitterErrs, "mocks.emitter errors")
	go func() {
		// create a new context for each loop, cancel it before each loop
		var cancel context.CancelFunc = func() {}
		defer cancel()
		for {
			select {
			case snapshot := <-watch:
				// cancel any open watches from previous loop
				cancel()
				ctx, canc := context.WithCancel(opts.Ctx)
				cancel = canc
				err := el.syncer.Sync(ctx, snapshot)
				if err != nil {
					errs <- err
				}
			case <-opts.Ctx.Done():
				return
			}
		}
	}()
	return errs, nil
}