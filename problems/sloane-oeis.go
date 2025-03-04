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
	var inputNumber int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&inputNumber)

	if inputNumber <= 0 {
		fmt.Println("Angka tidak boleh negatif dan 0, masukkan angka positif!")
		return
	}

	sequence := lazyCatererSequence(inputNumber)
	for i, val := range sequence {
		if i == len(sequence)-1 {
			fmt.Printf("%d\n", val)
		} else {
			fmt.Printf("%d-", val)
		}
	}
}
