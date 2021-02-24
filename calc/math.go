package calc

import "errors"

//Add returns sum of two integers
func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return sum
	}

}
