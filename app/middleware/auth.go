package middleware

import (
	"net/http"
	controller "peduli-covid/controllers"
	"peduli-covid/helpers/messages"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   int `json:"id"`
	Role int `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	Role            []int
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controller.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(userID, role int) string {
	claims := JwtCustomClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

func (jwtConf *ConfigJWT) VerifyRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		claims := &JwtCustomClaims{}

		tokenAuthHeader := ctx.Request().Header.Get("Authorization")
		if !strings.Contains(tokenAuthHeader, "Bearer") {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrInvalidBearerToken)
		}

		tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)

		_, err = jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtConf.SecretJWT), nil
		})
		if err != nil {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, err)
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrExpiredToken)
		}

		var isValidRole bool
		for _, role := range jwtConf.Role {
			if role == claims.Role {
				isValidRole = true
			}
		}

		if !isValidRole {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrInvalidRole)
		}

		ctx.Set("user", claims)

		return next(ctx)
	}
}

// GetUser from jwt ...
func GetUser(c echo.Context) (res *JwtCustomClaims) {
	user := c.Get("user")
	if user != nil {
		res = user.(*JwtCustomClaims)
	}
	return res
}
