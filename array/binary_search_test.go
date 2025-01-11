package array_test

import (
	"testing"

	"leetcode/array"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "基本测试-找到目标值",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15},
			target:   7,
			expected: 3,
		},
		{
			name:     "基本测试-目标值不存在",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15},
			target:   8,
			expected: -1,
		},
		{
			name:     "边界测试-空数组",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "边界测试-查找第一个元素",
			arr:      []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "边界测试-查找最后一个元素",
			arr:      []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "边界测试-单个元素数组",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "特殊测试-重复元素",
			arr:      []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 2, // 返回其中任意一个位置即可
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
