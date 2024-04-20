// Package app 应用信息
package app

import "go-api/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "Production"
}
func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
