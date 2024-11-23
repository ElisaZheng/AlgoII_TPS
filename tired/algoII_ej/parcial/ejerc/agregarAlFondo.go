package ejerc

import "tdas/pila"

func AgregarAlFondo[T any](p pila.Pila[T], elem T) {
	if !p.EstaVacia() {
		valor := p.Desapilar()
		AgregarAlFondo(p, elem)
		p.Apilar(valor)
	} else {
		p.Apilar(elem)
	}
}

// la complejidad de la funcion va a ser O(n), porque tiene que desapilar todos los elementos que estan en la pila 
