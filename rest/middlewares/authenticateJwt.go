package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"miniShop/config"
	"miniShop/util"
	"net/http"
	"strings"
	"time"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			util.SendError(w, http.StatusUnauthorized, "missing token")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.SendError(w, http.StatusUnauthorized, "invalid token format")
			return
		}

		tokenParts := strings.Split(parts[1], ".")
		if len(tokenParts) != 3 {
			util.SendError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		headerPayload := tokenParts[0] + "." + tokenParts[1]
		signature := tokenParts[2]

		cnf := config.GetConfig()

		mac := hmac.New(sha256.New, []byte(cnf.JWTSecretKey))
		mac.Write([]byte(headerPayload))
		expectedSig := base64URLEncode(mac.Sum(nil))

		if !hmac.Equal([]byte(expectedSig), []byte(signature)) {
			util.SendError(w, http.StatusUnauthorized, "invalid signature")
			return
		}

		payloadBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).
			DecodeString(tokenParts[1])
		if err != nil {
			util.SendError(w, http.StatusUnauthorized, "invalid payload")
			return
		}

		var claims util.Claims
		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			util.SendError(w, http.StatusUnauthorized, "invalid claims")
			return
		}

		now := time.Now().UTC().Unix()

		if now < claims.Nbf {
			util.SendError(w, http.StatusUnauthorized, "token not active yet")
			return
		}

		if now > claims.Exp {
			util.SendError(w, http.StatusUnauthorized, "token expired")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
