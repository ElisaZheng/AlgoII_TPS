package rpl

func ParteEnteraRaiz(n int) int {
    return _parteEntera(n, n / 2)
}

func _parteEntera(n, medio int) int {
    if medio  * medio == n {
        return medio
    }
    if medio * medio < n && (medio + 1) * (medio + 1) > n {
        return medio
    }
    if (medio + 1) * (medio + 1) == n {
        return medio + 1
    }
    if medio * medio > n {
        return _parteEntera(n, medio / 2)
    } else {
        return _parteEntera(n, medio + (medio - 1) / 2)
    }
}
