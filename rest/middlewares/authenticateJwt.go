package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"miniShop/config"
	"net/http"
	"strings"
)

// AuthenticateJWT verifies the JWT from the Authorization header before allowing the request to reach the next handler.
//
// WHY:
// - Ensures only authenticated requests access protected routes
// - Keeps authentication logic centralized (middleware pattern)
func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Step 1: Read the Authorization header
		// WHY: JWT is typically sent as "Authorization: Bearer <token>"
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized request", http.StatusUnauthorized)
			return
		}

		// Step 2: Split "Bearer <token>"
		// WHY: We must extract only the JWT part
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := authParts[1]

		// Step 3: Split JWT into header.payload.signature
		// WHY: JWT must always have exactly 3 parts
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		encodedHeader := tokenParts[0]
		encodedPayload := tokenParts[1]
		receivedSignature := tokenParts[2]

		// Step 4: Rebuild the unsigned token
		// WHY: Signature is calculated over "header.payload"
		unsignedToken := encodedHeader + "." + encodedPayload

		// Load application configuration
		// WHY: Secret key should come from config, not hardcoded
		cnf := config.GetConfig()

		// Convert secret and message to bytes
		// WHY: HMAC works on byte slices
		secretKeyBytes := []byte(cnf.JWTSecretKey)
		messageBytes := []byte(unsignedToken)

		// Step 5: Recalculate signature using same algorithm & secret
		// WHY: If token was altered, signature will not match
		signingMAC := hmac.New(sha256.New, secretKeyBytes)
		signingMAC.Write(messageBytes)
		signatureBytes := signingMAC.Sum(nil)
		calculatedSignature := encodeBase64URLNoPadding(signatureBytes)

		// Step 6: Compare calculated signature with received signature
		// WHY: This verifies token integrity and authenticity
		if calculatedSignature != receivedSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Step 7: Token is valid → allow request to proceed
		next.ServeHTTP(w, r)
	})
}

// encodeBase64URLNoPadding encodes data using Base64 URL encoding without padding.
//
// WHY:
// - JWT requires URL-safe Base64 encoding
// - Padding is explicitly forbidden by JWT spec
func encodeBase64URLNoPadding(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
