package unit_test

import (
	"avito_shop/internal/services"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSendCoinService(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	tests := []struct {
		name       string
		fromUserId uint
		toUser     string
		amount     int
		mockFn     func()
		wantErr    bool
	}{
		{
			name:       "SuccessfulTransfer",
			fromUserId: 1,
			toUser:     "recipient",
			amount:     100,
			mockFn: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("UPDATE users").
					WithArgs(100, "recipient").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(2, "recipient", "hashedpass", 1100))

				mock.ExpectQuery("UPDATE users").
					WithArgs(-100, uint(1)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(1, "sender", "hashedpass", 900))

				mock.ExpectQuery("INSERT INTO coin_transfers").
					WithArgs(uint(1), uint(2), 100).
					WillReturnRows(sqlmock.NewRows([]string{"id", "from_user_id", "to_user_id", "amount"}).
						AddRow(1, 1, 2, 100))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:       "InsufficientFunds",
			fromUserId: 1,
			toUser:     "recipient",
			amount:     1000000,
			mockFn: func() {
				mock.ExpectBegin()

				mock.ExpectQuery("UPDATE users").
					WithArgs(1000000, "recipient").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "balance"}).
						AddRow(2, "recipient", "hashedpass", 1100000))

				mock.ExpectQuery("UPDATE users").
					WithArgs(-1000000, uint(1)).
					WillReturnError(sql.ErrNoRows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			err := services.SendCoinService(db, tt.fromUserId, tt.toUser, tt.amount)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
