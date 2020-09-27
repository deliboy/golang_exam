package main

import "fmt"

func Move(n int, from, to, via string) {

	if n <= 0 {
		return
	}

	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}

func Hanoi(n int) {
	fmt.Println("Number of diks:", n)
	Move(n, "From", "To", "Via")
}

func main() {
	Hanoi(3)
}
