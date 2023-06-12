package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Two small positive ints",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "Small positive dividend, small negative dividend",
			args: args{
				dividend: 10,
				divisor:  -3,
			},
			want: -3,
		},
		{
			name: "Small positive dividend, small negative dividend",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "Equal positive dividend and divisor",
			args: args{
				dividend: 1,
				divisor:  1,
			},
			want: 1,
		},
		{
			name: "Equal negative dividend and divisor",
			args: args{
				dividend: -1,
				divisor:  -1,
			},
			want: 1,
		},
		{
			name: "Max negative int32 dividend",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		},
		{
			name: "Perf test 1",
			args: args{
				dividend: 2147483647,
				divisor:  3,
			},
			want: 715827882,
		},
		// {
		// 	name: "Perf test 2",
		// 	args: args{
		// 		dividend: 2147483647,
		// 		divisor:  2,
		// 	},
		// 	want: 1073741823,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
