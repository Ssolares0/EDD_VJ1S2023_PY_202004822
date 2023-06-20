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
	r.HandleFunc("/ViewTree", MostrarArbol).Methods("GET")

	//Recibimos datos del frontend
	r.HandleFunc("/AddTree", AgregarTree).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", r))

}

func MostrarArbol(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tree)
}

func AgregarTree(w http.ResponseWriter, req *http.Request) {
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
