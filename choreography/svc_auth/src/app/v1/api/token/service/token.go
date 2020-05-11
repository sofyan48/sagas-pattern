package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	clientRepo "github.com/sofyan48/svc_auth/src/app/v1/api/client/repository"
	"github.com/sofyan48/svc_auth/src/app/v1/api/token/entity"
	"github.com/sofyan48/svc_auth/src/utils/database"
	"github.com/sofyan48/svc_auth/src/utils/openssl"
	"github.com/sofyan48/svc_auth/src/utils/redis"
)

// TokenService ...
type TokenService struct {
	SSL    openssl.OpensslInterface
	TrxDB  *gorm.DB
	Redis  redis.RedisLibInterface
	Client clientRepo.ClientRepositoryInterface
}

// TokenServiceHandler ...
func TokenServiceHandler() *TokenService {
	return &TokenService{
		SSL:    openssl.OpensslHandler(),
		TrxDB:  database.GetTransactionConnection(),
		Redis:  redis.RedisLibHandler(),
		Client: clientRepo.ClientRepositoryHandler(),
	}
}

// TokenServiceInterface ...
type TokenServiceInterface interface {
	ClientCredential(payload *entity.ClientCredentialRequest) (interface{}, error)
}

// ClientCredential ...
func (service *TokenService) ClientCredential(payload *entity.ClientCredentialRequest) (interface{}, error) {
	clientData, err := service.Client.GetClientByKey(payload.ClientKey)
	if err != nil {
		return nil, err
	}
	signBytes, err := ioutil.ReadFile("." + clientData.ClientPrivateKey)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Println("INI")
		return nil, err
	}

	sessionData := &entity.AuthDataModels{}
	sessionData.Level = "client"
	sessdata, _ := json.Marshal(sessionData)
	tokenExpires := time.Now().Add(time.Hour * 1).Unix()
	claims := &entity.Claims{}
	claims.Session = string(sessdata)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: tokenExpires,
		Issuer:    clientData.ClientKey,
		Id:        clientData.ClientSecret,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	dataClientObj, _ := json.Marshal(clientData)
	service.Redis.RowsCached(tokenString, dataClientObj, tokenExpires)
	credentialResponse := &entity.ClientCredentialResponse{}
	credentialResponse.AccessToken = tokenString
	credentialResponse.Expires = tokenExpires
	credentialResponse.Type = "bearer"
	return credentialResponse, nil
}
