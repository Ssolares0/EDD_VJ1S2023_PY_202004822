package main

import (
	"EDD_VJ1S2023_PY_202004822/Estructuras"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var ListaNuevaEmpleados = Estructuras.New_Lista()
var ListaNuevaImagenes = Estructuras.New_ListaDoble()
var ListaNuevaClientes = Estructuras.New_Lista_circularSimp()
var ListaNuevaClientesPend = Estructuras.New_ListaCola()
var ListaNuevaPila = Estructuras.New_Lista_pila()

var ListaNuevaDispersa = Estructuras.NewMatriz()
var ListaNuevaPedidos = Estructuras.New_Lista_pila()

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
		if opc == 2 {
			fmt.Println("Saliendo del sistema")
			os.Exit(1)
		}

	}

}

func login() {
	fmt.Println("Bienvenido al sistema de login!!")
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
		IDEMPLEADO, _ := strconv.Atoi(nombre)

		comprobar := ListaNuevaEmpleados.BuscarEmpleado(IDEMPLEADO, password)
		if comprobar == nil {
			fmt.Println("Usuario o contraseña incorrectos, intente de nuevo!!")

		} else {
			fmt.Println("Se ha logeado con exito!! Bienvenido: ", IDEMPLEADO)
			menu_empleado()

		}

	}

}

func menu_empleado() {
	var opc int = 0
	for opc != 5 {
		fmt.Println("******MENU EMPLEADO******")
		fmt.Println("1. Ver imagenes Cargadas")
		fmt.Println("2. Realizar Pedido")
		fmt.Println("3. IMAGEN FINAL")
		fmt.Println("4. Resetear Matriz")
		fmt.Println("5. Salir del apartado de empleado")
		fmt.Println("6. Salir del sistema")
		fmt.Println("Ingrese una opción: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			fmt.Println("****Ver imagenes Cargadas****")
			mostarImagenesCargadas()

		case 2:
			fmt.Println("Realizar Pedido")
			realizarPedido(ListaNuevaClientesPend, ListaNuevaClientes, ListaNuevaImagenes)

		case 3:
			GenerarImagenFinal()

		case 4:
			reset_matriz()

		case 5:
			menu_login()

		}

	}
}

