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

func modoRecursivo(image_width int, image_height int, pixel_width int, pixel_height int) {
	tamanio := []int{image_width, image_height}
	fmt.Println("Tamaño:  ", tamanio)
}
