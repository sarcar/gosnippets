package main


/*
Determine the square root of a number using Newton's bisection method
*/


import (
	"math"
	"fmt"
	"os"
	"strconv"
)

/*
Bisection is successive approximation
z[n+1] = z[n] - (z[n] squared - origninal number)/ 2z[n]
*/

func bisect(z float64, x float64) float64 {
	var result = z - (((z*z) - (x))/(2 * z))
	return result
}

/*
Determine the square root
*/

func sqrt (x float64) float64 {
	i:=1
	var interim float64
	var max_iterations = 10
	z := x/2
	for i < max_iterations {
		interim = bisect(z, x)
		z = interim
		fmt.Printf("x = %f, interim root = %f\n",x, z)
		i++
	}
	return interim
}

func print_help() {
	//TODO: Args[0] needs to be parsed to extract only the program name, not fullpathname
	fmt.Printf("Usage:%s <number>\n", os.Args[0])
	fmt.Printf("\tnumber\tthe number whose square root you would like to compute\n")
	os.Exit(1)
}


func main() {
	if len(os.Args) != 2 {
		print_help()
	}

	var x float64
	var err error

	// Silently dropping error
	x, err = strconv.ParseFloat(os.Args[1], 64)

	f := sqrt(x)
	g := math.Sqrt(x)
	fmt.Printf("Square root computed by Go library = %f\n", g)
	fmt.Printf("Square root computed by my program = %f\n", f)
}
