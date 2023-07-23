package router

import (
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	e := gin.Default()

	return e
}
