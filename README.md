# awesome-tencentcloud-go

简单方便的腾讯云 SDK，自动注入 ClientToken，自动重试网络错误，触发频率限制时自动延时。

Smart and awesome tencentcloud sdk on golang, which automatic retry request when network issues occurred, or cross over
rate-limit when request limit exceeded.

### 示例 Quick View

```bash
go get github.com/vincenthcui/awesome-tencentcloud-go@v0.2.1
```

```golang
package main

import (
	"context"
	"time"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tc "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"
	action "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions/cvm/2017-03-12"
)

func main() {
	client := tc.NewClient(tc.WithSecret("YOUR_SECRET_ID", "YOUR_SECRET_KEY"))

	req := cvm.NewDescribeInstancesRequest()
	resp := cvm.NewDescribeInstancesResponse()

	context.WithTimeout(context.Background(), 3*time.Second)
	if err := client.Send(context.TODO(), action.DescribeInstances, req, resp); err != nil {
		panic(err)
	}
}
```

### 特性 Features

- 自动重试 Automatic retry
- 注入客户端凭证 Automatic inject ClientToken
- 跨越频限 Cross through rate-limit
- 基于 context 的超时控制 Context-based timeout control

### 为什么要重新做一套 Why

[官方SDK](https://github.com/TencentCloud/tencentcloud-sdk-go)
为每个服务设置了不同的 Client 结构，在每个 Client 结构上设置不同方法，用于请求具体的 API。 这种结构让我们很难在 SDK 层实现重试和频率限制。

The [official tencentcloud sdk](https://github.com/TencentCloud/tencentcloud-sdk-go)
use different Client type for every service, and every Client has different method to perform actually api request. That
make it hard to do retry or rate-limit work in sdk level.

### 更多特性 More Features

需要其他特性，请直接提交 issue。

Feel free if you need any feature not supported now.

- [x] 自动生成请求 Automatic generate Actions
- [ ] 兼容 GET 请求和其他加密方式 Backward compatibility for GET request and other algorithm