//go:build unit

package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTask(t *testing.T) {
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
		want    *TaskRepository
		wantErr bool
	}{
		{
			name: "Expect error when initializing without db",
			args: args{
				db: nil,
			},
			want:    &TaskRepository{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				db: sqlxDB,
			},
			want:    &TaskRepository{db: sqlxDB},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
