package persistentvolumeclaim

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
)

const (
	pvc            string = "PersistentVolumeClaim"
	volumeSnapshot string = "VolumeSnapshot"
)

// DropDisabledFields removes disabled fields from the pvc spec.
// This should be called from PrepareForCreate/PrepareForUpdate for all resources containing a pvc spec.
func DropDisabledFields(pvcSpec, oldPVCSpec *core.PersistentVolumeClaimSpec) {
	if !dataSourceIsEnabled(pvcSpec) && !dataSourceInUse(oldPVCSpec) {
		pvcSpec.DataSource = nil
	}
}

func dataSourceInUse(oldPVCSpec *core.PersistentVolumeClaimSpec) bool {
	if oldPVCSpec == nil {
		return false
	}
	if oldPVCSpec.DataSource != nil {
		return true
	}
	return false
}

func dataSourceIsEnabled(pvcSpec *core.PersistentVolumeClaimSpec) bool {
	if pvcSpec.DataSource != nil {
		if utilfeature.DefaultFeatureGate.Enabled(features.AnyVolumeDataSource) {
			return true
		}

		apiGroup := ""
		if pvcSpec.DataSource.APIGroup != nil {
			apiGroup = *pvcSpec.DataSource.APIGroup
		}
		if pvcSpec.DataSource.Kind == pvc &&
			apiGroup == "" {
			return true

		}

		if utilfeature.DefaultFeatureGate.Enabled(features.VolumeSnapshotDataSource) &&
			pvcSpec.DataSource.Kind == volumeSnapshot &&
			apiGroup == "snapshot.storage.k8s.io" {
			return true
		}
	}
	return false
}
