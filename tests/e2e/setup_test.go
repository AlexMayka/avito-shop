package e2e

import (
	"avito_shop/internal/routes"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var Server *httptest.Server
var DB *sql.DB
var Token string

func TestMain(m *testing.M) {
	cfg := LoadConfig()

	dbConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.Db)

	var err error
	DB, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to test DB: %s", err)
	}

	createTestTables(DB)

	r := routes.SetupRouter(DB, cfg)
	Server = httptest.NewServer(r)

	Token = getAuthToken()

	code := m.Run()

	cleanupTestTables(DB)

	Server.Close()
	DB.Close()

	os.Exit(code)
}

func getAuthToken() string {
	authPayload := `{"username": "testuser", "password": "testpassword"}`
	req, _ := http.NewRequest("POST", Server.URL+"/api/auth", bytes.NewBuffer([]byte(authPayload)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error during authentication request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Auth failed: status=%d", resp.StatusCode)
	}

	var authResp struct {
		Token string `json:"token"`
	}
	json.NewDecoder(resp.Body).Decode(&authResp)

	return authResp.Token
}
