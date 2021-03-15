package exercises

import "testing"

func Test_fibOne(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "name fib of 0 is 1", args: args{number: 0}, want: 1},
		{name: "name fib of 1 is 1", args: args{number: 1}, want: 1},
		{name: "name fib of 3 is 2", args: args{number: 3}, want: 2},
		{name: "name fib of 10 is 55", args: args{number: 10}, want: 55},
		{name: "name fib of 20 is 6765", args: args{number: 20}, want: 6765},
		{name: "name fib of 45 is huge", args: args{number: 45}, want: 1134903170},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibOne(tt.args.number); got != tt.want {
				t.Errorf("fibOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
