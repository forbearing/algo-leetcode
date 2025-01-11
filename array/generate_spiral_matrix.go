package array

// GenerateSpiralMatrix 生成螺旋矩阵
//
// 1. 算法原理：
//   - 使用模拟法，按顺时针方向填充数字
//   - 使用四个边界变量控制填充范围
//   - 按照右、下、左、上的顺序循环填充
//
// 2. 时间复杂度分析：
//   - O(n^2)：需要填充 n×n 个格子
//   - 每个格子只被访问一次
//
// 3. 空间复杂度分析：
//   - O(1)：除了返回的矩阵外，只使用常量额外空间
//   - 不考虑返回值的空间占用
//
// 4. 关键实现点：
//   - 使用四个指针标记边界
//   - 按照固定的方向顺序填充
//   - 每填充完一条边，更新相应的边界
//
// 5. 适用场景：
//   - 需要生成螺旋矩阵
//   - 需要模拟螺旋遍历的场景
//
// 6. 注意事项：
//   - 输入必须是正整数
//   - 需要准确控制边界条件
//   - 需要正确处理每一轮的起止位置
func GenerateSpiralMatrix(n int) [][]int {
	// 创建 n×n 的矩阵
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 定义四个边界
	left, right := 0, n-1 // 左右边界
	top, bottom := 0, n-1 // 上下边界
	num := 1              // 要填充的数字

	// 当 num 小于等于 n^2 时继续填充
	for num <= n*n {
		// 从左到右填充上边
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++ // 上边界下移

		// 从上到下填充右边
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right-- // 右边界左移

		// 从右到左填充下边
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom-- // 下边界上移

		// 从下到上填充左边
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++ // 左边界右移
	}

	return matrix
}

func GenerateSpiralMatrix2(n int) [][]int {
	startI, startJ := 0, 0
	loop, mid := n/2, n/2
	offset, count := 1, 1

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := startI, startJ

		for j = startJ; j < n-offset; j++ {
			matrix[i][j] = count
			count++
		}
		for i = startI; i < n-offset; i++ {
			matrix[i][j] = count
			count++
		}
		for j = n - offset; j > startJ; j-- {
			matrix[i][j] = count
			count++
		}
		for i = n - offset; i > startI; i-- {
			matrix[i][j] = count
			count++
		}

		offset++
		loop--
		startI++
		startJ++
	}
	if n%2 == 1 {
		matrix[mid][mid] = count
	}
	return matrix
}
