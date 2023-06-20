package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Base case: n = 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Base case: n = 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "n = 4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "n = 5",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "n = 6",
			args: args{
				n: 6,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
