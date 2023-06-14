package Estructuras

import (
	"fmt"
	"strconv"
)

type Matriz struct {
	Raiz *NodoMatriz
	/*
		Pixel_width  int
		Pixel_height int
		Image_width  int
		Image_height int
	*/
}
type NodoMatriz struct {
	Arriba    *NodoMatriz
	Abajo     *NodoMatriz
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Adelante  *NodoMatriz
	Atras     *NodoMatriz
	CoorX     int
	CoorY     int
	Color     string
}

func (l *Matriz) buscarcolumna(x int) *NodoMatriz {
	aux := l.Raiz
	for aux != nil {
		if aux.CoorX == x {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (l *Matriz) buscarfila(y int) *NodoMatriz {
	aux := l.Raiz
	for aux != nil {
		if aux.CoorY == y {
			return aux
		}
		aux = aux.Abajo
	}
	return nil
}

func (l *Matriz) InsertarColumna(newNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	tmp := nodoRaiz
	pivote := false

	for {
		if tmp.CoorX == newNodo.CoorX {
			tmp.CoorY = newNodo.CoorY
			tmp.Color = newNodo.Color
			return tmp

		} else if tmp.CoorX > newNodo.CoorX {
			pivote = true
			break

		}
		if tmp.Siguiente != nil {
			tmp = tmp.Siguiente
		} else {
			break
		}
	}
	if pivote {
		newNodo.Siguiente = tmp
		tmp.Anterior.Siguiente = newNodo
		newNodo.Anterior = tmp.Anterior
		tmp.Anterior = newNodo

	} else {
		tmp.Siguiente = newNodo
		newNodo.Anterior = tmp
	}
	return newNodo
}
func (l *Matriz) InsertarFila(newNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	tmp := nodoRaiz
	pivote := false

	for {
		if tmp.CoorY == newNodo.CoorY {
			tmp.CoorX = newNodo.CoorX
			tmp.Color = newNodo.Color
			return tmp

		} else if tmp.CoorY > newNodo.CoorY {
			pivote = true
			break

		}
		if tmp.Abajo != nil {
			tmp = tmp.Abajo
		} else {
			break
		}
	}
	if pivote {
		newNodo.Abajo = tmp
		tmp.Arriba.Abajo = newNodo
		newNodo.Arriba = tmp.Arriba
		tmp.Arriba = newNodo

	} else {
		tmp.Abajo = newNodo
		newNodo.Arriba = tmp
	}
	return newNodo
}
func (l *Matriz) MostrarMatrizd() {
	aux := l.Raiz
	for aux != nil {
		aux2 := aux
		for aux2 != nil {
			fmt.Println("X: ", aux2.CoorX, "Y: ", aux2.CoorY, "Color: ", aux2.Color)
			aux2 = aux2.Siguiente
		}
		aux = aux.Abajo
	}
}

func (l *Matriz) newColumna(x int) *NodoMatriz {
	colum := "C" + strconv.Itoa(x)
	newNodo := &NodoMatriz{CoorX: x, CoorY: -1, Color: colum}
	columna := l.InsertarColumna(newNodo, l.Raiz)
	return columna
}

func (l *Matriz) newFila(y int) *NodoMatriz {
	colum := "F" + strconv.Itoa(y)
	newNodo := &NodoMatriz{CoorX: -1, CoorY: y, Color: colum}
	fila := l.InsertarFila(newNodo, l.Raiz)
	return fila
}

func (l *Matriz) AgregarElementos(x int, y int, color string) {
	newNodo := &NodoMatriz{CoorX: x, CoorY: y, Color: color}
	nodo_Columna := l.buscarcolumna(x)
	nodo_Fila := l.buscarfila(y)

	/*casos
	1) columna y fila no exista
	2) columna exista y fila no
	3) fila exista y columna no
	comprobamos cada caso
	*/
	if nodo_Columna == nil && nodo_Fila == nil {
		fmt.Println("caso 1")

		nodo_Columna = l.newColumna(x)
		nodo_Fila = l.newFila(y)
		newNodo = l.InsertarColumna(newNodo, nodo_Fila)
		newNodo = l.InsertarFila(newNodo, nodo_Columna)

	} else if nodo_Columna != nil && nodo_Fila == nil {
		fmt.Println("caso 2")
		nodo_Fila = l.newFila(y)
		newNodo = l.InsertarColumna(newNodo, nodo_Fila)
		newNodo = l.InsertarFila(newNodo, nodo_Columna)

	} else if nodo_Columna == nil && nodo_Fila != nil {
		fmt.Println("caso 3")
		nodo_Columna = l.newColumna(x)
		newNodo = l.InsertarColumna(newNodo, nodo_Fila)
		newNodo = l.InsertarFila(newNodo, nodo_Columna)

	} else if nodo_Columna != nil && nodo_Fila != nil {
		fmt.Println("caso 4")
		nodo_Columna = l.newColumna(x)
		newNodo = l.InsertarColumna(newNodo, nodo_Fila)
		newNodo = l.InsertarFila(newNodo, nodo_Columna)

	} else {
		fmt.Println("Ocurrio Un Error!")
	}

}
func (l *Matriz) ReporteGraphviz(nombreCapa string, numeroCapa string) {

	texto := ""
	//name_archivo := "Resultados/madis.dot"
	//name_imagen := "Resultados/madis.jpg"
	name_archivo := "Resultados/" + nombreCapa + "/capa" + numeroCapa + ".dot"
	name_imagen := "Resultados/" + nombreCapa + "/capa" + numeroCapa + ".jpg"
	aux1 := l.Raiz
	aux2 := l.Raiz
	aux3 := l.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"
		/** Creacion de los nodos actuales */
		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.CoorX+1) + strconv.Itoa(aux1.CoorY+1) + "[label=\"" + aux1.Color + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.CoorX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			for aux1 != nil {
				texto += "nodo" + strconv.Itoa(aux1.CoorX+1) + strconv.Itoa(aux1.CoorY+1) + "[label=\"" + aux1.Color + "\" ,group=" + strconv.Itoa(aux1.CoorX+1) + "]; \n"
				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		/** Conexiones entre los nodos de la matriz */
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.CoorX+1) + strconv.Itoa(aux1.CoorY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.CoorX+1) + strconv.Itoa(aux1.Siguiente.CoorY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.CoorX+1) + strconv.Itoa(aux1.CoorY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.CoorX+1) + strconv.Itoa(aux1.Abajo.CoorY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}
	//fmt.Println(texto)
	createArch(name_archivo)
	escribirEnArch(texto, name_archivo)
	run(name_imagen, name_archivo)

}
func NewMatriz() *Matriz {
	return &Matriz{Raiz: &NodoMatriz{CoorX: -1, CoorY: -1, Color: "Raiz"}}
}
