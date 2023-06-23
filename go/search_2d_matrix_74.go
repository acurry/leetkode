package main

func searchMatrix(matrix [][]int, target int) bool {
	target_row_index := 0
	found := false

	for i := 0; i < len(matrix); i++ {
		head := matrix[i][0]
		tail := matrix[i][len(matrix[0])-1]
		if target == head || target == tail {
			found = true
			break
		} else if target > head && target < tail {
			target_row_index = i
		}
	}

	// target not one of the first elements in each row
	if !found {
		for j := 0; j < len(matrix[target_row_index]); j++ {
			curr := matrix[target_row_index][j]

			if target == curr {
				found = true
				break
			}
		}
	}

	return found
}
