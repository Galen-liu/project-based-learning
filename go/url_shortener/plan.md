
## 依赖
- redigo
- gin

## 对外接口
- 获取短链接
```md
- url
POST /api/v1/short-urls

- body
{
	url: "http://www.example.com/xxx/yyy"
}  

- response
{
	"originUrl": "http://www.example.com/xxx/yyy",
	"shortenedUrl": "http://url_shortener_host/eujekxueke"
}
```

- 解析短链接(redirect)
```md
- url
GET http://url_shortener_host/eujekxueke

- reponse
Redirect to http://www.example.com/xxx/yyy
```

## 待确定实现疑问
### 怎么生成 shortenedUrl
> uuid.NewString()
 
通过 uuid 包生成唯一 id

### 如何使用redigo缓存
```go
package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
}
```

