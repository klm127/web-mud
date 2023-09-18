package token

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pwsdc/web-mud/arg"
)

const (
	TokenIDKey  = "i"
	TokenExpKey = "x"
)

// used by JWT, it confirms the signing method and returns the secret key for parsing.
func keyfunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("couldn't parse token: bad signing method of " + token.Method.Alg())
	}
	return []byte(arg.Config.Http.JWTKey()), nil
}

// Returns a signed JWT token with the appropriate claim and expiration time.
func GenerateToken(userid int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[TokenExpKey] = time.Now().Add(60 * time.Minute).Format(time.RFC3339)
	claims[TokenIDKey] = fmt.Sprint(userid)
	tstring, err := token.SignedString([]byte(arg.Config.Http.JWTKey()))
	if err != nil {
		return "", err
	}
	return tstring, nil
}

// Gets the actor ID by parsing the cookie containing the JWT token
func ExtractID(c *gin.Context) (int64, error) {
	tok := c.Request.Header.Get(arg.Config.Http.AuthCookie())
	if tok == "" {
		return 0, errors.New("no token set")
	}
	token, err := jwt.Parse(tok, keyfunc)
	if err != nil {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	expires := claims[TokenExpKey].(string)
	t, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return 0, errors.New("badly formatted time")
	}
	if time.Since(t) > 0 {
		return 0, errors.New("token expired")
	}
	id_s := claims[TokenIDKey].(string)
	raw_id, err := strconv.Atoi(id_s)
	if err != nil {
		return 0, errors.New("badly formatted id")
	}
	return int64(raw_id), nil
}

// Writes the cookie header with the user ID
func WriteID(id int64, c *gin.Context) error {
	tok, err := GenerateToken(id)
	if err != nil {
		return err
	}
	c.Header(arg.Config.Http.AuthCookie(), tok)
	return nil
}
