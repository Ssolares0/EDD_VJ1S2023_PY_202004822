package main

import (
	"FASE2/BACKEND/Estructuras2"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tree *Estructuras2.Arbol
var simple *Estructuras2.Lista
var cola *Estructuras2.Lista_cola

type ResponseImage struct {
	ImageBase64 string
	Name        string
}

type TipoFiltro struct {
	Filtro string
}

type User struct {
	Username string
	Password string
	Admin    bool
}

type Rutas struct {
	Empleados string
	Pedidos   string
}

// Estructura para el reporte de pedidos
type Pedido struct {
	IDCliente int    `json:"id_cliente"`
	Imagen    string `json:"imagen"`
	Cliente   string `json:"cliente"`
}
type Datos struct {
	Pedidos []Pedido `json:"pedidos"`
}

type DatosFactura struct {
	Fecha       string
	Id_Empleado int
	Id_Cliente  int
	Pago        float64
}

var UserNew User
var RutasNew Rutas
var DatosNew Datos
var TipoFiltroNew TipoFiltro
var ListaNuevaSimple = Estructuras2.New_Lista()
var ListaNuevaCola = Estructuras2.New_ListaCola()
var ListaNuevaArbol = Estructuras2.New_Arbol()
var ListaNuevaDispersa = Estructuras2.NewMatriz()

func main() {
	tree = &Estructuras2.Arbol{Raiz: nil}
	simple = &Estructuras2.Lista{Inicio: nil}

	/*
		app := fiber.New()
		app.Use(cors.New())

		app.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(&fiber.Map{
				"status": "ok",


			})

		})

		app.Post("/AddTree", func(c *fiber.Ctx) error {
			var newNodo Estructuras2.NodoArbol
			c.BodyParser(&newNodo)
			tree.Insertar(newNodo)
			return c.JSON(&fiber.Map{
				"status": "ok",
			})

		})


	*/

	r := mux.NewRouter()

	//Mostramos el arbol
	r.HandleFunc("/ViewTree", ViewTree).Methods("GET")

	//Recibimos datos del frontend
	//r.HandleFunc("/AddTree", AddTree).Methods("POST")

	r.HandleFunc("/ReporteTree", SendReporte).Methods("GET")

	//Reseteamos el arbol
	r.HandleFunc("/ResetTree", ResetTree).Methods("DELETE")

	//Login
	r.HandleFunc("/Login", Login).Methods("POST")

	//Carga masiva
	r.HandleFunc("/CargaMasiva", CargaMasiva).Methods("POST")

	//Mostrar rutas
	r.HandleFunc("/MostrarRutas", MostrarRutas).Methods("GET")

	//Obtener datos del login
	r.HandleFunc("/ObtenerDatosLogin", ObtenerDatosLogin).Methods("GET")

	//Obtener datos empleados
	r.HandleFunc("/Empleados", MostrarEmpleados).Methods("GET")

	//Crear Espejo X
	r.HandleFunc("/Filtros", Filtros).Methods("POST")

	r.HandleFunc("/MostrarFiltros", mostrarFiltros).Methods("GET")

	r.HandleFunc("/GenerarFactura", GenerarFactura).Methods("POST")

	//menu adminisrador
	//r.HandleFunc("/MenuAdmin", MenuAdmin).Methods("GET")

	http.ListenAndServe(":3001",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(r))

	log.Fatal(http.ListenAndServe(":3001", r))

}

// Funcion que nos permite loguearnos
func Login(w http.ResponseWriter, req *http.Request) {

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &UserNew)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if UserNew.Username == "admin" && UserNew.Password == "admin" {

		UserNew.Admin = true
		json.NewEncoder(w).Encode(UserNew)
	} else {
		IDEMPLEADO, _ := strconv.Atoi(UserNew.Username)

		fmt.Println(IDEMPLEADO, UserNew.Password)
		comprobar := ListaNuevaSimple.BuscarEmpleado(IDEMPLEADO, UserNew.Password)

		if comprobar == nil {
			fmt.Println("No se encontro el usuario")
		} else {
			w.WriteHeader(http.StatusCreated)
			UserNew.Admin = false
			json.NewEncoder(w).Encode(UserNew)

		}

	}

}

func ObtenerDatosLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserNew)

}

func CargaMasiva(w http.ResponseWriter, req *http.Request) {
	ListaNuevaArbol = Estructuras2.New_Arbol()
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &RutasNew)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)

	file, err := os.Open(RutasNew.Empleados)
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

		ListaNuevaSimple.AgregarEmpleado(record[1], sv, record[2], record[3])
		fmt.Println(record[0], record[1], record[2], record[3])

	}

	archJson, err := os.Open(RutasNew.Pedidos)
	if err != nil {
		fmt.Println(err)
	}
	defer archJson.Close()

	var data Datos
	decoder := json.NewDecoder(archJson)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	for _, pedido := range data.Pedidos {
		fmt.Println(pedido.IDCliente)
		ListaNuevaArbol.InsertarElemento(pedido.IDCliente, pedido.Imagen)

	}
	ListaNuevaArbol.Grafico()
	ListaNuevaArbol.Inorder()
	ListaNuevaCola = Estructuras2.New_ListaCola()
	list := ListaNuevaArbol.Raiz.ValorColado
	list2 := ListaNuevaArbol.Raiz.ImagenColado
	for x := 0; x < len(list); x++ {
		//fmt.Println(list[x], list2[x])
		ListaNuevaCola.Colar(list[x], list2[x])
	}
	ListaNuevaCola.GraficarCola()

}

