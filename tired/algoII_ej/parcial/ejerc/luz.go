package ejerc

/*
Se tiene un arreglo con las distancias en años luz entre la Tierra y distintos cuerpos celestes en la Via Lactea y
otras galaxias ceranas descubiertos hasta el momento. Debido a un error cometido por el sysadmin del Planetario,
estas distancias hoy están completamente desordenadas. 
Esta persona necesita volver a ordenar los registros antes de que descubran el error para no perder su trabajo. 
Como no cursó Algoritmos 2 no sabe cómo hacer, por lo que nos pide ayuda a nosotros. 
Diseñar un algoritmo de ordenamiento que pueda ordenar 
un arreglo de CuerpoCeleste donde cada uno tiene un campo distancia int que funcione lo más rápido posible. 
Indicar y justificar el orden del algoritmo.
De los datos, se sabe lo siguiente:
• Hay alrededor de 500 mil de registros
• Las distancias son muy dispares, yendo desde unos pocos años luz (como Alfa Centauri, la estrella mas cercana a
la Tierra que está a 4 años luz) a un poco más de 20 millones (como ADS 7251,
	 una de las estrellas que conforman la constelación Osa Mayor).
*/

func OrdenarDistancia(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	medio := len(arr) / 2
	parte_izq := OrdenarDistancia(arr[:medio])
	parte_der := OrdenarDistancia(arr[medio:])

	return mezclar(parte_izq, parte_der)
}

func mezclar(izq, der []int) []int {
	resultado := make([]int, 0, len(izq)+len(der))
	i, j := 0, 0

	for i < len(izq) && j < len(der) {
		if izq[i] <= der[j] {
			resultado = append(resultado, izq[i])
			i++
		} else {
			resultado = append(resultado, der[j])
			j++
		}
	}
	resultado = append(resultado, izq[i:]...)
	resultado = append(resultado, der[j:]...)
	return resultado
}
