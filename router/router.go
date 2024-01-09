package router

import (
	"github.com/gin-gonic/gin"
	"ini/handler/login_and_signup"
	"ini/handler/upload"
)

func RouterInit() *gin.Engine {
	e := gin.Default()

	LoginGroup := e.Group("api/v1/")
	{
		LoginGroup.POST("/signup", login_and_signup.Signup)
		LoginGroup.POST("/pwd_login", login_and_signup.LoginWithPwd)
		LoginGroup.POST("/send_mail", login_and_signup.SendCode)
		LoginGroup.POST("/code_login", login_and_signup.LoginWithCode)
	}

	e.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	e.POST("/upload", upload.UploadImg)

	return e
}
