package unit_test

import (
	"avito_shop/internal/services"
	"avito_shop/pkg"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	tests := []struct {
		name     string
		username string
		password string
		mockFn   func()
		wantErr  bool
	}{
		{
			name:     "NewUserSuccess",
			username: "newuser",
			password: "password123",
			mockFn: func() {
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("newuser").
					WillReturnError(sql.ErrNoRows)

				mock.ExpectQuery("INSERT INTO users").
					WithArgs("newuser", sqlmock.AnyArg(), int64(1000)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(1, "newuser", "hashedpass", 1000))
			},
			wantErr: false,
		},
		{
			name:     "ExistingUserSuccess",
			username: "existinguser",
			password: "password123",
			mockFn: func() {
				hashedPass, _ := pkg.HashPassword("password123")
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("existinguser").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(1, "existinguser", hashedPass, 1000))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			_, err := services.AuthService(db, "secret", tt.username, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
