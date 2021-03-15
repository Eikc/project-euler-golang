package exercises

import "testing"

func Test_numberOfMatchesInTournament(t *testing.T) {
	type args struct {
		participants int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "number of matches in a tournament of four", args: args{participants: 4}, want: 6},
		{name: "number of matches in a tournament of 10.000", args: args{participants: 10000}, want: 49995000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatchesInTournament(tt.args.participants); got != tt.want {
				t.Errorf("numberOfMatchesInTournament() = %v, want %v", got, tt.want)
			}
		})
	}
}
