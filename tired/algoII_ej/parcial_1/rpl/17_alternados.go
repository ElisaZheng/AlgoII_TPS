package rpl



func Alternar(arr []int) []int {
	_modificarArreglo(arr, 0, len(arr)-1)
	return arr
}

func _modificarArreglo(arr []int, inicio, fin int) {
    if inicio == fin {
        return
    }

    medio := (inicio + fin) / 2

    // Dividir el arreglo en dos mitades
    _modificarArreglo(arr, inicio, medio)
    _modificarArreglo(arr, medio+1, fin)

    // Intercambiar los elementos de las dos mitades
    intercambiar(arr, inicio, medio, fin)
}

func intercambiar(arr []int, inicio, medio, fin int) {
    // Calcular el tamaño de las mitades
    tam := medio - inicio + 1

    // Intercambiar los elementos de las dos mitades
    for i := 0; i < tam; i++ {
        arr[medio+i+1], arr[inicio+2*i+1] = arr[inicio+2*i+1], arr[medio+i+1]
    }
}