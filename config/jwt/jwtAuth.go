package jwt

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/comic-go/lib"
	"github.com/golang-jwt/jwt"
)

var signKey *rsa.PrivateKey
var verifyKey *rsa.PublicKey

func CreateToken() (string, error) {
	claims := jwt.MapClaims{
		"iss": "jwthost",
		"sub": "user_id1234",
		"exp": time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}

	// ヘッダーとペイロード生成
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signBytes, err := ioutil.ReadFile("jwt-key.pem")
	if err != nil {
		return "", err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}

	// トークンに署名を付与
	tokenString, err := token.SignedString(signKey)

	return tokenString, err
}

func VerifyToken(tokenString string) (string, error) {

	varifyBytes, err := ioutil.ReadFile("jwt-key-pub.pem")
	if err != nil {
		return "", err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(varifyBytes)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		parts := strings.Split(tokenString, ".")
		err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], verifyKey)
		if err != nil {
			fmt.Println(err)
		}
		// なんか上手くいかなかった
		// if _, err := token.Method.(*jwt.SigningMethodHMAC); err {
		// 	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// }
		return verifyKey, nil
	})

	fmt.Println(token.Valid)

	if !token.Valid {
		fmt.Println("token is invalid")
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return lib.ToString(claims["sub"]), err
}
