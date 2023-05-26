package main

import (
	"k8s.io/client-go/kubernetes"
	"kubernetes-devops/kubectl-plugins/pod/lib"
)

var client *kubernetes.Clientset

func main() {
	lib.RunCmd()
}
