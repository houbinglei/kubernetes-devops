package main

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-devops/kubectl-plugins/pod/lib"
	"os"
)

var client *kubernetes.Clientset

func run(c *cobra.Command, args []string) error {
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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名称", "命名空间", "IP", "状态"})

	for _, pod := range list.Items {
		table.Append([]string{pod.Name, pod.Namespace, pod.Status.PodIP,
			string(pod.Status.Phase)})
	}
	table.Render()
	return nil
}

func main() {
	client = lib.InitClient()
	lib.RunCmd(run)
}
