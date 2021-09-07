package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string,error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		//strcov：uuidからintに変換してくれる
		Issuer: issuer,
		//jwtの有効期限
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1day
	})
	return claims.SignedString([]byte(SecretKey))

}

func ParseJwt(cookie string) (string, error) {

	//jwtトークンを解析して欲しい部分だけに分割してくれる
	token,err :=jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey),nil
	})

	//認証エラーハンドリング
	if err!=nil ||!token.Valid{
		return "",err
	}
	claims :=token.Claims.(*jwt.StandardClaims)

	return claims.Issuer,nil
}