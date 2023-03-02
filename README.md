# go-gpt

一个OpenAI Gpt客户端,目前支持3.5模型,之后接入其他模型,让你不用导入更多包,争取最多跑一次.

**目前支持及等待支持能力：**
* [x] 支持Gpt3.5模型
* (未支持)支持Gpt3模型
* (未支持)再命令行聊天
* (未支持)微信快捷部署
* (未支持)支持语耳模型


##  <span id="915">1.导入:</span>
`go get github.com/AlmazDelDiablo/gpt3-5-turbo-g`
##  <span id="915">2.食用方法：</span>
```go
package main

import (
	"github.com/China-Chris/gpt/client"
	"github.com/China-Chris/gpt/gpt35"
	"github.com/China-Chris/gpt/model/gpt3_5"
)

func main() {
	apiKey := "YouApiKey"
	resp, err := gpt35.GetGPTResponse(apiKey, gpt3_5.TurboEngin, client.RoleUser, "content", gpt35.OptionalParams{})
	if err != nil {
		panic(err)
	}
	println(resp)
}
```
##  <span id="915">3.鸣谢：</span>
- [go-gpt3]https://github.com/PullRequestInc/go-gpt3
- [gpt3-5-turbo-go]https://github.com/AlmazDelDiablo/gpt3-5-turbo-go