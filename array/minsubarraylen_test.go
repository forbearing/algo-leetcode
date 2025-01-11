package array_test

import (
	"testing"

	"leetcode/array"
)

func TestMinSubArray(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "示例1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "没有满足条件的子数组",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "单个元素满足条件",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "所有元素刚好满足条件",
			target: 15,
			nums:   []int{5, 5, 5},
			want:   3,
		},
		{
			name:   "空数组",
			target: 100,
			nums:   []int{},
			want:   0,
		},
		{
			name:   "较大的目标值",
			target: 11,
			nums:   []int{1, 2, 3, 4, 5},
			want:   3,
		},
		{
			name:   "最小的可能目标值",
			target: 1,
			nums:   []int{1, 1, 1},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.MinSubMarrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("MinSubArrayLen(%v, %v) = %v, want %v",
					tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
