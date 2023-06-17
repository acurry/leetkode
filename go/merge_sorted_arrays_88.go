package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	for i := 0; i < m+n; i++ {
		curr_nums1 := nums1[i]
		for j := 0; j < n; j++ {
			curr_nums2 := nums2[j]
			if curr_nums2 < curr_nums1 {
				insert_at_index(nums1, i, curr_nums2)
			}
		}
	}
}

func insert_at_index(n []int, index, val int) {
	n = n[0:index]
	m := n[index:]
	n = append(n, val)
	n = append(n, m...)
}
