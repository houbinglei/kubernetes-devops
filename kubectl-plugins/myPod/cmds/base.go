package cmds

import (
	"github.com/spf13/cobra"
	"kubernetes-devops/kubectl-plugins/myPod/cache"
	"kubernetes-devops/kubectl-plugins/myPod/utils"
	"log"
)

func addCmdFlags() {
	//用来支持 是否 显示标签
	promptCmd.Flags().StringVar(&utils.Namespace, "namespace", "", "kubectl pods --namespace=default")
}

func RunCmd() {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "start for root cmd ",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
	}

	addCmdFlags()
	cmd.AddCommand(promptCmd, cacheCmd)
	cache.InitClient("prod")
	cache.InitCache()
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
