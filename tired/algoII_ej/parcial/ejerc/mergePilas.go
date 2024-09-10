package ejerc

import pila "tdas/pila"

func MergePilas(pila1, pila2 pila.Pila[int]) []int {
	resultado := make([]int, 0)
	aux := pila.CrearPilaDinamica[int]()
	for !pila1.EstaVacia() || !pila2.EstaVacia() {
		elem1 := pila1.VerTope()
		elem2 := pila2.VerTope()
		if elem1 == elem2 {
			if !aux.EstaVacia() {
				if aux.VerTope() == elem1 {
					pila2.Desapilar()
					pila1.Desapilar()
					continue
				}
			}
			pila1.Desapilar()
		} else if elem1 > elem2 {
			aux.Apilar(pila2.Desapilar())
		} else {
			aux.Apilar(pila1.Desapilar())
		}
	}

	for !pila1.EstaVacia() {
		elem := pila1.Desapilar()
		if aux.EstaVacia() || aux.VerTope() != elem {
			aux.Apilar(elem)
		}
	}

	for !pila2.EstaVacia() {
		elem := pila2.Desapilar()
		if aux.EstaVacia() || aux.VerTope() != elem {
			aux.Apilar(elem)
		}
	}
	for !aux.EstaVacia(){
		resultado = append(resultado, aux.Desapilar())
	}
	return resultado
}