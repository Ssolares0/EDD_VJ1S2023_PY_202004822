package Estructuras

import (
	//"io/ioutil"
	"fmt"
	"os"
	//"os/exec"
)

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

func escribirEnArch() {
	fmt.Println("escribir en archivo")
}
