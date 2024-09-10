package ejer

import (
	"math"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDAHash "tdas/diccionario"
	TDAPila "tdas/pila"
)

/*
Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.
*/

func Buscar_no_repetido(arr []int) int {
	return buscar_no_repetido(arr, arr[0])
}

func buscar_no_repetido(arr []int, k int) int {
	if len(arr) < 1 {
		return -1
	}
	medio := len(arr) / 2
	if arr[medio] != k {
		if arr[medio - 1] == k {
			return arr[medio]
		} else {
			return k
		}
	}
	izq := buscar_no_repetido(arr[:medio], k)
	if izq != -1 {
		return izq
	}
	return buscar_no_repetido(arr[medio + 1:], k)
}

/*
Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos del arreglo A
entre las posición [0;i] (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.
*/

func Mulplicar_primeros_k_maximos(arr []int, k int) []int {
	heap := TDAHeap.CrearHeap(func (a, b int) int { return a - b}) // Crear un heap minimo
	nuevo := make([]int, len(arr))
	mas_chico := 0
	mul := 1
	for i, elem := range arr {
		nuevo[i] = -1 
		mul *= elem
		heap.Encolar(elem)
		if mas_chico > elem {
			mas_chico = elem
		}
		if i == k - 1 {
			break
		}
	}
	if len(arr) < k {
		return nuevo
	}
	for i := k; i < len(arr); i ++ {
		minimo := heap.VerMax()
		heap.Encolar(arr[k])
		if heap.VerMax() == minimo {
			mul /= minimo
			mul *= arr[k]
		}
		nuevo[k] = mul
	}
	return nuevo
}

/*
4. Implementar un algoritmo que reciba dos cadenas (strings) y determine si son anagramas entre sí. Indicar y justificar
la complejidad del algortmo implementado.
*/

func Es_anagrama(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, elem := range s1 {
		if elem != rune(s2[len(s2) - 1 - i]) {
			return false
		}
	}
	return true
}

/*
5. Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves sean los
niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. Indicar y justificar la
complejidad del algoritmo implementado.
*/


//Implementar una primitiva que reciba un arbol binario izquierdista, y la cantidad de nodos que tiene, y devuelva el dato del elemento mas a
// abajo y a la derecha del arbol. Para que el ejercicio se pueda
// considerar como aprobable, debe resolverse en no  mas de O(n), 
//sin contar con otros errores. Para que se considere completamente bien, debe
// ejecutar en O(log n). Justif el orden del algoritmo implementado.

type nodoAbb[K comparable, V any] struct {
	clave     K
	dato      V
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
}


type ab[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      func(K, K) int
	cantidad int
}

func (arbol ab[K, V]) Buscar_el_mas_derecha(cantidad int) V {
	cantidad -= 1
	pila := TDAPila.CrearPilaDinamica[int]() // 0 signifca izq, 1 significa der
	for cantidad > 0 {
		if cantidad % 2 == 1 {
			pila.Apilar(0)
			cantidad -= 1
		} else {
			pila.Apilar(1)
			cantidad -= 2
		}
		cantidad /= 2
	}
	resul := arbol.raiz
	for !pila.EstaVacia() {
		if pila.Desapilar() == 0 {
			resul = resul.izquierdo
		} else {
			resul = resul.derecho
		}
	}
	return resul.dato
}
// la complejidad es O(log(n))


// Implementar un algoritmo que ordene un arreglo con puntos (valores (x, y)) que se encuentran dentro del circulo unitario
// (x² + y² < 1), sabiendo que la distribucion de los puntos es uniforme en dicho dominio. El criterio para ordenar es de
// menor a mayor norma (distancia al origen). Tener en cuenta que los numeros pueden tener "infinitos" decimales. El
// algoritimo debe ejecutar en tiempo lineal a la cantidad de elementos del arreglo a ordenar. 
//Justificar el orden del algoritmo propuesto (editado) 



type Puntos struct {
	x float64
	y float64
	distancia float64
}

func Ordenar_puntos(arr []Puntos) []Puntos {
	for _, punto := range arr {
		punto.distancia = math.Sqrt((punto.x * punto.y))
	}
	heap := TDAHeap.CrearHeapArr(arr, func (a, b Puntos) int {
		return int(a.distancia - b.distancia)
	}) // Crear un heap de minimo
	resul := make([]Puntos, len(arr))
	for i := range resul {
		resul[i] = heap.Desencolar()
	}
	return resul
}

/*
4. Implementar un algoritmo que dado un texto, devuelva cuál es la palabra más frecuente del mismo. Indicar y justificar
la complejidad del algoritmo implementado. Nota: recordar que existe la función split(cadena, separador), que
funciona en O(m), siendo m el largo de la cadena.
*/

func Buscar_palabras_mas_frecuencia(texto string) string {
	palabras := strings.Split(texto, " ") // O(n)
	disc := TDAHash.CrearHash[string, int]()
	contador := 0
	var palabra_mas_frecuencia string
	for _, palabra := range palabras {
		veces := 1
		if disc.Pertenece(palabra) {
			veces += disc.Obtener(palabra)
		}
		disc.Guardar(palabra, veces)
		if veces > contador {
			contador = veces
			palabra_mas_frecuencia = palabra
		}
	}
	return palabra_mas_frecuencia
}