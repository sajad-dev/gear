package developer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGit_Changes(t *testing.T) {
	st := AutoCompile{
		ProjectPath: "./sample",
	}

	chenge, err := st.Chenges()
	assert.NoError(t, err)
	assert.True(t, chenge)
}
