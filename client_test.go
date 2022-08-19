package moex

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	if statusCode, err := func() (statusCode int, err error) {
		req, err := http.NewRequest("GET", baseURL+"/index.xml", nil)
		if err != nil {
			return
		}

		client := http.Client{Timeout: 3 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		statusCode = resp.StatusCode

		return
	}(); err != nil || statusCode != http.StatusOK {
		t.Log(statusCode, err)
		t.SkipNow()
	}

	client := NewClient()
	rate, err := client.GetRate(USDRUB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, rate, float64(1))
}
