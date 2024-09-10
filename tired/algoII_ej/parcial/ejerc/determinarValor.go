package ejerc

func DeterminarValor(arr []int) bool {
	valor := -len(arr) * len(arr)
	return _determinar(arr, valor)
}

func _determinar(arr []int, valor int) bool {
	if len(arr) == 1 {
		return arr[0] == valor
	}
	medio := len(arr) / 2
	if arr[medio] == valor {
		return true
	} else if arr[medio] > valor {
		return _determinar(arr[medio:], valor)
	} else {
		return _determinar(arr[:medio], valor)
	}
}
/* inicializar un variable valor -> O(1)
entrar a la llamada recusiva -> 1
el arrglo original se parte a mitad -> 2
el costo de juntar y unir -> 0
uso t de Maestro
T(n) = AT(n/B) + O(n^C)
logB A = 0 = C -> O(log(n))
*/