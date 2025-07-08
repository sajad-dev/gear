package developer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TeatTestErr_RunTest(t *testing.T) {

	st := AutoCompile{
		TestPath: "./sample/...",
	}

	err := st.RunTest()
	assert.NoError(t, err)
}
