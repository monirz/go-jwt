package utils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GetAccessToken(authorization string) (string, error) {

	if len(authorization) < 1 {
		return "", errors.New("invalid Authorization")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid Authorization")
	}

	return strings.Split(authorization, " ")[1], nil
}

func ValidateTokenData(payload *CustomClaims) (bool, error) {

	if payload == nil || payload.Issuer == "" || payload.Subject == "" || payload.Audience == "" || payload.Issuer != os.Getenv("TOKEN_ISSUER") || payload.Audience != os.Getenv("TOKEN_AUDIENCE") || payload.Prm == "" {

		err := errors.New("invalid Access Token")
		return false, err
	}

	log.Println("payload ", payload.Subject)

	return true, nil
}

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateTokens(uid, accessTokenKey string, refreshTokenKey string) (signedToken string, signedRefreshToken string, err error) {

	token := &Token{}

	accessTokenValidityStr := os.Getenv("ACCESS_TOKEN_VALIDITY")
	refreshTokenValidityStr := os.Getenv("REFRESH_TOKEN_VALIDITY")

	accessTokenValidity, err := strconv.Atoi(accessTokenValidityStr)

	if err != nil {
		return "", "", err
	}

	refreshTokenValidity, err := strconv.Atoi(refreshTokenValidityStr)

	if err != nil {
		return "", "", err
	}

	//generate key pair
	accessTokenPayload := NewJWTPayload(os.Getenv("TOKEN_ISSUER"), os.Getenv("TOKEN_AUDIENCE"), uid, accessTokenKey, accessTokenValidity)
	refreshTokenPayload := NewJWTPayload(os.Getenv("TOKEN_ISSUER"), os.Getenv("TOKEN_AUDIENCE"), uid, refreshTokenKey, refreshTokenValidity)

	jwt := &JWT{}
	token.AccessToken, err = jwt.Encode(accessTokenPayload)

	if err != nil {
		log.Println(err)
		return "", "", err
	}
	token.RefreshToken, err = jwt.Encode(refreshTokenPayload)

	return token.AccessToken, token.RefreshToken, err
}
