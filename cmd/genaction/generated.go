package main

import (
	"fmt"
	"strings"
)

const (
	fileTpl = `
package %s

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	%s
)`
)

type instance struct {
	Service string
	Version string
	Action  string
}

func (ins instance) Lines() []string {
	return []string{
		fmt.Sprintf("%s = tencentcloud.Action{", ins.Action),
		fmt.Sprintf("\tService: \"%s\",", ins.Service),
		fmt.Sprintf("\tVersion: \"%s\",", ins.Version),
		fmt.Sprintf("\tAction: \"%s\",", ins.Action),
		"}",
	}
}

type instances []instance

func (inss instances) Len() int {
	return len(inss)
}

func (inss instances) Swap(i, j int) {
	inss[i], inss[j] = inss[j], inss[i]
}

func (inss instances) Less(i, j int) bool {
	return strings.Compare(inss[i].Action, inss[j].Action) >= 0
}

type file struct {
	service   string
	instances instances
}

func (f file) Lines() []string {
	lines := []string{
		fmt.Sprintf("package %s", f.service),
		"",
		`import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"`,
		"",
		"var (",
	}
	for _, ins := range f.instances {
		for _, line := range ins.Lines() {
			lines = append(lines, "\t"+line)
		}
	}
	lines = append(lines, []string{
		")",
	}...)

	return lines
}

func (f file) String() string {
	return strings.Join(f.Lines(), "\n")
}
