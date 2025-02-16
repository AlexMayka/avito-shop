package unit_test

import (
	"avito_shop/internal/models"
	"avito_shop/internal/repositories"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	q := &MockQuerier{db: db, mock: mock}

	tests := []struct {
		name     string
		username string
		mockFn   func()
		want     *models.User
		wantErr  bool
	}{
		{
			name:     "Success",
			username: "testuser",
			mockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
					AddRow(1, "testuser", "hashedpass", 1000)
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("testuser").
					WillReturnRows(rows)
			},
			want: &models.User{
				ID:       1,
				Username: "testuser",
				Password: "hashedpass",
				Balance:  1000,
			},
			wantErr: false,
		},
		{
			name:     "UserNotFound",
			username: "nonexistent",
			mockFn: func() {
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("nonexistent").
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			got, err := repositories.GetUserByUsername(q, tt.username)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
