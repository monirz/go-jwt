package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/monirz/gojwt/response"
	"github.com/monirz/gojwt/utils"

	jw "github.com/dgrijalva/jwt-go"
)

func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jwtToken := r.Header.Get("Authorization")

		// log.Println("jwt ", jwtToken)
		if jwtToken == "" {
			response.ResponseError(w, http.StatusUnauthorized, "Missing Auth token", nil)
			return
		}

		jwt := &utils.JWT{}
		jwtToken = strings.TrimSpace(jwtToken)
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)

		claims, err := jwt.Validate(jwtToken)

		if err != nil {
			v, _ := err.(*jw.ValidationError)

			if v.Errors == jw.ValidationErrorExpired {
				log.Println("token is expired")
				response.ResponseError(w, http.StatusUnauthorized, "token is expired", nil)
				return
			}

			log.Println("Invalid Auth token", err)
			// w.WriteHeader(http.StatusUnauthorized)

			response.ResponseError(w, http.StatusUnauthorized, "Invalid Auth token", nil)

			return
		}

		// r.Header.Set("email", claims.Email)
		// log.Println("subject user_id", claims.Subject)
		r.Header.Set("user_id", claims.Subject)
		next.ServeHTTP(w, r)
	})
}
