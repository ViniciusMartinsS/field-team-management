package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetRole(t *testing.T) {
	type fields struct {
		ID     int64
		RoleID int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Expect user to be technician",
			fields: fields{
				ID:     1,
				RoleID: 2,
			},
			want: Technician,
		},
		{
			name: "Expect user to be manager",
			fields: fields{
				ID:     1,
				RoleID: 1,
			},
			want: Manager,
		},
		{
			name: "Expect not find user type",
			fields: fields{
				ID:     1,
				RoleID: 3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:     tt.fields.ID,
				RoleID: tt.fields.RoleID,
			}
			assert.Equalf(t, tt.want, u.GetRole(), "GetRole()")
		})
	}
}
