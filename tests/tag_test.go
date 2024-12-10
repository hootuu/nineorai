package tests

import (
	"nineorai/domains"
	"testing"
)

func TestTag_All(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		tag  domains.Tag
		args args
		want bool
	}{
		{
			name: "Add-ONE",
			tag:  domains.NewTag(),
			args: args{"ONE"},
			want: true,
		}, {
			name: "RE-Add-ONE",
			tag:  domains.NewTag(),
			args: args{"ONE"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tag.Append(tt.args.t)
			if got := tt.tag.Exists(tt.args.t); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestTag_Remove(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		tag  domains.Tag
		args args
	}{
		{
			name: "RV-ONE",
			tag:  domains.NewTag("ONE", "TWO"),
			args: args{
				"ONE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tag.Remove(tt.args.t)
			if got := tt.tag.Exists(tt.args.t); got != false {
				t.Errorf("Remove Failed() = %v, want %v", got, false)
			}
		})
	}
}
