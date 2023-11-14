package notifier

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	channel := amqp.Channel{}
	queue := "test"

	type args struct {
		queueName  string
		connection *amqp.Channel
	}
	tests := []struct {
		name    string
		args    args
		want    *Notifier
		wantErr bool
	}{
		{
			name: "Expect error when initializing without queueName",
			args: args{
				queueName:  "",
				connection: &channel,
			},
			want:    &Notifier{},
			wantErr: true,
		},
		{
			name: "Expect error when initializing without connection",
			args: args{
				queueName:  queue,
				connection: nil,
			},
			want:    &Notifier{},
			wantErr: true,
		},
		{
			name: "Expect success",
			args: args{
				queueName:  queue,
				connection: &channel,
			},
			want: &Notifier{
				queueName:  queue,
				connection: &channel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNotifier(tt.args.queueName, tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNotifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
