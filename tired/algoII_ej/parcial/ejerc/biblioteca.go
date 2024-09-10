package ejerc

type Libro struct {
	Titulo string
	Ubicacion string
}

type Autor struct {
	Nombre string
	Libro []Libro
}

func BuscarLibro(nombreAutor string, tituloLibro string, autores []Autor) string {
	ubicacion := ""
	ind := BuscarAutor(nombreAutor, autores)
	if ind == -1 {
		return "no existe el autor"
	}
	ubicacion = BuscarUbicacion(tituloLibro, autores[ind].Libro)
	return ubicacion
}

func BuscarAutor(nombreAutor string, autores []Autor) int {
	if len(autores) == 0 {
		return -1
	}
	if len(autores) == 1 {
		if autores[0].Nombre == nombreAutor {
			return 0
		}
		return -1
	}
	medio := len(autores) / 2
	if autores[medio].Nombre > nombreAutor {
		return BuscarAutor(nombreAutor, autores[:medio])
	} else if autores[medio].Nombre < nombreAutor {
		return BuscarAutor(nombreAutor, autores[medio+1:])
	} else {
		return medio
	}
}

func BuscarUbicacion(tituloLibro string, libros []Libro) string {
	if len(libros) == 0 {
		return ""
	}
	if len(libros) == 1 {
		if libros[0].Titulo == tituloLibro {
			return libros[0].Ubicacion
		}
		return ""
	}
	medio := len(libros) / 2
	if libros[medio].Titulo > tituloLibro {
		return BuscarUbicacion(tituloLibro, libros[:medio])
	} else if libros[medio].Titulo < tituloLibro {
		return BuscarUbicacion(tituloLibro, libros[medio+1:])
	}
	return libros[medio].Ubicacion
}
/*
verificar si esta vacio el arreglo autor O(1)
entrar a la llamada recursiva : figuran 3, pero solo entrar 1
el arreglo original se parte en 2 
y el costo de unir y juntar es O(1) 
por eso podemos usar T de Maestro 
T(n) = AT(n/B) +O(n^C)
A = 1
B = 2
C = 0
logB(A) = 0 = C -> la complejidad es O(log(n))
*/

/* func main() {
	autores := []Autor{
		{
			Nombre: "George Orwell",
			Libro: []Libro{
				{Titulo: "1984", Ubicacion: "Sección B, biblio. 1"},
				{Titulo: "Rebelión en la granja", Ubicacion: "Sección B, biblio. 1"},
			},
		},
		{
			Nombre: "J.K. Rowling",
			Libro: []Libro{
				{Titulo: "HP y la cámara secreta", Ubicacion: "Sección A, biblio. 4"},
			},
		},
	}

	// Ejemplo 1: Buscar ubicación de un libro por título y autor
	ubicacion := buscarLibro("George Orwell", "1984", autores)
	fmt.Println("La ubicación del libro es:", ubicacion) // Devuelve "Sección B, biblio. 1"
	nombre := BuscarAutor("George Orwell", autores)
	fmt.Println(nombre)

}*/


