package rpl

/*
Se tiene un arreglo tal que [1, 1, 1, …, 0, 0, …] (es decir, 
unos seguidos de ceros). Se pide una función de orden O(log(n)) que encuentre el índice del primer 0. 
Si no hay ningún 0 (solo hay unos), debe devolver -1.
*/

func IndicePrimeroCero(arr []int) int {
    return _indice(arr, 0, len(arr)-1)
}

func _indice(arr []int, inicio, fin int) int {
    if inicio > fin {
        return -1
    }
    
    medio := (inicio + fin) / 2
    
    // Verificar si el medio es el primer 0
    if (medio == 0 || arr[medio-1] == 1) && arr[medio] == 0 {
        return medio
    }
    
    // Si el medio es 1, el primer 0 debe estar a la derecha
    if arr[medio] == 1 {
        return _indice(arr, medio+1, fin)
    }
    
    // Si no, buscar en la mitad izquierda
    return _indice(arr, inicio, medio-1)
}