package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func getAccessTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header not found")
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("Bearer token not found in authorization header")
	}

	accessToken := splitToken[1]
	return accessToken, nil
}
