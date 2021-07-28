package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/jwtRS256.key" // openssl genrsa -out app.rsa keysize
)

var (
	signKey *rsa.PrivateKey
)

func CreateToken(userId uint64) (string, error) {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return "", err
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, err := at.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
