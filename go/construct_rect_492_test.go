package main

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1: 4 => [2,2]",
			args: args{area: 4},
			want: []int{2, 2},
		},
		{
			name: "t2 prime: 37 => [37,1]",
			args: args{area: 37},
			want: []int{37, 1},
		},
		{
			name: "t3: 122122 => [427,286]",
			args: args{area: 122122},
			want: []int{427, 286},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
