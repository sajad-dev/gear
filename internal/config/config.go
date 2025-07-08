package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfigSt struct {
	Db     *gorm.DB
	Tables []any
	Http   func(int, *gorm.DB) (*gin.Engine, error)

	APP_NAME    string
	DESCRIPTION string
}

var Config ConfigSt
