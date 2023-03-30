package main

func candy(sweetness []int) int {
	n := len(sweetness)
	dp := make([]int, n)
	dp[0] = sweetness[0]
	dp[1] = max(sweetness[0], sweetness[1])
	dp[2] = max(sweetness[0], max(sweetness[1], sweetness[2]))
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-3]+sweetness[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
