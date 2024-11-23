package main

import ejer "parcial_1/rpl"
import "fmt"


func main (){
	grupo := [] ejer.Persona{
        ejer.Persona{Nombre: "contagiado"},
        ejer.Persona{Nombre: "persona2"},
        ejer.Persona{Nombre: "peronsa3"},
        ejer.Persona{Nombre: "persona4"},
        ejer.Persona{Nombre: "persona5"},
    }

    contagiado := ejer.Detectar_contagiado(grupo)
    fmt.Printf("La persona contagiada es: %s\n", contagiado.Nombre)
}
