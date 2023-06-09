package Estructuras

import (
	"strconv"
)

// "fmt"
// "strconv"

type Clientes struct {
	id     string
	nombre string
}

// creamos el nodo cola
type NodoCola struct {
	cliente   *Clientes
	siguiente *NodoCola
}

// creamos la estructura de la cola

type Lista_cola struct {
	inicio   *NodoCola
	longitud int
}

//validar si esta vacia

func (l *Lista_cola) EstaVacia() bool {
	if l.longitud == 0 {
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
		l.inicio = nuevoNodo(nuevoCliente)
		l.longitud++

	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		nuevoCliente := &Clientes{id, nombre}
		aux.siguiente = nuevoNodo(nuevoCliente)
		aux.siguiente.siguiente = nil
		l.longitud++
	}
}

func (l *Lista_cola) Descolar() {
	if l.EstaVacia() {
		println("La lista esta vacia")
	} else {
		l.inicio = l.inicio.siguiente
		l.longitud--
	}
}

func MostrarCola(l *Lista_cola) {
	aux := l.inicio
	for aux != nil {
		println("Id: " + aux.cliente.id + " Nombre: " + aux.cliente.nombre)
		aux = aux.siguiente
	}
}

func (c *Lista_cola) GraficarCola() {
	nombre_archivo := "Reportes/cola.dot"
	nombre_imagen := "Reportes/cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.inicio
	contador := 0
	if c.longitud == 1 {
		return
	}
	for aux != nil {
		texto = texto + "nodo" + strconv.Itoa(contador) + "[label=\"{" + aux.cliente.id + "|" + aux.cliente.nombre + " }\"];\n"
		aux = aux.siguiente
		contador += 1
	}

	contador = 0
	aux = c.inicio

	for aux != nil {
		texto = texto + "nodo" + strconv.Itoa(contador) + "->nodo" + strconv.Itoa(contador+1) + ";\n"
		aux = aux.siguiente
		contador += 1
		if aux.siguiente == nil {
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
