package jwt

import (
	"fmt"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testKey = "my_test_key"

func TestNew(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name    string
		args    args
		want    domain.Authenticator
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				key: "",
			},
			want:    &authenticator{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				key: testKey,
			},
			want: &authenticator{
				key: testKey,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_authenticator_GenerateAccessToken(t *testing.T) {
	type dependencies struct {
		key string
	}

	type args struct {
		user domain.User
	}

	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		want         string
		wantErr      bool
	}{
		{
			name: "happy",
			dependencies: dependencies{
				key: testKey,
			},
			args: args{
				domain.User{
					ID:     1,
					Email:  "test@test.com",
					RoleID: 1,
				},
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJyb2xlX2lkIjoxLCJ1c2VyX2lkIjoxfQ.styQ3v1Bw-Obn3Vg51PtSfOZgbuYrWsdM5nLzZTnlyg",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authenticator{
				key: tt.dependencies.key,
			}
			got, err := a.GenerateAccessToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_authenticator_IsAccessTokenValid(t *testing.T) {
	type dependencies struct {
		key string
	}

	type args struct {
		token string
	}

	type want struct {
		want   bool
		claims map[string]any
	}

	tests := []struct {
		name         string
		dependencies dependencies
		args         args
		want         want
		wantErr      bool
	}{
		{
			name: "error invalid token",
			dependencies: dependencies{
				key: testKey,
			},
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			},
			wantErr: true,
		},
		{
			name: "happy",
			dependencies: dependencies{
				key: testKey,
			},
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJyb2xlX2lkIjoxLCJ1c2VyX2lkIjoxfQ.styQ3v1Bw-Obn3Vg51PtSfOZgbuYrWsdM5nLzZTnlyg",
			},
			want: want{
				want: true,
				claims: map[string]any{
					"user_id": 1,
					"email":   "test@test.com",
					"role_id": 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authenticator{
				key: tt.dependencies.key,
			}

			got, claims, err := a.IsAccessTokenValid(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAccessTokenValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.want, got)
			assert.Equal(t, fmt.Sprint(tt.want.claims), fmt.Sprint(claims))
		})
	}
}
