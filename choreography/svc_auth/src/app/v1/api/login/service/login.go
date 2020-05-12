package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	clientRepo "github.com/sofyan48/svc_auth/src/app/v1/api/client/repository"
	"github.com/sofyan48/svc_auth/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_auth/src/middleware"
	"github.com/sofyan48/svc_auth/src/utils/redis"
	"github.com/sofyan48/svc_auth/src/utils/requester"
)

// LoginService ...
type LoginService struct {
	Middleware middleware.DefaultMiddlewareInterface
	Requester  requester.RequesterInterface
	Client     clientRepo.ClientRepositoryInterface
	Redis      redis.RedisLibInterface
}

// LoginServiceHandler ...
func LoginServiceHandler() *LoginService {
	return &LoginService{
		Middleware: middleware.DefaultMiddlewareHandler(),
		Requester:  requester.RequesterHandler(),
		Client:     clientRepo.ClientRepositoryHandler(),
		Redis:      redis.RedisLibHandler(),
	}
}

// LoginServiceInterface ...
type LoginServiceInterface interface {
	PostLogin(payload *entity.LoginRequest, ctx *gin.Context) (*entity.ClientCredentialResponse, error)
}

// PostLogin ...
func (service *LoginService) PostLogin(payload *entity.LoginRequest, ctx *gin.Context) (*entity.ClientCredentialResponse, error) {
	claims, err := service.Middleware.GetSessionClaim(ctx)
	if err != nil {
		return nil, err
	}
	svcUsersURI := os.Getenv("USER_SERVICE_URL") + "/v1/login/session"
	data, _ := json.Marshal(payload)
	userResult, err := service.Requester.POST(svcUsersURI, service.Middleware.GetToken(ctx), data)
	if err != nil {
		return nil, err
	}
	// sessionModel := &entity.UserModelSession{}
	// err = json.Unmarshal([]byte(userResult), sessionModel)
	// if err != nil {
	// 	return nil, err
	// }

	clientData, err := service.Client.GetClientByKey(fmt.Sprintf("%s", claims["iss"]))
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

	tokenExpires := time.Now().Add(time.Hour * 1).Unix()
	claimsUser := &entity.Claims{}
	claimsUser.Session = string(userResult)
	claimsUser.StandardClaims = jwt.StandardClaims{
		ExpiresAt: tokenExpires,
		Issuer:    clientData.ClientKey,
		Id:        clientData.ClientSecret,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claimsUser)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return nil, err
	}
	dataClientObj, _ := json.Marshal(clientData)
	service.Redis.RowsCached(tokenString, dataClientObj, tokenExpires)
	credentialResponse := &entity.ClientCredentialResponse{}
	credentialResponse.AccessToken = tokenString
	credentialResponse.RefreshToken = ""
	credentialResponse.Expires = tokenExpires
	credentialResponse.Type = "bearer"
	return credentialResponse, nil
}
