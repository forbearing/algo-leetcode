package array

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CalculateRangeSum 使用前缀和计算区间和
//
// 1. 算法原理：
//   - 先计算前缀和数组
//   - 使用前缀和数组快速计算区间和
//   - 区间和 = prefixSum[right] - prefixSum[left-1]
//
// 2. 时间复杂度分析：
//   - 构建前缀和：O(n)
//   - 每次查询：O(1)
//   - 总体：O(n + q)，q为查询次数
//
// 3. 空间复杂度分析：
//   - O(n)：需要存储前缀和数组
//   - 额外使用常量空间处理输入输出
//
// 4. 关键实现点：
//   - 使用前缀和优化查询
//   - 处理输入输出的细节
//   - 正确计算区间和
func CalculateRangeSum() {
	// 创建带缓冲的读写器
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 读取数组长度
	var n int
	fmt.Fscanln(reader, &n)

	// 读取数组元素
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &nums[i])
	}

	// 计算前缀和
	prefixSum := make([]int, n+1) // 多一个元素方便计算
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	// 处理区间查询
	for {
		// 读取一行
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 解析区间
		var left, right int
		fmt.Sscanf(line, "%d %d", &left, &right)

		// 计算并输出区间和
		// 注意：输入的区间是从0开始的，前缀和数组是从1开始的
		sum := prefixSum[right+1] - prefixSum[left]
		fmt.Fprintln(writer, sum)
	}
}

// CalculateRangeSum2 另一种实现方式：使用字符串分割处理输入
func CalculateRangeSum2() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 读取数组长度
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 读取数组元素
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}

	// 计算前缀和
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	// 处理查询
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var left, right int
		fmt.Sscanf(line, "%d %d", &left, &right)
		sum := prefixSum[right+1] - prefixSum[left]
		fmt.Fprintln(writer, sum)
	}
}
