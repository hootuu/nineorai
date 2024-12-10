package tests

import (
	"fmt"
	"github.com/hootuu/nineorai/domains"
	"testing"
)

func TestCtrl_All(t *testing.T) {
	type args struct {
		pos int
		val bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "0->T",
			args: args{
				pos: 0,
				val: true,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "112->T",
			args: args{
				pos: 112,
				val: true,
			},
			want:    true,
			wantErr: false,
		},

		{
			name: "1120->T",
			args: args{
				pos: 1120,
				val: true,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := domains.NewCtrl(128)
			_ = c.Set(tt.args.pos, tt.args.val)
			got, err := c.Check(tt.args.pos)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
			c.Iter(func(pos int, value bool) {
				fmt.Println("[", tt.name, "]: ", pos+1, "=>", value)
			})
		})
	}
}
