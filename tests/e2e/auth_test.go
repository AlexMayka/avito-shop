package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	authPayload := `{"username": "testuser", "password": "testpassword"}`
	req, _ := http.NewRequest("POST", Server.URL+"/api/auth", bytes.NewBuffer([]byte(authPayload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var authResp struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&authResp)

	assert.NotEmpty(t, authResp.Token)
}
