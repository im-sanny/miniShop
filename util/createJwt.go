package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// Header represents the JWT header section.
// WHY: The header tells the receiver how the token was created and which algorithm must be used to verify its signature.
type Header struct {
	Alg string `json:"alg"` // WHY: So the verifier knows which signing algorithm to use
	Typ string `json:"typ"` // WHY: Identifies the token as a JWT
}

// Payload represents the JWT payload (claims).
// WHY: This is the actual data we want to securely transmit between client and server.
type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// CreateJwt creates a signed JWT string.
// WHY: A JWT allows stateless authentication—no server-side session storage needed.
func CreateSignedJwt(secret string, claims Payload) (string, error) {

	// Create the header object
	jwtHeader := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Convert header struct into JSON bytes
	// WHY: JWT header must be JSON before Base64URL encoding
	headerJSONBytes, err := json.Marshal(jwtHeader)
	if err != nil {
		return "", err
	}

	// Encode header JSON using Base64 URL encoding
	// WHY: JWT must be safe to transmit via HTTP headers
	encodedHeader := encodeBase64URLNoPadding(headerJSONBytes)

	// Convert payload (claims) into JSON bytes
	// WHY: Claims must be represented as JSON in JWT
	payloadJSONBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err // WHY: Token cannot exist without payload
	}

	// Encode payload JSON using Base64 URL encoding
	// WHY: Ensures safely travel over HTTP
	encodedPayload := encodeBase64URLNoPadding(payloadJSONBytes)

	// Combine encoded header and payload
	// WHY: JWT signature is calculated over "header.payload"
	unsignedToken := encodedHeader + "." + encodedPayload

	// Create HMAC signer using SHA-256 and secret key
	// WHY: Ensures the token cannot be modified without the secret
	signingMAC := hmac.New(sha256.New, []byte(secret))

	// Write unsigned token into the HMAC hash
	// WHY: This is the data we want to protect from tampering
	signingMAC.Write([]byte(unsignedToken))

	// Generate raw signature bytes
	// WHY: Signature proves authenticity and integrity
	signatureBytes := signingMAC.Sum(nil)

	// Encode signature using Base64 URL encoding
	// WHY: JWT signature must also be URL-safe
	encodedSignature := encodeBase64URLNoPadding(signatureBytes)

	// Build the final signed JWT
	// WHY: JWT standard format is "header.payload.signature"
	signedJWT := unsignedToken + "." + encodedSignature

	// Return the signed token
	// WHY: Caller sends this token to the client
	return signedJWT, nil
}

// encodeBase64URLNoPadding encodes bytes using Base64 URL encoding without padding.
// Converts binary data into a URL-safe string.
// WHY:
// JWT specification requires Base64 URL encoding.
// Padding characters are not allowed in JWTs.
func encodeBase64URLNoPadding(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
