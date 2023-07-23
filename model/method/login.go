package method

import (
	"encoding/base64"
	"ini/dao/mysql"
	model "ini/model/user_struct"
	"ini/router/middleware"
	"log"
)

// 账号密码登录
func UserLogin1(LoginUser model.User) (string, error, int) {
	enc_str := base64.StdEncoding.EncodeToString([]byte(LoginUser.Password))
	err := mysql.DB.Where("username = ? AND password = ?", LoginUser.Username, enc_str).First(&LoginUser).Error
	if err != nil {
		return "", err, 1
	}
	signedToken, errr := middleware.GenerateToken(LoginUser.Username)
	if err != nil {
		log.Println(errr)
		return "", errr, 2
	}
	return signedToken, nil, 0
}
