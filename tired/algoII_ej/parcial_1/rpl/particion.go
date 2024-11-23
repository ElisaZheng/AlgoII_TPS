package rpl

import (
	"fmt"
)

func DeterminarParejasDe(tuercas[]int, tor[]int) ([]int, []int){
	res := comparar(tor, 0, len(tuercas) - 1, tuercas)
	return res, tor
	
}

func comparar(tor []int, inicio, fin int, tuercas []int) []int {
	resultado := make([]int, len(tor))
	for _, t := range tuercas {
		p := partition(tor, inicio, fin, t)
		fmt.Println(p)
		resultado[p] = t
	}
	return resultado
} 

func partition(arr []int, inicio, fin, t int) int {
	i := inicio 
	for j := inicio; j < fin; i++ {
		if arr[j] < t {
			arr[i], arr[j] = arr[j], arr[i]
			i ++
		} 
	}
	arr[i], arr[fin] = arr[fin], arr[i]
	return i
}