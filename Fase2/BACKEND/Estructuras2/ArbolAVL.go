package Estructuras2

import (
	"math"
)

type NodoArbol struct {
	Izquierdo  *NodoArbol
	Derecho    *NodoArbol
	Valor      int
	Altura     int
	Equilibrio int
}

type Arbol struct {
	Raiz *NodoArbol
}

func (a *Arbol) InsertarElemento(valor int) {
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *Arbol) rotationIzq(raiz *NodoArbol) *NodoArbol {
	root_Derecho := raiz.Derecho
	hijo_Izquierdo := root_Derecho.Izquierdo
	root_Derecho.Izquierdo = raiz
	raiz.Derecho = hijo_Izquierdo
	numMax := math.Max(float64(a.Height(raiz.Izquierdo)), float64(a.Height(raiz.Derecho)))
	raiz.Altura = 1 + int(numMax)
	numMax = math.Max(float64(a.Height(root_Derecho.Izquierdo)), float64(a.Height(root_Derecho.Derecho)))
	root_Derecho.Altura = 1 + int(numMax)
	raiz.Equilibrio = a.Balance(raiz)
	root_Derecho.Equilibrio = a.Balance(root_Derecho)
	return root_Derecho
}

func (a *Arbol) rotationDer(raiz *NodoArbol) *NodoArbol {
	root_Izquierdo := raiz.Izquierdo
	hijo_Derecho := root_Izquierdo.Derecho
	root_Izquierdo.Derecho = raiz
	raiz.Izquierdo = hijo_Derecho
	numMax := math.Max(float64(a.Height(raiz.Izquierdo)), float64(a.Height(raiz.Derecho)))
	raiz.Altura = 1 + int(numMax)
	numMax = math.Max(float64(a.Height(root_Izquierdo.Izquierdo)), float64(a.Height(root_Izquierdo.Derecho)))
	root_Izquierdo.Altura = 1 + int(numMax)
	raiz.Equilibrio = a.Balance(raiz)
	root_Izquierdo.Equilibrio = a.Balance(root_Izquierdo)
	return root_Izquierdo
}

func (a *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	numMax := math.Max(float64(a.Height(raiz.Izquierdo)), float64(a.Height(raiz.Derecho)))
	raiz.Altura = 1 + int(numMax)

	swinging := a.Balance(raiz)
	raiz.Equilibrio = swinging
	//Rotacion simple a la izquierda
	if swinging > 1 && nuevoNodo.Valor > raiz.Derecho.Valor {
		return a.rotationIzq(raiz)

	}

	if swinging < -1 && nuevoNodo.Valor < raiz.Izquierdo.Valor {
		return a.rotationDer(raiz)
	}

	if swinging > 1 && nuevoNodo.Valor < raiz.Derecho.Valor {
		raiz.Derecho = a.rotationDer(raiz.Derecho)
		return a.rotationIzq(raiz)
	}

	if swinging < -1 && nuevoNodo.Valor > raiz.Izquierdo.Valor {
		raiz.Izquierdo = a.rotationIzq(raiz.Izquierdo)
		return a.rotationDer(raiz)
	}

	return raiz
}

// Funcion que nos permite obtener el factor de equilibrio
func (a *Arbol) Balance(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (a.Height(raiz.Derecho) - a.Height(raiz.Izquierdo))
}

// Funcion que nos permite obtener la altura del arbol
func (a *Arbol) Height(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}
