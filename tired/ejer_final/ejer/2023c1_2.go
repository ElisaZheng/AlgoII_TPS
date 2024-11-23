package ejer

import TDAHash "tdas/diccionario"

/*
¿Qué implica que un algoritmo de ordenamiento sea estable? Explicar detalladamente por qué el algoritmo de
ordenamiento auxiliar de RadixSort debe ser sí o sí estable. Dar algún ejemplo en el que se evidencie esta necesidad.
*/

// Es estable: significa cuando ordena los elemenros, se mantienen las posiciones relativas

/*
Implementar en Go una primitiva para la lista enlazada Filter(func(T) bool) Lista[T] que reciba una función
por parámetro, y devuelva una lista nueva que contenga los elementos de la original para los cuales la función pasada
por parámetro devuelve true. El orden relativo de los elementos debe ser el mismo que en la lista original. La lista
original debe quedar en el mismo estado que el original. Indicar y justificar la complejidad de la primitiva.
*/

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodo[T any](elem T) *nodoLista[T] {
	return &nodoLista[T]{dato: elem, siguiente: nil}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func (lista listaEnlazada[T]) Filter(func(T), bool) listaEnlazada[T] {
	nuevo := listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
	act := lista.primero
	nuevo_act := nuevo.primero
	for act != nil {
		if nuevo_act == nil {
			nuevo.primero = act
		} else {
			nuevo_act.siguiente = act
			nuevo_act = nuevo_act.siguiente
		}
		nuevo.ultimo = nuevo_act
		nuevo.largo++
		act = act.siguiente
	}
	return nuevo
}

/*
Implementar un algoritmo que cuente la mínima cantidad de elementos a eliminar de un arreglo para que todos los
elementos sean iguales. Indicar y justificar la complejidad de la función.
*/

func Eliminar_minima_cantidad[T comparable](arr []T) int {
	dic := TDAHash.CrearHash[T, int]()
	maximo := 0
	for _, elem := range arr {
		cant := 1
		if dic.Pertenece(elem) {
			cant += dic.Obtener(elem)
		}
		dic.Guardar(elem, cant)
		if cant > maximo {
			maximo = cant
		}
	}
	return len(arr) - maximo
}

/*
Sabemos que, en Go, concatenar 2 strings (str1 + str2) es O(n + m) siendo n y m el largo de las cadenas. También
sabemos que podemos obtener el array de caracteres (runes) que corresponde a un string ([]rune(str)) en tiempo
lineal, y podemos luego acceder a una posición de un arreglo en tiempo constante. Asimismo, podemos obtener
nuevamente el string en tiempo lineal (string(arrRunes)).
Queremos implementar un algoritmo que reciba un arreglo de k strings, y devuelva un string con todo junto, sin
separaciones. Alan implementó el algoritmo del dorso, que efectivamente devuelve lo pedido. Bárbara lo utilizó, y
demoraba demasiado, y le pidió “amablemente” que lo corrija. Indicar cuál es el problema del algoritmo, y cómo lo
corregirías para que funcione en un tiempo acorde. Para tu análisis podés considerar que todas las cadenas son del
mismo largo m (→ caracteres del arreglo final es n = k · m). Aclaración: no usar strings.Join, por supuesto.
*/

func Concatener_strings(s1 []string) string {
	res := make([]rune, len(s1) * len(s1[0]))
	indice := 0
	for _, i := range s1 {
		for _, letra := range i {
			res[indice] = rune(letra)
			indice ++
		}
	}
	return string(res)
}

/*
Existe una estructura llamada dequeue (Double-Ended Queue), que es como una pila y una cola en simultáneo: permite
insertar al principio y al final, y eliminar tanto al principio como al final. Todas esas operaciones, en O(1). ¿Cómo
implementarías dicha estructura? Definir detalladamente.
*/

// Puede definirlo como una lista