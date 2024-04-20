package bootstrap

import (
	"go-api/pkg/config"
	"go-api/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_age"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
