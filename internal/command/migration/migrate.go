package migration

import (
	"github.com/fatih/color"
	"github.com/sajad-dev/gear/config"
)

func Migrate() error {
	err := config.Config.Db.AutoMigrate(config.Config.Tables...)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	color.Green("Sucssfuly migrate :) - ;)")
	return nil
}
