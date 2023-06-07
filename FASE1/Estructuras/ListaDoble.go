package Estructuras

type Empleado2 struct {
	nombre   string
	id       int
	cargo    string
	password string
}

type NodoDoble struct {
	empleado  *Empleado
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

func (l *ListaDoble) AgregarEmpleado(nombre string, id int, cargo string, password string) {
	nuevoEmpleado := &Empleado{nombre, id, cargo, password}

	if l.validar_vacia() {
		l.inicio = &NodoDoble{nuevoEmpleado, nil, nil}
		l.longitud++

	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoDoble{nuevoEmpleado, nil, aux}
		l.longitud++

	}
}
