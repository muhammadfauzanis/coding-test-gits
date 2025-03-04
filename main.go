package main

import (
	"fmt"
	"technical-test-gits/problems"
)

func main() {
	fmt.Println("--- Problem No 1 Sloane OEIS ---")
	problems.RunSloaneOeis()
	fmt.Println()
	fmt.Println("--- Problem No 2 Dense Rank ---")
	problems.RunDenseRank()
	fmt.Println()
	fmt.Println("--- Problem No 3 Bracket Tests ---")
	problems.RunBalancedBracketTests()
}
