package rpl


type Alumno struct {
    nombre string
    p1, p2, p3 int
}

func OrdenarAlumnosRadixSort(arr []Alumno) []Alumno {
    CoutingSort(arr, 11, 3)
    CoutingSort(arr, 11, 2)
    CoutingSort(arr, 11, 1)
    CoutingSort(arr, 11, 4)
    return arr
}

func CoutingSort(arr []Alumno, rango, digito int) {
    frecuencia := make([]int, rango)
    acumulado := make([]int, rango)
    resultado := make([]Alumno, len(arr))

    for _ ,elem := range arr {
        frecuencia[obtener_nota(elem, digito)] ++
    }

    for i := 1 ; i < rango ; i ++ {
        acumulado[i] = acumulado [i - 1] + frecuencia [i - 1]
    }

    for _, elem := range arr {
        valor := obtener_nota(elem, digito)
        posicion := acumulado[valor]
        resultado[posicion] = elem
        acumulado[valor] ++ 
    }

    for i := 1 ; i < len(arr) ; i ++  {
        arr[i] = resultado[i]
    }
}

func obtener_nota(a Alumno, digito int) int {
    switch digito {
    case 1:
        return a.p1
    case 2:
        return a.p2
    case 3:
        return a.p3
    case 4:
        return (a.p1 + a.p2 + a.p3) / 3
    }
    return 0
}
