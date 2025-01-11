package array

// MinSubArrayLen 寻找和大于等于目标值的最短连续子数组的长度
//
// 1. 算法原理：
//   - 使用滑动窗口法
//   - 窗口右边界不断右移来扩大窗口
//   - 当窗口内元素和大于等于目标值时，尝试移动左边界以缩小窗口
//   - 在这个过程中维护最小长度
//
// 2. 时间复杂度分析：
//   - O(n)：每个元素最多被访问两次（加入窗口和移出窗口）
//   - 左右指针各遍历数组一次
//
// 3. 空间复杂度分析：
//   - O(1)：只需要常量级额外空间
//   - 仅使用几个变量来维护窗口状态和结果
//
// 4. 关键实现点：
//   - 使用两个指针维护滑动窗口
//   - 通过比较窗口和与目标值的关系来移动指针
//   - 实时更新最小长度
//
// 5. 适用场景：
//   - 需要查找连续子数组
//   - 子数组需要满足特定的和条件
//   - 要求最小长度的场景
//
// 6. 注意事项：
//   - 所有输入数字均为正整数
//   - 如果不存在满足条件的子数组，返回0
//   - 需要考虑边界情况
//   - 注意初始最小长度的设置
func MinSubMarrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	n++

	minLen := n
	l, sum := 0, 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			currLen := r - l + 1
			if currLen < minLen {
				minLen = currLen
			}
			sum -= nums[l]
			l++
		}
	}
	if minLen == n {
		return 0
	} else {
		return minLen
	}
}
