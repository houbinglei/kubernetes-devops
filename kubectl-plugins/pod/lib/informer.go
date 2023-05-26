package lib

import "k8s.io/client-go/informers"

var fact informers.SharedInformerFactory

type PodHandler struct {
}

func (p *PodHandler) OnAdd(obj interface{}) {
}

func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
}

func (p *PodHandler) OnDelete(obj interface{}) {
}

func InitCache() {
	fact = informers.NewSharedInformerFactory(client, 0)
	fact.Core().V1().Pods().Informer().AddEventHandler(&PodHandler{})
	fact.Core().V1().Events().Informer().AddEventHandler(&PodHandler{})
	fact.Core().V1().Namespaces().Informer().AddEventHandler(&PodHandler{})
	ch := make(chan struct{})
	fact.Start(ch)
	fact.WaitForCacheSync(ch)
}
