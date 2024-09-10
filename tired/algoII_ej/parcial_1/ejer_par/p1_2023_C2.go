package ejerpar

/*
Implementar en Go un algoritmo que, utilizando división y conquista, reciba un arreglo de enteros de tamaño n y
determine cuál es el par de números vecinos cuya suma es la mayor. Indicar y justificar el orden del algoritmo.
Ejemplo: para el arreglo [4, 8, 9, 1, 3, 6, 7, 1, 11] el par es (8, 9)
*/

func Buscar_suma_mayor(arr []int) int {
	if len(arr) == 2 {
		return arr[0] + arr[1]
	}
	medio := len(arr) / 2
	izq := Buscar_suma_mayor(arr[:medio])
	der := Buscar_suma_mayor(arr[medio + 1:])
	if izq > der {
		return izq
	}
	return der
}


