package main

import (
	"FASE2/BACKEND/Estructuras2"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tree *Estructuras2.Arbol
var ListaNuevaEmpleados = Estructuras2.New_Lista()

type ResponseImage struct {
	Base64 string
	Name   string
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

var UserNew User
var RutasNew Rutas

func main() {
	tree = &Estructuras2.Arbol{Raiz: nil}
	r := mux.NewRouter()

	//Mostramos el arbol
	r.HandleFunc("/ViewTree", ViewTree).Methods("GET")

	//Recibimos datos del frontend
	r.HandleFunc("/AddTree", AddTree).Methods("POST")

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
		w.WriteHeader(http.StatusCreated)
		UserNew.Admin = false
		json.NewEncoder(w).Encode(UserNew)
	}

}

func ObtenerDatosLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserNew)

}

func CargaMasiva(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &RutasNew)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RutasNew)

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
		ListaNuevaEmpleados.AgregarEmpleado(record[1], sv, record[2], record[3])

	}

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

// Funcion que nos permite agregar un nuevo nodo al arbol
func AddTree(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	var newNodo Estructuras2.NodoArbol
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &newNodo)
	tree.InsertarElemento(newNodo.Valor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNodo)
}

// Funcion que nos permite enviar el reporte del arbol
func SendReporte(w http.ResponseWriter, req *http.Request) {
	var image ResponseImage = ResponseImage{Name: "cola.jpg"}

	imageBytes, err := ioutil.ReadFile(image.Name)
	if err != nil {
		fmt.Fprintf(w, "La imagen no es valida")
	}
	image.Base64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
	fmt.Fprintf(w, "%+v", image)
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
