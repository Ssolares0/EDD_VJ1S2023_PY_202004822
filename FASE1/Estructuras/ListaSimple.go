package Estructuras

import (
	"fmt"
	"strconv"
)

/*
Lista Simple enlazada para los empleados
*/
type Empleado struct {
	nombre   string
	id       int
	cargo    string
	password string
}

type Nodo struct {
	empleado  *Empleado
	siguiente *Nodo
}

/*
Creamos la estructura de la lista simple
consta de 2 atributos, inicio que se refiere al primer elemento de la lista de tipo Nodo
y longitud de tipo entero que se refiere al tamaño de la lista actual
*/
type Lista struct {
	inicio   *Nodo
	longitud int
}

func NewNodo(empleado *Empleado) *Nodo {
	return &Nodo{empleado, nil}
}

/*
Creamos la funcion estaVacia que nos va devolver un valor booleano
para saber si la lista esta vacia o tiene elementos
*/

func (l *Lista) EstaVacia() bool {
	if l.longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *Lista) AgregarEmpleado(nombre string, id int, cargo string, password string) {
	comprobacion := l.EstaVacia()
	if comprobacion {
		nuevoEmpleado := &Empleado{nombre, id, cargo, password}
		l.inicio = NewNodo(nuevoEmpleado)
		l.longitud++
	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		nuevoEmpleado := &Empleado{nombre, id, cargo, password}
		aux.siguiente = NewNodo(nuevoEmpleado)
		aux.siguiente.siguiente = nil
		l.longitud++
	}
}

/**Creamos la funcion MostrarLista para imprimir los valores de la lista*/
func MostrarLista(lista *Lista) {
	auxiliar := lista.inicio
	for auxiliar != nil {
		fmt.Println(auxiliar.empleado.nombre)
		fmt.Println(auxiliar.empleado.id)
		fmt.Println(auxiliar.empleado.cargo)
		fmt.Println(auxiliar.empleado.password)
		fmt.Println("------------------------")
		auxiliar = auxiliar.siguiente
	}
}

func (l *Lista) ValidarEmpleado(id int, password string) {
	comprobacion := l.EstaVacia()
	fmt.Println(comprobacion)
	auxiliar := l.inicio
	for auxiliar != nil {
		if auxiliar.empleado.id == id {
			if auxiliar.empleado.password == password {
				fmt.Println("Bienvenido usuario")
			}

		}

	}

}

func (l *Lista) BuscarEmpleado(id int, password string) *Empleado {
	comprobacion := l.EstaVacia()
	if comprobacion {
		return nil
	} else {
		auxiliar := l.inicio
		for auxiliar != nil {
			if auxiliar.empleado.id == id && auxiliar.empleado.password == password {
				return auxiliar.empleado
			}
			auxiliar = auxiliar.siguiente
		}
		return nil
	}
}
func (l *Lista) Grafico() {
	fmt.Println("Generando grafico")
	name_archivo := "ListaEmpleados.dot"

	name_imagen := "ListaEmpleados.jpg"
	txt := "digraph ListaSimple{\n"
	txt += "rankdir=LR;\n"
	txt += "node[shape = oval];\n"
	txt += "nodonull2[label=\"null\"];\n"
	aux := l.inicio
	count := 0

	for i := 0; i < l.longitud; i++ {
		txt = txt + "nodo" + strconv.Itoa(i) + "[label=\"" + aux.empleado.nombre + "\n" + "ID: " + strconv.Itoa(aux.empleado.id) + " \"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
		c := i + 1
		txt += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		count += 1
	}
	txt += "nodo" + strconv.Itoa(count) + "->nodonull2;\n"
	txt += "}"
	createArch(name_archivo)
	escribirEnArch(txt, name_archivo)
	run(name_imagen, name_archivo)

}

func New_Lista() *Lista {
	return &Lista{nil, 0}
}
