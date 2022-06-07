package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"goweb_gin/controller"
	"goweb_gin/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		panic(err.Error())
	}

	tool.InitRedisStore(context.Background())

	app := gin.Default()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	app.Use(cors.New(corsConfig))

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
