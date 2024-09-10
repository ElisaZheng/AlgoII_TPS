package ejerc2


type ab struct {
    izq *ab
    der *ab
    clave string
}

func ReconstruirParticular() *ab {
    pre := []string{"E","U","R","M","A","O","N","D","V","S","Z","T"}
    in := []string{"M","R","A","U","O","Z","S","V","D","N","E","T"}
    ab := &ab{}
    reconstruccion(pre, in, ab)
    return ab

}
func reconstruccion(pre, in []string, arbol *ab) {
    if len(pre) == 0 {
        return 
    }
    arbol.clave = pre[0]
    ind := 0
    for i , elem := range in {
        if elem == pre[0] {
            ind = i
            break
        }
    }
    izq := in[:ind]
    der := in[ind+1:]
    if len(izq) > 0 {
        arbol.izq = &ab{}
        reconstruccion(pre[1:1+len(izq)], izq, arbol.izq)
    }
    if len(der) > 0 {
        arbol.der = &ab{}
        reconstruccion(pre[1+len(izq):], der, arbol.der)
    }
}