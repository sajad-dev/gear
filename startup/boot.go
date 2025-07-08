package startup

import "github.com/sajad-dev/gear/internal/config"

// type Bootstrap struct {
// 	Config config.ConfigSt
	
// }

func Boot(bootstrap config.ConfigSt) {
	config.Config = bootstrap
}
