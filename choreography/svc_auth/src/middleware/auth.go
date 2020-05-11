package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/entity"
)

// GetToken params
// @context: *gin.Context
// return gin.HandlerFunc
func (m *DefaultMiddleware) GetToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"]
	if len(token) < 1 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token Not Compatible",
		})
		context.Abort()
		return "event"
	}
	return token[0]
}

// AuthToken ...
func (m *DefaultMiddleware) AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := m.GetToken(ctx)
		tokens := strings.ReplaceAll(tokenString, "bearer ", "")
		rdsData, err := m.Redis.GetRowsCached(tokens)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
		}
		clientData := &entity.Clients{}
		err = json.Unmarshal([]byte(rdsData), clientData)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
		}
		signBytes, err := ioutil.ReadFile("." + clientData.ClientPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
		}
		signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("RS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return signKey, nil
		})

		if token == nil && err != nil {
			ctx.Next()
		}
	}

}

// GetSessionClaim params
func (m *DefaultMiddleware) GetSessionClaim(ctx *gin.Context) (map[string]interface{}, error) {
	tokens := strings.ReplaceAll(m.GetToken(ctx), "bearer ", "")
	rdsData, err := m.Redis.GetRowsCached(tokens)
	if err != nil {
		return nil, err
	}
	clientData := &entity.Clients{}
	err = json.Unmarshal([]byte(rdsData), clientData)
	if err != nil {
		return nil, err
	}
	signBytes, err := ioutil.ReadFile("." + clientData.ClientPrivateKey)
	if err != nil {
		return nil, err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	cekParse, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	return cekParse.Claims.(jwt.MapClaims), nil
}
