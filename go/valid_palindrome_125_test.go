package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "race a car",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "Empty string ' '",
			args: args{
				s: " ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
