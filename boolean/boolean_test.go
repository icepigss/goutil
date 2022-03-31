package boolean

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToBool(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.True(t, ToBool(1, 0))
		assert.False(t, ToBool(0, 0))
	})
	t.Run("float", func(t *testing.T) {
		assert.True(t, ToBool(1.0, 0.0))
		assert.False(t, ToBool(0.0, 0.0))
	})
	t.Run("string", func(t *testing.T) {
		assert.True(t, ToBool("a", ""))
		assert.False(t, ToBool("", ""))
	})
	t.Run("pointer", func(t *testing.T) {
		type st struct {
			a int
		}
		p := &st{1}
		var p1 *st
		assert.True(t, ToBool(p, nil))
		assert.False(t, ToBool(p1, nil))
	})
}

func TestToInt(t *testing.T) {
	a := true
	b := false
	assert.Equal(t, ToInt(a), 1)
	assert.Equal(t, ToInt(b), 0)
}
