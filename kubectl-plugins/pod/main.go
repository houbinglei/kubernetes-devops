package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-devops/kubectl-plugins/pod/lib"
)

func main() {
	client := lib.InitClient()
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods ",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, pod := range list.Items {
				fmt.Println(pod.Name)
			}
			return nil
		},
	}
	lib.MergeFlags(cmd)
	cmd.Execute()
}
