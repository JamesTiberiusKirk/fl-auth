package controllers

import (
	"fl-auth/db"
	"fl-auth/lib"
	"fl-auth/lib/auth"
	"fl-auth/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

/* Error messages. */
const (
	MSG_LOGGED_IN           = "Logged in"
	MSG_UNAUTHORIZED        = "Login failed"
	MSG_JWT_ERR             = "JWT error"
	MSG_JWT_INVALID         = "JWT invalid"
	MSG_USER_EXISTS         = "User already exists"
	MSG_DATABASE_ERR        = "Database error"
	MSG_INTERNAL_SERVER_ERR = "Internal server error"
)

/* Success messages. */
const (
	MSG_JWT_VALID               = "JWT valid"
	MSG_SUCCESSFULLY_REGISTERED = "Successesfully registered"
	MSG_OK                      = "OK"
)

/* Controller to user registration. */
func Register(c echo.Context) error {
	db := c.Get("db").(*db.Client)

	// Struct binding
	var nUser models.User
	if bindErr := c.Bind(&nUser); bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr)
	}
	defaultRoles := make([]string, 1)
	defaultRoles[0] = "user"
	nUser.Roles = defaultRoles

	// Input validatin
	if validationErr := nUser.IsValid(); validationErr != nil {
		log.Fatalln(validationErr)
		return c.JSON(http.StatusBadRequest, validationErr)
	}

	//Check for user
	existingUser, dbErr := db.CheckUser(nUser.Email)
	if dbErr != nil {
		log.Fatalln(dbErr)
		return c.String(http.StatusInternalServerError, MSG_DATABASE_ERR)
	}
	if existingUser {
		return c.String(http.StatusBadRequest, MSG_USER_EXISTS)
	}

	// Hash the password
	hashedPass, hashErr := lib.Hash(nUser.Password)
	if hashErr != nil {
		log.Fatalln(hashErr)
		return c.String(http.StatusInternalServerError, MSG_INTERNAL_SERVER_ERR)
	}
	nUser.Password = hashedPass

	// Database insertion
	if dbErr := db.AddUser(nUser); dbErr != nil {
		log.Fatalln(dbErr)
		return c.String(http.StatusInternalServerError, MSG_DATABASE_ERR)
	}

	return c.String(http.StatusOK, MSG_SUCCESSFULLY_REGISTERED)
}

/* Controller to log in the user and return JWT token. */
func Login(c echo.Context) error {
	db := c.Get("db").(*db.Client)

	// Struct binding
	var userLogin models.UserLoginForm
	if bindErr := c.Bind(&userLogin); bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr)
	}

	// Query hashed db pass
	dbUser, dbErr := db.GetUserByEmail(userLogin.Email)
	if dbErr != nil {
		return c.String(http.StatusInternalServerError, MSG_DATABASE_ERR)
	}

	// Compare hash and plaintext
	if !lib.VerifyHash(dbUser.Password, userLogin.Password) {
		return c.String(http.StatusUnauthorized, MSG_UNAUTHORIZED)

	}

	// Create JWT
	userJwt, jwtErr := auth.GenerateJWT(dbUser)
	if jwtErr != nil {
		return c.String(http.StatusInternalServerError, MSG_JWT_ERR)
	}

	resp := models.LoginResponseDto{
		Message: MSG_LOGGED_IN,
		Jwt:     userJwt,
	}

	return c.JSON(http.StatusOK, resp)
}

/* Controller to verify the jwt tokens. */
func VerifyMe(c echo.Context) error {
	var userJwt models.JwtDto
	if bindErr := c.Bind(&userJwt); bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}
	if validateErr := userJwt.IsValid(); validateErr != nil {
		return c.JSON(http.StatusBadRequest, validateErr.Error())
	}

	_, jwtErr := auth.ValidateToken(userJwt.Jwt)
	if jwtErr != nil {
		return c.String(http.StatusUnauthorized, MSG_JWT_INVALID)
	}

	return c.String(http.StatusOK, MSG_OK)
}
