package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/middleware"
	"go-blog/route"
	"log"
)

func Run() {
	// 初始化配置
	err := config.Init()
	if err != nil {
		log.Fatal("config load error", err)
	}

	eng := gin.Default()
	rGroup := new(route.RouterGroup)
	var routers = eng.Group("rest").Use((new(middleware.Authorize)).Auth)
	rGroup.InitRouter(routers)

	ip := config.Get("server.ip")
	if ip == "" {
		ip = "127.0.0.1"
	}

	port := config.Get("server.port")
	if port == "" {
		port = "8000"
	}

	addr := fmt.Sprintf("%s:%s", ip, port)
	eng.Run(addr)
}
