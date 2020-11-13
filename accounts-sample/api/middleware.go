package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/drbyronw/accounts/service"
)

func (wa *WebApp) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if !strings.HasPrefix(auth, "Bearer") {
			http.Error(w, "invalid authorization content", http.StatusUnauthorized)
			return
		}

		t := strings.Split(auth, " ")[1]

		claims, err := service.VerifyJWT(t)
		if err != nil {
			http.Error(w, "invalid authorization content", http.StatusUnauthorized)
			return
		}

		if !claims.VerifyAudience(os.Getenv("AUDIENCE"), true) {
			http.Error(w, "Unable to use this token in this environment", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
