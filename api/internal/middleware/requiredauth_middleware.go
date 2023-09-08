package middleware

import (
	"github.com/Hudayberdyyev/go-zero-playground/api/domain"
	"net/http"
)

type RequiredAuthMiddleware struct {
	auth domain.Authorization
}

func NewRequiredAuthMiddleware(auth domain.Authorization) *RequiredAuthMiddleware {
	return &RequiredAuthMiddleware{
		auth: auth,
	}
}

func (m *RequiredAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
