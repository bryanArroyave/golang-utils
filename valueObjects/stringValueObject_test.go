package valueobjects

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValueObject(t *testing.T) {

	t.Run("invalid max length", func(t *testing.T) {
		j := NewStringValueObject("mock", "JourneyName").MinLength(5).MaxLength(10)
		val, err := j.Value()

		assert.Equal(t, val, "")
		assert.False(t, j.isValid())
		assert.Error(t, err)
	})

	t.Run("invalid min length", func(t *testing.T) {

		j := NewStringValueObject("mock", "J").MinLength(5).MaxLength(10)
		val, err := j.Value()
		assert.Equal(t, val, "")
		assert.False(t, j.isValid())
		assert.Error(t, err)

	})

	t.Run("invalid Regex", func(t *testing.T) {
		j := NewStringValueObject("mock", "123").Pattern("^[a-zA-Z]+$")
		val, err := j.Value()
		assert.False(t, j.isValid())
		assert.Error(t, err)
		assert.Equal(t, val, "")
	})

	t.Run("valid", func(t *testing.T) {

		j := NewStringValueObject("mock", "Journey").MinLength(5).MaxLength(10).Pattern("^[a-zA-Z]+$")
		val, err := j.Value()

		expected := "Journey"
		assert.Equal(t, val, expected)
		assert.Empty(t, err)
	})

	t.Run("no validations invalid min length", func(t *testing.T) {

		j := NewStringValueObject("mock", "JourneyName")
		val, err := j.Value()
		expected := "JourneyName"
		assert.Equal(t, val, expected)
		assert.Empty(t, err)
	})

	t.Run("no validations invalid max length", func(t *testing.T) {

		j := NewStringValueObject("mock", "J")
		val, err := j.Value()
		expected := "J"

		assert.Equal(t, val, expected)
		assert.Empty(t, err)
	})

	t.Run("optional value", func(t *testing.T) {

		j := NewStringValueObject("mock", "").MinLength(5).MaxLength(10).Optional()
		val, err := j.Value()
		assert.True(t, j.isValid())
		assert.Nil(t, err)
		assert.Equal(t, val, "")

	})

	t.Run("include values", func(t *testing.T) {
		j := NewStringValueObject("mock", "Journey").Include("Journey", "JourneyName")
		val, err := j.Value()
		assert.True(t, j.isValid())
		assert.Nil(t, err)
		assert.Equal(t, val, "Journey")
	})

	t.Run("invalid include values", func(t *testing.T) {
		j := NewStringValueObject("mock", "Journey").Include("JourneyName", "Prueba")
		val, err := j.Value()
		assert.False(t, j.isValid())
		assert.Error(t, err)
		fmt.Println(err.Error())
		assert.Equal(t, val, "")
	})
}
