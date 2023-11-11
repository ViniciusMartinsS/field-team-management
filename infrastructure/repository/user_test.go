package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *UserRepository
		wantErr bool
	}{
		{
			name: "error - nil db",
			args: args{
				db: nil,
			},
			want:    &UserRepository{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				db: sqlxDB,
			},
			want:    &UserRepository{db: sqlxDB},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
