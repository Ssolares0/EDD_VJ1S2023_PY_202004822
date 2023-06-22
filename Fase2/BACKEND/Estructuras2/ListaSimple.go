package Estructuras2

import (
	"fmt"
	"strconv"
)

/*
Lista Simple enlazada para los empleados
*/
type Empleado struct {
	Nombre   string
	Id       int
	cargo    string
	Password string
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
	Inicio   *Nodo
	Longitud int
}

func NewNodo(empleado *Empleado) *Nodo {
	return &Nodo{empleado, nil}
}

/*
Creamos la funcion estaVacia que nos va devolver un valor booleano
para saber si la lista esta vacia o tiene elementos
*/

func (l *Lista) EstaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *Lista) AgregarEmpleado(nombre string, id int, cargo string, password string) {
	comprobacion := l.EstaVacia()
	if comprobacion {
		nuevoEmpleado := &Empleado{nombre, id, cargo, password}
		l.Inicio = NewNodo(nuevoEmpleado)
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		nuevoEmpleado := &Empleado{nombre, id, cargo, password}
		aux.siguiente = NewNodo(nuevoEmpleado)
		aux.siguiente.siguiente = nil
		l.Longitud++
	}
}

/**Creamos la funcion MostrarLista para imprimir los valores de la lista*/
func MostrarLista(lista *Lista) {
	auxiliar := lista.Inicio
	for auxiliar != nil {
		fmt.Println(auxiliar.empleado.Nombre)
		fmt.Println(auxiliar.empleado.Id)
		fmt.Println(auxiliar.empleado.cargo)
		fmt.Println(auxiliar.empleado.Password)
		fmt.Println("------------------------")
		auxiliar = auxiliar.siguiente
	}
}

func (l *Lista) ValidarEmpleado(id int, password string) {
	comprobacion := l.EstaVacia()
	fmt.Println(comprobacion)
	auxiliar := l.Inicio
	for auxiliar != nil {
		if auxiliar.empleado.Id == id {
			if auxiliar.empleado.Password == password {
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
		auxiliar := l.Inicio
		for auxiliar != nil {
			if auxiliar.empleado.Id == id && auxiliar.empleado.Password == password {
				return auxiliar.empleado
			}
			auxiliar = auxiliar.siguiente
		}
		return nil
	}
}
func (l *Lista) Grafico() {
	fmt.Println("Generando grafico")
	name_archivo := "Reportes/ListaEmpleados.dot"

	name_imagen := "Reportes/ListaEmpleados.jpg"
	txt := "digraph ListaSimple{\n"
	txt += "rankdir=LR;\n"
	txt += "node[shape = oval];\n"
	txt += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	count := 0

	for i := 0; i < l.Longitud; i++ {
		txt = txt + "nodo" + strconv.Itoa(i) + "[label=\"" + aux.empleado.Nombre + "\n" + "ID: " + strconv.Itoa(aux.empleado.Id) + " \"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
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
