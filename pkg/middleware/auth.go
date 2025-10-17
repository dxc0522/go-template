package middleware

import (
	"github.com/go-template/pkg/constant"
	"github.com/go-template/pkg/jwts"

	"net/http"
	"time"
)

type AuthInterceptorMiddleware struct {
	AccessSecret string
	AccessExpire int64
}

func NewAuthInterceptorMiddleware(AccessSecret string, AccessExpire int64) *AuthInterceptorMiddleware {
	return &AuthInterceptorMiddleware{
		AccessSecret: AccessSecret,
		AccessExpire: AccessExpire,
	}
}

func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(constant.AdminUserHeaderTokenKey)
		if tokenString == "" {
			cookie, err := r.Cookie(constant.AdminUserTokenKey)
			if err == nil {
				tokenString = cookie.Value
			}
		}
		if tokenString == "" {
			return
		}
		claims, err := jwts.ParseToken(tokenString, m.AccessSecret)
		if err != nil {
			return
		}
		AccessExpire := time.Duration(m.AccessExpire) * time.Second
		if time.Now().After(claims.ExpiresAt.Add(-AccessExpire)) {
			newTokenString, err := jwts.GenToken(claims.JwtPayLoad, m.AccessSecret, m.AccessExpire)
			if err != nil {
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:   constant.AdminUserTokenKey,
				Value:  newTokenString,
				Path:   "/",
				MaxAge: constant.AdminUserLoginMaxAge,
			})
		}
		next(w, r)
	}
}
