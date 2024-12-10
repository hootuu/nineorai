package tests

import (
	"github.com/hootuu/nineorai/domains"
	"testing"
)

func TestMeta_All(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		wantValue int32
	}{
		{
			name: "int-int",
			args: args{
				key:   "int",
				value: "1234",
			},
			want:      true,
			wantValue: 1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meta, _ := domains.NewMeta()
			err := meta.Set(tt.args.key, tt.args.value)
			if err != nil {
				t.Errorf("Set Failed: %v", err)
			}
			if got := meta.Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
			getInt, _ := meta.GetInt32(tt.args.key)
			if getInt != tt.wantValue {
				t.Errorf("GetInt32() = %v, want %v", getInt, tt.args.value)
			}
		})
	}
}
