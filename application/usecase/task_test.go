//go:build unit

package usecase

import (
	"context"
	"errors"
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	creator := domain.NewMockTaskCreator(ctrl)
	retriever := domain.NewMockTaskRetriever(ctrl)
	updater := domain.NewMockTaskUpdater(ctrl)
	remover := domain.NewMockTaskRemover(ctrl)
	encryptor := domain.NewMockSummaryEncryptor(ctrl)
	notifier := domain.NewMockTaskNotifier(ctrl)

	type args struct {
		creator   domain.TaskCreator
		retriever domain.TaskRetriever
		updater   domain.TaskUpdater
		remover   domain.TaskRemover
		encryptor domain.SummaryEncryptor
		notifier  domain.TaskNotifier
	}

	tests := []struct {
		name    string
		args    args
		want    domain.TaskUsecase
		wantErr bool
	}{
		{
			name: "Expect error when initializing without creator",
			args: args{
				creator:   nil,
				retriever: retriever,
				updater:   updater,
				remover:   remover,
				encryptor: encryptor,
				notifier:  notifier,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without retriever",
			args: args{
				creator:   creator,
				retriever: nil,
				updater:   updater,
				remover:   remover,
				encryptor: encryptor,
				notifier:  notifier,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without user updater",
			args: args{
				creator:   creator,
				retriever: retriever,
				updater:   nil,
				remover:   remover,
				encryptor: encryptor,
				notifier:  notifier,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without user remover",
			args: args{
				creator:   creator,
				retriever: retriever,
				updater:   updater,
				remover:   nil,
				encryptor: encryptor,
				notifier:  notifier,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without user encryptor",
			args: args{
				creator:   creator,
				retriever: retriever,
				updater:   updater,
				remover:   remover,
				encryptor: nil,
				notifier:  notifier,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without user notifier",
			args: args{
				creator:   creator,
				retriever: retriever,
				updater:   updater,
				remover:   remover,
				encryptor: encryptor,
				notifier:  nil,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				creator:   creator,
				retriever: retriever,
				updater:   updater,
				remover:   remover,
				encryptor: encryptor,
				notifier:  notifier,
			},
			want: &taskUseCase{
				creator:   creator,
				retriever: retriever,
				updater:   updater,
				remover:   remover,
				encryptor: encryptor,
				notifier:  notifier,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.creator, tt.args.retriever, tt.args.updater, tt.args.remover, tt.args.encryptor, tt.args.notifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_ListByUser(t *testing.T) {
	type dependencies struct {
		retriever *domain.MockTaskRetriever
		encryptor *domain.MockSummaryEncryptor
	}

	var (
		tasks = []domain.Task{
			{
				ID:      1,
				Summary: "task summary test",
				UserID:  1,
			},
		}
		managerUser = domain.User{
			ID:     1,
			RoleID: 1,
		}
		technicalUser = domain.User{
			ID:     2,
			RoleID: 2,
		}
	)

	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            []domain.Task
		wantErr         bool
	}{
		{
			name: "Expect error when user ID is missing",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					RoleID: technicalUser.RoleID,
				},
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error when user role ID is missing",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					ID: technicalUser.ID,
				},
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by List when listing tasks for manager",
			args: args{
				ctx:  context.Background(),
				user: managerUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().List(context.Background()).Return([]domain.Task{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by ListByUserID when listing tasks for technician",
			args: args{
				ctx:  context.Background(),
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByUserID(context.Background(), technicalUser.ID).Return([]domain.Task{}, errors.New("err"))
			},
			want:    []domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect success when listing tasks for manager",
			args: args{
				ctx:  context.Background(),
				user: managerUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().List(context.Background()).Return(tasks, nil)
				d.encryptor.EXPECT().Decrypt(tasks[0].Summary).Return(tasks[0].Summary, nil)
			},
			want:    tasks,
			wantErr: false,
		},
		{
			name: "Expect success when listing tasks for technician",
			args: args{
				ctx:  context.Background(),
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByUserID(context.Background(), technicalUser.ID).Return(tasks, nil)
				d.encryptor.EXPECT().Decrypt(tasks[0].Summary).Return(tasks[0].Summary, nil)
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
				retriever: domain.NewMockTaskRetriever(ctrl),
				encryptor: domain.NewMockSummaryEncryptor(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				retriever: d.retriever,
				encryptor: d.encryptor,
			}

			got, err := u.ListByUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByUser() error = %v, wantErr %v", err, tt.wantErr)
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
		user domain.User
	}

	var (
		task = domain.Task{
			ID:      1,
			Summary: "task summary test",
			UserID:  2,
		}
		technicalUser = domain.User{
			ID:     2,
			RoleID: 2,
		}
		technicalUser2 = domain.User{
			ID:     3,
			RoleID: 2,
		}
	)

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            domain.Task
		wantErr         bool
	}{
		{
			name: "Expect error when task user ID is missing",
			args: args{
				ctx: context.Background(),
				task: domain.Task{
					Summary: task.Summary,
				},
				user: technicalUser,
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error when user ID is missing",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: domain.User{
					RoleID: technicalUser.RoleID,
				},
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error when task user role ID is missing",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: domain.User{
					ID: technicalUser.ID,
				},
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect forbidden error when user is tries to add a task for somebody else",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser2,
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by Encrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by Add",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.creator.EXPECT().Add(context.Background(), task).Return(task.ID, errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by Decrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.creator.EXPECT().Add(context.Background(), task).Return(task.ID, nil)
				d.encryptor.EXPECT().Decrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
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

			got, err := u.Add(tt.args.ctx, tt.args.task, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_Add_Notify_Concurrent(t *testing.T) {
	var (
		ctx  = context.Background()
		date = time.Now()
		task = domain.Task{
			ID:      1,
			Summary: "task summary test",
			Date:    &date,
			UserID:  2,
		}
		user = domain.User{
			ID:     2,
			RoleID: 2,
		}
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	encryptor := domain.NewMockSummaryEncryptor(ctrl)
	encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
	encryptor.EXPECT().Decrypt(task.Summary).Return(task.Summary, nil)

	creator := domain.NewMockTaskCreator(ctrl)
	creator.EXPECT().Add(ctx, task).Return(task.ID, nil)

	var wg sync.WaitGroup
	wg.Add(1)

	notifier := domain.NewMockTaskNotifier(ctrl)
	notifier.EXPECT().SendNotification(ctx, gomock.Any()).
		Do(func(arg0, arg1 interface{}) interface{} {
			defer wg.Done()
			return nil
		})

	u := &taskUseCase{
		creator:   creator,
		encryptor: encryptor,
		notifier:  notifier,
	}
	got, err := u.Add(ctx, task, user)

	wg.Wait()
	assert.Equal(t, nil, err)
	assert.Equal(t, task, got)
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
		user domain.User
	}

	var (
		task = domain.Task{
			ID:      1,
			Summary: "task summary test",
			UserID:  1,
		}
		technicalUser = domain.User{
			ID:     1,
			RoleID: 2,
		}
	)

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		want            domain.Task
		wantErr         bool
	}{
		{
			name: "Expect error when task ID is missing",
			args: args{
				ctx: context.Background(),
				task: domain.Task{
					Summary: task.Summary,
					UserID:  task.UserID,
				},
				user: technicalUser,
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error when user ID is missing",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: domain.User{
					RoleID: technicalUser.RoleID,
				},
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by ListByIDAndUserID",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(domain.Task{}, errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by Encrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(task, nil)
				d.encryptor.EXPECT().Encrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect error thrown by Update",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
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
			name: "Expect error thrown by Decrypt",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
			},
			setDependencies: func(d *dependencies) {
				d.retriever.EXPECT().ListByIDAndUserID(context.Background(), task.ID, task.UserID).Return(task, nil)
				d.encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
				d.updater.EXPECT().Update(context.Background(), task).Return(nil)
				d.encryptor.EXPECT().Decrypt(task.Summary).Return("", errors.New("err"))
			},
			want:    domain.Task{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				ctx:  context.Background(),
				task: task,
				user: technicalUser,
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

			got, err := u.Update(tt.args.ctx, tt.args.task, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_taskUseCase_Update_Notify_Concurrent(t *testing.T) {
	var (
		ctx  = context.Background()
		date = time.Now()
		task = domain.Task{
			ID:      1,
			Summary: "task summary test",
			Date:    &date,
			UserID:  2,
		}
		user = domain.User{
			ID:     2,
			RoleID: 2,
		}
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	encryptor := domain.NewMockSummaryEncryptor(ctrl)
	encryptor.EXPECT().Encrypt(task.Summary).Return(task.Summary, nil)
	encryptor.EXPECT().Decrypt(task.Summary).Return(task.Summary, nil)

	updater := domain.NewMockTaskUpdater(ctrl)
	updater.EXPECT().Update(ctx, task).Return(nil)

	retriever := domain.NewMockTaskRetriever(ctrl)
	retriever.EXPECT().ListByIDAndUserID(ctx, task.ID, task.UserID).Return(task, nil)

	var wg sync.WaitGroup
	wg.Add(1)

	notifier := domain.NewMockTaskNotifier(ctrl)
	notifier.EXPECT().SendNotification(ctx, gomock.Any()).
		Do(func(arg0, arg1 interface{}) interface{} {
			defer wg.Done()
			return nil
		})

	u := &taskUseCase{
		retriever: retriever,
		updater:   updater,
		encryptor: encryptor,
		notifier:  notifier,
	}
	got, err := u.Update(ctx, task, user)

	wg.Wait()
	assert.Equal(t, nil, err)
	assert.Equal(t, task, got)
}

func Test_taskUseCase_Remove(t *testing.T) {
	type dependencies struct {
		remover *domain.MockTaskRemover
	}

	type args struct {
		ctx  context.Context
		id   int64
		user domain.User
	}

	var (
		id          = int64(1)
		managerUser = domain.User{
			ID:     1,
			RoleID: 1,
		}
		technicalUser = domain.User{
			ID:     2,
			RoleID: 2,
		}
	)

	tests := []struct {
		name            string
		args            args
		setDependencies func(d *dependencies)
		wantErr         bool
	}{
		{
			name: "Expect error when task ID is missing",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "Expect error when user ID is missing",
			args: args{
				ctx: context.Background(),
				id:  id,
				user: domain.User{
					RoleID: managerUser.RoleID,
				},
			},
			wantErr: true,
		},
		{
			name: "Expect error when task user role ID is missing",
			args: args{
				ctx: context.Background(),
				id:  id,
				user: domain.User{
					ID: managerUser.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "Expect forbidden error when user technician tries to remove a task",
			args: args{
				ctx:  context.Background(),
				id:   id,
				user: technicalUser,
			},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				ctx:  context.Background(),
				id:   id,
				user: managerUser,
			},
			setDependencies: func(d *dependencies) {
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
				remover: domain.NewMockTaskRemover(ctrl),
			}

			if tt.setDependencies != nil {
				tt.setDependencies(&d)
			}

			u := &taskUseCase{
				remover: d.remover,
			}

			err := u.Remove(tt.args.ctx, tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
