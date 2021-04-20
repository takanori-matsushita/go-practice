package sum

func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) int {
	sum := 0

	// リファクタ前
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// リファクタ後
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// リファクタ前
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// リファクタ後
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers[1:]))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return
}
