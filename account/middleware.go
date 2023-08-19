package account

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"lending-service/constant"
	"lending-service/utility/response"
	"lending-service/utility/wraped_error"
)

type middleware struct {
}

func newMiddleware() middleware {
	return middleware{}
}

func (m middleware) validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if len(authorizationHeader) <= 0 {
			response.ErrorWrapped(w, wraped_error.WrapError(fmt.Errorf("header Authorization empty"), http.StatusBadRequest))
			return
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			response.ErrorWrapped(w, wraped_error.WrapError(fmt.Errorf("headaer Authorization invalid"), http.StatusBadRequest))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return constant.JWT_KEY, nil
		})
		if err != nil {
			if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
				response.ErrorWrapped(w, wraped_error.WrapError(fmt.Errorf("token expired"), http.StatusBadRequest))
				return
			}
			logrus.WithField("error_message", err.Error()).Error("invalid token")
			response.ErrorWrapped(w, wraped_error.WrapError(fmt.Errorf("invalid token"), http.StatusBadRequest))
			return
		}

		byteClaims, err := json.Marshal(token.Claims)
		if err != nil {
			response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusInternalServerError))
			return
		}

		var account Account
		if err := json.Unmarshal(byteClaims, &account); err != nil {
			response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusInternalServerError))
			return
		}

		ctx := context.WithValue(context.Background(), "account", account)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
