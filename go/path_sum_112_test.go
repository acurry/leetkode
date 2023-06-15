package main

import (
	"testing"
)

var simpleCaseOneTree TreeNode

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "Simple case 1",
		// 	args: args{
		// 		root:
		// 	}
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}

}
