package awesome_tencentcloud_go

import "net/http"

type client struct {
	client *http.Client

	region   string
	language string
}
