package array_test

import (
	"reflect"
	"testing"

	"leetcode/array"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "示例1",
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "示例2",
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			name:     "全负数",
			nums:     []int{-5, -3, -2, -1},
			expected: []int{1, 4, 9, 25},
		},
		{
			name:     "全正数",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 4, 9, 16, 25},
		},
		{
			name:     "只有零",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.SortedSquares(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("对于输入 %v, 期望得到 %v, 但得到 %v",
					tt.nums, tt.expected, result)
			}
		})
	}
}
