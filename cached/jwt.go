package cached

import (
	"fmt"
	"time"

	"github.com/linqiurong2021/gin-book-frontend/config"
	"github.com/linqiurong2021/gin-book-frontend/models"

	"github.com/dgrijalva/jwt-go"
)

// Claims 自定义
type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// Create 创建
func Create(user *models.User) (string, error) {
	//
	mySigningKey := []byte(config.Conf.JWTSignKey)
	//
	now := time.Now().Unix()
	expiresAt := now + config.Conf.TokenExpireMinutes*60

	fmt.Println(expiresAt, "expiresAt")
	claims := Claims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singString, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", singString, err)
	return singString, err
}

// Parse 解析Token
func Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JWTSignKey), nil
	})
	return token, err
}

// Check Token校验
func Check(jwtToken *jwt.Token) (*Claims, bool) {
	claims, ok := jwtToken.Claims.(*Claims)
	return claims, ok
}
