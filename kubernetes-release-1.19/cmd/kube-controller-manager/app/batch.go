package app

import (
	"fmt"

	"net/http"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/controller/cronjob"
	"k8s.io/kubernetes/pkg/controller/job"
)

func startJobController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "jobs"}] {
		return nil, false, nil
	}
	go job.NewController(
		ctx.InformerFactory.Core().V1().Pods(),
		ctx.InformerFactory.Batch().V1().Jobs(),
		ctx.ClientBuilder.ClientOrDie("job-controller"),
	).Run(int(ctx.ComponentConfig.JobController.ConcurrentJobSyncs), ctx.Stop)
	return nil, true, nil
}

func startCronJobController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: "batch", Version: "v1beta1", Resource: "cronjobs"}] {
		return nil, false, nil
	}
	cjc, err := cronjob.NewController(
		ctx.ClientBuilder.ClientOrDie("cronjob-controller"),
	)
	if err != nil {
		return nil, true, fmt.Errorf("error creating CronJob controller: %v", err)
	}
	go cjc.Run(ctx.Stop)
	return nil, true, nil
}
