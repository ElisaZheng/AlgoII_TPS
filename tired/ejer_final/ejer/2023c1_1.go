package ejer

import TDAHeap "tdas/cola_prioridad"

/*
Definimos a la reducción de una secuencia de números como la operación de sacar los dos elementos más pequeños de la
misma, y volver a guardar el resultado de 2 * mínimo + segundo_mínimo. Esto sólo puede hacerse si hay al menos
dos elementos. Ejemplo: [1, 7, 2, 3] -> [7, 3, 4]. De esta forma, la secuencia queda con un elemento menos.
Implementar un algoritmo que reciba un arreglo y devuelva el único valor que quedaría en el arreglo si aplicaramos
dicha reducción hasta que quede un único elemento en el arreglo. Indicar y justificar la complejidad del algoritmo.
*/

func Reducir_num(arr []int) int {
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int {return a - b}) // Crear un heap de minimo (O(n))
	for heap.Cantidad() > 1 {
		n1 := heap.Desencolar()
		n2 := heap.Desencolar()
		heap.Encolar(n1 * 2 + n2)
	}
	return heap.Desencolar()
}

/*
Así como en Go tenemos slices de arreglos, nos gustaría tener un slice sobre una lista enlazada. Es decir, poder tener
“una porción” de dicha lista. Por ejemplo, si yo hago slice := lista.Slice(5, 9) y luego hago slice.VerPrimero(),
dicha operación me devuelva el equivalente a haber iterado hasta la posición 5 en la original, y devolver ese elemento
(pero lo hace en tiempo constante). También podemos usar las primitivas que modifican la lista, obtener el largo (en
el ejemplo sería 4), usar el iterador interno y externo sobre dicho slice, así como crear otro slice para ese slice. Es
decir, Slice debe implementar Lista. Explicar cómo implementarías dicha estructura slice (campos, y cómo serían
las primitivas) para que esto funcione como se indica, de forma eficiente. No deben hacerse menciones a modificar
la lista enlazada (salvo, por supuesto, el agregado de la primitiva Slice), o su iterador externo, ya que esto no es
ni debe ser necesario. No es necesario implementar, si mencionar los puntos claves sobre referencias que se tendrían,
cómo funcionarían las primitivas que modifiquen la lista (tal que esta quede correctamente actualizada) y cómo sería la
iteración externa e interna. Se puede asumir que, tal como con los iteradores, no se pueden tener dos slices con uno
modificando la lista, y lo mismo con un slice y un iterador, ni tampoco usar las primitivas de modificación de la lista
mientras haya un slice en uso.
*/

type Slice [T any]struct {
	primero *nodo[T]
	ultimo *nodo[T]
	largo int
}

type nodo [T any]struct {
	dato T
	siguiente *nodo[T]
}

func (lista *Slice[T]) Slice(ini, fin int) Slice[T] {
	act := 0
	primero := lista.primero
	slice := Slice[T] {primero: nil, ultimo: nil, largo: 0}
	nodo_act := slice.primero
	for primero.siguiente != nil {
		if act >= ini && act <= fin {
			if nodo_act == nil {
				slice.primero = primero
			} else {
				nodo_act.siguiente = primero
			}
			primero = primero.siguiente
			nodo_act = nodo_act.siguiente
			slice.ultimo = nodo_act
		}
	}
	return slice
}

/*
Implementar un algoritmo que permita ordenar canciones, por su respectivo año de publicación, de forma eficiente,
tomando en cuenta canciones desde el año 1800. Indicar y justificar la complejidad del algoritmo (no se toman como
válidas respuestas parciales). Si quisiéramos considerar hasta incluso las canciones que escuchaban los dinosaurios, ¿el
algoritmo propusto seguiría siendo eficiente? Si lo es, justificar. Si no lo es, mencionar otro algoritmo que sea mejor.
*/

// Se pude usar radix sort

