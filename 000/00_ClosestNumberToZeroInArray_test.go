package exercises

import "testing"

func Test_getClosesNumberToZeroIn(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "closest to zero in array of 1, -9, 8, 2, 5, 9 ", args: args{numbers: []int{-9, 8, 1, 2, 5, 9}}, want: 1},
		{name: "closest to zero in array of  -9, 8, 2, 5, 9 ", args: args{numbers: []int{-9, 8, 2, 5, 9}}, want: 2},
		{name: "closest to zero in array of  -3, 8, 3, 5, 9 ", args: args{numbers: []int{-3, 8, 3, 5, 9}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClosesNumberToZeroIn(tt.args.numbers); got != tt.want {
				t.Errorf("getClosesNumberToZeroIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
