package talentio

import (
	"testing"
	"time"

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

func TestAmbiguousTime(t *testing.T) {

	t.Run("Time", func(t *testing.T) {
		v := AmbiguousTime{false}
		assert.True(t, v.Time().IsZero())
		v = AmbiguousTime{"2016-03-23T11:28:41+09:00"}
		assert.False(t, v.Time().IsZero())
	})

	t.Run("Before", func(t *testing.T) {
		v := AmbiguousTime{false}
		assert.False(t, v.Before(time.Now()))
		v = AmbiguousTime{"2016-03-23T11:28:41+09:00"}
		assert.False(t, v.Before(time.Time{}))
	})

	t.Run("After", func(t *testing.T) {
		v := AmbiguousTime{false}
		assert.False(t, v.After(time.Now()))
		v = AmbiguousTime{"2016-03-23T11:28:41+09:00"}
		assert.True(t, v.After(time.Time{}))
	})
}
