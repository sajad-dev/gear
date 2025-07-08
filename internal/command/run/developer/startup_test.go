package developer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartUp_GenerateFile(t *testing.T) {
	st := AutoCompile{
		ProjectPath: "./sample",
		BuildPath:   "./sample/build",
	}
	assert.NoError(t, st.GenerateFile())
}
