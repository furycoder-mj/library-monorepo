package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"

	"epam.com/web-services/library-management/library-service/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const ( // openssl genrsa -out app.rsa keysize
	pubKeyPath = "keys/jwtRS256.key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
)

func getPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

type UserTokenClaims struct {
	*jwt.StandardClaims
	TokenType string
	models.UserToken
}

func ValidateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		publicKey, pubErr := getPublicKey()
		if pubErr != nil {
			return nil, pubErr
		}
		return publicKey, nil
	})
	// If the token is missing or invalid, return error
	if err != nil {
		return nil, err
	}
	return token, nil
}
