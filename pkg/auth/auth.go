package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return JwtAuthReq()

}

func JwtAuthReq() func(c *gin.Context) {
	return jwtAuthReq
}
func jwtAuthReq(c *gin.Context) {
	//extract token from cookies
	//validate token
	//if token is valid, set user id in context
	//if token is invalid, return 401
	//if token is missing, return 401
	//if refreshToken is expired, return 401
	c.Set("userId", "")

	token, err := c.Cookie("token")
	if err != nil || token == "" {
		c.JSON(401, gin.H{"error": "token missing"})
		c.Abort()
		return
	}
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil || refreshToken == "" {
		c.JSON(401, gin.H{"error": "refresh token missing"})
		c.Abort()
		return
	}

	claims, err := validateToken(token, refreshToken)
	if err != nil {
		if errors.Is(err, ErrorTokenExpired) {
			token, refreshToken, err = refreshTokens(claims)
			if err != nil {
				c.JSON(401, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			c.SetCookie("token", token, int(time.Minute*15), "/", "", false, true)
			c.SetCookie("refreshToken", refreshToken, int(time.Hour*24), "/", "", false, true)
		} else {
			ClearCookies(c)
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
	//set user id in context
	c.Set("userId", claims["sub"])
	c.Next()
}

func ClearCookies(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
}

func SetCookies(c *gin.Context, token, refreshToken string) {
	c.SetCookie("token", token, int(time.Minute*15), "/", "", false, true)
	c.SetCookie("refreshToken", refreshToken, int(time.Hour*24), "/", "", false, true)
}

func validateToken(token, refreshToken string) (jwt.MapClaims, error) {
	//validate token
	var tokenClaims jwt.MapClaims
	var refreshTokenClaims jwt.MapClaims
	_, tokenErr := jwtParser.ParseWithClaims(token, &tokenClaims, keyFunc)
	_, refreshTokenErr := jwtParser.ParseWithClaims(refreshToken, &refreshTokenClaims, keyFunc)
	if tokenErr != nil {
		if errors.Is(tokenErr, jwt.ErrTokenExpired) {
			if refreshTokenErr != nil {
				if errors.Is(refreshTokenErr, jwt.ErrTokenExpired) {
					return jwt.MapClaims{}, ErrorRefreshTokenExpired
				} else if errors.Is(refreshTokenErr, jwt.ErrTokenSignatureInvalid) {
					return jwt.MapClaims{}, ErrorRefreshTokenInvalid
				} else {
					return jwt.MapClaims{}, refreshTokenErr
				}
			}
			return refreshTokenClaims, ErrorTokenExpired
		} else if errors.Is(tokenErr, jwt.ErrTokenSignatureInvalid) {
			return jwt.MapClaims{}, ErrorTokenInvalid
		} else {
			return jwt.MapClaims{}, tokenErr
		}
	}
	return tokenClaims, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return jwtKey, nil
}

func refreshTokens(claims jwt.MapClaims) (string, string, error) {
	username := claims["sub"].(string)
	return generateTokenPair(username)
}

func GenerateTokenPair(c *gin.Context, username string) error {
	token, refreshToken, err := generateTokenPair(username)
	if err != nil {
		return err
	}
	SetCookies(c, token, refreshToken)
	return nil
}

func generateTokenPair(username string) (string, string, error) {
	refreshTokenUUID := uuid.NewString()
	refreshTokenClaims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "ToDoApp",
		Subject:   username,
		ID:        refreshTokenUUID,
	}

	tokenClaims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		Issuer:    "ToDoApp" + refreshTokenUUID,
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

var (
	jwtKey                   = []byte("x;=w.g*eK@v5]<DZsHM^kd,VB2N[-hA3}b8zECWfUt!m_a4:cX")
	jwtParser                = jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}))
	ErrorTokenExpired        = errors.New("token expired")
	ErrorTokenInvalid        = errors.New("token invalid")
	ErrorRefreshTokenExpired = errors.New("refresh token expired")
	ErrorRefreshTokenInvalid = errors.New("refresh token invalid")
)
