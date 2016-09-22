package main
import "fmt"

type fizzbuzz int

func (x fizzbuzz) String() string {
	switch [2]bool{x%3 == 0, x%5 == 0} {
	case [2]bool{true, true}:
		return "FizzBuzz"
	case [2]bool{true, false}:
		return "Fizz"
	case [2]bool{false, true}:
		return "Buzz"
	default:
		return fmt.Sprint(int(x))
	}
}

func main() {
	for x := fizzbuzz(1); x <= 100; x++ {
		fmt.Println(x)
	}
}