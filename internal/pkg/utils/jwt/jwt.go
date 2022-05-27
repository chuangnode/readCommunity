package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
// 生成token
func GenerateToken(username, password string) (string, error) {
	signKey := []byte(viper.GetString("JwtSecret"))
	timeNow := time.Now()
	expireTime := timeNow.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "readCommunity",
		},
	}
	
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(signKey)
	if err != nil {
		fmt.Printf("GenerateToken failed,err:%v\n",err)
	}
	return token, err
}

// 解析token


func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JwtSecret")), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
		return claims, nil
	} else {
		fmt.Println(err)
	}
	return nil, err
}