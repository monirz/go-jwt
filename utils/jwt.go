package utils

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// SignedDetails
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var (
	privateKeyPath = "keys/private.pem"
	publicKeyPatth = "keys/public.pem"

	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type JWTPayload struct {
	Aud string
	Sub string
	Iss string
	Iat int64
	Exp int64
	Prm string
}

var (
	expireRule = time.Minute * 30
)

func NewJWTPayload(issuer string, audience string, subject string, param string, validity int) *JWTPayload {

	issueAt := time.Now()
	expiresAt := issueAt.Add(time.Duration(validity) * time.Second)

	return &JWTPayload{
		Sub: subject,
		Iss: issuer,
		Iat: issueAt.Unix(),
		Exp: expiresAt.Unix(),
		// Prm: param,
	}

}

type JWT struct {
}

func (j *JWT) Encode(payload *JWTPayload) (string, error) {
	//encode the token

	//for test only
	privateKeyPath = "keys/private.pem"

	signBytes, err := ioutil.ReadFile(privateKeyPath)

	if err != nil {
		//handle error
		return "", err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		log.Println("err >>", err)
		return "", err
	}

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			IssuedAt:  payload.Iat,
			ExpiresAt: int64(payload.Exp),
			Issuer:    payload.Iss,
			Audience:  payload.Aud,
			Subject:   payload.Sub,
		},
		payload.Prm,
	}

	return t.SignedString(signKey)
}

// Decode the token with public key and get claims
func (j *JWT) Decode(token string) (*CustomClaims, error) {

	verifyBytes, err := ioutil.ReadFile(publicKeyPatth)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	//decode the token with publicKey
	decodedPayload, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, err
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println("Invalid Token and/or Signature")
			return nil, err
		}
	}

	claims := decodedPayload.Claims.(*CustomClaims)
	return claims, err
}

func (j *JWT) Validate(token string) (*CustomClaims, error) {

	verifyBytes, err := ioutil.ReadFile(publicKeyPatth)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Println("Invlaid public key", err)
		return nil, err
	}

	//verify the pub key in the token
	tkn, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		return verifyKey, err
	})

	if err != nil {
		log.Println("err parsing token: ", err)
		return nil, err
	}

	claims := tkn.Claims.(*CustomClaims)

	return claims, err
}

func (j *JWT) readPublicKey() (string, error) {
	file, err := os.Open("../keys/public.pem")

	if err != nil {
		log.Println(err)

		return "", err
	}

	b := make([]byte, 1024)
	n, err := file.Read(b)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println(n, string(b))

	return string(b), nil

}

func (j *JWT) readPrivateKey() (string, error) {
	file, err := os.Open("../keys/private.pem")

	if err != nil {
		log.Println(err)

		return "", err
	}

	b := make([]byte, 1024)
	n, err := file.Read(b)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println(n, string(b))
	return string(b), nil
}

type CustomClaims struct {
	*jwt.StandardClaims

	Prm string
}
