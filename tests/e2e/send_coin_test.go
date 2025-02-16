package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendCoinBetweenUsers(t *testing.T) {
	senderPayload := `{"username": "sender", "password": "senderpass"}`
	reqSender, err := http.NewRequest("POST", Server.URL+"/api/auth", bytes.NewBuffer([]byte(senderPayload)))
	assert.NoError(t, err)
	reqSender.Header.Set("Content-Type", "application/json")

	respSender, err := http.DefaultClient.Do(reqSender)
	assert.NoError(t, err)
	defer respSender.Body.Close()
	assert.Equal(t, http.StatusOK, respSender.StatusCode)

	var senderResp struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(respSender.Body).Decode(&senderResp)
	assert.NoError(t, err)
	assert.NotEmpty(t, senderResp.Token)

	recipientPayload := `{"username": "recipient", "password": "recipientpass"}`
	reqRecipient, err := http.NewRequest("POST", Server.URL+"/api/auth", bytes.NewBuffer([]byte(recipientPayload)))
	assert.NoError(t, err)
	reqRecipient.Header.Set("Content-Type", "application/json")

	respRecipient, err := http.DefaultClient.Do(reqRecipient)
	assert.NoError(t, err)
	defer respRecipient.Body.Close()
	assert.Equal(t, http.StatusOK, respRecipient.StatusCode)

	var recipientResp struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(respRecipient.Body).Decode(&recipientResp)
	assert.NoError(t, err)
	assert.NotEmpty(t, recipientResp.Token)

	sendCoinPayload := `{"toUser": "recipient", "amount": 50}`
	reqSend, err := http.NewRequest("POST", Server.URL+"/api/sendCoin", bytes.NewBuffer([]byte(sendCoinPayload)))
	assert.NoError(t, err)
	reqSend.Header.Set("Content-Type", "application/json")
	reqSend.Header.Set("Authorization", "Bearer "+senderResp.Token)

	respSend, err := http.DefaultClient.Do(reqSend)
	assert.NoError(t, err)
	defer respSend.Body.Close()
	assert.Equal(t, http.StatusOK, respSend.StatusCode)

	var sendCoinResp struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(respSend.Body).Decode(&sendCoinResp)
	assert.NoError(t, err)
	assert.Equal(t, "Coin transfer successful", sendCoinResp.Message)
}
