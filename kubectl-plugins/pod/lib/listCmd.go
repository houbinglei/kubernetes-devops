package lib

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
	"regexp"
)

var client = InitClient()

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "list pods for listCmd ",
	Example:      "kubectl pods list [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		if ns == "" {
			ns = "default"
		}
		list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{
			//LabelSelector: "app=reviews",
			LabelSelector: Labels,
			FieldSelector: Fields,
		})
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(InitHeader())
		podRow := make([]string, 0)
		for _, pod := range list.Items {
			if Name != "" {
				b, _ := json.Marshal(pod)
				ret := gjson.Get(string(b), "metadata.name")
				ok, _ := regexp.MatchString(Name, ret.String())
				if ok {
					podRow = []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}
				} else {
					continue
				}

			}
			podRow = []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}
			if ShowLabels {
				podRow = append(podRow, Map2String(pod.Labels))
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
	},
}

//var listCmd = &cobra.Command{
//	Use:          "list",
//	Short:        "list pods ",
//	Example:      "kubectl pods list [flags]",
//	SilenceUsage: true,
//	RunE: func(c *cobra.Command, args []string) error {
//		ns, err := c.Flags().GetString("namespace")
//		if err != nil {
//			return err
//		}
//		if ns == "" {
//			ns = "default"
//		}
//
//		var list = &v12.PodList{}
//
//		list, err = client.CoreV1().Pods(ns).List(context.Background(),
//			v1.ListOptions{LabelSelector: Labels, FieldSelector: Fields})
//		if err != nil {
//			return err
//		}
//
//		//FilterListByJSON(list)
//		//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
//		table := tablewriter.NewWriter(os.Stdout)
//		//设置头
//
//		table.SetHeader(InitHeader())
//		for _, pod := range list.Items {
//			podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP,
//				string(pod.Status.Phase)}
//			if ShowLabels {
//				podRow = append(podRow, Map2String(pod.Labels))
//			}
//			table.Append(podRow)
//		}
//		setTable(table)
//		table.Render()
//		return nil
//	},
//}
