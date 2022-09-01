package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/victorananias/go-auth-api/models"
	"github.com/victorananias/go-auth-api/repositories"
	"github.com/victorananias/go-auth-api/requests"
	"github.com/victorananias/go-auth-api/settings"
	"golang.org/x/crypto/bcrypt"
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
	Payload   JwtPayload
	Header    JwtHeader
	Signature string
	Token     string
}

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService() *AuthService {
	service := &AuthService{}
	service.userRepository = repositories.NewUserRepository()
	return service
}

func (service *AuthService) Register(request requests.RegisterRequest) error {
	user, err := service.userRepository.FindUserByUsername(request.Username)
	if err == nil {
		return errors.New("username already in use")
	}
	user = models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Username:  request.Username,
	}
	user.PasswordHash = service.hashPassword(request.Password)
	user.CreatedAt = time.Now().Format("YYYY-MM-DD hh:mm:ss")
	service.userRepository.Create(user)
	return nil
}

func (service *AuthService) Login(username, password string) (string, error) {
	user, err := service.userRepository.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("failed to login")
	}
	logged := service.compareHashAndPassword(user.PasswordHash, password)
	if !logged {
		return "", errors.New("failed to login")
	}
	jwt := service.NewJwt(username)
	return jwt.Token, nil
}

func (service *AuthService) NewJwt(username string) Jwt {
	jwt := Jwt{}
	jwt.Header = JwtHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	jwt.Payload = JwtPayload{
		Username: username,
		Exp:      int(time.Now().Add(time.Minute).Unix()),
	}
	jwt.Signature = service.generateSignature(jwt.Header, jwt.Payload)
	jwt.Token = strings.Join([]string{AsBase64URL(jwt.Header), AsBase64URL(jwt.Payload), jwt.Signature}, ".")
	return jwt
}

func NewJwtFromToken(token string) *Jwt {
	jwt := Jwt{Token: token}
	s := strings.Split(token, ".")
	header, payload, signature := s[0], s[1], s[2]
	FromBase64(&jwt.Header, header)
	FromBase64(&jwt.Payload, payload)
	jwt.Signature = signature
	return &jwt
}

func FromBase64(object interface{}, stringBase64 string) {
	bytesJson, err := base64.RawStdEncoding.DecodeString(stringBase64)
	reader := bytes.NewReader(bytesJson)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(reader).Decode(&object)
	if err != nil {
		log.Fatal(err)
	}
}

func (service *AuthService) Validade(token string) bool {
	// s := strings.Split(token, ".")
	// payload, header, signature := s[0], s[1], s[2]
	return false
	// validSignature := service.generateSignature(payload, header)
	// return signature == validSignature
}

func (jwt *Jwt) isExpired() bool {
	return int(time.Now().Unix()) > jwt.Payload.Exp
}

func (service *AuthService) generateSignature(header JwtHeader, payload JwtPayload) string {
	settings, _ := settings.NewSettings()
	secret := settings.Jwt.Secret

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(strings.Join([]string{jsonToBase64Url(header), jsonToBase64Url(payload)}, ".")))

	return hex.EncodeToString(hash.Sum(nil))
}

func AsBase64URL(object interface{}) string {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString([]byte(jsonBytes))
}

func jsonToBase64Url(j interface{}) string {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString([]byte(jsonBytes))
}

func (service *AuthService) compareHashAndPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return nil == err
}

func (service *AuthService) hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}
