package lista

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

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		anteriorPrimero := lista.primero
		lista.primero = nuevoNodo
		nuevoNodo.siguiente = anteriorPrimero
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
		lista.ultimo = nuevoNodo
	} else {
		anteriorUltimo := lista.ultimo
		anteriorUltimo.siguiente = nuevoNodo
		lista.ultimo = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	datoPrimero := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return datoPrimero
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{lista: lista, actual: lista.primero, anterior: nil}
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func (i *iteradorLista[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.dato
}

func (i *iteradorLista[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iteradorLista[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)
	if i.anterior != nil {
		i.anterior.siguiente = nuevoNodo
		nuevoNodo.siguiente = i.actual
		if i.actual == nil {
			i.lista.ultimo = nuevoNodo
		}
	} else {
		nuevoNodo.siguiente = i.actual
		if i.lista.primero == i.lista.ultimo {
			i.lista.ultimo = nuevoNodo
		}
		i.lista.primero = nuevoNodo
	}
	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iteradorLista[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := i.actual.dato
	if i.anterior != nil {
		i.anterior.siguiente = i.actual.siguiente
		if i.actual.siguiente == nil {
			i.lista.ultimo = i.anterior
			i.actual = i.anterior
		} else {
			i.actual = i.anterior.siguiente
		}
	} else {
		i.lista.primero = i.actual.siguiente
		i.actual = i.lista.primero
	}
	i.lista.largo--
	return dato
}
