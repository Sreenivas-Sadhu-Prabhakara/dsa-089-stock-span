package main

import (
	"reflect"
	"testing"
)

func TestSpans(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
		{[]int{10, 20, 30}, []int{1, 2, 3}},
		{[]int{30, 20, 10}, []int{1, 1, 1}},
		{[]int{5, 5, 5}, []int{1, 2, 3}}, // equal prices extend the span
		{[]int{7}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := Spans(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Spans(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
