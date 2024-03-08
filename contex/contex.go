package contex

func Sum(numbers []int, c chan int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	c <- sum
}
