package middleware

import (
	"backend/dao"
	"backend/utils"
	"backend/utils/errmsg"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

// JwtGenerateToken 创造TOKEN
/**
token的创建需要
	1.签发时间
	2.结束时间
	3.签发人
	前三条可放在StandardClaims模板结构中，第四条需要自建userclaims
	4.Token所携带的内容 这里是name，代表是谁登录的
*/
func JwtGenerateToken(name string, d time.Duration) (token string, code uint) {
	// 过期时间
	expireTime := time.Now().Add(d)
	// 创建模板
	registeredClaims := jwt.RegisteredClaims{
		// 设置过期时间 在当前基础上 添加一个小时后 过期
		//ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.TokenExpire") * time.Millisecond)),
		// 颁发时间 也就是生成时间
		//IssuedAt: jwt.NewNumericDate(time.Now()),
		//主题
		Subject:   "Token",
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "app",
	}
	// 创建userclaims
	claims := UserClaims{
		RegisteredClaims: registeredClaims,
		UserName:         name,
	}
	// 获取token
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//fmt.Println(Token)

	// 将token转换为字符串
	token, err := Token.SignedString(JwtKey)
	if err != nil {
		fmt.Printf("token to string fail:%s", err)
		code = errmsg.TOKEN_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

// JwtParseUser 解析TOKEN 将token转换为claims包装的对象并将其中用户信息提取
func JwtParseUser(tokenString string) (username string, err error) {
	if tokenString == "" {
		err = errors.New("token is empty")
		username = ""
	}
	// 通用方法 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	// 断言token的claims
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		username = claims.UserName
		err = nil
	} else {
		username = ""
		err = errors.New("claims not UserClaims")
	}
	return
}

// JwtToken jwt的中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var role uint
		role = 0
		var code uint
		// 获取权限等级Role
		//if "" == c.Request.Header.Get("role") {
		//	role = 0
		//} else {
		//	role, _ = strconv.Atoi(c.Request.Header.Get("role"))
		//}
		//// 此页面无需权限
		//if role == 0 {
		//	c.Next()
		//	return
		//}

		// 获取验证信息（Token）
		Authorization := c.Request.Header.Get("Authorization")

		// 请求头是否有token
		if Authorization == "" {
			code = errmsg.INEXISTENCE_TOKEN
			c.JSON(http.StatusOK, gin.H{
				"state":   code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 检查token格式
		Authorizations := strings.SplitN(Authorization, " ", 2)
		if Authorizations[0] != "Bearer" || len(Authorizations) < 2 {
			code = errmsg.TOKEN_TYPE_ERROR
			c.JSON(http.StatusOK, gin.H{
				"state":   code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 解析token
		userName, err := JwtParseUser(Authorizations[1])
		if err != nil {
			code = errmsg.TOKEN_PARSE_ERROR
			c.JSON(http.StatusOK, gin.H{
				"state":   code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 验证权限
		ok := dao.ExamineRole(userName, role)
		if !ok {
			code = errmsg.INSUFFICIENT_ROLE
			c.JSON(http.StatusOK, gin.H{
				"state":   code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 设置用户信息
		c.Set("userName", userName)
		c.Next()
	}
}
