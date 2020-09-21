# go-zero 文件服务
go-zero本身支持文件服务,但是我们需要写相关的handler文件,本文目的在于


+ 不写任何一个和文件相关的handler
+ 如果有新的文件,直接把文件模板到某个特定目录就好,不要动任何go代码 


需求在这里,开撸吧

# 在代码开始前,你可能需要阅读
[golang微服务框架go-zero系列-1:在go-zero中使用XormV2](https://mp.weixin.qq.com/s/NQMDvxvE1kH6MrpW50SUJg)
[golang微服务框架go-zero系列-2:在go-zero中使用jwt-token鉴权实践](https://mp.weixin.qq.com/s/defBD048957Qr1DH3MJpbQ)
[golang微服务框架go-zero系列-3:扩展go-zero,使之支持html模板解析自动化](https://mp.weixin.qq.com/s/iY1kTa41dpVti7L2auzefA)


# 注意
微服务讲究资源分离,实际生产过程中尽量使用专业的文件服务器或者OSS等第三方存储平台

# file服务实现思路
在`gin`中有专门的`static file`服务封装,`go-zero`目前并没有提供。目前`go-zero`提供非常严格的路径匹配,如
访问
`/asset/l1.jpg` 将映射到 `/asset/:1`对应的handlerlv1
`/asset/l1/l2.jpg` 将映射到 `/asset/:1/:2`对应的handlerlv2
这有如下俩种情况

## 映射指定路径到单个文件

比如我们需要访问`favourite.ico`,系统指向`./www/favourite.ico`文件,代码如下
```golang
//处理函数,传入文件地址
func filehandler(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath)
	}
}
```
在router里面直接调用`AddRoute`方法添加单个路由
```golang
func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {

//这里直接添加单个
engine.AddRoute(
				rest.Route{
					Method:  http.MethodGet,
					Path:    "/favourite.ico",
					Handler: filehandler("./www/favourite.ico"),
				})
}
```

## 映射指定目录并提供服务
实际过程中我们需要对外暴露某一个目录,比如`/assets/`目录,该目录下存放一些资源文件如`css`,`js`,`img`等
```bash
tree /f
+---assets                                     
|   +---css                                    
|   +---fonts                                  
|   +---images                                 
|   +---js                                     
|   \---plugins                                
|       +---font-awesome                       
|       |   +---css                            
|       |   \---fonts                          
|       +---fontawesome                        
|       |   +---css                            
|       |   \---fonts                          
|       +---ionicons                           
|       |   +---css                            
|       |   \---fonts                          
|       +---jquery.contextmenu                 
|       |   \---images                         
|       +---jquery.pin                         
|       |   +---css                            
|       |   \---images                         
|       +---jqueryui-1.12.1                    
|       |   +---external                       
|       |   |   \---jquery                     
|       |   \---images                         
|       \---swiper-4.5.3                       
|           +---css                            
|           \---js                             
```
如果使用单个文件的方式来实现,肯定不合理,因为router会非常大,怎么解决这个问题?我们可以使用如下方法实现文件夹服务
```golang
//
func dirhandler(patern, filedir string) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)

	}
}
```
如上函数的核心是`http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))`函数,这个函数的核心功能是将映`patern`格式映射到某一个目录`filedir`
+ patern:请求路径格式`/assets/:1`,`/assets/:1/:2`这种
+ filedir:映射对应的文件夹`./assets/`这种

那么我们只需要构建多级文件访问格式和`dirhandler`的映射关系即可
```golang
func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {

			//这里注册
			dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
			patern := "/asset/"
			dirpath := "./assets/"
			for i := 1; i < len(dirlevel); i++ {
				path := prefix + strings.Join(dirlevel[:i], "/")
				//最后生成 /asset
				engine.AddRoute(
					rest.Route{
						Method:  http.MethodGet,
						Path:    path,
						Handler: dirhandler(patern,dirpath),
					})

				logx.Infof("register dir  %s  %s", path,dirpath)
			}
}
```

## 404

404可以在main函数中配置
```golang
rt := router.NewPatRouter()
	rt.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//这里内容可以定制
		w.Write([]byte("服务器开小差了,这里可定制"))
	}))
	server := rest.MustNewServer(c.RestConf, rest.WithRouter(rt))
```
此时请求`http://127.0.0.1:8888/hello`,系统响应
`服务器开小差了,这里可定制`



## 测试
启动系统后运行
```bash
E:\workspace@go\gozero\file>go run file.go
2020/09/05 20:18:24 {"@timestamp":"2020-09-05T20:18:24.682+08","level":"info","content":"{{{file-api { console logs info false 0 100} pro  { 0 }} 0.0.0.0 8081 false 10000 1048576 3000 900 {false 0s []}} [/asset/=./assets]}"} 
{"@timestamp":"2020-09-05T20:18:24.682+08","level":"info","content":"register dir  /asset/:1  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.683+08","level":"info","content":"register dir  /asset/:1/:2  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.683+08","level":"info","content":"register dir  /asset/:1/:2/:3  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.683+08","level":"info","content":"register dir  /asset/:1/:2/:3/:4  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.697+08","level":"info","content":"register dir  /asset/:1/:2/:3/:4/:5  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.697+08","level":"info","content":"register dir  /asset/:1/:2/:3/:4/:5/:6  ./assets"}
{"@timestamp":"2020-09-05T20:18:24.698+08","level":"info","content":"register dir  /asset/:1/:2/:3/:4/:5/:6/:7  ./assets"}

```

访问系统都能正常响应

`http://127.0.0.1:8888/asset/images/avatar.jpg`
`http://127.0.0.1:8888/asset/js/test.js`
`http://127.0.0.1:8888/asset/js/lv2/test.js`

>注意,请求的是`/asset/**`  不是`/assets/**`


## 思考一下
我们可以在NotFoundHandler中根据req.URL.path来实现文件服务,如何实现呢?


# 本文代码获取
关注公众号`betaidea` 输入`file`即可获得本文相关代码
关注公众号`betaidea` 输入`html`即可获得html解析相关代码
关注公众号`betaidea` 输入`jwt`即可获得gozero集成jwt-token相关代码
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