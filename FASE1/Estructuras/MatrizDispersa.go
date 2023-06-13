package Estructuras

import (
	"fmt"
)

type Matriz struct {
	inicial      *NodoMatriz
	pixel_width  int
	pixel_height int
	image_width  int
	image_height int
}
type NodoMatriz struct {
	Arriba    *NodoMatriz
	Abajo     *NodoMatriz
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Adelante  *NodoMatriz
	Atras     *NodoMatriz
	CoorX     int
	CoorY     int
	Color     string
}

func (l *Matriz) AgregarValues(image_width int, image_height int, pixel_width int, pixel_height int) {
	tamanio := []int{image_width, image_height}
	fmt.Println("Tamaño:  ", tamanio)
}

func NewMatriz() *Matriz {
	return &Matriz{inicial: nil, pixel_width: 0, pixel_height: 0, image_width: 0, image_height: 0}
}
