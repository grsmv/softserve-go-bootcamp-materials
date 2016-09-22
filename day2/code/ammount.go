package main

func add(amount int, paymants ...int) (sum int) {
	sum = len(paymants)
	for _, i := range paymants {
		sum += i
	}
	return
}

func main() {
	println("Total:", add(10, 20, 30, -20))
}
