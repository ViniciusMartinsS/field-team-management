package api

import (
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	authenticator := domain.NewMockAuthenticator(ctrl)
	taskUsecase := domain.NewMockTaskUsecase(ctrl)
	r := gin.Default()

	type args struct {
		r             *gin.Engine
		authenticator domain.Authenticator
		taskUsecase   domain.TaskUsecase
	}
	tests := []struct {
		name    string
		args    args
		want    *TaskAPIHandler
		wantErr bool
	}{
		{
			name: "Expect error when initializing without router",
			args: args{
				r:             nil,
				authenticator: authenticator,
				taskUsecase:   taskUsecase,
			},
			want:    &TaskAPIHandler{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without authenticator",
			args: args{
				r:             r,
				authenticator: nil,
				taskUsecase:   taskUsecase,
			},
			want:    &TaskAPIHandler{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without taskUsecase",
			args: args{
				r:             r,
				authenticator: authenticator,
				taskUsecase:   nil,
			},
			want:    &TaskAPIHandler{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				r:             r,
				authenticator: authenticator,
				taskUsecase:   taskUsecase,
			},
			want: &TaskAPIHandler{
				router:        r,
				authenticator: authenticator,
				taskUsecase:   taskUsecase,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.r, tt.args.authenticator, tt.args.taskUsecase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
