package middleware

import "net/http"

type OtherMiddleware struct {
}

func NewOtherMiddleware() *OtherMiddleware {
	return &OtherMiddleware{}
}

func (m *OtherMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Passthrough to next handler if need
		next(w, r)
	}
}
