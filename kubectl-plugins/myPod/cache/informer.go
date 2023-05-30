package cache

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubernetes-devops/kubectl-plugins/myPod/typed"
)

var Fact informers.SharedInformerFactory
var Client *kubernetes.Clientset

func InitClient(filename string) *kubernetes.Clientset {
	switch filename {
	case "prod":
		filename = typed.Prod
	case "tools":
		filename = typed.Tools
	case "hci-uat":
		filename = typed.HciUat
	case "uat":
		filename = typed.Uat
	case "hci-prod":
		filename = typed.HciProd
	}
	kubeConfig, _ := clientcmd.LoadFromFile(filename)
	restConfig, _ := clientcmd.NewDefaultClientConfig(*kubeConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	Client, _ = kubernetes.NewForConfig(restConfig)
	return Client
}

func InitCache() {
	Fact = informers.NewSharedInformerFactory(Client, 0)
	Fact.Core().V1().Pods().Informer().AddEventHandler(&typed.PodHandler{})
	Fact.Core().V1().Events().Informer().AddEventHandler(&typed.PodHandler{})
	Fact.Core().V1().Namespaces().Informer().AddEventHandler(&typed.PodHandler{})
	ch := make(chan struct{})
	Fact.Start(ch)
	Fact.WaitForCacheSync(ch)
}
