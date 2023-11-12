package encryption

import (
	"github.com/firdasafridi/gocrypt"
	"github.com/stretchr/testify/assert"
	"log"
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

func Test_encryptor_Encrypt(t *testing.T) {
	des, err := gocrypt.NewDESOpt("123456789123456789123456")
	if err != nil {
		log.Fatal(err)
	}

	type dependencies struct {
		des *gocrypt.DESOpt
	}
	type args struct {
		value string
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
				des: des,
			},
			args: args{
				value: "This is a encryption test",
			},
			want:    "juU8nhDbtBgrJ_Ief8qligIwXymOANa1hy3uyLMEjxM=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encryptor{
				des: tt.dependencies.des,
			}
			got, err := e.Encrypt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_encryptor_Decrypt(t *testing.T) {
	des, err := gocrypt.NewDESOpt("123456789123456789123456")
	if err != nil {
		log.Fatal(err)
	}

	type dependencies struct {
		des *gocrypt.DESOpt
	}
	type args struct {
		value string
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
				des: des,
			},
			args: args{
				value: "juU8nhDbtBgrJ_Ief8qligIwXymOANa1hy3uyLMEjxM=",
			},
			want:    "This is a encryption test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encryptor{
				des: tt.dependencies.des,
			}
			got, err := e.Decrypt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
