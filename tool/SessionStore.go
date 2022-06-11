package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(engine *gin.Engine) {
	config := GetConfig().Redis
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, "", []byte(config.Password))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	engine.Use(sessions.Sessions("mysession", store))
}

func SetSession(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

func GetSession(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
