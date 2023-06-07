package Estructuras

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/*
Este archivo create.go se encargara de crear los archivos .dot y sobrescibir en el, para
luego poder graficarlos con graphviz

*/

/*
funcion createarch se encarga de crear el archivo .dot
*/
func createArch(nombre string) {
	var _, err = os.Stat(nombre)

	if os.IsNotExist(err) {
		var file, err = os.Create(nombre)
		if err != nil {
			return
		}
		defer file.Close()
	} else {
		var filee, err = os.Create(nombre)
		filee.WriteString("")
		if err != nil {
			return
		}
		defer filee.Close()

	}
}

/*
funcion escribirEnArch se encarga de sobrescribir en el archivo .dot
*/
func escribirEnArch(texto string, name_archivo string) {
	var file, err = os.OpenFile(name_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Escribimos texto linea por linea
	_, err = file.WriteString(texto)
	if err != nil {
		return
	}
	// Guarda cambios del archivo
	err = file.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo modificado existosamente.")
}

func run(name_imagen string, name_archivo string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", name_archivo).Output()
	mode := 0777
	_ = ioutil.WriteFile(name_imagen, cmd, os.FileMode(mode))

}
