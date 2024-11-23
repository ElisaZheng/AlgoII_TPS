package ejerc

func SumarVecinoEsMayor(arr []int) ([]int, int) {
	result, valor := _sumarVecinoEsMayor(arr, []int {0, 0})
	return result, valor
}

func _sumarVecinoEsMayor(arr []int, sumo[]int) ([]int, int) {
	medio := len(arr) / 2
	mayor := sumo[0] + sumo[1] // O[1]
	if len(arr) <= 2 { // O[1]
		return arr, mayor
	}
	if arr[medio] + arr[medio-1] > mayor { //O[1]
		sumo[0], sumo[1] = arr[medio-1], arr[medio]
		mayor = arr[medio-1] + arr[medio]
	}
	
	max, valor1 := _sumarVecinoEsMayor(arr[:medio], sumo)
	max2, valor2 :=_sumarVecinoEsMayor(arr[medio:], sumo)
	if valor1 > mayor && valor1 > valor2 {
		return max, valor1
	} else if valor2 > mayor && valor2 > valor1 {
		return max2, valor2
	}
	return sumo, mayor
}

// porque el algoritmo lo resolvi en division y conquista, puedo usar Torema de Maestro para justificar la complejidad.
// A, la cantidad de llamados recursivos, va a ser 2, 
// B, proporcion del tamano original con el que llamamos recursivamos, va a ser 2(cada vez se divide en 2),
// C, el costo de partir y juntar, va a ser O(1) * 5 = O(5) = O(1), entonces C = 0, 
// la formula es :T(n) = AT(n/B) + O(n^C)
// log2(2) = 1 > C = 0 -> O(n^1)