package e2e

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyMerch(t *testing.T) {
	itemName := "test-item"
	req, _ := http.NewRequest("GET", Server.URL+"/api/buy/"+itemName, nil)
	req.Header.Set("Authorization", "Bearer "+Token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var buyResp struct {
		Message string `json:"message"`
		Item    string `json:"item"`
		Price   int    `json:"price"`
		Balance int    `json:"balance"`
	}
	json.NewDecoder(resp.Body).Decode(&buyResp)

	assert.Equal(t, "Purchase successful", buyResp.Message)
	assert.Equal(t, itemName, buyResp.Item)
	assert.GreaterOrEqual(t, buyResp.Balance, 0, "Balance should not be negative")
}
