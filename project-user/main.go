package main

import (
	_ "dzhxka.com/project-user/api/user"
	"dzhxka.com/project-user/router"
	"github.com/gin-gonic/gin"
	common "project-common"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	common.Run(r, "Goproject", ":8080") //优雅启停项目
}
