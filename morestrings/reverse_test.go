package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{s: "out"}, want: "tuo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRunes(tt.args.s); got != tt.want {
				t.Errorf("ReverseRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
