package user

import (
	"dzhxka.com/project-user/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init router success")
	router.Register(&RouterUser{}) //把url注册进去路由里
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
