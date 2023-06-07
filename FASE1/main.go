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
		fmt.Println("Se ha logeado con exito, bienvenido admin!!")
		menu_administrador()

	} else {
		id, _ := strconv.Atoi(nombre)

		comprobar := ListaNuevaEmpleados.BuscarEmpleado(id, password)
		if comprobar == nil {
			fmt.Println("Usuario o contraseña incorrectos, intente de nuevo!!")

		} else {
			fmt.Println("Se ha logeado con exito!! Bienvenido: ", id)
			menu_empleado()

		}

	}

}
func menu_administrador() {
	var opc int = 0
	for opc != 7 {
		fmt.Println("******MENU ADMINISTRADOR******")
		fmt.Println("1. Cargar Empleados")
		fmt.Println("2. Cargar Imagenes")
		fmt.Println("3. Cargar Usuarios")
		fmt.Println("4. Actualizar Cola")
		fmt.Println("5. Reportes Estructuras")
		fmt.Println("6. Salir del apartado de administrador")
		fmt.Println("Ingrese una opción: ")
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
		case 6:
			menu_login()

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
	if err != nil {
		return
	}
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
	ListaNuevaEmpleados.Grafico()

}
func menu_empleado() {
	var opc int = 0
	for opc != 4 {
		fmt.Println("******MENU EMPLEADO******")
		fmt.Println("1. Ver imagenes Cargadas")
		fmt.Println("2. Realizar Pedido")
		fmt.Println("3. Salir del apartado de empleado")
		fmt.Println("4. Salir del sistema")
		fmt.Println("Ingrese una opción: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			fmt.Println("Ver imagenes Cargadas")
		case 2:
			fmt.Println("Realizar Pedido")
		case 3:
			menu_login()

		}
	}
}

func main() {
	menu_login()

}
