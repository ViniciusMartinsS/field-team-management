//go:build unit

package usecase

import (
	"github.com/ViniciusMartinss/field-team-management/application/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	creator := domain.NewMockTaskCreator(ctrl)
	retriever := domain.NewMockTaskRetriever(ctrl)

	type args struct {
		creator   domain.TaskCreator
		retriever domain.TaskRetriever
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
				creator:   nil,
				retriever: retriever,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "error - nil retriever",
			args: args{
				creator:   creator,
				retriever: nil,
			},
			want:    &taskUseCase{},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				creator:   creator,
				retriever: retriever,
			},
			want:    &taskUseCase{creator: creator, retriever: retriever},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.creator, tt.args.retriever)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
