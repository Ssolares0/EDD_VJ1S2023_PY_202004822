package main

import (
	"FASE2/BACKEND/Estructuras2"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tree *Estructuras2.Arbol

func main() {
	tree = &Estructuras2.Arbol{Raiz: nil}
	r := mux.NewRouter()
	//Mostramos el arbol
	r.HandleFunc("/ViewTree", ViewTree).Methods("GET")

	//Recibimos datos del frontend
	r.HandleFunc("/AddTree", AddTree).Methods("POST")

	//Reseteamos el arbol
	r.HandleFunc("/ResetTree", ResetTree).Methods("DELETE")

	//Login
	r.HandleFunc("/Login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", r))

}

// Funcion que nos permite loguearnos
func Login(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	fmt.Fprintf(w, "%+v", string(reqBody))
	/*var user Estructuras2.Usuario
	if err != nil {
		fmt.Fprintf(w, "No es valido!!")
	}
	json.Unmarshal(reqBody, &user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	*/
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

// Funcion que nos permite resetear el arbol
func ResetTree(w http.ResponseWriter, req *http.Request) {
	tree = &Estructuras2.Arbol{Raiz: nil}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tree)
}
