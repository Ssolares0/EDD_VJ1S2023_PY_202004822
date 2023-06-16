package Estructuras

import (
	"fmt"
	"strconv"
)

type Imagen struct {
	Nombre string
	capas  int
}

type NodoDoble struct {
	Imagen    *Imagen
	Siguiente *NodoDoble
	anterior  *NodoDoble
}

type ListaDoble struct {
	Inicio   *NodoDoble
	Longitud int
}

func (l *ListaDoble) validar_vacia() bool {
	return l.Longitud == 0
}

func (l *ListaDoble) AgregarImagen(nombre string, capas int) {
	nuevoImagen := &Imagen{nombre, capas}

	if l.validar_vacia() {
		l.Inicio = &NodoDoble{nuevoImagen, nil, nil}
		l.Longitud++

	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoDoble{nuevoImagen, nil, aux}
		l.Longitud++

	}
}

func MostrarListaDoble(l *ListaDoble) {
	count := 1
	auxiliar := l.Inicio
	for auxiliar != nil {
		fmt.Printf("Numero: %d, Nombre de la imagen: %s, Capa: %d \n", count, auxiliar.Imagen.Nombre, auxiliar.Imagen.capas)
		println("------------------------")
		auxiliar = auxiliar.Siguiente
		count++
	}

}

func MostrarNormal(l *ListaDoble) {
	auxiliar := l.Inicio
	for auxiliar != nil {
		fmt.Printf("Nombre de la imagen: %s, Capa: %d \n", auxiliar.Imagen.Nombre, auxiliar.Imagen.capas)
		println("------------------------")
		auxiliar = auxiliar.Siguiente
	}
}

func (l *ListaDoble) GraficoDoble() {
	name_archivo := "Reportes/ListaImagenes.dot"
	name_imagen := "Reportes/ListaImagenes.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.Imagen.Nombre + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	createArch(name_archivo)
	escribirEnArch(texto, name_archivo)
	run(name_imagen, name_archivo)

}

func New_ListaDoble() *ListaDoble {
	return &ListaDoble{nil, 0}
}
