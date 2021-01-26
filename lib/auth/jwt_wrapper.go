package auth

import (
	"errors"
	"fl-auth/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ErrClaimParse = errors.New("Couldn't parse claim")
var ErrJwtExpired = errors.New("JWT has expired")
var ErrJwtInvalid = errors.New("JWT invalid")

const (
	SERVICE_NAME         = "fl-auth"
	JWT_SECRET_ENV       = "JWT_SECRET"
	JWT_EXPIRATION_HOURS = 24
)

/* JwtWrapper wraps the signing key and the issuer. */
// type JwtWrapper struct {
// 	SecretKey       string
// 	Issuer          string
// 	ExpirationHours int64
// }

/* JwtClaim adds email and roles as a claim to the token. */
type JwtClaim struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func getJwtSecretFromEnv() string {
	return os.Getenv(JWT_SECRET_ENV)
}

/* Function to generate JWT token. */
func GenerateJWT(user models.User) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: user.Email,
		Roles: user.Roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(JWT_EXPIRATION_HOURS)).Unix(),
			Issuer:    SERVICE_NAME,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(getJwtSecretFromEnv()))
	if err != nil {
		return
	}

	return
}

/* Function to validate the JWT token and return the custom claims. */
func ValidateToken(userToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		userToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getJwtSecretFromEnv()), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = ErrClaimParse
		return
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		err = ErrJwtExpired
		return
	}

	return
}
