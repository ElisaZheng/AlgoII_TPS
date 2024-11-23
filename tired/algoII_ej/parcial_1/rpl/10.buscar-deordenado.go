package rpl


/*
Implementar, por división y conquista,
una función que dado un arreglo sin elementos repetidos y casi ordenado (todos los elementos se encuentran ordenados,
	salvo uno), obtenga el elemento fuera de lugar.
Indicar y justificar el orden.
*/

/* func ElementoDesordenado(arr []int) int {
	if len(arr) == 2 {
        if arr[0] > arr[1] {
			return arr[0]
		} else {
			return -1
		}
    }
	if len(arr) <= 1{
		return -1
	}
	medio := len(arr) / 2
    
    if arr[medio-1] < arr[medio] && arr[medio] < arr[medio + 1]  {
        izq := ElementoDesordenado(arr[:medio])
        der := ElementoDesordenado(arr[medio + 1:])
        if izq != -1 {
            return izq
        } else {
            return der
        }
    } else if arr[medio-1] > arr[medio] {
		return arr[medio]
	} else {
		return arr[medio + 1]
	}
}*/
func ElementoDesordenado(arr []int) int {
    return _desordenado(arr, 0, len(arr) - 1)
}

func _desordenado(arr []int, inicio, fin int) int {
    medio := (inicio + fin) / 2
    if inicio > fin {
        return - 1
    }
    if medio == 0 && arr[medio] > arr[medio + 1] {
        return arr[medio]
    }

    if arr[medio] > arr[medio + 1] {
        return arr[medio + 1]
    }
    izq := _desordenado(arr, inicio, medio - 1)
    if izq != -1 {
        return izq
    }
    return _desordenado(arr, medio + 1, fin)

}