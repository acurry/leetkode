package main

import "testing"

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple Flush",
			args: args{
				ranks: []int{13, 1, 2, 3, 4},
				suits: []byte{'a', 'a', 'a', 'a', 'a'},
			},
			want: "Flush",
		},
		{
			name: "Three of a kind",
			args: args{
				ranks: []int{4, 4, 2, 4, 4},
				suits: []byte{'d', 'a', 'a', 'b', 'c'},
			},
			want: "Three of a Kind",
		},
		{
			name: "Pair",
			args: args{
				ranks: []int{10, 10, 2, 12, 9},
				suits: []byte{'a', 'b', 'c', 'a', 'd'},
			},
			want: "Pair",
		},
		{
			name: "Pair and three of a kind",
			args: args{
				ranks: []int{1, 1, 1, 9, 9},
				suits: []byte{'a', 'b', 'c', 'a', 'd'},
			},
			want: "Three of a Kind",
		},
		{
			name: "High Card",
			args: args{
				ranks: []int{8, 10, 2, 12, 9},
				suits: []byte{'a', 'b', 'c', 'a', 'd'},
			},
			want: "High Card",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestHand(tt.args.ranks, tt.args.suits); got != tt.want {
				t.Errorf("bestHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
