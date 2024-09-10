package rpl

type Persona struct {
    Nombre string
    // otros campos relevantes
}

// Simulación de una función PCR que determina si hay un contagiado en el grupo
func pcr(grupo []Persona) bool {
    // Implementación de la lógica para detectar si hay un contagiado en el grupo
    // Esta función debe retornar true si hay al menos una persona contagiada en el grupo
    // Aquí está solo una simulación
    for _, persona := range grupo {
        if persona.Nombre == "contagiado" {
            return true
        }
    }
    return false
}

// Función para detectar a la persona contagiada utilizando división y conquista
func Detectar_contagiado(grupo []Persona) Persona {
    if len(grupo) == 1 {
        return grupo[0]
    }

    medio := len(grupo) / 2

    if pcr(grupo[:medio]) {
        return Detectar_contagiado(grupo[:medio])
    } else {
        return Detectar_contagiado(grupo[medio:])
    }
}
