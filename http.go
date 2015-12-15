package rupert

import (
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.POST("/user/register")
	return engine
}


