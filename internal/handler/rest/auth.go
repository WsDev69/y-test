package rest

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"y-test/internal/security"
)

type AuthenticationHandler struct {
	s *security.Jwt
}

func NewJwtAuth(s *security.Jwt) *AuthenticationHandler {
	return &AuthenticationHandler{s: s}
}

func (ah *AuthenticationHandler) Authentication(f http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r*http.Request) {
		errAuth := ah.s.TokenValid(r)
		if errAuth != nil {
			logrus.Trace("User token is invalid:", errAuth)
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		f.ServeHTTP(rw,r)
	})

}



