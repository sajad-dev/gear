package developer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange_ChangeHandel(t *testing.T) {
	st := AutoCompile{
		ProjectPath: "./sample",
		BuildPath:   "./sample/build",
	}

	chenge, err := st.HandelChange()
	assert.NoError(t, err)
	assert.True(t, chenge)
}
