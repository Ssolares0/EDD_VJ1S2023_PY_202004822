package Estructuras

import (
	"fmt"
	"strconv"
)

//"fmt"

type Cliente struct {
	id     int
	nombre string
}

type Nodo_circularSimp struct {
	cliente   *Cliente
	siguiente *Nodo_circularSimp
}

func (l *Lista_circularSimp) validar_vacia2() bool {
	return l.longitud == 0
}

// creamos la esctructura de la lista circular simple
type Lista_circularSimp struct {
	inicio   *Nodo_circularSimp
	longitud int
}

func (l *Lista_circularSimp) AgregarCliente(id int, nombre string) {
	nuevoCliente := &Cliente{id, nombre}
	if l.validar_vacia2() {
		l.inicio = &Nodo_circularSimp{cliente: nuevoCliente, siguiente: nil}
		l.inicio.siguiente = l.inicio
		l.longitud++

	} else {

		if l.longitud == 1 {
			l.inicio.siguiente = &Nodo_circularSimp{cliente: nuevoCliente, siguiente: l.inicio}
			l.longitud++
		} else {
			auxiliar := l.inicio
			for i := 0; i < l.longitud-1; i++ {
				auxiliar = auxiliar.siguiente
			}
			auxiliar.siguiente = &Nodo_circularSimp{cliente: nuevoCliente, siguiente: l.inicio}
			l.longitud++
		}

	}

}

func MostrarListaCircular(l *Lista_circularSimp) {
	auxiliar := l.inicio
	for i := 0; i < l.longitud; i++ {
		fmt.Printf("Nombre del Cliente: %s, Id: %d \n", auxiliar.cliente.nombre, auxiliar.cliente.id)
		println("------------------------")
		auxiliar = auxiliar.siguiente
	}
}

func New_Lista_circularSimp() *Lista_circularSimp {
	return &Lista_circularSimp{nil, 0}
}

func (l *Lista_circularSimp) GraficoCircular() {
	fmt.Println("Generando grafico")
	name_archivo := "Reportes/ListaClientes.dot"

	name_imagen := "Reportes/ListaClientes.jpg"
	txt := "digraph ListaSimple{\n"
	txt += "rankdir=LR;\n"
	txt += "node[shape = oval];\n"
	aux := l.inicio
	count := 0

	for i := 0; i < l.longitud; i++ {
		txt = txt + "nodo" + strconv.Itoa(i) + "[label=\"" + aux.cliente.nombre + "\n" + "ID: " + strconv.Itoa(aux.cliente.id) + " \"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
		c := i + 1
		txt += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		count += 1
	}
	txt += "nodo" + strconv.Itoa(count) + "->nodo0;\n"
	txt += "}"
	createArch(name_archivo)
	escribirEnArch(txt, name_archivo)
	run(name_imagen, name_archivo)

}
