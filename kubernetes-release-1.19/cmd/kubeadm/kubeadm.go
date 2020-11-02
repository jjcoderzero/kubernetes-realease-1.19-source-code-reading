package main

import (
	"k8s.io/kubernetes/cmd/kubeadm/app"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
)

func main() {
	kubeadmutil.CheckErr(app.Run())
}
