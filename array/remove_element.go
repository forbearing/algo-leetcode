package array

// RemoveElement 移除数组中指定的元素
//
// 1. 算法原理：
//   - 使用快慢双指针遍历数组
//   - 快指针负责遍历数组寻找不等于目标值的元素
//   - 慢指针指向待更新的位置
//   - 当快指针找到不等于目标值的元素时，将其移动到慢指针位置
//
// 2. 时间复杂度分析：
//   - O(n)：需要遍历整个数组一次
//   - 每个元素最多被操作一次
//
// 3. 空间复杂度分析：
//   - O(1)：只使用了两个指针变量
//   - 在原数组上进行修改，不需要额外空间
//
// 4. 关键实现点：
//   - 使用快慢指针避免了使用额外空间
//   - 通过覆盖的方式实现原地删除
//   - 只遍历一次数组即可完成操作
//
// 5. 适用场景：
//   - 需要原地删除数组中特定元素
//   - 对空间复杂度有严格要求的场景
//   - 元素顺序可以改变的场景
//
// 6. 注意事项：
//   - 原数组会被修改
//   - 返回值表示新数组的长度
//   - 不需要考虑超出新长度后的元素
//   - 元素的相对顺序可能会改变
func RemoveElement(nums []int, val int) int {
	// 慢指针，指向当前需要更新的位置
	slow := 0
	// 快指针，用于遍历数组
	for fast := 0; fast < len(nums); fast++ {
		// 当快指针指向的元素不等于目标值时
		if nums[fast] != val {
			// 将快指针指向的元素移动到慢指针位置
			nums[slow] = nums[fast]
			// 慢指针向前移动
			slow++
		}
	}
	return slow
}

// RemoveElement2 移除数组中指定的元素（相向双指针法）
//
// 1. 算法原理：
//   - 使用左右双指针从数组两端向中间移动
//   - 左指针寻找等于目标值的元素
//   - 右指针寻找不等于目标值的元素
//   - 当找到时，用右指针元素覆盖左指针元素
//   - 这种方法相比同向双指针，可以减少数组元素的移动次数
//
// 2. 时间复杂度分析：
//   - O(n)：需要遍历数组一次
//   - 但实际运行时比同向双指针更快，因为移动次数更少
//
// 3. 空间复杂度分析：
//   - O(1)：只使用了两个指针变量
//   - 在原数组上进行修改，不需要额外空间
//
// 4. 关键实现点：
//   - 使用首尾双指针，从两端向中间移动
//   - 通过交换的方式减少元素移动次数
//   - 当左指针超过右指针时结束
//
// 5. 适用场景：
//   - 需要原地删除数组中特定元素
//   - 对空间复杂度有严格要求的场景
//   - 元素顺序可以改变的场景
//
// 6. 注意事项：
//   - 原数组会被修改
//   - 元素的相对顺序会改变
//   - 不需要考虑超出新长度后的元素
//   - 特别适合目标元素较少的情况
func RemoveElement2(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// 从左向右找等于val的元素
		for left <= right && nums[left] != val {
			left++
		}
		// 从右向左找不等于val的元素
		for left <= right && nums[right] == val {
			right--
		}
		// 如果左指针依然小于等于右指针，则进行交换
		if left <= right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
