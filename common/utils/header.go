package utils

import (
	"net/http"
	"strings"
)

// GetToken 我们Token哪里获取?
// 1. URL Query String ?
// 2. Custom Header ?
// 3. Authorization Header
// test go mod cache
func GetToken(r *http.Request) (accessToken string) {
	auth := r.Header.Get("Authorization")
	al := strings.Split(auth, " ")
	if len(al) > 1 {
		accessToken = al[1]
	} else {
		// 兼容 Authorization <token>
		accessToken = auth
	}
	return
}
