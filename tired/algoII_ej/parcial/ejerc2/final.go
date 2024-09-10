package ejerc2

/*
Dado un arreglo de n enteros positivos que casi representa una progresión aritmética creciente (es una progresión
aritmética a la que le falta un elemento), implementar un algoritmo que devuelva el elemento faltante de manera eficiente
(complejidad logarítmica). Se puede suponer que el arreglo tiene al menos 4 elementos. Justificar la complejidad del
algoritmo implementado. Por ejemplo, si la sucesión es [5, 8, 14, 17, 20, 23] tiene que devolver 11.
*/

func Buscar_faltante(arr []int) int {
	medio := len(arr) / 2
	diferencia := (arr[len(arr) - 1] - arr[0]) / len(arr) + 1
	valor_medio := arr[0] + diferencia * medio
	if arr[medio] - arr[medio - 1] != 3 {
		return arr[medio] - 3
	}
	if valor_medio < arr[medio] {
		return Buscar_faltante(arr[:medio])
	}
	return Buscar_faltante((arr[medio + 1:]))
}
