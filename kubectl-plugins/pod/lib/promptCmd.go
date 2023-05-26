package lib

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
	"log"
	"os"
	"regexp"
	"strings"
)

func executorCmd(cmd *cobra.Command) func(in string) {
	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		args := []string{}
		if len(blocks) > 1 {
			args = blocks[1:]
		}
		switch blocks[0] {
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		case "list":
			err := cacheCmd.RunE(cmd, args)
			if err != nil {
				log.Fatalln(err)
			}
		case "get":
			// getPodDetail(args,cmd)
			runtea(args, cmd)
		case "use":
			// todo
			setNameSpace(cmd, args)
		case "ns":
			ns, _ := cmd.Flags().GetString("namespace")
			if ns == "" {
				ns = "default"
			}
			fmt.Println(ns)
		}
	}
}

func setNameSpace(cmd *cobra.Command, args []string) {
	err := cmd.Flags().Set("namespace", args[0])
	if err != nil {
		log.Println("设置namespace失败:", err.Error())
	}
	fmt.Println("设置namespace成功")
}

var suggestions = []prompt.Suggest{
	// Command
	{"list", "显示Pods列表"},
	{"exit", "退出交互式窗口"},
	{"get", "获取POD详细"},
	{"ns", "获取当前namespace"},
	{"use", "设置当前namespace"},
}

func getPodsList(cmd *cobra.Command) (ret []prompt.Suggest) {
	ns, _ := cmd.Flags().GetString("namespace")
	if ns == "" {
		ns = "default"
	}
	pods, err := fact.Core().V1().Pods().Lister().
		Pods(ns).List(labels.Everything())
	if err != nil {
		return
	}
	for _, pod := range pods {
		ret = append(ret, prompt.Suggest{
			Text: pod.Name,
			Description: "节点:" + pod.Spec.NodeName + " 状态:" +
				string(pod.Status.Phase) + " IP:" + pod.Status.PodIP,
		})
	}
	return
}

func getNamespaceList() (ret []prompt.Suggest) {
	ns, err := fact.Core().V1().Namespaces().Lister().List(labels.Everything())
	if err != nil {
		return
	}
	for _, n := range ns {
		ret = append(ret, prompt.Suggest{
			Text:        n.Name,
			Description: "ClusterName:" + n.ClusterName,
		})
	}
	return
}

func parseCmd(w string) (string, string) {
	w = regexp.MustCompile("\\s+").ReplaceAllString(w, " ")
	l := strings.Split(w, " ")
	if len(l) >= 2 {
		return l[0], strings.Join(l[1:], " ")
	}
	return w, ""
}

func completerWrapper(command *cobra.Command) func(in prompt.Document) []prompt.Suggest {
	return func(in prompt.Document) []prompt.Suggest {
		w := in.GetWordBeforeCursor()
		if w == "" {
			return []prompt.Suggest{}
		}
		cmd, opt := parseCmd(in.TextBeforeCursor())
		if cmd == "get" {
			return prompt.FilterHasPrefix(getPodsList(command), opt, true)
		}
		if cmd == "use" {
			return prompt.FilterHasPrefix(getNamespaceList(), opt, true)
		}
		return prompt.FilterHasPrefix(suggestions, w, true)
	}
}

var promptCmd = &cobra.Command{
	Use:          "prompt",
	Short:        "prompt pods for promptCmd ",
	Example:      "kubectl pods prompt",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		InitCache()
		p := prompt.New(
			executorCmd(c),
			completerWrapper(c),
			prompt.OptionPrefix(">>> "),
		)
		p.Run()
		return nil
	},
}
