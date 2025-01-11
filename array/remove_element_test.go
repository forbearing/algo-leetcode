package array_test

import (
	"reflect"
	"testing"

	"leetcode/array"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		result   []int
	}{
		{
			name:     "示例1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			result:   []int{2, 2},
		},
		{
			name:     "示例2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			result:   []int{0, 1, 3, 0, 4},
		},
		{
			name:     "空数组",
			nums:     []int{},
			val:      1,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "全部是目标值",
			nums:     []int{1, 1, 1},
			val:      1,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "没有目标值",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			expected: 4,
			result:   []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.nums))
			copy(original, tt.nums)
			length := array.RemoveElement(tt.nums, tt.val)
			if length != tt.expected {
				t.Errorf("期望长度为 %d, 但得到 %d", tt.expected, length)
			}
			if length > 0 {
				result := tt.nums[:length]
				if !reflect.DeepEqual(result, tt.result) {
					t.Errorf("对于输入 %v, 期望结果为 %v, 但得到 %v",
						original, tt.result, result)
				}
			}
		})
	}
}

func TestRemoveElement2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			name:     "示例2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			name:     "空数组",
			nums:     []int{},
			val:      1,
			expected: 0,
		},
		{
			name:     "全部是目标值",
			nums:     []int{1, 1, 1},
			val:      1,
			expected: 0,
		},
		{
			name:     "没有目标值",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			expected: 4,
		},
		{
			name:     "单个元素-等于目标值",
			nums:     []int{1},
			val:      1,
			expected: 0,
		},
		{
			name:     "单个元素-不等于目标值",
			nums:     []int{1},
			val:      2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.nums))
			copy(original, tt.nums)

			length := array.RemoveElement2(tt.nums, tt.val)

			if length != tt.expected {
				t.Errorf("输入数组 %v, 目标值 %d, 期望长度为 %d, 但得到 %d",
					original, tt.val, tt.expected, length)
			}

			// 验证前length个元素中不包含目标值
			for i := 0; i < length; i++ {
				if tt.nums[i] == tt.val {
					t.Errorf("结果数组中仍包含目标值 %d", tt.val)
				}
			}
		})
	}
}
