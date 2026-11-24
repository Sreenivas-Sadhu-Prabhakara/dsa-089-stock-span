package main

import "fmt"

// Spans returns, for each day, how many consecutive days ending today had a
// price less than or equal to today's. A monotonic stack of indexes with
// strictly greater prices turns the lookback into O(1) amortized.
func Spans(prices []int) []int {
	out := make([]int, len(prices))
	stack := []int{} // indexes whose prices strictly decrease upward
	for i, p := range prices {
		for len(stack) > 0 && prices[stack[len(stack)-1]] <= p {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			out[i] = i + 1
		} else {
			out[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return out
}

func main() {
	fmt.Println(Spans([]int{100, 80, 60, 70, 60, 75, 85}))
}
