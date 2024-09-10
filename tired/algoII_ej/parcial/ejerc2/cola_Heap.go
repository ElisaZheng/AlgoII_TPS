package ejerc2

type colaSobreHeap[T any] struct {
    primero *nodo[T]
	ultimo *nodo[T]
	cantidad int
}

type nodo[T any] struct {
	dato T
	siguiente *nodo[T]
}

func CrearColaSobreColaPrioridad[T any]() ColaPrioridad[T] {
	return &colaSobreHeap[T]{primero: nil, cantidad: 0}
}

func crearNodo[T any] (dato T) *nodo[T] {
	return &nodo[T]{dato: dato, siguiente: nil}
}

func (cola *colaSobreHeap[T]) Encolar(dato T) {
	nuevo := crearNodo(dato)
	if cola.primero == nil {
		cola.primero = nuevo
	} else {
		cola.ultimo.siguiente = nuevo
	}
	cola.ultimo = nuevo
	cola.cantidad ++
}

func (cola *colaSobreHeap[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic ("la cola esta vacia")
	} 
	cola.cantidad --
	valor := cola.primero.dato
	cola.primero = cola.primero.siguiente
	return valor
}

func (cola *colaSobreHeap[T]) EstaVacia() bool {
    return cola.ultimo == nil
}

func (cola *colaSobreHeap[T]) VerMax() T {
	if cola.EstaVacia() {
		panic ("la cola esta vacia")
	} 
	valor := cola.primero.dato
	return valor
}

func (cola *colaSobreHeap[T]) Cantidad() int {
	return cola.cantidad
}
