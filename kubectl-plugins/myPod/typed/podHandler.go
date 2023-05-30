package typed

import (
	"fmt"
	"os"
)

type PodHandler struct {
}

func (p *PodHandler) OnAdd(obj interface{}) {
}

func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
}

func (p *PodHandler) OnDelete(obj interface{}) {
}

var HelperInfo = `可用命令选项：
- exit/q/Q  --> 退出
- list  --> 显示pods 列表
- get --> 显示pods 详情
- use --> 设置当前namespace
- ns --> 显示当前namespace
- jump --> 跳转其他集群（支持： prod/uat/tools/hci-prod/hci-uat/）
`

var HomeDir string

func init() {
	HomeDir = os.Getenv("HOME")
	HciProd = fmt.Sprintf("%s/.kube/hci-prod", HomeDir)
	Prod = fmt.Sprintf("%s/.kube/prod", HomeDir)
	Tools = fmt.Sprintf("%s/.kube/tools", HomeDir)
	HciUat = fmt.Sprintf("%s/.kube/hci-uat", HomeDir)
	Uat = fmt.Sprintf("%s/.kube/uat", HomeDir)
}

var (
	Prod         string
	HciProd      string
	HciUat       string
	Tools        string
	Uat          string
	PodLogType   = "_log_path"
	PodEventType = "_event_path"
)
