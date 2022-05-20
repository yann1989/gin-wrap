# gin-wrap

包装gin

# 状态码说明

- 前两位为服务码
- 中两位为模块
- 后两位为错误码
- 10xxxxx为通用码

# 快速开始

启动服务器

```go
import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    ginwrap "github.com/yann1989/gin-wrap"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type User struct {
    Name string `json:"name" form:"name" uri:"name" binding:"required" err:"用户名称不能为空"`
}

func (u User) DecodeType() ginwrap.DecodeType {
    return ginwrap.DecodeTypeQuery
}

func main() {
    engine := gin.New()
    wrap := ginwrap.NewEngineWrap(engine, ginwrap.EngineWrapLoggerOption, ginwrap.PrintReqParamsOption, ginwrap.PrintRespParamsOption)
    wrap.GET("", &User{}, func (ctx *gin.Context, base ginwrap.IBase) *ginwrap.Response {
        time.Sleep(time.Second)
        var user = base.(*User)
        return ginwrap.Response200(user.Name)
    })
    server := &http.Server{
        Addr:    ":80",
        Handler: engine,
    }
    go func () {
        quit := make(chan os.Signal)
        signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
        select {
        case s := <-quit:
        fmt.Println("收到停止信号:", s)
        }
        server.Shutdown(context.Background())
    }()
if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    panic("启动异常: " + err.Error())
}
    fmt.Println("正常退出")
}

```

## 成功示例
请求

```shell
curl --location --request GET '127.0.0.1?name=yann'
```
打印
```
请求参数: {"name":"yann"}
响应参数: {"code":100000,"message":"成功","data":"yann","timestamp":1003}
[请求1返回] 2022-05-20 14:02:54 | 200 | 1.003468505s |  100000  |  成功  | / 
```

响应

```json
{
  "code": 100000,
  "message": "成功",
  "data": "yann",
  "timestamp": 1005
}
```

# 错误示例
请求

```shell
curl --location --request GET '127.0.0.1'
```
打印
```
响应参数: {"code":100400,"message":"用户名称不能为空","timestamp":0}
[请求1返回] 2022-05-20 14:07:48 | 200 | 372.116µs |  100400  |  用户名称不能为空  | / 
```

响应

```json
{
  "code": 100400,
  "message": "用户名称不能为空",
  "timestamp": 0
}
```
