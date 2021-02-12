package tencentcloud

import "fmt"

// Action define a api request
type Action interface {
	// Service name, e.g. CVM/CDB/CLB
	Service() string
	// Action name, e.g. DescribeInstances
	Action() string
	// Version of api, e.g. 2017-03-12
	Version() string
	// FullDomainName return full name domain with specific postfix
	FullDomainName(postfix string) string
}

type action struct {
	service string
	action  string
	version string
}

func (act action) Service() string {
	return act.service
}

func (act action) Action() string {
	return act.action
}

func (act action) Version() string {
	return act.version
}

func (act action) FullDomainName(postfix string) string {
	return fmt.Sprintf("%s.%s", act.service, postfix)
}

var (
	CVMDescribeInstances = action{
		service: "cvm",
		action:  "DescribeInstances",
		version: "2017-03-12",
	}
)
