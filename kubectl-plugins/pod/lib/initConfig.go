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

func MergeFlags(cmd *cobra.Command) {
	configFlags.AddFlags(cmd.Flags())
}

func RunCmd(f func(c *cobra.Command, args []string) error) {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods ",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE:         f,
	}
	MergeFlags(cmd)
	cmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "kubectl pods --show-labels")
	cmd.Flags().StringVar(&Labels, "labels", "", "kubectl pods --lables app=ngx or kubectl pods --lables=\"app=ngx,version=v1\"")

	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
