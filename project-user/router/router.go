package router

import (
	"github.com/gin-gonic/gin"
)

// 定义了一个路由的接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}
func (RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	//rg := New()
	////以后的模块路由在这进行注册
	//rg.Route(&user.RouterUser{}, r)
	////以后的模块路由在这进行注册
	//通过一个循环来读取每次写的路由
	for _, i := range routers {
		i.Route(r)
	}
}

//… 其实是go的一种语法糖。
//
//第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
//第二个用法是slice可以被打散进行传递。\

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
