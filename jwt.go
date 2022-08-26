package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type JwtHeader struct {
	Alg string
	Typ string
}

type JwtPayload struct {
	Username string
	Exp      int
}

type Jwt struct {
}

func (jwt *Jwt) build(username string) string {
	header := jwt.generateHeader("HS256", "JWT")
	payload := jwt.generatePayload(username, 1547974082)
	signature := jwt.generateSignature(header, payload)

	return strings.Join([]string{header, payload, signature}, ".")
}

func (jwt *Jwt) generateHeader(algorithm string, typ string) string {
	jwtHeader := JwtHeader{
		Alg: algorithm,
		Typ: typ,
	}
	return jwt.jsonToBase64Url(jwtHeader)
}

func (jwt *Jwt) generatePayload(username string, expirationTime int) string {
	jwtPayload := JwtPayload{
		Username: username,
		Exp:      expirationTime,
	}
	return jwt.jsonToBase64Url(jwtPayload)
}

func (jwt *Jwt) jsonToBase64Url(j interface{}) string {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString([]byte(jsonBytes))
}

func (jwt *Jwt) generateSignature(header string, payload string) string {
	settings, _ := newSettings()
	secret := settings.Jwt.Secret

	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(strings.Join([]string{header, payload}, ".")))

	return hex.EncodeToString(h.Sum(nil))
}
