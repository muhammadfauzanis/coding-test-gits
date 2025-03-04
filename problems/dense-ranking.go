package problems

import "fmt"

func DenseRanking(scores []int, gitsScores []int) []int {
	rankMap := make(map[int]int) // Map to store the rank of each unique score
	rank := 1
	prevScore := -1

	// Generate dense ranking for leaderboard scores
	for _, score := range scores {
		if score != prevScore {
			rankMap[score] = rank
			rank++
		}
		prevScore = score
	}

	// Find rank for each GITS score
	result := []int{}
	for _, gitsScore := range gitsScores {
		for i := 0; i < len(scores); i++ {
			if gitsScore >= scores[i] {
				result = append(result, rankMap[scores[i]])
				break
			}
			if i == len(scores)-1 {
				result = append(result, rank)
			}
		}
	}

	return result
}

func RunDenseRank() {
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	gitsScores := []int{5, 25, 50, 120}

	result := DenseRanking(scores, gitsScores)
	fmt.Println(result)
}
