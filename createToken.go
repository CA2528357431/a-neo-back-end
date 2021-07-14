package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


const (
	key = "myownsecret"
	salt =  "hyperloop,tesla,elon"
)



type jwtCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}



func createToken(id string) string {

	claims := jwtCustomClaims{
		id + salt,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)


	token, err := t.SignedString([]byte("myownsecret"))

	checkErr(err)

	return token

}



func beta(token *jwt.Token) (interface{}, error) {
	return []byte(key), nil
}
func parserToken(token string)  (*jwtCustomClaims,error){


	t, err := jwt.ParseWithClaims(token,&jwtCustomClaims{},beta)

	if t == nil{
		return nil,err
	}

	claims, ok := t.Claims.(*jwtCustomClaims)

	if ok && t.Valid{

		return claims, nil
	}

	return nil, err

}



func tokenCheck(getToken string,getId string) bool {
	tokenData,err := parserToken(getToken)
	checkErr(err)

	tokenId:=tokenData.Id
	tokenTime:=tokenData.ExpiresAt
	currentTime:=time.Now().Unix()

	judge := tokenId==(getId+salt)&&currentTime<tokenTime

	return judge

}