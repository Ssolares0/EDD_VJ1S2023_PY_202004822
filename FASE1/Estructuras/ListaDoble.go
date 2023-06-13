package Estructuras

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Imagen struct {
	nombre string
	capas  int
}

type NodoDoble struct {
	imagen    *Imagen
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

func (l *ListaDoble) AgregarImagen(nombre string, capas int) {
	nuevoImagen := &Imagen{nombre, capas}

	if l.validar_vacia() {
		l.inicio = &NodoDoble{nuevoImagen, nil, nil}
		l.longitud++

	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoDoble{nuevoImagen, nil, aux}
		l.longitud++

	}
}

func MostrarListaDoble(l *ListaDoble) {
	auxiliar := l.inicio
	for auxiliar != nil {
		fmt.Printf("Nombre de la imagen: %s, Capa: %d \n", auxiliar.imagen.nombre, auxiliar.imagen.capas)
		println("------------------------")
		auxiliar = auxiliar.siguiente
	}
	GenerarImagen()

}

func MostrarNormal(l *ListaDoble) {
	auxiliar := l.inicio
	for auxiliar != nil {
		fmt.Printf("Nombre de la imagen: %s, Capa: %d \n", auxiliar.imagen.nombre, auxiliar.imagen.capas)
		println("------------------------")
		auxiliar = auxiliar.siguiente
	}
}

func (l *ListaDoble) GraficoDoble() {
	name_archivo := "Reportes/ListaImagenes.dot"
	name_imagen := "Reportes/ListaImagenes.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.imagen.nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	createArch(name_archivo)
	escribirEnArch(texto, name_archivo)
	run(name_imagen, name_archivo)

}
func GenerarImagen() {
	Layer := []int{}
	File := []string{}
	//Config := []string{}
	//Value := []int{}
	image_width := 0
	image_height := 0
	pixel_width := 0
	pixel_height := 0

	fmt.Println("Ingrese el nombre de la imagen que desea seleccionar")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seleccionada := scanner.Text()

	ruta := "csv/" + seleccionada + "/inicial.csv"
	fmt.Println("La ruta es:  " + ruta)

	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println(err)
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}
	for _, record := range records {
		if (record[0] == "Layer") || (record[0] == "layer") {
			continue
		}

		sv, _ := strconv.Atoi(record[0])
		Layer = append(Layer, sv)
		File = append(File, record[1])

	}
	fmt.Println("Layer:  ", Layer)
	fmt.Println("File:  ", File)

	for i := 0; i < len(Layer); i++ {
		if Layer[i] == 0 {
			config := File[i]
			fmt.Println("La config es:  " + config)

			rutaconfig := "csv/" + seleccionada + "/" + config
			fmt.Println("La ruta de la configuracion es :  " + rutaconfig)

			file2, err2 := os.Open(rutaconfig)
			if err2 != nil {
				fmt.Println(err2)
			}
			records2, err2 := csv.NewReader(file2).ReadAll()
			if err2 != nil {
				return
			}
			for _, record2 := range records2 {
				if (record2[0] == "config") || (record2[0] == "Config") {
					continue
				}
				if record2[0] == "image_width" {
					image_width, _ = strconv.Atoi(record2[1])
				}

				if record2[0] == "image_height" {
					image_height, _ = strconv.Atoi(record2[1])
				}
				if record2[0] == "pixel_width" {
					pixel_width, _ = strconv.Atoi(record2[1])
				}
				if record2[0] == "pixel_height" {
					pixel_height, _ = strconv.Atoi(record2[1])
				}

			}
			//NewMatriz.AgregarValues(image_width, image_height, pixel_width, pixel_height)
			fmt.Println(image_width, image_height, pixel_width, pixel_height)

		} else {
			fmt.Println("La capa es:  ", Layer[i], File[i])
		}

	}

}

func New_ListaDoble() *ListaDoble {
	return &ListaDoble{nil, 0}
}
