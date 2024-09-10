package rpl

func PosicionPico(v []int, ini, fin int) int {
    medio := (ini + fin) / 2
    if v[medio - 1] < v[medio] && v[medio] > v[medio + 1] {
        return medio
    } else if v[medio - 1] < v [medio] {
        return PosicionPico(v, medio, fin)
    } else {
        return PosicionPico(v, ini, medio)
    }
}