func menu_administrador() {
	var opc int = 0
	for opc != 7 {
		fmt.Println("******MENU ADMINISTRADOR******")
		fmt.Println("1. Cargar Empleados")
		fmt.Println("2. Cargar Imagenes")
		fmt.Println("3. Cargar Clientes")
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
			cargar_Imagenes()
		case 3:
			fmt.Println("Cargar Clientes")
			cargar_Clientes()
		case 4:
			fmt.Println("Actualizar Cola")
			actualizar_Cola()
		case 5:
			fmt.Println("Reportes Estructuras")
			ListaNuevaEmpleados.Grafico()
			ListaNuevaImagenes.GraficoDoble()
			ListaNuevaClientes.GraficoCircular()
			ListaNuevaClientesPend.GraficarCola()
			ListaNuevaPila.GraficarPila()
			crearJSON(ListaNuevaPila)

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

func cargar_Imagenes() {
	fmt.Println("carga masiva de Imagenes")
	fmt.Println("ingrese la ruta del archivo csv")
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
		if (record[0] == "Imagen") || (record[0] == "imagen") {
			continue
		}

		sv, _ := strconv.Atoi(record[1])
		ListaNuevaImagenes.AgregarImagen(record[0], sv)
	}

}

func cargar_Clientes() {
	fmt.Println("carga masiva de Clientes")
	fmt.Println("ingrese la ruta del archivo csv")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()
	file, err := os.Open(nombre)

	if err != nil {
		return
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
		ListaNuevaClientes.AgregarCliente(sv, record[1])
	}
	mostrarClientesCargados()
}

func actualizar_Cola() {

	//Nuevos := []string{}
	//NuevosId := []int{}
	fmt.Println("Actualizando Cola")
	fmt.Println("ingrese la ruta del archivo csv")

	fmt.Println("carga masiva de Clientes en cola")
	fmt.Println("ingrese la ruta del archivo csv")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nombre := scanner.Text()
	file, err := os.Open(nombre)

	if err != nil {
		return
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}
	for _, record := range records {
		if record[0] == "id" {
			continue
		}

		ListaNuevaClientesPend.Colar(record[0], record[1])

	}
	mostrarClientesencola()

}

func mostrarEmpleados() {
	fmt.Println("*****************************************")

	Estructuras.MostrarLista(ListaNuevaEmpleados)

}
func mostarImagenesCargadas() {
	fmt.Println("*****************************************")
	Estructuras.MostrarListaDoble(ListaNuevaImagenes)
	GenerarImagen()

}
func mostrarClientesCargados() {
	fmt.Println("*****************************************")
	Estructuras.MostrarListaCircular(ListaNuevaClientes)

}

func mostrarClientesencola() {
	fmt.Println("*****************************************")
	Estructuras.MostrarCola(ListaNuevaClientesPend)

}

func realizarPedido(cActual *Estructuras.Lista_cola, cli *Estructuras.Lista_circularSimp, cimagen *Estructuras.ListaDoble) {
	fmt.Println("*****************************************")
	fmt.Println("Realizar Pedido")
	// Inicializar el generador de números aleatorios con una semilla única
	rand.Seed(time.Now().UnixNano())

	// Generar un número aleatorio de 4 dígitos
	numeroAleatorio := rand.Intn(9000) + 1000

	existente := verificarCola(ListaNuevaClientesPend, ListaNuevaClientes)
	if existente && cActual.Inicio != nil {

		fmt.Println("El usuario actual es: ", cActual.Inicio.Cliente.Nombre)
		Estructuras.MostrarListaDoble(ListaNuevaImagenes)

		fmt.Println("Eliga una pelicula: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nombrepelicula := scanner.Text()
		ListaNuevaPila.Push(cActual.Inicio.Cliente.Id, nombrepelicula)
		cActual.Descolar()

	} else if !existente && cActual.Inicio != nil {
		fmt.Println("El usuario actual es: ", cActual.Inicio.Cliente.Nombre)

		//Agregamos el cliente nuevo a la lista circular
		ListaNuevaClientes.AgregarCliente(numeroAleatorio, cActual.Inicio.Cliente.Nombre)

		//mostramos las imagenes disponibles
		Estructuras.MostrarListaDoble(ListaNuevaImagenes)

		fmt.Println("Eliga una pelicula: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nombrepelicula := scanner.Text()

		numeroAleatorioString := strconv.Itoa(numeroAleatorio)
		ListaNuevaPila.Push(numeroAleatorioString, nombrepelicula)
		cActual.Descolar()

	}

}

func verificarCola(cActual *Estructuras.Lista_cola, cli *Estructuras.Lista_circularSimp) bool {
	fmt.Println("verificando cola")
	aux := cActual.Inicio
	aux2 := cli.Inicio
	if aux != nil {

		if aux.Cliente.Id != "x" {

			for i := 0; i < cli.Longitud; i++ {
				sv, _ := strconv.Atoi(aux.Cliente.Id)
				if sv == aux2.Cliente.Id {
					fmt.Println("Cliente encontrado")

					return true
				}

			}
			//aux2 = aux2.Siguiente

		} else {

			return false

		}
		//aux = aux.Siguiente
	}

	return false
}

func crearJSON(l *Estructuras.Lista_pila) {
	fmt.Println("Creando JSON")

	//Escribir en el archivo
	archivoJSON := "Reportes/" + "reporteJSON" + ".json"
	contenidoJSON := ("{\n" + "\"Pedidos\":[\n")
	for i := 0; i < l.Longitud; i++ {
		contenidoJSON += ("{\n" + "\"IdCliente\":\"" + l.Inicio.ClientePila.Idpila + "\",\n")
		contenidoJSON += ("\"NombreCliente\":\"" + l.Inicio.ClientePila.Nombrepila + "\" \n")
		contenidoJSON += ("},\n")
		l.Inicio = l.Inicio.Siguiente
	}
	contenidoJSON += "{ \n"
	contenidoJSON += "} \n"
	contenidoJSON += "] \n"
	contenidoJSON += "}"

	var _, err = os.Stat(archivoJSON)

	if os.IsNotExist(err) {
		var file, err = os.Create(archivoJSON)
		if err != nil {
			return
		}
		defer file.Close()
	} else {
		var filee, err = os.Create(archivoJSON)
		filee.WriteString("")
		if err != nil {
			return
		}
		defer filee.Close()

	}

	var file2, err2 = os.OpenFile(archivoJSON, os.O_RDWR, 0644)
	if err2 != nil {
		return
	}
	defer file2.Close()
	// Escribimos texto linea por linea
	_, err2 = file2.WriteString(contenidoJSON)
	if err2 != nil {
		return
	}
	// Guarda cambios del archivo
	err2 = file2.Sync()
	if err2 != nil {
		return
	}
	fmt.Println("Archivo modificado existosamente.")
	//Guardar cambios

}
func reset_matriz() {
	ListaNuevaDispersa = Estructuras.NewMatriz()
}

func GenerarImagen() {
	//matriz := &Estructuras.Matriz{Raiz: &Estructuras.NodoMatriz{CoorX: -1, CoorY: -1, Color: "RAIZ"}}

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

	fmt.Println("Ingrese la capa que desea empezando de la 1: ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	capaSeleccionada := scanner2.Text()

	capaSeleccionadaInt, _ := strconv.Atoi(capaSeleccionada)

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
	/*fmt.Println("Layer:  ", Layer)
	fmt.Println("File:  ", File)*/

	for i := 0; i < len(Layer); i++ {
		if Layer[i] == 0 {
			config := File[i]
			fmt.Println("La config es:  " + config)

			rutaconfig := "csv/" + seleccionada + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, seleccionada)
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else if Layer[i] == capaSeleccionadaInt {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + seleccionada + "/" + File[i]

			file, err := os.Open(rutaCapa)
			if err != nil {
				fmt.Println("No pude abrir el archivo")
				return
			}
			defer file.Close()

			lectura := csv.NewReader(file)
			lectura.Comma = ','
			x := 0
			y := 0
			for {
				linea, err := lectura.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("No pude leer la linea del csv")
					continue
				}
				for i := 0; i < len(linea); i++ {
					if linea[i] != "x" {
						//ListaNuevaDispersa.AgregarElementos(x, y, linea[i])
						//matriz.AgregarElementos(x, y, linea[i])
						ListaNuevaDispersa.AgregarElementos(x, y, linea[i])

					}
					x++
				}
				x = 0
				y++
			}

		}

	}
	ListaNuevaDispersa.ReporteGraphviz(seleccionada, capaSeleccionada)
	reset_matriz()

}

func GenerarImagenFinal() {
	Estructuras.MostrarListaDoble(ListaNuevaImagenes)
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
	/*fmt.Println("Layer:  ", Layer)
	fmt.Println("File:  ", File)*/

	for i := 0; i < len(Layer); i++ {
		if Layer[i] == 0 {
			config := File[i]
			fmt.Println("La config es:  " + config)

			rutaconfig := "csv/" + seleccionada + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, seleccionada)
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + seleccionada + "/" + File[i]

			file, err := os.Open(rutaCapa)
			if err != nil {
				fmt.Println("No pude abrir el archivo")
				return
			}
			defer file.Close()

			lectura := csv.NewReader(file)
			lectura.Comma = ','
			x := 0
			y := 0
			for {
				linea, err := lectura.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("No pude leer la linea del csv")
					continue
				}
				for i := 0; i < len(linea); i++ {
					if linea[i] != "x" {
						//ListaNuevaDispersa.AgregarElementos(x, y, linea[i])
						//matriz.AgregarElementos(x, y, linea[i])
						ListaNuevaDispersa.AgregarElementos(x, y, linea[i])

					}
					x++
				}
				x = 0
				y++
			}

		}

	}

	ListaNuevaDispersa.Css()
	reset_matriz()
}

func main() {
	menu_login()

}