// mostrar Empleados cargados
func MostrarEmpleados(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ListaNuevaSimple)
}

func MostrarRutas(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RutasNew)
}

// Funcion que nos muestra el arbol
func ViewTree(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tree)
}

// Funcion que nos permite enviar el reporte del arbol
func SendReporte(w http.ResponseWriter, req *http.Request) {

	var image ResponseImage = ResponseImage{Name: "arbolAVL.jpg"}

	imageBytes, err := ioutil.ReadFile(image.Name)
	if err != nil {
		http.Error(w, "Error al leer la imagen", http.StatusInternalServerError)
		return
	}
	//image.Base64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
	image.ImageBase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(image)
}

// Funcion que nos permite resetear el arbol
func ResetTree(w http.ResponseWriter, req *http.Request) {
	tree = &Estructuras2.Arbol{Raiz: nil}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tree)
}

func Filtros(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}

	json.Unmarshal(reqBody, &TipoFiltroNew)

	Id, image := ListaNuevaCola.MostrarPrimerValor()
	fmt.Println(Id, image)

	generarMatrizOriginal(Id, image)

	//descolamos
	//ListaNuevaCola.Descolar()

	if TipoFiltroNew.Filtro == "Negativo" {
		fmt.Println("Estas en Negativo")
		generarMatrizNegativo(Id, image)

	}
	if TipoFiltroNew.Filtro == "EscalaGrises" {
		fmt.Println("Estas en EscalaGrises")
		generarMatrizGray(Id, image)
	}

	if TipoFiltroNew.Filtro == "EspejoX" {
		fmt.Println("Estas en EspejoX")
		generarMatrizX(Id, image)
	}
	if TipoFiltroNew.Filtro == "EspejoY" {
		fmt.Println("Estas en EspejoY")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(TipoFiltroNew)
}

func mostrarFiltros(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(TipoFiltroNew)
}

func generarMatrizOriginal(id string, image string) {
	ListaNuevaDispersa = Estructuras2.NewMatriz()
	Layer := []int{}
	File := []string{}
	//Config := []string{}
	//Value := []int{}
	image_width := 0
	image_height := 0
	pixel_width := 0
	pixel_height := 0

	ruta := "csv/" + image + "/inicial.csv"
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

			rutaconfig := "csv/" + image + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, image)
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + image + "/" + File[i]

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
}

func generarMatrizX(id string, image string) {
	ListaNuevaDispersa = Estructuras2.NewMatriz()
	Layer := []int{}
	File := []string{}
	//Config := []string{}
	//Value := []int{}
	image_width := 0
	image_height := 0
	pixel_width := 0
	pixel_height := 0

	ruta := "csv/" + image + "/inicial.csv"
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

			rutaconfig := "csv/" + image + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, image+"EspejoX")
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + image + "/" + File[i]

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

	ListaNuevaDispersa.CssX()

}

func generarMatrizGray(id string, image string) {
	ListaNuevaDispersa = Estructuras2.NewMatriz()
	Layer := []int{}
	File := []string{}
	//Config := []string{}
	//Value := []int{}
	image_width := 0
	image_height := 0
	pixel_width := 0
	pixel_height := 0

	ruta := "csv/" + image + "/inicial.csv"
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

			rutaconfig := "csv/" + image + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, image+"Gray")
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + image + "/" + File[i]

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

	ListaNuevaDispersa.CssGray()

}

func generarMatrizNegativo(id string, image string) {
	ListaNuevaDispersa = Estructuras2.NewMatriz()
	Layer := []int{}
	File := []string{}
	//Config := []string{}
	//Value := []int{}
	image_width := 0
	image_height := 0
	pixel_width := 0
	pixel_height := 0

	ruta := "csv/" + image + "/inicial.csv"
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

			rutaconfig := "csv/" + image + "/" + config

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
			ListaNuevaDispersa.MandarData(image_width, image_height, pixel_width, pixel_height, image+"Negativo")
			fmt.Println("image_width:  ", image_width, "image_height:  ", image_height, "pixel_width:  ", pixel_width, "pixel_height:  ", pixel_height)

		} else {
			fmt.Println("Si entrooo")

			rutaCapa := "csv/" + image + "/" + File[i]

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

	ListaNuevaDispersa.CssNegativo()
}

func GenerarFactura(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &DatosNew)

	ListaNuevaCola.Descolar()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(DatosNew)

}
