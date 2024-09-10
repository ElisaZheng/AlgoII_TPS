package ejer

import TDAHeap "tdas/cola_prioridad"
import TDAHash "tdas/diccionario"

/*
Se quiere implementar una búsqueda similar a la Búsqueda Binaria: la Búsqueda Ternaria. Esta, en vez de partir el
arreglo en 2, parte el arreglo en 3 tercios. Verifica si el elemento buscado está en la posición del primer tercio, así como
también en la posición del segundo tercio (en vez de únicamente en la mitad, como hace Búsqueda binaria). Cuando
no se trata de ninguno de estos, llama recursivamente para el segmento que corresponda (el primer tercio, segundo o
tercero, según cómo sea el elemento buscado respecto al elemento del primer tercio y del segundo).
a. Determinar y justificar el orden de la Búsqueda Ternaria.
b. Si en vez de dividir en 3 partes, ahora decidiéramos dividir en n partes (siendo n el tamaño de arreglo), ¿cuál sería
la complejidad del algoritmo? ¿A qué algoritmo se asemeja dicha implementación?
c. Dado los resultados anteriores, ¿tiene sentido implementar la búsqueda K-aria, para k > 2? Justificar.
*/

/*
a. La complejidad segun la teorema de maestro :
	A: la cantidad de llama cada vez -> 2 o 1
	B: la cantidad de porcion -> 3
	C: 0 (supongamos)
	log b (a) = log 3(2) > 0   o  log 3(1) == 0
	=>  O(log n) o O(n^log b (a))
b. T(n)=T(n/n)+O(n) = O(n)
c. La búsqueda binaria ya es óptima con una complejidad  O(logn).

*/

////////////////////////////////////////////////////////////////////////////


/*
Implementar un algoritmo que reciba un Grafo con características de árbol (no árbol binario, sino referido a árbol
de teoría de grafos) y devuelva una lista con los puntos de articulación de dicho árbol. Indicar y justificar la complejidad
del algoritmo implementado. Importante: aprovechar las características del grafo que se recibe para que la solución sea
lo más simple posible.

caracteristica:
1. es aciclo y conexo
2. |V| = |E - 1|

def puntos_articulos_arbol(grafo) :
	puntos = set()
	for v in grafo.obtener_vertices():
		if len(grafo.adyacentes(v)) > 1 :
			puntos.add(v)
	return puntos
*/

////////////////////////////////////////////////////////////////////////////

/*
Carlos es nuevo en la empresa en la que trabajan Alan y Bárbara. Alan va a ser el mentor de Carlos, quien debe
implementar un nuevo TDA Gatito. Alan, revisando el trabajo que hizo Carlos, nota que este agregó una primitiva
Redimensionar, pública en la interfaz Gatito, para que la use Bárbara. Alan lo increpa a Carlos, preguntando para qué
es dicha primitiva, y este le contesta “Tal como dice la documentación, es para que Bárbara me diga cómo redimensionar
el arreglo de pelos que tiene el gatito”. Alan, que conoce bien el temperamento de Bárbara, decide evitar que echen a
Carlos en su segunda semana de trabajo. En este ejercicio, te toca hacer de Alan.
Escribir una explicación de por qué esto que está haciendo Carlos está mal. Considerá que Carlos es muy testarudo
(incluso, a pesar de su propio bien), así que tu argumentación deberá ser muy clara y contundente.
*/

/*
Al exponer Redimensionar, estás delegando esta responsabilidad a los usuarios del TDA, 
quienes podrían no tener el contexto completo o la intención adecuada para realizar esta operación correctamente.
Esto puede llevar a fallos en el comportamiento del TDA y hacer que sea menos robusto.
*/

////////////////////////////////////////////////////////////////////////////

/*
Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede
formar con dichos dígitos.
*/

func BuscarNumMasGrande(arr []int) int {
	// Crear un heap de maximo
	heap := TDAHeap.CrearHeapArr(arr, func(n1, n2 int) int {return n1 - n2}) 
	num := 0
	for !heap.EstaVacia(){
		num *= 10
		num += heap.Desencolar()
	}
	return num
}

/*
Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un número K y determinar en
O(n) si existe un par de elementos en el arreglo que sumen exactamente K.
*/

func SumarEsK(arr []int, k int) bool {
	dic := TDAHash.CrearHash[int, int]()
	for _, n := range arr {
		valor_necesita := k - n
		if dic.Pertenece(valor_necesita) {
			return true
		}
		dic.Guardar(n, n)
	}
	return false
}