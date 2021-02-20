package deployment

import (
	"github.com/kubernetes/dashboard/src/app/backend/api"
	metricapi "github.com/kubernetes/dashboard/src/app/backend/integration/metric/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	"github.com/kubernetes/dashboard/src/app/backend/resource/event"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

// The code below allows to perform complex data section on Deployment

type DeploymentCell apps.Deployment

// GetProperty is used to get property of the deployment
func (self DeploymentCell) GetProperty(name dataselect.PropertyName) dataselect.ComparableValue {
	switch name {
	case dataselect.NameProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Name)
	case dataselect.CreationTimestampProperty:
		return dataselect.StdComparableTime(self.ObjectMeta.CreationTimestamp.Time)
	case dataselect.NamespaceProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Namespace)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}

func (self DeploymentCell) GetResourceSelector() *metricapi.ResourceSelector {
	return &metricapi.ResourceSelector{
		Namespace:    self.ObjectMeta.Namespace,
		ResourceType: api.ResourceKindDeployment,
		ResourceName: self.ObjectMeta.Name,
		Selector:     self.Spec.Selector.MatchLabels,
		UID:          self.UID,
	}
}

func toCells(std []apps.Deployment) []dataselect.DataCell {
	cells := make([]dataselect.DataCell, len(std))
	for i := range std {
		cells[i] = DeploymentCell(std[i])
	}
	return cells
}

func fromCells(cells []dataselect.DataCell) []apps.Deployment {
	std := make([]apps.Deployment, len(cells))
	for i := range std {
		std[i] = apps.Deployment(cells[i].(DeploymentCell))
	}
	return std
}

func getStatus(list *apps.DeploymentList, rs []apps.ReplicaSet, pods []v1.Pod, events []v1.Event) common.ResourceStatus {
	info := common.ResourceStatus{}
	if list == nil {
		return info
	}

	for _, deployment := range list.Items {
		matchingPods := common.FilterDeploymentPodsByOwnerReference(deployment, rs, pods)
		podInfo := common.GetPodInfo(deployment.Status.Replicas, deployment.Spec.Replicas, matchingPods)
		warnings := event.GetPodsEventWarnings(events, matchingPods)

		if len(warnings) > 0 {
			info.Failed++
		} else if podInfo.Pending > 0 {
			info.Pending++
		} else {
			info.Running++
		}
	}

	return info
}

func getConditions(deploymentConditions []apps.DeploymentCondition) []common.Condition {
	conditions := make([]common.Condition, 0)

	for _, condition := range deploymentConditions {
		conditions = append(conditions, common.Condition{
			Type:               string(condition.Type),
			Status:             condition.Status,
			Reason:             condition.Reason,
			Message:            condition.Message,
			LastTransitionTime: condition.LastTransitionTime,
			LastProbeTime:      condition.LastUpdateTime,
		})
	}

	return conditions
}
