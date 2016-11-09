package talentio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {

	t.Run("Bool", func(t *testing.T) {
		t.Parallel()
		v := true
		assert.Equal(t, v, *Bool(v))
	})

	t.Run("Int", func(t *testing.T) {
		t.Parallel()
		v := 1
		assert.Equal(t, v, *Int(v))
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		v := "foo"
		assert.Equal(t, v, *String(v))
	})
}
