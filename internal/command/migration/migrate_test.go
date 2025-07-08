package migration

import (
	"testing"

	"github.com/sajad-dev/gear/internal/config"
	"github.com/sajad-dev/gear/startup"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
}

func TestMigration_Migrate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	startup.Boot(config.ConfigSt{
		Db:     db,
		Tables: []any{&User{}},
	})

	assert.NoError(t, Migrate())

	assert.True(t, db.Migrator().HasTable(&User{}))
}
