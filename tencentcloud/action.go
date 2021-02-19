package tencentcloud

// Action define a api request
type Action interface {
	// Service name, e.g. CVM/CDB/CLB
	Service() string
	// Version of api, e.g. 2017-03-12
	Version() string
	// Action name, e.g. DescribeInstances
	Action() string
}

type action struct {
	service string
	version string
	action  string
}

func (act action) Service() string {
	return act.service
}

func (act action) Version() string {
	return act.version
}

func (act action) Action() string {
	return act.action
}
