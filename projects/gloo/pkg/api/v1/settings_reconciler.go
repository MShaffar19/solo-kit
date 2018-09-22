package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionSettingsFunc func(original, desired *Settings) (bool, error)

type SettingsReconciler interface {
	Reconcile(namespace string, desiredResources SettingsList, transition TransitionSettingsFunc, opts clients.ListOpts) error
}

func settingssToResources(list SettingsList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, settings := range list {
		resourceList = append(resourceList, settings)
	}
	return resourceList
}

func NewSettingsReconciler(client SettingsClient) SettingsReconciler {
	return &settingsReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type settingsReconciler struct {
	base reconcile.Reconciler
}

func (r *settingsReconciler) Reconcile(namespace string, desiredResources SettingsList, transition TransitionSettingsFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "settings_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Settings), desired.(*Settings))
		}
	}
	return r.base.Reconcile(namespace, settingssToResources(desiredResources), transitionResources, opts)
}