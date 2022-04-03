package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: "secret-key-super-secreta",
		issuer:    "http://localhost:8080",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
	Roles []string `json:"roles,omitempty"`
}

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
	ip              = "http://localhost:8080"
)

func (s *JwtService) GenerateToken(claims *Claim, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}

	return tokenString, nil
	//claim := &Claim{
	//	Sum: id,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	//		Issuer:    s.issuer,
	//		IssuedAt:  time.Now().Unix(),
	//	},

	//Roles: roleNames,
}

func VerifyToken(tokenString string) (bool, *Claim) {
	claims := &Claim{}
	token, _ := getTokenFromString(tokenString, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}

	return false, claims
}

func getTokenFromString(tokenString string, claims *Claim) (*jwt.Token, error) {
	// O método ParseWithClaims vai preencher o meu ponteiro claims com as claims do usuário
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jWTPrivateToken), nil
	})
}

func GetClaims(tokenString string) Claim {
	claims := &Claim{}

	_, err := getTokenFromString(tokenString, claims)
	if err == nil {
		return *claims
	}
	return *claims
}
