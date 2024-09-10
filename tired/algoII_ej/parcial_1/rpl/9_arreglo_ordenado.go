package rpl

/*
Implementar, por división y conquista, 
una función que dado un arreglo y su largo, 
determine si el mismo se encuentra ordenado. 
Indicar y justificar el orden.
*/


func EstaOrdenado(arr []int) bool {
    if len(arr) <= 1 {
        return true
    }
	inicio := 0
	fin := len(arr)
	return _estaOrdenado(arr, inicio, fin)
}

func _estaOrdenado(arr []int, inicio, fin int) bool {
    if fin - inicio <= 1 {
		return true
	}
	medio := (inicio + fin) / 2
	
	if arr[medio-1] > arr[medio] {
		return false
	} else {
        return _estaOrdenado(arr, inicio, medio) && _estaOrdenado(arr, medio,  fin)
    }
}

func EstaOrdenado_1(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	medio := len(arr) / 2
	if arr[medio - 1] > arr[medio] {
		return false
	} else {
		return EstaOrdenado_1(arr[:medio]) && EstaOrdenado_1(arr[medio:])
	}
}