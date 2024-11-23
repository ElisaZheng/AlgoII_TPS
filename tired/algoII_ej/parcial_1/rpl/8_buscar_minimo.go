package rpl

/*
Implementar, por división y conquista, una función que determine el mínimo de un arreglo.
Indicar y justificar el orden.
*/


// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento
func BuscarMinimo_1(arr []int) int {
    inicio := 0
    fin := len(arr) - 1
    medio := (inicio + fin) / 2
    min := arr[medio]
    return _buscarMinimo(arr, inicio, fin, min)
}

func _buscarMinimo(arr []int, inicio, fin, min int) int {
    if inicio >= fin {
        return min
    }
    medio := (inicio + fin) / 2
    if arr[medio] < min {
        min = arr[medio]
    }
    min = _buscarMinimo(arr, inicio, medio - 1, min)
    min = _buscarMinimo(arr, medio + 1, fin, min)
    return min
}

//puedo calcular la complejidad con la teorema Maestro, 
//A: cantidad de llamados recursivos = 2, 
//B: la proporcion del tamano original = 2 , 
//C = 0 => log 2(2) = 1 > 1 => O(n)

func BuscarMinimo_2(arr []int) int {
    inicio := 0
    fin := len(arr) - 1
    return _buscarMinimo_2(arr, inicio, fin)
}

func _buscarMinimo_2(arr []int, inicio, fin int) int {
    if inicio == fin {
        return arr[inicio]
    }
	medio := (inicio + fin) / 2
    izq := _buscarMinimo_2(arr, inicio, medio)
    der := _buscarMinimo_2(arr, medio + 1, fin)
    if izq > der {
		return der
	}
	return izq
}

func BuscarMinimo_3(arr []int) int {
    inicio := 0
    fin := len(arr) - 1
    medio := (inicio + fin) / 2
    min := arr[medio]
    return _buscarMinimo_3(arr, inicio, fin, min)
}

func _buscarMinimo_3(arr []int, inicio, fin, min int) int {
    if inicio >= fin {
        return min
    }
    medio := (inicio + fin) / 2
    if arr[medio] < min {
        min = arr[medio]
    }
    min = _buscarMinimo_3(arr, inicio, medio - 1, min)
    min = _buscarMinimo_3(arr, medio + 1, fin, min)
    return min
}