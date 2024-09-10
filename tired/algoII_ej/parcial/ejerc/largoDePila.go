package ejerc

import "tdas/pila"

func LargoDePila[T any] (p pila.Pila[T]) int {
	contador := 0
	if p.EstaVacia() {
		return contador
	}
	valor := p.Desapilar()
	contador = 1 + LargoDePila(p)
	p.Apilar(valor)
	return contador
}

/*
•{Orden del algoritmo}•
iniciarlizar variable contador -> O(1)
llamada recusiva hasta la pila esta vacia
ver si la pila esta vacia, si no desapilar, contador + 1 -> O(1)
ingresar a la llamada resursiva para desapilar hasta la pila esta vacia , luego apilar los valores (va a hacer n veces) -> O(n)
*/

/*
•{Orden del algoritmo}•
(Para un n, siendo n la cantidad de elementos en la pila)
Inicializar variable len -> O(1)
Primer while -> se ejecuta n veces
    Ver si la pila está vacía, apilar, desapilar e incrementar a len -> O(1)
    ->> n.O(1) ->> O(n)
Segundo while -> se ejecuta n veces
    Ver si la pila está vacía, apilar y desapilar -> O(1)
    ->> n.O(1) ->> O(n)

Total =>   T(n) = O(1) + O(1) + O(n) + O(n)
O(1) << O(n) -> = 2.O(n)
Finalmente => T(n) = O(n) -> Lineal
*/

