package ejerc

func SumarEsMayor(arr []int) []int {
	maxi := make([]int, 2)
	maxi[0] = arr[0]
	maxi[1] = arr[1]
	_sumar(arr, 0, len(arr) - 1, maxi)
	return maxi
}

func _sumar(arr []int, inicio, fin  int, maxi []int ) {
	medio := (inicio + fin) / 2
	maximo := maxi[0] + maxi[1]
	if fin - inicio < 2 {
		return
	}
	
	resul := arr[medio - 1] + arr[medio]
	if resul > maximo {
		maxi[0] = arr[medio - 1]
		maxi[1] = arr[medio]
	}
	if arr[medio] + arr[medio+1] > maximo {
		maxi[0] = arr[medio]
		maxi[1] = arr[medio + 1]
	}
	_sumar(arr, inicio, medio, maxi)
	_sumar(arr, medio, fin, maxi)
}