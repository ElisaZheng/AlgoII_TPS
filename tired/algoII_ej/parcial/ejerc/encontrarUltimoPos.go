package ejerc
/*
Dado un array ordenado y un número N, 
implementar una función por división y conquista que permita encontrar la posición de la última aparición de N. 
Justificar el orden del algoritmo.
Ejemplo:
[1,2,2,2, 3,3, 4, 4,4,5] N = 2 →> 3
[1,2,2,2,3,3,4,4,4,5] N = 5 →> 9
[1,2,2,2,3,3,4,4,4,5] N = 6 →> -1
*/
func EncontrarUltimoPos(arr []int, N int) int {
	if len(arr) <= 1 {
		if arr[0] != N {
			return -1
		} else {
			return 0
		}
	}
	inicio := 0
	final := len(arr) - 1
	for inicio <= final {
		medio := (inicio + final) / 2
		sig := (medio + 1) % len(arr)
		if arr[medio] == N {
			if arr[sig] > N || inicio == len(arr)-1 {
				return medio
			} else {
				inicio = medio + 1
			}
		} else if arr[medio] < N {
			inicio = medio + 1
		} else {
			final = medio - 1
		}
	}
	return -1
}

/*
inicializar dos variables inicio y final 2*O(1)
iniciar un for {
	el arreglo se parte a la mitad,  si es valor en posi medio es menor, verifica la parte der,
	si es mayor, verifica la parte izq, si son igual, se verifica si es ultimo N, sino verifica la parte izq
	eso puedo usar T de Maestro T(n) = AT(n/B) + O(n ^ C)
	A : cantidad de llamados recusivos = 1
	B : Porporcion de tomano = 2
	C : el costo de partir y juntar = 0
	entonces : logB(A) = 0 = C -> O(log n) 
}
*/