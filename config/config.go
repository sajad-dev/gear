package config

import (
	"gorm.io/gorm"
)

type ConfigSt struct {
	Db     *gorm.DB
	Tables []any
	Http   func(int, *gorm.DB, chan struct{}) error

	APP_NAME    string
	DESCRIPTION string
}

var Config ConfigSt
