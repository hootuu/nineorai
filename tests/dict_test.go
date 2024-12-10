package tests

import (
	"nineorai/domains"
	"testing"
)

func TestDict_All(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string-string",
			args: args{
				key:   "a",
				value: "b",
			},
			want: true,
		},
		{
			name: "bool-bool",
			args: args{
				key:   "bool",
				value: "true",
			},
			want: true,
		},
		{
			name: "int-int",
			args: args{
				key:   "int",
				value: "1234",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict, _ := domains.NewDict()
			err := dict.Set(tt.args.key, tt.args.value)
			if err != nil {
				t.Errorf("Set Failed: %v", err)
			}
			if got := dict.Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
			getStr, _ := dict.GetString(tt.args.key)
			if getStr != tt.args.value {
				t.Errorf("GetString() = %v, want %v", getStr, tt.args.value)
			}
		})
	}
}
