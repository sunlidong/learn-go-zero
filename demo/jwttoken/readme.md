# 在go-zero中使用jwt-token鉴权实践

# 阅读本文前你需要阅读
。。。

# 创建项目
## 生成go.mod文件
以如下指令创建项目
```bash
mkdir jwttoken
cd jwttoken
go mod init  jwttoken
```
## 定义user.api
本文设计API如下
|描述|格式|方法|参数|返回|是否需要鉴权|
|----|----|----|----|----|----|
|用户登录|/open/authorization|post|mobile:手机号,passwd:密码,code:图片验证码|id:用户ID,token:用户token|否|
|更新用户信息|/user/update|post|mobile:用户手机号|token:用户新的token|是|

根据以上描述,书写api的模板文件如下

```golang

type (
	UserOptReq struct {
		mobile string `form:"mobile"`
		passwd string `form:"passwd"`
		code   string `form:"code,optional"`
	}

	UserOptResp struct {
		id    uint   `json:"id"`
		token string `json:"token"`
	}
	//修改
	UserUpdateReq struct {
		id     uint   `form:"id"`
		mobile string `form:"mobile,optional"`
	}
)

service user-api {
	@server(
		handler: authorizationHandler
		folder: open
	)
	post /open/authorization(UserOptReq) returns(UserOptResp)

	@server(
		handler: edituserHandler
		folder: user
	)
	post /user/update(UserUpdateReq) returns(UserOptResp)
	
}

```
注意
+ 一个文件里面只能有一个service
+ 工具最后会以type里面模型为样板生成各种结构体,所以参数和结构体保持一致即可
+ 如果我们需要分文件夹管理业务, 可以用folder属性来定义
## 生成代码
采用如下指令生成代码
```bash
goctl api  go   -api   user.api   -dir  .
```

运行一下
```bash
go run open.go
```
测试一下
```bash
curl http://127.0.0.1:8888/open/authorization -X POST -d "mobile=15367151352&passwd=123rte&code=asasa"\"passwd\":\"testpwd\",\"code\":\"asdf\"}
{"id":0,"token":""}
```

# 中间件实现鉴权
在`handler`下新建auth.go文件,关键代码如下
```golang

//鉴权白名单,在这里面的是不需要鉴权的
var whiteList []string = []string{
	"/open/",
}

//鉴权中间件
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "auth")
		uri := r.RequestURI
		//默认不在
		isInWhiteList := false
		//判断请求是否包含白名单中的元素
		for _, v := range whiteList {
			if strings.Contains(uri, v) {
				isInWhiteList = true
			}
		}
		//如果爱白名单里面直接通过
		if isInWhiteList {
			next(w, r)
			return
		}
		//否则获取前端header 里面的X-Token字段,这个就是token	
		token := r.Header.Get("X-Token")
		//工具类见util\jwttoken.go
		_, err := utils.DecodeJwtToken(token)
		//如果有错直接返回error
		if err != nil {
			httpx.Error(w, err)
			return
		}
		//没报错就继续
		next(w, r)
	}
}


```
关于配置文件,系统内置了一部分关键字 如Cache,资料不多;基本上可以随便配置,然后在Conf中定义同名变量即可。

## 生成jwttoken
在`logic\open\authorizationlogic.go`中实现jwttoken的获取
```golang
func (l *AuthorizationLogic) Authorization(req types.UserOptReq) (*types.UserOptResp, error) {
	//这个是生成jwttoken的工具类
	token, err := utils.EncodeJwtToken(map[string]interface{}{
		"role": "kefu",
		"id":   "10086",
	})
	return &types.UserOptResp{
		Token: token,
	}, err
}

```


## 测试
### 不携带token时访问
```bash
>curl http://127.0.0.1:8888/user/update -X POST -d "mobile=15367151352&id=123"
鉴权失败,缺少鉴权参数
```
### 获取token
```bash
>curl http://127.0.0.1:8081/open/authorization -X POST -d "mobile=15367151352&passwd=123rte&code=asasa"
{"id":1599063149,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkzMjIzNDksImlkIjoiMTUzNjcxNTEzNTIifQ.jcdg3c2rdigPO5ZTxcDilVGERAuMIdY9BUmMNX3ZA9c"}
```
### 携带token时访问
```bash
>curl http://127.0.0.1:8888/user/update -X POST -H "X-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkzMjIzNDksImlkIjoiMTUzNjcxNTEzNTIifQ.jcdg3c2rdigPO5ZTxcDilVGERAuMIdY9BUmMNX3ZA9c" -d "mobile=15367151352&id=123"
# 请求成功
{"id":123,"token":""}
```
### 携带错误的token时访问
```bash
>curl http://127.0.0.1:8888/user/update -X POST -H "X-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkzMjIzNDksImlkIjoiMTUzNjcxNTEzNTIifQ.jcdg3c2rdigPO5ZTxcDilVGERAuMIdY9BUmMNX3ZA9c0000" -d "mobile=15367151352&id=123"
# 返回签名无效
signature is invalid
```


# 本文代码获取
关注公众号`betaidea` 输入`jwt`即可获得
关注公众号`betaidea` 输入`gozero`即可gozero入门代码

# 广而告之
送福利了uniapp用户福音来啦！
历经数十万用户考验,我们的客服系统终于对外提供服务了。
你还在为商城接入客服烦恼吗?只需一行代码,即可接入啦!!
只需一行代码!!!!

```html
/*kefu.vue*/
<template>
	<view>
		<IdeaKefu :siteid="siteId"  ></IdeaKefu>
	</view>
</template>

<script>
	import IdeaKefu from "@/components/idea-kefu/idea-kefu.vue"
    export default {
		components:{
			IdeaKefu
		},
		data() {
			return {
				siteId:2
			}
		}
    }   
```
效果杠杠的
![客服效果](http://kefu.techidea8.com/html/wiki/assets/image/vistor-1.png)

开发文档地址
[http://kefu.techidea8.com/html/wiki/](http://kefu.techidea8.com/html/wiki/)