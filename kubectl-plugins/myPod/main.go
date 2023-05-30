package main

import (
	"kubernetes-devops/kubectl-plugins/myPod/cmds"
	_ "kubernetes-devops/kubectl-plugins/myPod/typed"
)

func main() {
	cmds.RunCmd()
}
