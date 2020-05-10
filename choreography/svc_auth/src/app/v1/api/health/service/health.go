package service

import (
	"time"

	"github.com/sofyan48/svc_auth/src/app/v1/api/health/entity"
)

// HealthService ...
type HealthService struct{}

// HealthServiceHandler ...
func HealthServiceHandler() *HealthService {
	return &HealthService{}
}

// HealthServiceInterface ...
type HealthServiceInterface interface {
	HealthService() *entity.HealtResponse
}

// HealthService ...
func (service *HealthService) HealthService() *entity.HealtResponse {
	now := time.Now()
	result := &entity.HealtResponse{}
	result.Status = "OK"
	result.CreatedAt = &now
	return result
}
