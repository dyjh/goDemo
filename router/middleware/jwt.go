package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jailer "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
)

var MyJwtSecret = []byte("123456789")

func GetJwtToken(UserId int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    UserId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 添加过期时间为2个小时
	})
	tokenString, _ := token.SignedString(MyJwtSecret)
	return tokenString
}

func ParseClaimsConfig(ctx iris.Context) UserInfo {
	mapClaims := ctx.Values().Get(jailer.DefaultContextKey).(*jwt.Token).Claims.(jwt.MapClaims)
	id, _ := mapClaims["id"].(int64)
	userInfo := UserInfo{Id:id}
	return userInfo
}

func JwtHandler() *jailer.Middleware {
	return jailer.New(jailer.Config{
		// 验证 jwt 的 token 的方法
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// 自己加密的密钥
			return MyJwtSecret, nil
		},

		SigningMethod: jwt.SigningMethodHS256,
		Expiration:true,
	})

}

type UserInfo struct {
	Id int64 `json:"id"`
}