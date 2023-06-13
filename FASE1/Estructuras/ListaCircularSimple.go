package Estructuras

import (
	"fmt"
	"strconv"
)

//"fmt"

type Cliente struct {
	Id     int
	nombre string
}

type Nodo_circularSimp struct {
	Cliente   *Cliente
	Siguiente *Nodo_circularSimp
}

func (l *Lista_circularSimp) validar_vacia2() bool {
	return l.Longitud == 0
}

// creamos la esctructura de la lista circular simple
type Lista_circularSimp struct {
	Inicio   *Nodo_circularSimp
	Longitud int
}

func (l *Lista_circularSimp) AgregarCliente(id int, nombre string) {
	nuevoCliente := &Cliente{id, nombre}
	if l.validar_vacia2() {
		l.Inicio = &Nodo_circularSimp{Cliente: nuevoCliente, Siguiente: nil}
		l.Inicio.Siguiente = l.Inicio
		l.Longitud++

	} else {

		if l.Longitud == 1 {
			l.Inicio.Siguiente = &Nodo_circularSimp{Cliente: nuevoCliente, Siguiente: l.Inicio}
			l.Longitud++
		} else {
			auxiliar := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				auxiliar = auxiliar.Siguiente
			}
			auxiliar.Siguiente = &Nodo_circularSimp{Cliente: nuevoCliente, Siguiente: l.Inicio}
			l.Longitud++
		}

	}

}

func MostrarListaCircular(l *Lista_circularSimp) {
	auxiliar := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Printf("Nombre del Cliente: %s, Id: %d \n", auxiliar.Cliente.nombre, auxiliar.Cliente.Id)
		println("------------------------")
		auxiliar = auxiliar.Siguiente
	}
}

func (l *Lista_circularSimp) BuscarCliente(id int) {
	aux := l.Inicio
	for aux != nil {
		if aux.Cliente.Id == id {
			fmt.Println("Cliente encontrado")
		}

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
	aux := l.Inicio
	count := 0

	for i := 0; i < l.Longitud; i++ {
		txt = txt + "nodo" + strconv.Itoa(i) + "[label=\"" + aux.Cliente.nombre + "\n" + "ID: " + strconv.Itoa(aux.Cliente.Id) + " \"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
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
