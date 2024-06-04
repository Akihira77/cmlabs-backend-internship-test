package main

import "fmt"

func main() {
	n := 100

	for i := 1; i <= n; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}

		if i%5 == 0 {
			output += "Buzz"
		}

		fmt.Println(i, output)
	}
}
