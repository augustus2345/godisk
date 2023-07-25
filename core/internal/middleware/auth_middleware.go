package middleware

import (
	"go_disk/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))

			return
		}
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			r.Header.Set("UserId", string(rune(uc.Id)))
			r.Header.Set("UserIdentity", uc.Identity)
			r.Header.Set("UserName", uc.Name)

			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
