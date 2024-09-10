package ejer

import TDAHash "tdas/diccionario"
import Pila "tdas/pila"

/*
3. Implementar un algoritmo que reciba dos arreglos desordenados y determine si ambos arreglos tienen los mismos
elementos (y en mismas cantidades). Indicar y justificar la complejidad del algoritmo implementado.
*/

func Determinar_si_tienen_mismos_elem[T comparable](a1, a2 []T) bool {
	dic := TDAHash.CrearHash[T, int]()
	for _, elem := range a1 { // O(n)
		veces := 1
		if dic.Pertenece(elem) {
			veces += dic.Obtener(elem)
		}
		dic.Guardar(elem, veces)
	}

	for _, elem := range a2 { // O(n)
		if !dic.Pertenece(elem) {
			return false
		}
		veces := dic.Obtener(elem) - 1
		if veces == 0 {
			dic.Borrar(elem)
		} else {
			dic.Guardar(elem, veces)
		}
	}
	return true
}

/*
Un algoritmo iterativo sencillo para obtener la potencia de un número (b ^ n) tiene complejidad O(n). Tal como vieron en

el secundario (esperamos), sabemos que b ^ 2 = (b ^ n/2 ) ^ 2

. Utilizar esta propiedad para implementar un algoritmo que calcule , en tiempo O(log n). Justificar la complejidad del algoritmo implementado. Recordar tener cuidado con el caso que
n sea un valor impar.
*/

func PotenciaDeUnNum(n, potencia int) int {
	if potencia == 1 {
		return n
	}
	resultado := n
	if potencia % 2 == 1 {
		resultado *= n
		potencia -= 1
	}
	for potencia >= 2 {
		resultado *= resultado
		potencia /= 2
	}

	return resultado
}

/*
Implementar una función que dada una pila, determine si la misma se encuentra ordenada (es decir, se ingresaron los
elementos de menor a mayor). La pila debe quedar en el mismo estado al original al terminar la ejecución de la función.
Indicar y justificar la complejidad de la función.
*/

func Determinar_es_orde(p Pila.Pila[int]) bool {
	if p.EstaVacia() {
		return true
	}
	act := p.Desapilar()
	if p.EstaVacia() {
		p.Apilar(act)
		return true
	}
	if act < p.VerTope() {
		return false
	}
	if !Determinar_es_orde(p) {
		p.Apilar(act)
		return false
	}
	p.Apilar(act)
	return true
}