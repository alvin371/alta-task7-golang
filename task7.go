package main

import (
	"fmt"
	"math"
)

// Fibbonacci Top Down
func fiboTopDown(n int) int {
	var arr map[int]int = map[int]int{}
	nilai, x := arr[n]

	if x {
		return nilai
	} else if n <= 1 {
		arr[n] = n
	} else {
		arr[n] = fiboTopDown(n-2) + fiboTopDown(n-1)
	}
	return arr[n]
}

// FibonacciBottomUp

func fiboBottumUp(n int) int {
	if n <= 1 {
		return n
	}
	arr := []int{0, 1, 0}
	for i := 2; i <= n; i++ {
		arr[2] = arr[0] + arr[1]
		arr[0] = arr[1]
		arr[1] = arr[2]
	}
	return arr[2]
}

// Frog
func pengurangan(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Frog(jumps []int) int {
	result := make(map[int]int)
	N := len(jumps) - 1
	result[0] = 0
	result[1] = pengurangan(jumps[0], jumps[1])
	for i := 2; i <= N; i++ {
		result[i] = int(math.Min(float64(result[i-2]+pengurangan(jumps[i], jumps[i-2])), float64(result[i-1]+pengurangan(jumps[i], jumps[i-1]))))
	}

	return result[N]
}
func main() {
	// fiboTopDown
	fmt.Println(fiboTopDown(0))  // 0
	fmt.Println(fiboTopDown(1))  // 1
	fmt.Println(fiboTopDown(2))  // 1
	fmt.Println(fiboTopDown(3))  // 2
	fmt.Println(fiboTopDown(5))  // 5
	fmt.Println(fiboTopDown(6))  // 8
	fmt.Println(fiboTopDown(7))  // 13
	fmt.Println(fiboTopDown(9))  // 34
	fmt.Println(fiboTopDown(10)) // 55
	fmt.Println(fiboTopDown(50)) // 55

	// fiboBotUp
	fmt.Println(fiboBottumUp(0))  // 0
	fmt.Println(fiboBottumUp(1))  // 1
	fmt.Println(fiboBottumUp(2))  // 1
	fmt.Println(fiboBottumUp(3))  // 2
	fmt.Println(fiboBottumUp(5))  // 5
	fmt.Println(fiboBottumUp(6))  // 8
	fmt.Println(fiboBottumUp(7))  // 13
	fmt.Println(fiboBottumUp(9))  // 34
	fmt.Println(fiboBottumUp(10)) // 55

	// Frog
	fmt.Println(Frog([]int{10, 30, 40, 20}))             // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))     // 40
	fmt.Println(Frog([]int{10, 30, 40, 20, 40}))         // 30
	fmt.Println(Frog([]int{10, 30, 40, 20, 40, 50, 40})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 60})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 20})) // 70
}
