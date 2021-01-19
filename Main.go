package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	randomMin := 1
	randomMax := 500
	max := 100

	arrNumber := []int{}
	for i := 1; i <= max; i++ {
		arrNumber = append(arrNumber, rand.Intn(randomMax-randomMin+1)+randomMin)
	}

	slice1 := arrNumber[:33]
	slice2 := arrNumber[33:66]
	slice3 := arrNumber[66:]

	fmt.Println("100 angka random dari 1-500 : ")
	fmt.Println(arrNumber)
	fmt.Println("Panjang : ", len(arrNumber))

	fmt.Println()
	fmt.Println("slice1\t\t: ", slice1)
	fmt.Println("Panjang\t\t: ", len(slice1))
	fmt.Println("Min\t\t: ", findMin(slice1))
	fmt.Println("Max\t\t: ", findMax(slice1))
	fmt.Println("Total\t\t: ", sum(slice1))
	rata1 := sum(slice1) / len(slice1)
	fmt.Println("Rata-rata\t: ", rata1)
	sort.Ints(slice1)
	fmt.Println("urut\t\t: ", slice1)

	fmt.Println()
	fmt.Println("slice2\t\t: ", slice2)
	fmt.Println("Panjang\t\t: ", len(slice2))
	fmt.Println("Min\t\t: ", findMin(slice2))
	fmt.Println("Max\t\t: ", findMax(slice2))
	fmt.Println("Total\t\t: ", sum(slice2))
	rata2 := sum(slice2) / len(slice2)
	fmt.Println("Rata-rata\t: ", rata2)
	sort.Ints(slice2)
	fmt.Println("urut\t\t: ", slice2)

	fmt.Println()
	fmt.Println("slice3\t\t: ", slice3)
	fmt.Println("Panjang\t\t: ", len(slice3))
	fmt.Println("Min\t\t: ", findMin(slice3))
	fmt.Println("Max\t\t: ", findMax(slice3))
	fmt.Println("Total\t\t: ", sum(slice3))
	rata3 := sum(slice3) / len(slice3)
	fmt.Println("Rata-rata\t: ", rata3)
	sort.Ints(slice3)
	fmt.Println("urut\t\t: ", slice3)
}

func findMin(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func findMax(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
