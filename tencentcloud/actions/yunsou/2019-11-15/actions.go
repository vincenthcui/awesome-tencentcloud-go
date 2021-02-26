package yunsou

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DataSearch = tencentcloud.Action{
		Service: "yunsou",
		Version: "2019-11-15",
		Action:  "DataSearch",
	}
	DataManipulation = tencentcloud.Action{
		Service: "yunsou",
		Version: "2019-11-15",
		Action:  "DataManipulation",
	}
)
