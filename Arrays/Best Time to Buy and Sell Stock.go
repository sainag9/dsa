package Arrays

// Two pointer technique find valley and peaks
// find valley and assign minimum value
// at peaks find the profit and if greater than previous profit, over wright.
func maxProfit(prices []int) int {

	p := prices[0]
	maxP := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < p {
			p = prices[i]
		} else if prices[i]-p > maxP {
			maxP = prices[i] - p
		}
	}
	return maxP
}
