package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("How many Twix bars did I eat today?", rand.Intn(1000000))
	fmt.Printf("How fat am I now? %g kg\n", math.Sqrt(1248410))
	fmt.Printf("The value of pi to 6 significant digits is: %.5f", math.Pi)
}
