package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "una lista vacía no tiene primero para eliminar")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "una lista vacía no tiene primero para ver")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "una lista vacía no tiene último para ver")
}

func TestInsertarPrimero(t *testing.T) {
	{
		enteros := TDALista.CrearListaEnlazada[int]()

		enteros.InsertarPrimero(8)
		require.Equal(t, 1, enteros.Largo())
		require.False(t, enteros.EstaVacia())
		require.Equal(t, 8, enteros.VerPrimero())
		require.Equal(t, 8, enteros.VerUltimo())

		enteros.InsertarPrimero(10)
		enteros.InsertarPrimero(14)
		require.Equal(t, 3, enteros.Largo())
		require.Equal(t, 14, enteros.VerPrimero())
		require.Equal(t, 8, enteros.VerUltimo())
	}
	{
		cadenas := TDALista.CrearListaEnlazada[string]()

		cadenas.InsertarPrimero("datos")
		require.Equal(t, 1, cadenas.Largo())
		require.False(t, cadenas.EstaVacia())
		require.Equal(t, "datos", cadenas.VerPrimero())
		require.Equal(t, "datos", cadenas.VerUltimo())

		cadenas.InsertarPrimero(" y estructuras de ")
		cadenas.InsertarPrimero("algoritmos")
		require.Equal(t, 3, cadenas.Largo())
		require.Equal(t, "algoritmos", cadenas.BorrarPrimero())
		require.Equal(t, " y estructuras de ", cadenas.VerPrimero())
	}
}

func TestInsertarUltimo(t *testing.T) {
	flotantes := TDALista.CrearListaEnlazada[float64]()

	flotantes.InsertarUltimo(8.14)
	require.Equal(t, 1, flotantes.Largo())
	require.False(t, flotantes.EstaVacia())
	require.Equal(t, 8.14, flotantes.VerPrimero())
	require.Equal(t, 8.14, flotantes.VerUltimo())

	flotantes.InsertarUltimo(3.152)
	require.Equal(t, 2, flotantes.Largo())
	require.Equal(t, 3.152, flotantes.VerUltimo())
	require.Equal(t, 8.14, flotantes.VerPrimero())
}

func TestBorrarPrimero(t *testing.T) {
	numeros := TDALista.CrearListaEnlazada[int]()

	require.PanicsWithValue(t, "La lista esta vacia", func() { numeros.BorrarPrimero() }, "una lista vacía no tiene primero para eliminar")

	numeros.InsertarPrimero(8)
	numeros.InsertarPrimero(15)
	numeros.InsertarUltimo(16)
	numeros.InsertarUltimo(18)

	// [15 8 16 18]
	require.Equal(t, 4, numeros.Largo())
	require.Equal(t, 15, numeros.BorrarPrimero())
	require.Equal(t, 8, numeros.BorrarPrimero())
	require.Equal(t, 16, numeros.BorrarPrimero())
	require.Equal(t, 1, numeros.Largo())

	require.Equal(t, 18, numeros.VerPrimero())
	require.Equal(t, 18, numeros.VerUltimo())

	require.Equal(t, 18, numeros.BorrarPrimero())

	require.Equal(t, 0, numeros.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { numeros.BorrarPrimero() }, "una lista vacía no tiene primero para eliminar")
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 10000; i++ {
		lista.InsertarUltimo(i)
	}
	for i := 0; i <= 10000; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
}

func TestBorrarTodo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("tipo")
	lista.InsertarUltimo("de")
	lista.InsertarUltimo("dato")
	lista.InsertarUltimo("abstracto")

	for !lista.EstaVacia() {
		lista.BorrarPrimero()
	}

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "una lista vacía no tiene primero para eliminar")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "una lista vacía no tiene primero para ver")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "una lista vacía no tiene último para ver")
}

func TestBorrarTodoEInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}

	for !lista.EstaVacia() {
		lista.BorrarPrimero()
	}

	lista.InsertarPrimero(20)
	lista.InsertarPrimero(22)

	require.Equal(t, 22, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
}

