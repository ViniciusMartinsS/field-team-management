//go:build unit

package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	retriever := domain.NewMockUserRetriever(ctrl)
	authenticator := domain.NewMockAuthenticator(ctrl)

	type args struct {
		authenticator domain.Authenticator
		retriever     domain.UserRetriever
	}
	tests := []struct {
		name    string
		args    args
		want    domain.AuthUsecase
		wantErr bool
	}{
		{
			name: "Expect error when initializing without retriever",
			args: args{
				authenticator: authenticator,
				retriever:     nil,
			},
			want:    &authUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without authenticator",
			args: args{
				authenticator: nil,
				retriever:     retriever,
			},
			want:    &authUseCase{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				authenticator: authenticator,
				retriever:     retriever,
			},
			want: &authUseCase{
				authenticator: authenticator,
				retriever:     retriever,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuth(tt.args.authenticator, tt.args.retriever)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_authUseCase_Authenticate(t *testing.T) {
	type dependencies struct {
		authenticator *domain.MockAuthenticator
		retriever     *domain.MockUserRetriever
	}

	var (
		user = domain.User{
			ID:       1,
			Email:    "example@example.com",
			Password: "$2a$10$CmFPyUT88OJo30aUJ0Vtw.0gpj0Af50swAWGC3tLhFwAKUkz3a1yG",
			RoleID:   2,
		}
		token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	)

	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            string
		wantErr         bool
	}{
		{
			name: "Expect error thrown by ListByEmail",
			args: args{
				ctx:      context.Background(),
				email:    user.Email,
				password: user.Password,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByEmail(context.Background(), user.Email).Return(domain.User{}, errors.New("err"))
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Expect error thrown by CompareHashAndPassword",
			args: args{
				ctx:      context.Background(),
				email:    user.Email,
				password: "invalid",
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByEmail(context.Background(), user.Email).Return(user, nil)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Expect error thrown by GenerateAccessToken",
			args: args{
				ctx:      context.Background(),
				email:    user.Email,
				password: "123456",
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByEmail(context.Background(), user.Email).Return(user, nil)
				d.authenticator.EXPECT().GenerateAccessToken(user).Return("", errors.New("err"))
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				ctx:      context.Background(),
				email:    user.Email,
				password: "123456",
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByEmail(context.Background(), user.Email).Return(user, nil)
				d.authenticator.EXPECT().GenerateAccessToken(user).Return(token, nil)
			},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			d := dependencies{
				authenticator: domain.NewMockAuthenticator(ctrl),
				retriever:     domain.NewMockUserRetriever(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &authUseCase{
				authenticator: d.authenticator,
				retriever:     d.retriever,
			}

			got, err := u.Authenticate(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
