package api

import (
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	authUsecase := domain.NewMockAuthUsecase(ctrl)
	r := gin.Default()

	type args struct {
		r           *gin.Engine
		authUsecase domain.AuthUsecase
	}
	tests := []struct {
		name    string
		args    args
		want    *AuthAPIHandler
		wantErr bool
	}{
		{
			name: "Expect error when initializing without router",
			args: args{
				r:           nil,
				authUsecase: authUsecase,
			},
			want:    &AuthAPIHandler{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without authUsecase",
			args: args{
				r:           r,
				authUsecase: nil,
			},
			want:    &AuthAPIHandler{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				r:           r,
				authUsecase: authUsecase,
			},
			want: &AuthAPIHandler{
				router:      r,
				authUsecase: authUsecase,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuth(tt.args.r, tt.args.authUsecase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
