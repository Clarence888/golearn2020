package app

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

/*
生成jwttoken
 */
func GenerateToken(appKey,appSecret string) (string,error)  {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

/*
解析和校验token
 */
func ParseToken(token string) (*Claims,error) {
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(),nil
	})

	if tokenClaims != nil {
		if claims,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid{
			return claims,nil
		}
	}

	return nil,err
}


