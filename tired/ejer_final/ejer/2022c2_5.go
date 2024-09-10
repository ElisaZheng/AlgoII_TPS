package ejer

import TDALista "tdas/lista"

/*
Una lista que representa un número es una lista en la que en cada lugar tiene un dígito, y dicho número se lee lugar a
lugar. Por ejemplo, la lista 1 -> 8 -> 1 -> 2 representa al número 1812. Implementar una función que reciba dos
listas que represeten números, y devuelva una lista que represente a la suma de estos. Por ejemplo, si recibe las listas
9 -> 1 -> 5 y 9 -> 6 debe devolver la lista 1 -> 0 -> 1 -> 1. Las listas recibidas por parámetro no deben verse
modificadas. Indicar y justificar la complejidad del algoritmo implementado.
*/

func Buscar_Suma_de_listas(l1, l2 TDALista.Lista[int]) TDALista.Lista[int] {
	res := TDALista.CrearListaEnlazada[int]()
	distancia := 0
	mas_largo := res
	if l1.Largo() > l2.Largo() {
		distancia = l1.Largo() - l2.Largo()
		mas_largo = l1
	} else {
		distancia = l2.Largo() - l1.Largo()
		mas_largo = l2
	}
	i1 := l1.Iterador()
	i2 := l2.Iterador()
	for i1.HaySiguiente() && i2.HaySiguiente() {
		if distancia != 0 {
			if mas_largo == l1 {
				res.InsertarUltimo(i1.VerActual())
				i1.Siguiente()
			} else {
				res.InsertarUltimo(i2.VerActual())
				i2.Siguiente()
			}
			distancia --
		} else {
			res.InsertarUltimo(i1.VerActual() + i2.VerActual())
			i1.Siguiente()
			i2.Siguiente()
		}
	}
	return res
}

