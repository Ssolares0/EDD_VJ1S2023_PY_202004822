package Estructuras

import (
	"fmt"
)

type Cliente struct {
	id     int
	nombre string
}

type Nodo_circularSimp struct {
	cliente   *Cliente
	siguiente *Nodo_circularSimp
}

// creamos la esctructura de la lista circular simple
type Lista_circularSimp struct {
	inicio *Nodo_circularSimp

	longitud int
}
