package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"oldshop/model"
)

var jwtSecret = []byte("my_secret")

//生成jwt token

func GenJwtToken(id uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//验证jwt token是否正确
func CheckJwtToken(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", err
	}
	id := uint(claims["id"].(float64))
	username := claims["username"].(string)
	return id, username, nil
}

//jwt token验证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(200, model.NewOkMsgWithData("token为空", false, nil))
			c.Abort()
			return
		}
		id, username, err := CheckJwtToken(tokenString)
		if err != nil {
			c.JSON(200, model.NewOkMsgWithData("token验证失败", false, nil))
			c.Abort()
			return
		}
		c.Set("id", id)
		c.Set("username", username)
	}
}
