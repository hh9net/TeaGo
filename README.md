# TeaGo - Go语言快速开发框架

## 定义不带参数的Action
*actions/hello/index.go*
~~~go
package hello

import "github.com/iwind/TeaGo/actions"

type IndexAction actions.Action

func (this *IndexAction) Run()  {
	this.Write("Hello")
}
~~~

## 定义带参数的Action
~~~go
package hello

import "github.com/iwind/TeaGo/actions"

type IndexAction actions.Action

func (this *IndexAction) Run(params struct {
	Name string
	Age  int
}) {
	this.WriteFormat("Name:%s, Age:%d",
		params.Name,
		params.Age)
}

~~~

## 使用Action
~~~go
package MyProject

import (
	"github.com/iwind/TeaGo"
	"github.com/iwind/MyProject/actions/hello/index"
)

func Start() {
	var server = TeaGo.NewServer()
	
	// 注册路由
	server.Get("/hello", new(hello.IndexAction))
	
	// 启动服务
	server.Start("0.0.0.0:8000")
}

~~~
