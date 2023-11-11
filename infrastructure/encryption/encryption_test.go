package encryption

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				key: "",
			},
			wantErr: true,
		},
		{
			name: "error invalid key",
			args: args{
				key: "invalid",
			},
			wantErr: true,
		},
		{
			name: "happy",
			args: args{
				key: "123456789123456789123456",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
