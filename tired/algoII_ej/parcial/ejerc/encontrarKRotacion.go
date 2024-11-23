package ejerc

/*3. A un array ordenado se lo rotó k posiciones.
 Implementar una función por división y conquista que permita encontrar la cantidad k de rotaciones que se le aplicó al array. 
 Justificar el orden del algoritmo. Ejemplo:
[1,2,3,4,5,6] → 0 rotaciones
[6,1,2,3,4,5] -> 1 rotación
[3,4,5,6,1,2] -> 4 rotaciones
*/

func EncontrarKRotacion(arr []int) int {
	contador := 0
	if len(arr) <= 1 {
		return contador
	}
	inicio := 0
	final := len(arr) -1
	for inicio < final {
		medio := (inicio + final) / 2
		sig := medio + 1
		if arr[medio] > arr[sig] {
			return medio
		}
		if arr[inicio] < arr[medio] {
			inicio = medio + 1
		} else {
			final = medio - 1
		}
	}
	return 0
}

/*
entrar a bucle "for" hasta inicio >= final 
lo cual ocurre aproximadamente logarimicamente al tomano de la entrada(busqueda binaria) -> O(log N)
*/