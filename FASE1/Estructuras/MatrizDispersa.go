package Estructuras

import (
	"fmt"
)

type Matriz struct {
	Raiz         *NodoMatriz
	Pixel_width  int
	Pixel_height int
	Image_width  int
	Image_height int
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
	return &Matriz{Raiz: nil, Pixel_width: 0, Pixel_height: 0, Image_width: 0, Image_height: 0}
}
