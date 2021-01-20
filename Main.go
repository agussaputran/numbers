package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func min(kelompok ...int) int {
	minimum := kelompok[0]
	for _, value := range kelompok { //mencari minimum
		if value < minimum {
			minimum = value
		}
	}
	return minimum
}

func max(kelompok ...int) int {
	maximum := kelompok[0]
	for _, value := range kelompok {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}

func sum(kelompok ...int) int {
	total := 0
	for i := 0; i < len(kelompok); i++ { //mencari total kelompok 1
		total += kelompok[i]
	}
	return total
}

func average(total int, kelompok ...int) float32 {
	return float32(total) / float32(len(kelompok))
}

func report(typeKelompok int, kelompok ...int) {
	total := sum(kelompok...)
	rata := average(total, kelompok...)
	min := min(kelompok...)
	max := max(kelompok...)

	fmt.Println("---------------------")
	fmt.Println("Kelompok", typeKelompok)
	fmt.Println("---------------------")

	fmt.Println("total\t:", total)
	fmt.Println("rata\t:", rata)
	fmt.Println("min\t:", min)
	fmt.Println("max\t:", max)
	sort.Ints(kelompok)
	fmt.Println("urut\t:", kelompok)
	fmt.Println()
	fmt.Println("---------------------")
}

func sliceKelompok(arr ...int) (kelompok1, kelompok2, kelompok3 []int) {
	fmt.Println(len(arr))
	for i, v := range arr {
		if i%3 == 0 {
			kelompok1 = append(kelompok1, v)
		} else if i%3 == 1 {
			kelompok2 = append(kelompok2, v)
		} else if i%3 == 2 {
			kelompok3 = append(kelompok3, v)
		}
	}
	return
}

func inisialisasiKelompok(min int, max int, maxNilai int) (nilai []int) {
	for i := 1; i <= max; i++ {
		nilai = append(nilai, rand.Intn(maxNilai-min)+min)
	}
	fmt.Println(nilai)
	return
}

func main() {
	arr := inisialisasiKelompok(20, 100, 100)
	kelompok1, kelompok2, kelompok3 := sliceKelompok(arr...)
	report(1, kelompok1...)
	report(2, kelompok2...)
	report(3, kelompok3...)
}
