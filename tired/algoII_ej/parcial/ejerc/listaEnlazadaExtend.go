package ejerc


type nodo[T any] struct {
    valor T
    siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
    primero *nodo[T]
    ultimo *nodo[T]
	largo int
}

type ListaEnlazada[T any] interface {

}

func CrearListaEnlazada[T any]() ListaEnlazada[T] {
	return &listaEnlazada[T]{primero :nil, ultimo: nil}
}


func (l *listaEnlazada[T]) Extender(otra *listaEnlazada[T]) {
	actual := l.primero
	for actual != nil && actual.siguiente != nil {
		actual = actual.siguiente
	}
	actual.siguiente = otra.primero
}
