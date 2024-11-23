package ejerc

type Helado struct {
	Nombre string
	Popularidad int
	Sabor string
	AptoCeliaco bool
}

func OrdenarHelados(h []Helado) []Helado {
	aptoC := make([]Helado, 0)
	noApto := make([]Helado, 0)
	for _, hel := range h { // O(n) porque va a revisar cada elem en el arreglo de Helado 
		if hel.AptoCeliaco {
			aptoC = append(aptoC, hel)
		} else {
			noApto = append(noApto, hel)
		}
	}
	aptoC = Ordenar(aptoC, 10, func(h Helado) int {return h.Popularidad}) // ordenar cada elemento usando counting sort
	noApto = Ordenar(noApto, 10, func(h Helado) int {return h.Popularidad})
	return append(aptoC, noApto...) // juntar los elem ->

}

func Ordenar(h []Helado, rango int , f func(Helado) int) []Helado { //es counting sort O(n+k)
	frecuencia := make([]int, rango)
	acumulada := make([]int, rango)
	result := make([]Helado, len(h))
	for _, helados := range h {
		frecuencia[f(helados)] ++
	}
	for i := 1; i < len(acumulada); i ++ {
		acumulada[i] = frecuencia[i-1] + acumulada[i-1]
	}
	for _, helados := range h {
		result[acumulada[f(helados)]] = helados
	}
	return result
}

// la complejidad va a ser 3*O(n) + 2* O(n+k) = O(3n+k) = O(n)