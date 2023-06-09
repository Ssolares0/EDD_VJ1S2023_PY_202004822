package Estructuras

type Cliente_PIla struct {
	id     int
	nombre string
}

type NodoPila struct {
	cliente   *Cliente
	siguiente *NodoPila
}

type Lista_pila struct {
	inicio   *NodoPila
	longitud int
}

func (l *Lista_pila) EstaVacia() bool {
	return l.longitud == 0
}

func (l *Lista_pila) Push(id int, nombre string) {
	nuevoCliente := &Cliente{id, nombre}
	if l.EstaVacia() {
		nuevoNodo2 := &NodoPila{cliente: nuevoCliente, siguiente: nil}
		l.inicio = nuevoNodo2
		l.longitud++
	} else {
		aux := l.inicio
		l.inicio = &NodoPila{cliente: nuevoCliente, siguiente: aux}
		l.longitud++
	}
}
