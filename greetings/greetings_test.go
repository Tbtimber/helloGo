package greetings

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "t1", args: args{name: ""}, want: "", wantErr: true},
		{name: "t1", args: args{name: "Matthieu"}, want: "Hello : Matthieu", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomFormat(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomFormat(); got != tt.want {
				t.Errorf("randomFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
