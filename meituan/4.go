package main

func candy2(sweetness []int, k int) int {
	// eat one candy every 2 days, but have k times to regret for eating a candy even if yesterday ate a candy
	// find the max sweetness you can get
	// dp[i][j] means the max sweetness you can get when you have eaten i candies and regretted j times
	n := len(sweetness)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = sweetness[0]
	dp[1][0] = max(sweetness[0], sweetness[1])
	if k > 0 {
		dp[1][1] = sweetness[1] + sweetness[0]
	}
	res := 0
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-2][0]+sweetness[i], dp[i-1][0])
		for j := 1; j <= k; j++ {
			dp[i][j] = max(max(dp[i-1][j-1]+sweetness[i], dp[i-2][j]+sweetness[i]), dp[i-1][j])
			// dp[i][j] = max(res+sweetness[i], dp[i-1][j])
			res = max(res, dp[i][j])
		}
	}
	return res
	// dp := make([][][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([][]int, k+1)
	// 	for j := 0; j <= k; j++ {
	// 		dp[i][j] = make([]int, 2)
	// 	}
	// }
	// dp[0][0][0] = 0
	// dp[0][0][1] = sweetness[0]
	// dp[1][0][0] = sweetness[0]
	// dp[1][0][1] = max(sweetness[0], sweetness[1])
	// if k > 0 {
	// 	dp[1][1][1] = sweetness[1] + sweetness[0]
	// }
	// res := 0
	// for i := 2; i < n; i++ {
	// 	dp[i][0][0] = dp[i-1][0][1]
	// 	dp[i][0][1] =
	// }
	// return res

}
