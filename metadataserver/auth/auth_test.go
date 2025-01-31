package auth

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"gotest.tools/assert"
)

func TestJWTAuthenticatorToken(t *testing.T) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	assert.NilError(t, err)

	auth := JWTAuthenticator{Key: key}

	token, err := auth.GenerateToken(60 * time.Second)
	assert.NilError(t, err)

	remaining, err := auth.VerifyToken(token)
	assert.NilError(t, err)
	assert.Assert(t, remaining > 58 && remaining < 60)
}

func TestJWTAuthenticatorExpiredToken(t *testing.T) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	assert.NilError(t, err)

	auth := JWTAuthenticator{Key: key}

	token, err := auth.GenerateToken(0)
	assert.NilError(t, err)

	time.Sleep(1 * time.Second)

	_, err = auth.VerifyToken(token)
	assert.Assert(t, err != nil)
}

func TestJWTAuthenticatorSetsAudience(t *testing.T) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	assert.NilError(t, err)

	auth := JWTAuthenticator{Key: key, Audience: "test-task"}

	token, err := auth.GenerateToken(60 * time.Second)
	assert.NilError(t, err)

	var claims jwt.StandardClaims
	_, _ = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	assert.Assert(t, claims.Audience == auth.Audience)
}

func TestJWTAuthenticatorVerifiesAudience(t *testing.T) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	assert.NilError(t, err)

	auth := JWTAuthenticator{Key: key, Audience: "audience-one"}
	otherAuth := JWTAuthenticator{Key: key, Audience: "audience-two"}

	token, err := auth.GenerateToken(60 * time.Second)
	assert.NilError(t, err)

	_, err = otherAuth.VerifyToken(token)
	assert.Assert(t, err != nil)
}
