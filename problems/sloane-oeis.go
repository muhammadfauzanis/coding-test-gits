package problems

import "fmt"

func lazyCatererSequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = (i*(i+1))/2 + 1
	}
	return sequence
}

func RunSloaneOeis() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Masukkan angka positif!")
		return
	}

	sequence := lazyCatererSequence(n)
	for i, val := range sequence {
		if i == len(sequence)-1 {
			fmt.Printf("%d\n", val)
		} else {
			fmt.Printf("%d-", val)
		}
	}
}
