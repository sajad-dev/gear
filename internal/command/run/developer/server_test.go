package developer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_Run(t *testing.T) {
	st := AutoCompile{
		ExecPath: "sample/main.go",
	}

	err := st.Run()
	assert.NoError(t,err)
}
