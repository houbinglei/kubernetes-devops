package lib

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"log"
)

var configFlags *genericclioptions.ConfigFlags

func InitClient() *kubernetes.Clientset {
	configFlags = genericclioptions.NewConfigFlags(true)

	config, err := configFlags.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		log.Fatalln(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	return client

}

var ShowLabels bool
var Labels string
var Fields string
var Name string

func MergeFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		configFlags.AddFlags(cmd.Flags())
	}
}

func addListCmdFlags() {
	//用来支持 是否 显示标签
	listCmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "kubectl pods --show-lables")
	listCmd.Flags().StringVar(&Labels, "labels", "", "kubectl pods --lables app=ngx or kubectl pods --lables=\"app=ngx,version=v1\"")
	listCmd.Flags().StringVar(&Fields, "fields", "", "kubectl pods --fields=\"status.phase=Running\"")
	listCmd.Flags().StringVar(&Name, "name", "", "kubectl pods --name=\"^my\"")
}

func RunCmd() {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "start for root cmd ",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
	}
	MergeFlags(cmd, listCmd, promptCmd, cacheCmd)
	addListCmdFlags()
	cmd.AddCommand(listCmd, promptCmd, cacheCmd)

	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
