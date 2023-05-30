package utils

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/json"
	"kubernetes-devops/kubectl-plugins/myPod/cache"
	"kubernetes-devops/kubectl-plugins/myPod/typed"
	"log"
	"os"
	"sigs.k8s.io/yaml"
)

func Map2String(m map[string]string) (str string) {
	for k, v := range m {
		str += fmt.Sprintf("%s=%s\n", k, v)
	}
	return
}

var ShowLabels bool
var Namespace string

//func MergeFlags(cmds ...*cobra.Command) {
//	for _, cmd := range cmds {
//		configFlags.AddFlags(cmd.Flags())
//	}
//}

// 初始化头
func InitHeader() []string {
	commonHeaders := []string{"名称", "命名空间", "IP", "状态"}
	if ShowLabels {
		commonHeaders = append(commonHeaders, "标签")
	}
	return commonHeaders
}

var eventHeaders = []string{"事件类型", "REASON", "所属对象", "消息"}

func printEvent(events []*v1.Event) {
	table := tablewriter.NewWriter(os.Stdout)
	//设置头
	table.SetHeader(eventHeaders)
	for _, e := range events {
		podRow := []string{e.Type, e.Reason,
			fmt.Sprintf("%s/%s", e.InvolvedObject.Kind, e.InvolvedObject.Name), e.Message}

		table.Append(podRow)
	}
	setTable(table)
	table.Render()
}
func setTable(table *tablewriter.Table) {
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
}

func getPodDetailByJSON(podName, path string, cmd *cobra.Command) {
	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		log.Println("error ns param")
		return
	}
	if ns == "" {
		ns = "default"
	}

	pod, err := cache.Fact.Core().V1().Pods().Lister().
		Pods(ns).Get(podName)
	if err != nil {
		log.Println(err)
		return
	}
	podEvents := []*v1.Event{}
	if path == typed.PodEventType {
		eventList, _ := cache.Fact.Core().V1().Events().Lister().Events(ns).List(labels.Everything())
		for _, e := range eventList {
			if pod.UID == e.InvolvedObject.UID {
				podEvents = append(podEvents, e)
			}
		}
		printEvent(podEvents)
		return
	}

	if path == typed.PodLogType {
		req := cache.Client.CoreV1().Pods(ns).GetLogs(pod.Name, &v1.PodLogOptions{})
		ret := req.Do(context.Background())
		b, err := ret.Raw()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(b))
		return
	}
	jsonStr, _ := json.Marshal(pod)
	//os.WriteFile("./pod.yaml", jsonStr, 0644)

	ret := gjson.Get(string(jsonStr), path)

	if !ret.Exists() {
		log.Println("无法找到对应的内容:" + path)
		return
	}
	if !ret.IsObject() && !ret.IsArray() { //不是对象不是 数组，直接打印
		fmt.Println(ret.Raw)
		return
	}
	var tempMap interface{}
	if ret.IsObject() {
		tempMap = make(map[string]interface{})
	}
	if ret.IsArray() {
		tempMap = []interface{}{}
	}

	err = yaml.Unmarshal([]byte(ret.Raw), &tempMap)
	if err != nil {
		log.Println(err)
		return
	}
	b, _ := yaml.Marshal(tempMap)
	fmt.Println(string(b))

}
