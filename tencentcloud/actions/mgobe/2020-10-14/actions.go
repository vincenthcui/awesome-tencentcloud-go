package mgobe

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	RemoveRoomPlayer = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "RemoveRoomPlayer",
	}
	ModifyRoom = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "ModifyRoom",
	}
	DismissRoom = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "DismissRoom",
	}
	DescribeRoom = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "DescribeRoom",
	}
	DescribePlayer = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "DescribePlayer",
	}
	ChangeRoomPlayerStatus = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "ChangeRoomPlayerStatus",
	}
	ChangeRoomPlayerProfile = tencentcloud.Action{
		Service: "mgobe",
		Version: "2020-10-14",
		Action:  "ChangeRoomPlayerProfile",
	}
)
