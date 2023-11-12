//go:build unit

package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/encryption"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	creator := domain.NewMockTaskCreator(ctrl)
	retriever := domain.NewMockTaskRetriever(ctrl)
	updater := domain.NewMockTaskUpdater(ctrl)
	remover := domain.NewMockTaskRemover(ctrl)
	userRetriever := domain.NewMockUserRetriever(ctrl)
	encryptor, _ := encryption.New("12345678912345678912345") // Change to mock

	type args struct {
		creator       domain.TaskCreator
		retriever     domain.TaskRetriever
		updater       domain.TaskUpdater
		remover       domain.TaskRemover
		userRetriever domain.UserRetriever
		encryptor     domain.SummaryEncryptor
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
				updater:       updater,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil retriever",
			args: args{
				creator:       creator,
				retriever:     nil,
				updater:       updater,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil user retriever",
			args: args{
				creator:       creator,
				retriever:     retriever,
				updater:       updater,
				remover:       remover,
				userRetriever: nil,
				encryptor:     encryptor,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil user updater",
			args: args{
				creator:       creator,
				retriever:     retriever,
				updater:       nil,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil user remover",
			args: args{
				creator:       creator,
				retriever:     retriever,
				updater:       updater,
				remover:       nil,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil user encryptor",
			args: args{
				creator:       creator,
				retriever:     retriever,
				updater:       updater,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     nil,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				creator:       creator,
				retriever:     retriever,
				updater:       updater,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			want: &taskUseCase{
				creator:       creator,
				retriever:     retriever,
				updater:       updater,
				remover:       remover,
				userRetriever: userRetriever,
				encryptor:     encryptor,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.creator, tt.args.retriever, tt.args.updater, tt.args.remover, tt.args.userRetriever, tt.args.encryptor)
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
			name: "error on ListByID when fetching user",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "error on ListByID when fetching tasks by manager",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{ID: 1, RoleID: 1}, nil)
				d.retriever.EXPECT().List(context.Background()).Return([]domain.Task{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "error on ListByID when fetching tasks by technician",
			args: args{
				ctx:    context.Background(),
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{ID: 1, RoleID: 2}, nil)
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
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{ID: 1, RoleID: 2}, nil)
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
				t.Errorf("ListByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_Add(t *testing.T) {
	type dependencies struct {
		creator   *domain.MockTaskCreator
		encryptor *domain.MockSummaryEncryptor
	}

	type args struct {
		ctx  context.Context
		task domain.Task
	}

	var task = domain.Task{
		ID:      1,
		Summary: "task summary test",
		UserID:  1,
	}

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            domain.Task
		wantErr         bool
	}{
		{
			name: "error on Encrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "error on Add",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.creator.EXPECT().Add(context.Background(), task).Return(task.ID, errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.creator.EXPECT().Add(context.Background(), task).Return(task.ID, nil)
				d.encryptor.EXPECT().Decrypt(task.Summary).Return(task.Summary, nil)
			},
			want:    task,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			d := dependencies{
				encryptor: domain.NewMockSummaryEncryptor(ctrl),
				creator:   domain.NewMockTaskCreator(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				creator:   d.creator,
				encryptor: d.encryptor,
			}

			got, err := u.Add(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_Update(t *testing.T) {
	type dependencies struct {
		retriever *domain.MockTaskRetriever
		updater   *domain.MockTaskUpdater
		encryptor *domain.MockSummaryEncryptor
	}

	type args struct {
		ctx  context.Context
		task domain.Task
	}

	var task = domain.Task{
		ID:      1,
		Summary: "task summary test",
		UserID:  1,
	}

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            domain.Task
		wantErr         bool
	}{
		{
			name: "error on Update without ID",
			args: args{
				ctx: context.Background(),
				task: domain.Task{
					Summary: task.Summary,
					UserID:  task.UserID,
				},
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "error on Update without User ID",
			args: args{
				ctx: context.Background(),
				task: domain.Task{
					ID:      task.ID,
					Summary: task.Summary,
				},
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "error on ListByIDAndUserID",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(domain.Task{}, errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "error on Encrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(task, nil)
				d.encryptor.EXPECT().Encrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "error on Update",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(task, nil)
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.updater.EXPECT().Update(context.Background(), task).Return(errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "hapy",
			args: args{
				ctx:  context.Background(),
				task: task,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(task, nil)
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.updater.EXPECT().Update(context.Background(), task).Return(nil)
				d.encryptor.EXPECT().Decrypt(task.Summary).Return(task.Summary, nil)
			},
			want:    task,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			d := dependencies{
				retriever: domain.NewMockTaskRetriever(ctrl),
				updater:   domain.NewMockTaskUpdater(ctrl),
				encryptor: domain.NewMockSummaryEncryptor(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				retriever: d.retriever,
				updater:   d.updater,
				encryptor: d.encryptor,
			}

			got, err := u.Update(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_Remove(t *testing.T) {
	type dependencies struct {
		userRetriever *domain.MockUserRetriever
		remover       *domain.MockTaskRemover
	}

	type args struct {
		ctx    context.Context
		id     int
		userID int
	}

	const id = 1

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		wantErr         bool
	}{
		{
			name: "error on Remove without ID",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			wantErr: true,
		},
		{
			name: "error on Remove without UserID",
			args: args{
				ctx: context.Background(),
				id:  id,
			},
			wantErr: true,
		},
		{
			name: "error on ListByID",
			args: args{
				ctx:    context.Background(),
				id:     id,
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{}, errors.New("err"))
			},
			wantErr: true,
		},
		{
			name: "error user forbidden",
			args: args{
				ctx:    context.Background(),
				id:     id,
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{ID: id, RoleID: 2}, nil)
			},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				ctx:    context.Background(),
				id:     id,
				userID: id,
			},
			setDependencies: func(d *dependencies) {
				d.userRetriever.EXPECT().ListByID(context.Background(), id).Return(domain.User{ID: id, RoleID: 1}, nil)
				d.remover.EXPECT().Remove(context.Background(), id).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			d := dependencies{
				userRetriever: domain.NewMockUserRetriever(ctrl),
				remover:       domain.NewMockTaskRemover(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				userRetriever: d.userRetriever,
				remover:       d.remover,
			}

			err := u.Remove(tt.args.ctx, tt.args.id, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
