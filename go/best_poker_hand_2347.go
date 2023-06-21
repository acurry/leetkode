package main

import "reflect"

type RankCount int

func bestHand(ranks []int, suits []byte) string {
	// check for Flush
	suitMap := make(map[byte]int, len(suits))
	for _, s := range suits {
		suitMap[s] += 1
	}
	keys := reflect.ValueOf(suitMap).MapKeys()

	if len(keys) == 1 {
		return "Flush"
	}

	rankMap := make(map[int]RankCount, len(ranks))
	for _, r := range ranks {
		rankMap[r] += 1
	}

	has_three_of_a_kind := false
	has_pair := false
	for _, rankCount := range rankMap {
		if rankCount >= 3 {
			has_three_of_a_kind = true
		} else if rankCount == 2 {
			has_pair = true
		}
	}

	if has_three_of_a_kind {
		return "Three of a Kind"
	} else if has_pair {
		return "Pair"
	}

	return "High Card"
}
