package persistentvolume

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
)

// DropDisabledFields removes disabled fields from the pv spec.
// This should be called from PrepareForCreate/PrepareForUpdate for all resources containing a pv spec.
func DropDisabledFields(pvSpec *api.PersistentVolumeSpec, oldPVSpec *api.PersistentVolumeSpec) {
	if !utilfeature.DefaultFeatureGate.Enabled(features.ExpandCSIVolumes) && !hasExpansionSecrets(oldPVSpec) {
		if pvSpec.CSI != nil {
			pvSpec.CSI.ControllerExpandSecretRef = nil
		}
	}
}

func hasExpansionSecrets(oldPVSpec *api.PersistentVolumeSpec) bool {
	if oldPVSpec == nil || oldPVSpec.CSI == nil {
		return false
	}

	if oldPVSpec.CSI.ControllerExpandSecretRef != nil {
		return true
	}
	return false
}
