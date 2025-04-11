package messagebroker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBrokerMessage(t *testing.T) {

	result, err := NewBrokerMessage("test")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "\"test\"", string(result.GetPayload()))
}
