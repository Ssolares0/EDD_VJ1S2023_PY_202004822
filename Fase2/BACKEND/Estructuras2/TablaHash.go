package Estructuras2

import (
	"math"
	"strconv"
)

type NodoHash struct {
	Key       int
	IdCliente string
	IdFactura string
}

type TablaHash struct {
	Tabla       [100]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) calcularIndice(idCliente int) int {
	indice := (22*idCliente + 2020204822) % t.Capacidad
	return indice
}

func (t *TablaHash) CapacidadTabla() {
	auxCapacidad := float64(t.Capacidad) * 0.60
	if t.Utilizacion > int(auxCapacidad) {
		t.Capacidad = 0
	}
}

func (t *TablaHash) newCapacidad() int {
	num := t.Capacidad + 1
	for !t.esPrimo(num) {
		num++
	}
	return num

}

func (t *TablaHash) esPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero == 2 {
		return true
	}
	if numero%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(numero))); i += 2 {
		if numero%i == 0 {
			return false
		}

	}
	return true
}
func (t *TablaHash) InsertarElemento(idCliente string, idFactura string) {
	numeroVar, _ := strconv.Atoi(idCliente)
	indice := t.calcularIndice(numeroVar)
	newNodo := &NodoHash{Key: indice, IdCliente: idCliente, IdFactura: idFactura}
	if indice < t.Capacidad {
		if t.Tabla[indice].Key == -1 {
			t.Tabla[indice] = *newNodo
			t.Utilizacion++

		}
	}
}
