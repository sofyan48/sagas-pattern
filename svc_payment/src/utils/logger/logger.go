package logger

import (
	"os"

	"github.com/sofyan48/go-cinlog/cinlog"
	v1 "github.com/sofyan48/go-cinlog/cinlog/v1"
	"github.com/sofyan48/go-cinlog/config"
	"github.com/sofyan48/go-cinlog/entity"
)

// Logger ...
type Logger struct{}

// LoggerHandler ...
func LoggerHandler() *Logger {
	return &Logger{}
}

// LoggerInterface ...
type LoggerInterface interface {
	Save(uuid, status string, data map[string]interface{}) (*entity.LoggerEventHistory, error)
	Find(uuid, action string) (*entity.LoggerEventHistory, error)
	Get(uuid string) (*entity.LoggerEventHistory, error)
}

func loggerInit() *v1.V1Session {
	cfg := config.NewConfig().SetClient(os.Getenv("LOGGER_URL"), os.Getenv("APP_NAME"))
	return cinlog.NewSession(cfg).V1()
}

// Save logger
func (log *Logger) Save(uuid, status string, data map[string]interface{}) (*entity.LoggerEventHistory, error) {
	logger := loggerInit()
	return logger.Save(uuid, status, data)

}

// Find logger
func (log *Logger) Find(uuid, action string) (*entity.LoggerEventHistory, error) {
	logger := loggerInit()
	return logger.Find(uuid, action)
}

// Get logger
func (log *Logger) Get(uuid string) (*entity.LoggerEventHistory, error) {
	logger := loggerInit()
	return logger.Get(uuid)
}
