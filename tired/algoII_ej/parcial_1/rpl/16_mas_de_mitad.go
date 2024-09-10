package rpl

func MasDeLaMitad(arr []int) bool {
    for  _, letra := range arr {
        contador := _masDeLaMitad(arr, letra, 0, len(arr) - 1, 0)
        if contador > len(arr) / 2 {
            return true
        }
    }
    return false
}

func _masDeLaMitad(arr []int, letra, inicio, fin, contador int) int {
    if inicio > fin {
        return contador 
	}
    medio := (inicio + fin) / 2
    if arr[medio] == letra {
        contador ++
    }
    contador = _masDeLaMitad(arr, letra, inicio, medio - 1, contador)
    contador = _masDeLaMitad(arr, letra, medio + 1, fin, contador)
    return contador
} 