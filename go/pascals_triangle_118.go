package main

func generatePascalsTriangle(numRows int) [][]int {
	rows := [][]int{}

	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			row = append(row, binomCoeff(i, j))
		}
		rows = append(rows, row)
	}

	return rows
}

/*
# See https://www.geeksforgeeks.org/space-and-time-efficient-binomial-coefficient/
# for details of this function
def binomialCoeff(n, k) :

	res = 1
	if (k > n - k) :
	    k = n - k
	for i in range(0 , k) :
	    res = res * (n - i)
	    res = res // (i + 1)

	return res
*/
func binomCoeff(n, k int) int {
	res := 1
	if k > n-k {
		k = n - k
	}
	for i := 0; i < k; i++ {
		res *= n - i
		res = int(float64(res / (i + 1)))
	}
	return res
}
