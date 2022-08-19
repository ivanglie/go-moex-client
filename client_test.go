package moex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate(USDRUB)
	assert.Nil(t, err.Error())
	assert.GreaterOrEqual(t, rate, float64(1))
}
