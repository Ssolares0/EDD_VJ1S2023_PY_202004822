package Estructuras2

import (
	"strconv"
)

// "fmt"
// "strconv"

type Clientes struct {
	Id     string
	Nombre string
}

// creamos el nodo cola
type NodoCola struct {
	Cliente   *Clientes
	Siguiente *NodoCola
}

// creamos la estructura de la cola

type Lista_cola struct {
	Inicio   *NodoCola
	Longitud int
}

//validar si esta vacia

func (l *Lista_cola) EstaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func nuevoNodo(cliente *Clientes) *NodoCola {
	return &NodoCola{cliente, nil}
}
func (l *Lista_cola) Colar(id string, nombre string) {
	comprobacion := l.EstaVacia()
	if comprobacion {
		nuevoCliente := &Clientes{id, nombre}
		l.Inicio = nuevoNodo(nuevoCliente)
		l.Longitud++

	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		nuevoCliente := &Clientes{id, nombre}
		aux.Siguiente = nuevoNodo(nuevoCliente)
		aux.Siguiente.Siguiente = nil
		l.Longitud++
	}

}

func (l *Lista_cola) Descolar() {
	if l.EstaVacia() {
		println("La lista esta vacia")
	} else {
		l.Inicio = l.Inicio.Siguiente
		l.Longitud--
	}
}

func MostrarCola(l *Lista_cola) {
	aux := l.Inicio
	for aux != nil {
		println("Id: " + aux.Cliente.Id + " Nombre: " + aux.Cliente.Nombre)
		aux = aux.Siguiente
	}
}

func (l *Lista_cola) MostrarPrimerValor() (string, string) {
	aux := l.Inicio
	return aux.Cliente.Id, aux.Cliente.Nombre
}

func (c *Lista_cola) GraficarCola() {
	nombre_archivo := "Reportes/cola.dot"
	nombre_imagen := "Reportes/cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.Inicio
	contador := 0
	if c.Longitud == 1 {
		return
	}
	for aux != nil {
		texto = texto + "nodo" + strconv.Itoa(contador) + "[label=\"{" + aux.Cliente.Id + "|" + aux.Cliente.Nombre + " }\"];\n"
		aux = aux.Siguiente
		contador += 1
	}

	contador = 0
	aux = c.Inicio

	for aux != nil {
		texto = texto + "nodo" + strconv.Itoa(contador) + "->nodo" + strconv.Itoa(contador+1) + ";\n"
		aux = aux.Siguiente
		contador += 1
		if aux.Siguiente == nil {
			break
		}
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	createArch(nombre_archivo)
	escribirEnArch(texto, nombre_archivo)
	run(nombre_imagen, nombre_archivo)
}
func New_ListaCola() *Lista_cola {
	return &Lista_cola{nil, 0}
}
