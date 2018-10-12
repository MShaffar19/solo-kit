package ratelimit

import (
	"context"

	"github.com/solo-io/solo-kit/projects/gloo/pkg/plugins"

	"github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	envoycache "github.com/solo-io/solo-kit/projects/gloo/pkg/control-plane/cache"
)

type Translator struct {
	rlplugin *Plugin
	xdsCache envoycache.SnapshotCache
}

func NewTranslator(plugins []plugins.Plugin, xdsCache envoycache.SnapshotCache) *Translator {
	// find the instance of our plugin
	for _, plug := range plugins {
		if rlplug, ok := plug.(*Plugin); ok {
			return &Translator{
				rlplugin: rlplug,
				xdsCache: xdsCache,
			}
		}
	}

	return &Translator{}
}

func (t *Translator) Sync(ctx context.Context, snap *v1.ApiSnapshot) error {
	if t.rlplugin == nil {
		return nil
	}

	if t.rlplugin.rlconfig == nil {
		return nil
	}

	resource := v1.NewRateLimitConfigXdsResourceWrapper(t.rlplugin.rlconfig)
	resources := []envoycache.Resource{resource}
	// TODO(yuval-k): un-hardcode
	rlsnap := envoycache.NewEasyGenericSnapshot("1", resources)
	t.xdsCache.SetSnapshot("ratelimit", rlsnap)
	// find our plugin
	return nil
}