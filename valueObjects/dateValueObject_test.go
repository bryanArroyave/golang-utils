package valueobjects

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateValueObject(t *testing.T) {

	t.Run("invalid min", func(t *testing.T) {
		j := NewDateValueObject(time.Date(2023, time.March, 23, 0, 0, 0, 0, time.UTC)).Min(time.Date(2024, time.March, 23, 0, 0, 0, 0, time.UTC)).Max(time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC))

		val, err := j.Value()
		assert.Equal(t, val, time.Time{})
		assert.Error(t, err)
		assert.False(t, j.isValid())
	})

	t.Run("invalid max", func(t *testing.T) {

		j := NewDateValueObject(time.Date(2026, time.March, 23, 0, 0, 0, 0, time.UTC)).Min(time.Date(2024, time.March, 23, 0, 0, 0, 0, time.UTC)).Max(time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC))
		val, err := j.Value()

		assert.False(t, j.isValid())
		assert.Error(t, err)
		assert.Equal(t, val, time.Time{})
	})

	t.Run("valid", func(t *testing.T) {

		j := NewDateValueObject(time.Date(2024, time.March, 25, 0, 0, 0, 0, time.UTC)).Min(time.Date(2024, time.March, 23, 0, 0, 0, 0, time.UTC)).Max(time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC))
		val, err := j.Value()

		expected := time.Date(2024, time.March, 25, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, val, expected)
		assert.Empty(t, err)
	})

	t.Run("no validations valid min", func(t *testing.T) {

		j := NewDateValueObject(time.Date(2023, time.March, 23, 0, 0, 0, 0, time.UTC))
		val, err := j.Value()

		expected := time.Date(2023, time.March, 23, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, val, expected)
		assert.Empty(t, err)

	})

	t.Run("no validations valid max", func(t *testing.T) {

		j := NewDateValueObject(time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC))
		expected := time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC)

		val, err := j.Value()

		assert.Equal(t, val, expected)
		assert.Empty(t, err)

	})

	t.Run("optional value", func(t *testing.T) {

		j := NewDateValueObject(time.Time{}).Min(time.Date(2024, time.March, 23, 0, 0, 0, 0, time.UTC)).Max(time.Date(2025, time.March, 23, 0, 0, 0, 0, time.UTC)).Optional()
		val, err := j.Value()
		assert.True(t, j.isValid())
		assert.Nil(t, err)
		assert.Empty(t, val)

	})

}
