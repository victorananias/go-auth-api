package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
)

type JwtCreator struct{}

type Jwt struct {
	Token string
}

func (jwtCreator *JwtCreator) CreateFromUser(user User) Jwt {
	header := struct {
		alg string
		typ string
	}{
		alg: "HS256",
		typ: "JWT",
	}
	payload := struct {
		email    string
		password string
	}{
		email:    user.Email,
		password: user.Password,
	}

	jsonHeader, _ := json.Marshal(header)
	jsonPayload, _ := json.Marshal(payload)

	base64Header := strings.Replace(base64.StdEncoding.EncodeToString(jsonHeader), "=", "", -1)
	base64Payload := strings.Replace(base64.StdEncoding.EncodeToString(jsonPayload), "=", "", -1)
	settings, err := newSettings()
	if err != nil {
		log.Fatalf(err.Error())
	}
	secret := settings.JwtSecret
	data := base64Header + "." + base64Payload
	h := hmac.New(sha256.New, []byte(secret))
	//h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum([]byte(data)))
	token = strings.Replace(token, "+", "-", -1)
	token = strings.Replace(token, "/+", "-", -1)
	token = strings.Replace(token, "\\", "_", -1)
	token = strings.Replace(token, "/", "_", -1)

	return Jwt{token}
}
