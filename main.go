package main

import (
	"ini/dao/mysql"
	"ini/router"
	"log"
)

// @title oyhx API
// @version 1.0
// @description 偶遇华夏API
// @termsOfService http://swagger.io/terms/
// @contact.name KitZhangYs
// @contact.email SJMbaiyang@163.com
// @host 127.0.0.1
// @BasePath /api/v1
func main() {
	mysql.MysqlInit()
	e := router.RouterInit()
	err := e.Run(":8080")
	if err != nil {
		log.Panic(err)
		return
	}
}
