package Estructuras

type Cliente_PIla struct {
	idpila     string
	nombrepila string
}

type NodoPila struct {
	clientePila *Cliente_PIla
	siguiente   *NodoPila
}

type Lista_pila struct {
	inicio   *NodoPila
	longitud int
}

func (l *Lista_pila) EstaVacia() bool {
	return l.longitud == 0
}

func (l *Lista_pila) Push(id string, nombre string) {
	if l.longitud == 0 {
		newNodo := &NodoPila{&Cliente_PIla{id, nombre}, nil}
		l.inicio = newNodo
		l.longitud++
	} else {
		newNodo := &NodoPila{&Cliente_PIla{id, nombre}, l.inicio}
		l.inicio = newNodo
		l.longitud++
	}
}

func (l *Lista_pila) GraficarPila() {
	name_archivo := "Reportes/pila.dot"
	name_imagen := "Reportes/pila.jpg"
	txt := "digraph pila{\n"
	txt += "rankdir=LR;\n"
	txt += "node[shape = record]"
	aux := l.inicio
	txt += "nodo0 [label=\""

	for i := 0; i < l.longitud; i++ {
		txt = txt + "|(ID: " + aux.clientePila.idpila + ", Imagen: " + aux.clientePila.nombrepila + ")"
		aux = aux.siguiente
	}
	txt += "\"]; \n}"
	createArch(name_archivo)
	escribirEnArch(txt, name_archivo)
	run(name_imagen, name_archivo)
}
func New_Lista_pila() *Lista_pila {
	return &Lista_pila{nil, 0}
}
