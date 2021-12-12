package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

/**
* @Author: kiki.zang
* @Date: 2021/12/12 5:50 PM
 */

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//自定义Claims
type CustomerClaims struct {
	UserId int64
	jwt.StandardClaims
}

/**
 *  生成token
 */
func GetJwt(userId int64, userName string) (tokenResult string) {
	//存活时间
	maxAge := 60 * 60 * 24

	//自定义claim
	customClaims := &CustomerClaims{
		UserId: userId, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    userName,                                                   // 用户名，
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)
	return tokenString
}

/**
 *  校验token
 */
func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("jwt")
		fmt.Sprintf("header jwt is %s", tokenString)
		//处理token
		token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRETKEY), nil
		})
		if err != nil {
			fmt.Println("check token is fail")
			c.JSON(-1, gin.H{
				"message": "check token is fail",
			})
		}

		if token.Valid {
			fmt.Println("token校验通过 开始搞事情")
			//执行之后的代码
			c.Next()
		}
		return
	}
}
