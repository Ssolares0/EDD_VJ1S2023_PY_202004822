package Estructuras

type Cliente_PIla struct {
	Idpila     string
	Nombrepila string
}

type NodoPila struct {
	ClientePila *Cliente_PIla
	Siguiente   *NodoPila
}

type Lista_pila struct {
	Inicio   *NodoPila
	Longitud int
}

func (l *Lista_pila) EstaVacia() bool {
	return l.Longitud == 0
}

func (l *Lista_pila) Push(id string, nombre string) {
	if l.Longitud == 0 {
		newNodo := &NodoPila{&Cliente_PIla{id, nombre}, nil}
		l.Inicio = newNodo
		l.Longitud++
	} else {
		newNodo := &NodoPila{&Cliente_PIla{id, nombre}, l.Inicio}
		l.Inicio = newNodo
		l.Longitud++
	}
}

func (l *Lista_pila) GraficarPila() {
	name_archivo := "Reportes/pila.dot"
	name_imagen := "Reportes/pila.jpg"
	txt := "digraph pila{\n"
	txt += "rankdir=LR;\n"
	txt += "node[shape = record]"
	aux := l.Inicio
	txt += "nodo0 [label=\""

	for i := 0; i < l.Longitud; i++ {
		txt = txt + "|(ID: " + aux.ClientePila.Idpila + ", Imagen: " + aux.ClientePila.Nombrepila + ")"
		aux = aux.Siguiente
	}
	txt += "\"]; \n}"
	createArch(name_archivo)
	escribirEnArch(txt, name_archivo)
	run(name_imagen, name_archivo)
}
func New_Lista_pila() *Lista_pila {
	return &Lista_pila{nil, 0}
}
