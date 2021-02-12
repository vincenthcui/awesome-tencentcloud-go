package tencentcloud

// Action define a api request
type Action interface {
	// Service name, e.g. CVM/CDB/CLB
	Service() string
	// Action name, e.g. DescribeInstances
	Action() string
	// Version of api, e.g. 2017-03-12
	Version() string
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

var (
	CVMDescribeInstances = action{
		service: "cvm",
		action:  "DescribeInstances",
		version: "2017-03-12",
	}
)
