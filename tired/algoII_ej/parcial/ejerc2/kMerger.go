package ejerc2


/*func KMerge(arr [][]int) []int {
    resultado := make([]int, len(arr) * len(arr[0]))
    vector := make([]compuesto, len(arr))
    for i := range arr {
        vector[i] = CrearCompuesto(i)
    }
    valores := make([]int, len(arr))
    for i := range valores {
        valores[i] = vector[i].valor 
    }
    //cola := heap.CrearHeap(valores, func (n1, n2 int) int {return n1 - n2})
    for i := 0; i < len(arr) * len(arr[0]); i ++ {
		//valor := cola.Desencolar()
        resultado[i] = valor
        //actualizar(vector, cola, valor)
    }
   return resultado   
}

type compuesto struct {
   vector int
   posicion int
   valor  int
}

func CrearCompuesto(vector int) compuesto {
    return compuesto{vector: vector, posicion : 0, valor :0}
}

func actualizar[T any](vector []compuesto, cola ColaPrioridad[int], valor int) {
    i := 0
    for i := range vector {
        if vector[i].valor ==  valor {
            vector[i].posicion += 1
            i = i
            break
        }
    }
    cola.Encolar(vector[i].posicion)
}


type ColaPrioridad[T any] interface {

	// EstaVacia devuelve true si la la cola se encuentra vacía, false en caso contrario.
	EstaVacia() bool

	// Encolar Agrega un elemento al heap.
	Encolar(T)

	// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con un mensaje
	// "La cola esta vacia".
	VerMax() T

	// Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con un
	// mensaje "La cola esta vacia"
	Desencolar() T

	// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
	Cantidad() int
}
*/