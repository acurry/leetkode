package main

import (
	"math"
	"sync"
	"sync/atomic"
)

const MAX_NEGATIVE int = -2_147_483_648
const MAX_POSITIVE int = 2_147_483_647

func divide(dividend int, divisor int) int {
	a := int(math.Abs(float64(divisor)))
	b := int(math.Abs(float64(dividend)))

	/* edge cases */
	/* MAX_POSITIVE / 1 */
	if dividend >= MAX_POSITIVE && divisor == 1 {
		return MAX_POSITIVE
		/* MAX_NEGATIVE / 1 */
	} else if dividend <= MAX_NEGATIVE && divisor == 1 {
		return MAX_NEGATIVE
	} else if dividend >= MAX_POSITIVE && divisor == -1 {
		return MAX_NEGATIVE
	} else if dividend <= MAX_NEGATIVE && divisor == -1 {
		return MAX_POSITIVE
	}

	var quotient int32
	if b >= MAX_POSITIVE {
		quotient = conc_divide(b, a)
	} else {
		sum := a
		quotient = 0
		for sum <= b {
			quotient += 1
			sum += a
		}
	}

	quotient_should_be_less_than_zero := (dividend < 0 && divisor > 0) ||
		(dividend > 0 && divisor < 0)

	if quotient_should_be_less_than_zero {
		return int(0 - quotient)
	} else {
		return int(quotient)
	}
}

func conc_divide(b, a int) int32 {
	var wg sync.WaitGroup

	var quotient int32

	sub_dividend := b / a

	for i := 0; i < a; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			sum := a
			sub_quotient := 0
			for sum <= sub_dividend {
				sub_quotient += 1
				sum += a
			}

			atomic.AddInt32(&quotient, int32(sub_quotient))
		}()
	}

	wg.Wait()

	return quotient
}

/* latest leetcode output */
/*
Wrong Answer

494 / 994 testcases passed
Input
dividend =
2147483647
divisor =
2
Use Testcase
Output
1073741822
Expected
1073741823
*/
