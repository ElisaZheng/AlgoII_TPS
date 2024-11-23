package rpl

func (arbol *arbol[K, V]) Reemplzar() {
	arbol.reemplazar()
}

//Va a revisar cada nodo, entonces es O(n)
func (arbol *arbol[K, V]) reemplzar() int {
	if arbol == nil { // O(1)
		return 0
	}
	cantidad_hijos := arbol.izq.reemplazar() + arbol.der.reemplazar()
	arbol.dato = cantidad_hijos
	return cantidad_hijos + 1
}