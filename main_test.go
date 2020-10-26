package main

import "testing"

func Test_addOne(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{value: 1}, want: 2},
		{name: "t2", args: args{value: -1}, want: 0},
		{name: "t3", args: args{value: 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOne(tt.args.value); got != tt.want {
				t.Errorf("addOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
