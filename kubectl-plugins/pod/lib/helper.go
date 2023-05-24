package lib

import (
	"fmt"
)

func Map2String(m map[string]string) (str string) {
	for k, v := range m {
		str += fmt.Sprintf("%s=%s\n", k, v)
	}
	return
}

// 初始化头
func InitHeader() []string {
	commonHeaders := []string{"名称", "命名空间", "IP", "状态"}
	if ShowLabels {
		commonHeaders = append(commonHeaders, "标签")
	}
	return commonHeaders
}
