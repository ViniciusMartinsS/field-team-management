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

func TestNewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	creator := domain.NewMockTaskCreator(ctrl)
	retriever := domain.NewMockTaskRetriever(ctrl)
	userRetriever := domain.NewMockUserRetriever(ctrl)

	type args struct {
		creator       domain.TaskCreator
		retriever     domain.TaskRetriever
		userRetriever domain.UserRetriever
	}

	tests := []struct {
		name    string
		args    args
		want    domain.TaskUsecase
		wantErr bool
	}{
		{
			name: "error - nil creator",
			args: args{
				creator:       nil,
				retriever:     retriever,
				userRetriever: userRetriever,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil retriever",
			args: args{
				creator:       creator,
				retriever:     nil,
				userRetriever: userRetriever,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil user retriever",
			args: args{
				creator:       creator,
				retriever:     retriever,
				userRetriever: nil,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				creator:       creator,
				retriever:     retriever,
				userRetriever: userRetriever,
			},
			want:    &taskUseCase{creator: creator, retriever: retriever, userRetriever: userRetriever},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.creator, tt.args.retriever, tt.args.userRetriever)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_ListByUserID(t *testing.T) {
	type dependencies struct {
		retriever     *domain.MockTaskRetriever
		userRetriever *domain.MockUserRetriever
	}

	var (
		id    = 1
		tasks = []domain.Task{
			{
				ID:      id,
				Summary: "task summary test",
				UserID:  id,
			},
		}
	)

	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            []domain.Task
		wantErr         bool
	}{
		{
			name: "error on ListByUserID when fetching user",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByUserID(context.Background(), id).Return(domain.User{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "error on ListByUserID when fetching tasks by manager",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByUserID(context.Background(), id).Return(domain.User{1, 1}, nil)
				d.retriever.EXPECT().List(context.Background()).Return([]domain.Task{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "error on ListByUserID when fetching tasks by technician",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByUserID(context.Background(), id).Return(domain.User{1, 2}, nil)
				d.retriever.EXPECT().ListByUserID(context.Background(), id).Return([]domain.Task{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByUserID(context.Background(), id).Return(domain.User{1, 2}, nil)
				d.retriever.EXPECT().ListByUserID(context.Background(), id).Return(tasks, nil)
			},
			want:    tasks,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			d := dependencies{
				retriever:     domain.NewMockTaskRetriever(ctrl),
				userRetriever: domain.NewMockUserRetriever(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				retriever:     d.retriever,
				userRetriever: d.userRetriever,
			}

			got, err := u.ListByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