func TestIteradorInternoSumar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}

	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 55, suma)

	sumaConCorte := 0
	lista.Iterar(func(v int) bool {
		if v >= 5 {
			sumaConCorte += v
			return true
		} else {
			return false
		}
	})
	require.Equal(t, 45, sumaConCorte)

	sumaPares := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			sumaPares += v
		}
		return true
	})
	require.Equal(t, 30, sumaPares)
}

func TestIteradorInternoIterarTodaLaLista(t *testing.T) {
	{
		lista := TDALista.CrearListaEnlazada[int]()

		for i := 1; i <= 10; i++ {
			lista.InsertarPrimero(i)
		}

		i := 10
		lista.Iterar(func(v int) bool {
			require.Equal(t, i, v)
			i--
			return true
		})
	}
	{
		lista := TDALista.CrearListaEnlazada[string]()
		arr := []string{"iterador", "tda", "interno", "lista"}
		for i := range 4 {
			lista.InsertarUltimo(arr[i])
		}

		res1 := []string{}
		res2 := []string{}
		cont1, cont2, indice := 0, 0, 0

		lista.Iterar(func(v string) bool {
			if cont1 < 2 && indice%2 == 0 {
				res1 = append(res1, arr[indice])
				indice++
				cont1++
				return true
			} else if cont2 < 2 {
				res2 = append(res2, arr[indice])
				indice++
				cont2++
				return true
			}
			return false
		})

		require.Equal(t, 2, len(res1))
		require.Equal(t, "iterador", res1[0])
		require.Equal(t, "interno", res1[1])

		require.Equal(t, 2, len(res2))
		require.Equal(t, "tda", res2[0])
		require.Equal(t, "lista", res2[1])
	}
}

func TestIterInsertarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(5)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
}

func TestInterInsertarAlPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 2, 3, 4, 5}
	valor := 0
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	iterador.Insertar(valor)
	require.Equal(t, valor, lista.VerPrimero())

	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestInterInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 2, 3, 4, 5, 6}
	valor := 6
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(valor)
	require.Equal(t, valor, lista.VerUltimo())

	iterador = lista.Iterador()
	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestInterInsertarAlMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 2, 100, 3, 4, 5}
	valor := 100
	pos := 0
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if pos == 3 {
			iterador.Insertar(valor)
		}
		pos++
		iterador.Siguiente()
	}

	iterador = lista.Iterador()
	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestIterInsertarMuchos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 10000; i >= 0; i-- {
		lista.InsertarPrimero(i)
	}
	iterador := lista.Iterador()
	for i := 0; i <= 10000; i++ {
		require.Equal(t, i, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestInterBorrarElPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{1, 2, 3, 4, 5}
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	require.Equal(t, 0, iterador.Borrar())
	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
	require.Equal(t, 1, lista.VerPrimero())
}

func TestInterBorrarElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 3, 4, 5}
	pos := 0
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		if pos == 2 {
			valor := iterador.Borrar()
			require.Equal(t, 2, valor)
		}
		pos++
		iterador.Siguiente()
	}

	iterador = lista.Iterador()
	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestBorrarElUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 2, 3, 4}
	pos := 0
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		if pos == 5 {
			dato := iterador.Borrar()
			require.Equal(t, 5, dato)
		}
		iterador.Siguiente()
	}
	require.False(t, iterador.HaySiguiente())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia())

	iterador = lista.Iterador()

	for _, valor := range valores {
		require.Equal(t, valor, iterador.VerActual())
		iterador.Siguiente()
	}
}

func TestIteradorBorraAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{0, 1, 2, 3, 4, 5}
	pos := 0
	for i := 0; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		require.Equal(t, valores[pos], iterador.Borrar())
		pos++
	}
	require.False(t, iterador.HaySiguiente())
	require.True(t, lista.EstaVacia())
}

func TestIteradorBorraListaConUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(0)
	iterador := lista.Iterador()

	require.Equal(t, 0, iterador.Borrar())
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
}

func TestIterInsertarYBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	for i := 5; i >= 0; i-- {
		iterador.Insertar(i)
	}
	for iterador.HaySiguiente() {
		iterador.Borrar()
	}
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(0)
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.False(t, lista.EstaVacia())
}
