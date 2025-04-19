package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntValueObject(t *testing.T) {

	t.Run("invalid min", func(t *testing.T) {
		j := NewIntValueObject("mock", 2).Min(5).Max(10)

		val, err := j.Value()
		assert.Equal(t, val, 0)
		assert.Error(t, err)
		assert.False(t, j.isValid())
	})

	t.Run("invalid max", func(t *testing.T) {

		j := NewIntValueObject("mock", 11).Min(5).Max(10)
		val, err := j.Value()

		assert.False(t, j.isValid())
		assert.Error(t, err)
		assert.Equal(t, val, 0)
	})

	t.Run("valid", func(t *testing.T) {

		j := NewIntValueObject("mock", 6).Min(5).Max(10)
		val, err := j.Value()

		expected := 6
		assert.Equal(t, val, expected)
		assert.Empty(t, err)
	})

	t.Run("no validations invalid min", func(t *testing.T) {

		j := NewIntValueObject("mock", 2)
		val, err := j.Value()

		expected := 2
		assert.Equal(t, val, expected)
		assert.Empty(t, err)

	})

	t.Run("no validations invalid max", func(t *testing.T) {

		j := NewIntValueObject("mock", 11)
		expected := 11

		val, err := j.Value()

		assert.Equal(t, val, expected)
		assert.Empty(t, err)

	})

	t.Run("optional value", func(t *testing.T) {

		j := NewIntValueObject("mock", 0).Min(5).Max(10).Optional()
		val, err := j.Value()
		assert.True(t, j.isValid())
		assert.Nil(t, err)
		assert.Equal(t, val, 0)

	})

}
