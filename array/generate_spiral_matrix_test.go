package array_test

import (
	"reflect"
	"testing"

	"leetcode/array"
)

func TestGenerateSpiralMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "3x3矩阵",
			n:    3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "1x1矩阵",
			n:    1,
			want: [][]int{{1}},
		},
		{
			name: "2x2矩阵",
			n:    2,
			want: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "4x4矩阵",
			n:    4,
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.GenerateSpiralMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSpiralMatrix(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := array.GenerateSpiralMatrix2(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSpiralMatrix2(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
