package ejerc


func IndicePrimeroCero(arr []int) int {
    return buscarIndiceDeCero(arr, 0, len(arr)-1)
}

func buscarIndiceDeCero(arr []int, inicio, final int) int {
    medio := (inicio + final + 1) / 2
    if final - inicio < 2 {
        if arr[medio] == 0 && arr[medio-1] == 1 {
                return medio
        } else if arr[medio] == 0 && arr[medio-1] == 0{
                return medio-1
        } else {
                return -1
        }
    }
    
    if arr[medio] == 0 && arr[medio-1] == 1{
        return medio
    } else if arr[medio] == 1 {
        return buscarIndiceDeCero(arr, medio, final)
    } else {
        return buscarIndiceDeCero(arr, inicio, medio)
    }
}