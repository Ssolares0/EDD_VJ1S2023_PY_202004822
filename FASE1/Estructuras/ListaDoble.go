package Estructuras

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Imagen struct {
	nombre string
	capas  int
}

type NodoDoble struct {
	imagen    *Imagen
	siguiente *NodoDoble
	anterior  *NodoDoble
}

type ListaDoble struct {
	inicio   *NodoDoble
	longitud int
}

func (l *ListaDoble) validar_vacia() bool {
	return l.longitud == 0
}

func (l *ListaDoble) AgregarImagen(nombre string, capas int) {
	nuevoImagen := &Imagen{nombre, capas}

	if l.validar_vacia() {
		l.inicio = &NodoDoble{nuevoImagen, nil, nil}
		l.longitud++

	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoDoble{nuevoImagen, nil, aux}
		l.longitud++

	}
}

func MostrarListaDoble(l *ListaDoble) {
	auxiliar := l.inicio
	for auxiliar != nil {
		fmt.Printf("Nombre de la imagen: %s, Capa: %d \n", auxiliar.imagen.nombre, auxiliar.imagen.capas)
		println("------------------------")
		auxiliar = auxiliar.siguiente
	}
	fmt.Println("Ingrese el nombre de la imagen que desea seleccionar")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seleccionada := scanner.Text()

	ruta := "C:/FASE1/csv/" + seleccionada + "/inicial.csv"
	fmt.Println("La ruta es:  " + ruta)

}

func (l *ListaDoble) GraficoDoble() {
	name_archivo := "Reportes/ListaImagenes.dot"
	name_imagen := "Reportes/ListaImagenes.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.imagen.nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
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
