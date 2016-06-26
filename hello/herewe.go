package main

import (
	"fmt"
	"time"
	"math"
)



func pow (x int, y int) int {
	result := x
	for i:=1; i < y; i++ {
		result = result * x
	}
	return result
}

func main() {
	const pi = math.Pi
	fmt.Printf("Pi equals roughly %.2f\n", pi)
	fmt.Println("The time is ", time.Now())
	fmt.Println("3 squared = ", pow(3,2))
	fmt.Println("3 cubed   = ", pow(3,3))
}


