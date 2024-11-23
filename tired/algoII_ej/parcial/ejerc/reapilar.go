package ejerc

import (
	pila "tdas/pila"
)

func Reapilar[T any](p pila.Pila[T]) {
	valor := p.Desapilar()
	if !p.EstaVacia() {
		 _reapilar(p)
	}
	p.Apilar(valor)
	
}

func _reapilar[T any](p pila.Pila[T]) *T {
	if p.EstaVacia() {
		return nil 
	}
	act := p.Desapilar() 
	siguiente := _reapilar(p)
	if siguiente == nil {
		return &act
	} else {
		p.Apilar(act)
		p.Apilar(*siguiente)
	}
	return nil
} 