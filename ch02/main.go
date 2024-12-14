package main

import "fmt"

const value = 60

func main() {
	i := 20
	f := float64(i)
	fmt.Println(int(i))
	fmt.Println(int(f))

	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	// Above are the legal max values of the variables (b, smallI, bigI)

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	// Once the variable exceeds it's max legal value, it's reset to the zero value of it's type

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
