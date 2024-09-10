package ejer

import "math"
import TDAHash "tdas/diccionario"

/*
Implementar en Go una primitiva Invertir() para el Heap, que haga que el heap se comporte con la función de
comparación contraria a la que venía utilizando hasta ese momento. El heap debe quedar en estado correcto para que
las operaciones siguientes sean válidas considerando la función de comparación, ahora invertida. No se puede modificar
la estructura del heap para implementar esta primitiva. Es decir, La implementación actual que tienen del heap debería
poder trabajar con esta primitiva, sin tener que modificar ninguna primitiva ni función auxiliar. Indicar y justificar la
complejidad del algoritmo implementado. Bajo estas condiciones, Si se llamara k veces a esta nueva primitiva, ¿podría
esto afectar a la complejidad de otras primitivas, como Encolar y Desencolar? Justificar.
*/

type HeapInvertido[T any] struct {
	heap []int
	cmp func (a, b int) int
}

//func (heap *Heap[T]) Invertir() {

/*
Se tiene un arreglo de n elementos, ordenado, cuyos valores van de 0 a log(n). Implementar un algoritmo que permita
obtener la frecuencia de todos los números entre 0 y log(n) en tiempo menor a O(n). Justificar el orden del algoritmo.
*/

func Obtener_frecuencia(arr []int) []int {
	resultados := make([]int, int(math.Log(float64(len(arr)))))
	for indice := range resultados { // O(log n)
		primer := buscar_pri_indice(arr, 0, len(arr) - 1, indice)
		ultimo := buscar_ult_indice(arr, 0, len(arr) - 1, indice)
		resultados[indice] = (ultimo - primer)
	}
	return resultados
}

func buscar_pri_indice(arr []int, ini, fin, n int) int { // O(log n)
	if ini == fin {return ini}
	medio := (ini + fin) / 2
	if arr[medio] >= n {
		return buscar_pri_indice(arr, ini, medio, n)
	}
	return buscar_pri_indice(arr, medio + 1, fin, n)
}


func buscar_ult_indice(arr []int, ini, fin, n int) int {// O(log n)
	if ini == fin {return ini}
	medio := (ini + fin) / 2
	if arr[medio] <= n {
		return buscar_pri_indice(arr, ini, medio, n)
	}
	return buscar_pri_indice(arr, medio + 1, fin, n)
}

// O(log^2 n)

/*
Bárbara está trabajando para una marca de ropa. Está analizando el inventario, y nota que los proveedores le trajeron
los guantes por unidades en vez de a pares, como sería lo esperable. El problema es que ahora ni siquiera sabe si van
a poder venderlos todos. Cuenta ahora con un arreglo de guantes, y necesita saber si para cada color puede formar
pares de guantes para poder venderlos (asumir que son todos de talle único). Implementar una función que reciba un
arreglo de guantes, y devuelva true si no queda ningún guante suelto sin formar pareja, o false en caso contrario.
No se pueden juntar guantes de colores diferentes. Suponer que cada guante es un struct con un campo color, que es
un enumerativo definido en algún lado. Indicar y justificar la complejidad de la función implementada. Si los guantes
pudieran ser de diferentes talles, ¿cómo modificarías el algoritmo implementado para que resuelva el problema (para
vender una pareja de guantes deben coincidir en talle y color)? ¿Cambiaría la complejidad del algoritmo?
*/

type Guantes struct {
	colores string
	taller int
}

func Es_pares(arr []Guantes) bool {
	dic := TDAHash.CrearHash[Guantes, int]()
	for _, guante := range arr {
		if dic.Pertenece(guante) {
			dic.Borrar(guante)
		} else {
			dic.Guardar(guante, 1)
		}
	}
	return dic.Cantidad() == 0
}

/*
Escribir un algoritmo que permita obtener los puntos de articulación de un grafo no dirigido. Indicar y justificar la
complejidad del mismo. Aplicar el algoritmo al grafo del dorso, comenzando desde el vértice G.
*/

/*def puntos_articulacion(grafo, g):
	visitado = set()
	padres = {}
	puntos = []
	_puntos_articulacion(grafo, g, visitado, padres, {}, {}, puntos, True)
	
def _puntos_articulacion(grafo, v, visitado, padres, mas_bajo, orden, puntos, es_raiz):
	visitado.add(v)
	mas_bajo[v] = orden[v]
	hijos = 0
	for w in grafo.adyacentes():
		if w not in visitado:
			hijos += 1
			orden[w] = orden[v] + 1
			padres[w] = v
			_puntos_articulacion(grafo, w, visitado, padres, mas_bajo, orden, False)
			if mas_bajo[w] >= mas_bajo[v] and not es_raiz:
				puntos.append(v)
			mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
		else if padres[v] != w:
			mas_bajo[v] = min(mas_bajo[v], orden[w])
	if es_raiz and hijos >= 2:
		puntos.append(v)

/*
En clase vimos que se puede implementar un heap con un árbol izquierdista, o su representación equivalente en arreglo.
Esta última, siendo mucho más sencilla de implementar. ¿Por qué no implementamos también el Árbol Binario de
Búsqueda con una representación en arreglo, en vez de implementarlo con, valga la redundancia, árboles?
*/
