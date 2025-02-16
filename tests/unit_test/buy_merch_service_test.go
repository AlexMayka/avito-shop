package unit_test

import (
	"avito_shop/internal/services"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuyMerchService(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	tests := []struct {
		name     string
		userId   uint
		itemName string
		mockFn   func()
		wantErr  bool
	}{
		{
			name:     "SuccessfulPurchase",
			userId:   1,
			itemName: "testitem",
			mockFn: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("SELECT (.+) FROM merch").
					WithArgs("testitem").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
						AddRow(1, "testitem", 100))

				mock.ExpectQuery("UPDATE users").
					WithArgs(-100, uint(1)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(1, "testuser", "hashedpass", 900))

				mock.ExpectQuery("INSERT INTO purchases").
					WithArgs(uint(1), uint(1), 100).
					WillReturnRows(sqlmock.NewRows([]string{"id", "merch_id", "price_bought"}).
						AddRow(1, 1, 100))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:     "ItemNotFound",
			userId:   1,
			itemName: "nonexistent",
			mockFn: func() {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT (.+) FROM merch").
					WithArgs("nonexistent").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			_, _, err := services.BuyMerchService(db, tt.userId, tt.itemName)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
