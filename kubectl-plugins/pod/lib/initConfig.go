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

func MergeFlags(cmd *cobra.Command) {
	configFlags.AddFlags(cmd.Flags())
}
