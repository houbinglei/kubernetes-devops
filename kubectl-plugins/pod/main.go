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
	list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{
		//LabelSelector: "app=reviews",
		LabelSelector: lib.Labels,
	})
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(lib.InitHeader())

	for _, pod := range list.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}
		if lib.ShowLabels {
			//fmt.Println(pod.Labels)
			// map[app:myapp pod-template-hash:6495489bdb]
			podRow = append(podRow, lib.Map2String(pod.Labels))
		}
		table.Append(podRow)
	}
	//table.SetAutoWrapText(false)
	//table.SetAutoFormatHeaders(true)
	//table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	//table.SetCenterSeparator("")
	//table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	//table.SetHeaderLine(false)
	//table.SetBorder(false)
	//table.SetTablePadding("\t") // pad with tabs
	//table.SetNoWhiteSpace(true)
	table.Render()
	return nil
}

func main() {
	client = lib.InitClient()
	lib.RunCmd(run)
}
