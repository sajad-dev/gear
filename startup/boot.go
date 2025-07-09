package startup

import (
	"github.com/sajad-dev/gear/internal/cli"
	"github.com/sajad-dev/gear/config"
)

// type Bootstrap struct {
// 	Config config.ConfigSt

// }

func Boot(bootstrap config.ConfigSt) {
	config.Config = bootstrap
	cli.Cli()
}
