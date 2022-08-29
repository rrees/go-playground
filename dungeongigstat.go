package main


func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func calculateDGStat(scores [3]int) int {
	if scores[0] != scores[1] && scores[1] != scores[2] {
		return scores[1]
	}
	return max(scores[1], scores[2])
}
