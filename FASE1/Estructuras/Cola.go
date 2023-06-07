package Estructuras

// "fmt"
// "strconv"

type Clientes struct {
	id     int
	nombre string
}

// creamos el nodo cola
type NodoCola struct {
	cliente   *Clientes
	siguiente *NodoCola
}

// creamos la estructura de la cola

type Cola struct {
	inicio   *NodoCola
	longitud int
}

//validar si esta vacia

func (l *Cola) EstaVacia() bool {
	if l.longitud == 0 {
		return true
	} else {
		return false
	}
}

func nuevoNodo(cliente *Clientes) *NodoCola {
	return &NodoCola{cliente, nil}
}
func (l *Cola) Agregar(id int, nombre string) {
	if l.EstaVacia() {
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
