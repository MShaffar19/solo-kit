package setup

import (
	"context"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/setup"
)

func Main(settingsDir string) error {
	settingsClient, err := setup.KubeOrFileSettingsClient(settingsDir)
	if err != nil {
		return err
	}
	cache := v1.NewSetupEmitter(settingsClient)
	ctx := contextutils.WithLogger(context.Background(), "gateway")
	eventLoop := v1.NewSetupEventLoop(cache, NewSetupSyncer())
	errs, err := eventLoop.Run([]string{"settings"}, clients.WatchOpts{
		Ctx:         ctx,
		RefreshRate: time.Second,
	})
	if err != nil {
		return err
	}
	for err := range errs {
		contextutils.LoggerFrom(ctx).Errorf("error in setup: %v", err)
	}
	return nil
}
