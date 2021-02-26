package yunsou

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DataSearch = tencentcloud.Action{
		Service: "yunsou",
		Version: "2018-05-04",
		Action:  "DataSearch",
	}
	DataManipulation = tencentcloud.Action{
		Service: "yunsou",
		Version: "2018-05-04",
		Action:  "DataManipulation",
	}
)
