package main

import (
	"fmt"
	"modulo/binarySearch"
	"modulo/recursion"
	"modulo/sortList"
)

func main() {
	list := []int{12, 14, 54, 65, 123, 324, 654, 1245, 1333, 1910, 2150, 4500}
	fmt.Println(binarySearch.Search(list, 1333))

	lista2 := []int{12, 2, 12, 4, 54, 132, 65, 4362, 543}
	fmt.Println(sortList.Sort(lista2))

	fatDeMenosDois := recursion.Fatorial(-2)
	fatDeZero := recursion.Fatorial(0)
	fatDeUm := recursion.Fatorial(1)
	fatDeCinco := recursion.Fatorial(5)
	fatDeDezoito := recursion.Fatorial(18)

	fmt.Println(fatDeMenosDois)
	fmt.Println(fatDeZero)
	fmt.Println(fatDeUm)
	fmt.Println(fatDeCinco)
	fmt.Println(fatDeDezoito)
}
