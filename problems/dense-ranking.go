package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DenseRanking(scores []int, gitsScores []int) []int {
	rankMap := make(map[int]int)
	rank := 1
	prevScore := -1

	for _, score := range scores {
		if score != prevScore {
			rankMap[score] = rank
			rank++
		}
		prevScore = score
	}

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
	reader := bufio.NewReader(os.Stdin)

	var player int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&player)

	reader.ReadString('\n')

	fmt.Print("Masukkan daftar skor: ")
	scoreInput, _ := reader.ReadString('\n')
	scoreStr := strings.Fields(scoreInput)
	scores := make([]int, len(scoreStr))
	for i, val := range scoreStr {
		scores[i], _ = strconv.Atoi(val)
	}

	var totalGames int
	fmt.Print("Masukkan jumlah permainan yang diikuti GITS: ")
	fmt.Scanln(&totalGames)

	fmt.Print("Masukkan skor yang didapat oleh GITS: ")
	inputGames, _ := reader.ReadString('\n')
	gitsStr := strings.Fields(inputGames)
	gitsScores := make([]int, len(gitsStr))
	for i, val := range gitsStr {
		gitsScores[i], _ = strconv.Atoi(val)
	}

	result := DenseRanking(scores, gitsScores)
	fmt.Println("Peringkat GITS:", result)
}
