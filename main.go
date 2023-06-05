package main

import (
	"EDD_VJ1S2023_PY_202004822/Estructuras"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var ListaNuevaEmpleados = Estructuras.New_Lista()

func menu_login() {
	var opc int = 0
	for opc != 2 {
		fmt.Println("******LOGIN******")
		fmt.Println("1. Iniciar sesión")
		fmt.Println("2. Salir del sistema")
		fmt.Println("Ingrese una opción: ")
		fmt.Scanln(&opc)

		if opc == 1 {
			login()
		}

	}

}
func login() {
	fmt.Print("Ingresa tu usuario: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()
	fmt.Print("Ingresa tu contraseña: ")
	scanner.Scan()
	password := scanner.Text()
	if nombre == "ADMIN_202004822" && password == "Admin" {
		fmt.Println("Se ha logeado con exito!!")
		menu_administrador()

	} else {
		fmt.Println("comprobando...")

	}

}
func menu_administrador() {
	var opc int = 0
	for opc != 6 {
		fmt.Println("******MENU ADMINISTRADOR******")
		fmt.Println("1. Cargar Empleados")
		fmt.Println("2. Cargar Imagenes")
		fmt.Println("3. Cargar Usuarios")
		fmt.Println("4. Actualizar Cola")
		fmt.Println("5. Reportes Estructuras")
		fmt.Println("6. Salir del sistema")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			fmt.Println("Cargar Empleados")
			cargar_Empleados()
		case 2:
			fmt.Println("Cargar Imagenes")
		case 3:
			fmt.Println("Cargar Usuarios")
		case 4:
			fmt.Println("Actualizar Cola")
		case 5:
			fmt.Println("Reportes Estructuras")

		}
	}
}
func cargar_Empleados() {
	fmt.Println("carga masiva de Empleados")
	fmt.Println("ingrese la ruta del archibo csv")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()
	fmt.Println(nombre)

	file, err := os.Open(nombre)
	if err != nil {
		fmt.Println(err)
	}
	records, err := csv.NewReader(file).ReadAll()
	for _, record := range records {
		if record[0] == "id" {
			continue
		}
		sv, _ := strconv.Atoi(record[0])
		ListaNuevaEmpleados.AgregarEmpleado(record[1], sv, record[2], record[3])

	}
	mostrarEmpleados()

}

func mostrarEmpleados() {
	fmt.Println("*****************************************")
	Estructuras.MostrarLista(ListaNuevaEmpleados)
}

func main() {
	menu_login()

}
