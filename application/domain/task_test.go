//go:build unit

package domain

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	type args struct {
		summary string
		date    *time.Time
		userID  int64
	}

	currentTime := time.Now()

	const (
		summary = "this is a summary"
		userID  = 1
	)

	tests := []struct {
		name    string
		args    args
		want    Task
		wantErr bool
	}{
		{
			name: "error when empty summary",
			args: args{
				summary: "",
				date:    &currentTime,
				userID:  userID,
			},
			want:    Task{},
			wantErr: true,
		},
		{
			name: "error when summary is too big",
			args: args{
				summary: strings.Repeat("a", 2600),
				date:    &currentTime,
				userID:  userID,
			},
			want:    Task{},
			wantErr: true,
		},
		{
			name: "error when no user",
			args: args{
				summary: summary,
				date:    &currentTime,
				userID:  0,
			},
			want:    Task{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				summary: summary,
				date:    &currentTime,
				userID:  userID,
			},
			want: Task{
				Summary: summary,
				Date:    &currentTime,
				UserID:  userID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.summary, tt.args.date, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
