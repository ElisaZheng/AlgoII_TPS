package ejer


/*
Dados dos arreglos ordenados A y B, donde B tiene “un elemento menos que A”, implementar un algoritmo de división y
conquista que permita obtener el valor faltante de A en B. Ejemplo, si A = {2, 4, 6, 8, 9, 10, 12} y B = {2, 4,
6, 8, 10, 12}, entonces la salida del algoritmo debe ser o bien la posición 4, o el valor 9 (lo que decidan que devuelva).
Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

func Buscar_faltante(a, b []int) int {
	inicio := 0
	fin := len(b) - 1
	for inicio < fin {
		medio := (inicio + fin) / 2 
		if b[medio] == a[medio] {
			inicio = medio + 1
		} else if b[medio] != a[medio] {
			fin = medio
		}
	}
	return a[inicio]
}

/*
Implementar una primitiva para el heap func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T])
ColaPrioridad[T], que reciba otro Heap y cree un nuevo Heap con los elementos del primero que no se
ecuentren en el segundo, y viceversa (es decir, la diferencia simétrica entre ambos). La función de comparación del
nuevo heap debe ser la del primer heap. Indicar y justificar la complejidad del algoritmo implementado.
*/

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

const (
	capInicial        int = 10
	factorRedimension int = 2
	factorProporcion  int = 2 * factorRedimension
)


func CrearHeap[T any](funcion_cmp func(T, T) int) heap[T] {
	return heap[T]{datos: make([]T, capInicial), cmp: funcion_cmp}
}

func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T]) heap[T] {
	h := CrearHeap[T](heap.cmp)
	for _, elem := range otro.datos {
		if !heap.buscar_elem(elem, heap.datos, 0) {
			h.Encolar(elem)
		}
	}
	for _, elem := range heap.datos {
		if !heap.buscar_elem(elem, otro.datos, 0) {
			h.Encolar(elem)
		}
	}
	return h
}

func upheap[T any](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	posPadre := (pos - 1) / 2
	if cmp(arr[posPadre], arr[pos]) < 0 {
		swap(&arr[posPadre], &arr[pos])
		upheap(arr, posPadre, cmp)
	}
}


func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func (heap *heap[T]) Encolar(elemento T) {
	heap.datos[heap.cant] = elemento
	upheap(heap.datos, heap.cant, heap.cmp)
	heap.cant++
}

func (heap *heap[T]) buscar_elem(elem T, d []T, indice int) bool {
	if heap.cmp(d[indice], elem) == 0 {
		return true
	}
	h1 := 2 * indice + 1
	h2 := 2 * indice + 2
	if h2 > len(d) - 1 {
		if h1 > len(d) - 1 {
			return false
		} else {
			if heap.cmp(d[indice], elem) == 0{
				return true
			} else {
				return false
			}
		}
	}
	if heap.cmp(d[h1], elem) == 0 || heap.cmp(d[h2], elem) == 0 {
		return true
	} else if heap.cmp(elem, d[h1]) < 1 || heap.cmp(elem, d[h2]) < 1 {
		return false
	} else if heap.cmp(elem, d[h1]) < 1 {
		return heap.buscar_elem(elem, d, h2)
	} 
	return heap.buscar_elem(elem, d, h1)
}

func Determinar_si_mas_de_mitad(arr[] int) bool {
	if len(arr) == 0 {
		return false
	}
	medio := len(arr) / 2
	valor_posible := arr[medio]
	prim := buscar_prim(arr[:medio + 1], valor_posible, 0, medio)
	ult := buscar_ultimo(arr, valor_posible, 0, len(arr) - 1)
	return (ult - prim + 1) > len(arr) / 2
}

func buscar_prim(arr[] int, n, ini, fin int) int {
	if ini == fin {
		return ini
	}
	medio := (ini + fin) / 2
	if arr[medio] < n {
		return buscar_prim(arr, n, medio + 1, fin)
	}
	return buscar_prim(arr, n, ini, medio)
}

func buscar_ultimo(arr[] int, n, ini, fin int) int {
	if ini == fin {
		return ini
	}
	medio := (ini + fin) / 2
	if arr[medio] > n {
		return buscar_ultimo(arr, n, ini, medio - 1)
	}
	return buscar_ultimo(arr, n, medio, fin)
}